/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package servicers

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"time"

	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/certifier"
	"magma/orc8r/cloud/go/services/magmad"

	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const ChallengeExpireTime = time.Minute * 5
const ChallengeLength = 512
const TimeLength = 8 // length of time encoded in byte array from int64
const MinKeyLength = 1024
const GatewayCertificateDuration = time.Hour * 97 // 4 days, lifetime of GW Certificate

type BootstrapperServer struct {
	privKey *rsa.PrivateKey
}

func NewBootstrapperServer(privKey *rsa.PrivateKey) (*BootstrapperServer, error) {
	srv := new(BootstrapperServer)
	if privKey.N.BitLen() < MinKeyLength {
		return nil, errorLogger(fmt.Errorf("Private key is too short"))
	}
	srv.privKey = privKey
	return srv, nil
}

// generate challenge in the format of [randomText : timestamp : signature]
// the format is designed mainly for demo/interface design, subjects to change in the future
func (srv *BootstrapperServer) GetChallenge(ctx context.Context, hwId *protos.AccessGatewayID) (*protos.Challenge, error) {
	// retrieve the challenge key type
	gatewayRecord, err := magmad.FindGatewayRecordWithHwId(hwId.Id)
	if err != nil {
		return nil, errorLogger(status.Errorf(codes.NotFound, "Failed to find gateway record: %s", err))
	}

	if gatewayRecord.Key.KeyType != protos.ChallengeKey_ECHO &&
		gatewayRecord.Key.KeyType != protos.ChallengeKey_SOFTWARE_RSA_SHA256 &&
		gatewayRecord.Key.KeyType != protos.ChallengeKey_SOFTWARE_ECDSA_SHA256 {
		return nil, errorLogger(status.Errorf(codes.Aborted, "Unsupported key type: %s", gatewayRecord.Key.KeyType))
	}

	// generate random text
	randText, err := generateRandomText(ChallengeLength - TimeLength - srv.signatureLength())
	if err != nil {
		return nil, errorLogger(status.Errorf(codes.Aborted, "Failed to generate random text: %s", err))
	}

	// generate timestamp
	timeBytes := make([]byte, TimeLength)
	binary.BigEndian.PutUint64(timeBytes, uint64(time.Now().UTC().Unix()))

	// generate challenge
	challenge := append(randText, timeBytes...)
	signature, err := srv.sign(challenge)
	if err != nil {
		err = status.Errorf(codes.Aborted, "Failed to sign the challenge: %s", err)
		return nil, errorLogger(err)
	}
	challenge = append(challenge, signature...)

	return &protos.Challenge{KeyType: gatewayRecord.Key.KeyType, Challenge: challenge}, nil
}

// verify the response by client and return signed certificate if response is correct
func (srv *BootstrapperServer) RequestSign(
	ctx context.Context, resp *protos.Response) (*protos.Certificate, error) {

	hwId := resp.HwId.Id
	gatewayRecord, err := magmad.FindGatewayRecordWithHwId(hwId)
	if err != nil {
		return nil, errorLogger(status.Errorf(
			codes.NotFound, "Failed to find gateway record: %s", err))
	}

	err = srv.verifyChallenge(resp.Challenge)
	if err != nil {
		return nil, errorLogger(status.Errorf(
			codes.Aborted, "Failed to verify challenge: %s", err))
	}

	// verify authentication / real response
	switch gatewayRecord.Key.KeyType {
	case protos.ChallengeKey_ECHO:
		err = verifyEcho(resp)
	case protos.ChallengeKey_SOFTWARE_RSA_SHA256:
		err = verifySoftwareRSASHA256(resp, gatewayRecord.Key.Key)
	case protos.ChallengeKey_SOFTWARE_ECDSA_SHA256:
		err = verifySoftwareECDSASHA256(resp, gatewayRecord.Key.Key)
	default:
		err = fmt.Errorf("Unsupported key type: %s", gatewayRecord.Key.KeyType)
	}
	if err != nil {
		return nil, errorLogger(status.Errorf(
			codes.Aborted, "Failed to verify response: %s", err))
	}

	// Ignore requested cert duration & overwrite it with our own
	if resp.Csr != nil {
		resp.Csr.ValidTime = ptypes.DurationProto(GatewayCertificateDuration)
	}
	cert, err := certifier.SignCSR(resp.Csr)
	if err != nil {
		return nil, errorLogger(status.Errorf(codes.Aborted, "Failed to sign csr: %s", err))
	}
	return cert, nil
}

// return the length of signature (number of bytes)
func (srv *BootstrapperServer) signatureLength() int {
	keyLength := srv.privKey.N.BitLen()
	return keyLength / 8
}

// authenticate text (challenge) with signature
func (srv *BootstrapperServer) sign(text []byte) ([]byte, error) {
	hashed := sha256.Sum256(text)
	return rsa.SignPKCS1v15(rand.Reader, srv.privKey, crypto.SHA256, hashed[:])
}

// verify the signature of text (challenge) to confirm the text is sent by server
func (srv *BootstrapperServer) verify(text, signature []byte) error {
	hashed := sha256.Sum256(text)
	return rsa.VerifyPKCS1v15(&srv.privKey.PublicKey, crypto.SHA256, hashed[:], signature)
}

// verify the challenge is the one sent by server and is not expired
// challenge[randomText : timestamp : signature]
func (srv *BootstrapperServer) verifyChallenge(challenge []byte) error {
	if n := len(challenge); n != ChallengeLength {
		return fmt.Errorf("Wrong length for challenge, expected %d, got %d", ChallengeLength, n)
	}

	// check time
	randLen := ChallengeLength - TimeLength - srv.signatureLength()
	timeBytes := challenge[randLen : randLen+TimeLength]
	issueTime := time.Unix(int64(binary.BigEndian.Uint64(timeBytes)), 0)
	expireTime := issueTime.Add(time.Duration(ChallengeExpireTime))
	now := time.Now().UTC()
	if issueTime.After(now) {
		return fmt.Errorf("Challenge is not valid yet")

	}
	if expireTime.Before(now) {
		return fmt.Errorf("Challenge has expired")
	}

	// verify signature
	signatureIdx := ChallengeLength - srv.signatureLength()
	err := srv.verify(challenge[:signatureIdx], challenge[signatureIdx:])
	if err != nil {
		return fmt.Errorf("Failed to veriry the authenticity of challenge")
	}
	return nil
}

// generate random byte slice given length
func generateRandomText(length int) ([]byte, error) {
	if length < 0 {
		return nil, fmt.Errorf("Incorrect length")
	}
	randText := make([]byte, length)
	_, err := rand.Read(randText)
	if err != nil {
		return nil, err
	}
	return randText, nil
}

// verify response with echo "encryption" method
func verifyEcho(resp *protos.Response) error {
	response := resp.GetEchoResponse() //.Response
	if response == nil {
		return fmt.Errorf("Wrong type of response, expected Echo")
	}
	if !bytes.Equal(response.Response, resp.Challenge) {
		return fmt.Errorf("Incorrect response")
	}
	return nil
}

// verify response with RSA signature and sha256 hash
func verifySoftwareRSASHA256(resp *protos.Response, key []byte) error {
	publicKey, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return fmt.Errorf("Failed to parse RSA public key: %s", err)
	}

	response := resp.GetRsaResponse()
	if response == nil {
		return fmt.Errorf("Wrong type of response, expected RSA")
	}

	hashed := sha256.Sum256(resp.Challenge)
	err = rsa.VerifyPKCS1v15(
		publicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], response.Signature)
	return err
}

// verify response with ecdsa signature ahd sha256 hash
func verifySoftwareECDSASHA256(resp *protos.Response, key []byte) error {
	publicKey, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return fmt.Errorf("Failed to parse ECDSA public key %s", err)
	}

	response := resp.GetEcdsaResponse()
	if response == nil {
		return fmt.Errorf("Wrong type of response, expected ECDSA")
	}

	var r, s big.Int
	r.SetBytes(response.R)
	s.SetBytes(response.S)
	hashed := sha256.Sum256(resp.Challenge)
	if !ecdsa.Verify(publicKey.(*ecdsa.PublicKey), hashed[:], &r, &s) {
		return fmt.Errorf("Wrong response")
	}
	return nil
}

func errorLogger(err error) error {
	log.Printf("Bootstrapper Error: %v", err)
	return err
}
