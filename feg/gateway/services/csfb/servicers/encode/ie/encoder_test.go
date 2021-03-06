/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package ie_test

import (
	"fmt"
	"testing"

	"magma/feg/gateway/services/csfb/servicers/decode"
	"magma/feg/gateway/services/csfb/servicers/encode/ie"

	"github.com/stretchr/testify/assert"
)

func TestEncodeIMSI(t *testing.T) {
	// successfully encode
	imsi := "111111111111"
	_, err := ie.EncodeIMSI(imsi)
	assert.NoError(t, err)

	// value field too short
	imsi = "1"
	encodedIMSI, err := ie.EncodeIMSI(imsi)
	mandatoryFieldLength := decode.LengthIEI + decode.LengthLengthIndicator
	errorMsg := fmt.Sprintf(
		"failed to encode IMSI, value field length violation, min length: %d, max length: %d, actual length: %d",
		decode.IELengthIMSIMin-mandatoryFieldLength,
		decode.IELengthIMSIMax-mandatoryFieldLength,
		len(imsi)/2+1,
	)
	assert.EqualError(t, err, errorMsg)
	assert.Equal(t, []byte{}, encodedIMSI)

	// value field too long
	imsi = "111111111111111111111"
	encodedIMSI, err = ie.EncodeIMSI(imsi)
	errorMsg = fmt.Sprintf(
		"failed to encode IMSI, value field length violation, min length: %d, max length: %d, actual length: %d",
		decode.IELengthIMSIMin-mandatoryFieldLength,
		decode.IELengthIMSIMax-mandatoryFieldLength,
		len(imsi)/2+1,
	)
	assert.EqualError(t, err, errorMsg)
	assert.Equal(t, []byte{}, encodedIMSI)
}

func TestEncodeMMEName(t *testing.T) {
	mmeName := ".mmec01.mmegi0001.mme.EPC.mnc001.mcc001.3gppnetwork.org"
	encodedMMEName, err := ie.EncodeMMEName(mmeName)
	assert.NoError(t, err)

	encodedMMEName2, err := ie.EncodeFixedLengthIE(
		decode.IEIMMEName,
		decode.IELengthMMEName,
		[]byte(mmeName),
	)

	// replace dots with length labels
	encodedMMEName2[2] = byte(0x06)
	encodedMMEName2[9] = byte(0x09)
	encodedMMEName2[19] = byte(0x03)
	encodedMMEName2[23] = byte(0x03)
	encodedMMEName2[27] = byte(0x06)
	encodedMMEName2[34] = byte(0x06)
	encodedMMEName2[41] = byte(0x0b)
	encodedMMEName2[53] = byte(0x03)

	assert.Equal(t, encodedMMEName, encodedMMEName2)
}

func TestEncodeFixedLengthIE(t *testing.T) {
	// successfully encode TMSI
	tmsi := []byte{byte(0x11), byte(0x12), byte(0x13), byte(0x14)}
	_, err := ie.EncodeFixedLengthIE(decode.IEITMSI, decode.IELengthTMSI, tmsi)
	assert.NoError(t, err)

	// wrong length
	encodedTMSI, err := ie.EncodeFixedLengthIE(decode.IEITMSI, decode.IELengthTMSI, tmsi[:2])
	mandatoryFieldLength := decode.LengthIEI + decode.LengthLengthIndicator
	errorMsg := fmt.Sprintf(
		"failed to encode TMSI, value field length violation, length of value field should be %d, actual length: %d",
		decode.IELengthTMSI-mandatoryFieldLength,
		2,
	)
	assert.EqualError(t, err, errorMsg)
	assert.Equal(t, []byte{}, encodedTMSI)
}

func TestEncodeLimitedLengthIE(t *testing.T) {
	// successfully encode NAS message container
	nasMessageContainer := []byte{byte(0x11), byte(0x12), byte(0x13), byte(0x14)}
	_, err := ie.EncodeLimitedLengthIE(
		decode.IEINASMessageContainer,
		decode.IELengthNASMessageContainerMin,
		decode.IELengthNASMessageContainerMax,
		nasMessageContainer,
	)
	assert.NoError(t, err)

	// value field too short
	nasMessageContainer = []byte{}
	encodedNASMessageContainer, err := ie.EncodeLimitedLengthIE(
		decode.IEINASMessageContainer,
		decode.IELengthNASMessageContainerMin,
		decode.IELengthNASMessageContainerMax,
		nasMessageContainer,
	)
	mandatoryFieldLength := decode.LengthIEI + decode.LengthLengthIndicator
	errorMsg := fmt.Sprintf(
		"failed to encode NASMessageContainer, value field length violation, min length: %d, max length: %d, actual length: %d",
		decode.IELengthNASMessageContainerMin-mandatoryFieldLength,
		decode.IELengthNASMessageContainerMax-mandatoryFieldLength,
		len(nasMessageContainer),
	)
	assert.EqualError(t, err, errorMsg)
	assert.Equal(t, []byte{}, encodedNASMessageContainer)

	// value field too long
	nasMessageContainer = make([]byte, decode.IELengthNASMessageContainerMax-mandatoryFieldLength+1)
	encodedNASMessageContainer, err = ie.EncodeLimitedLengthIE(
		decode.IEINASMessageContainer,
		decode.IELengthNASMessageContainerMin,
		decode.IELengthNASMessageContainerMax,
		nasMessageContainer,
	)
	mandatoryFieldLength = decode.LengthIEI + decode.LengthLengthIndicator
	errorMsg = fmt.Sprintf(
		"failed to encode NASMessageContainer, value field length violation, min length: %d, max length: %d, actual length: %d",
		decode.IELengthNASMessageContainerMin-mandatoryFieldLength,
		decode.IELengthNASMessageContainerMax-mandatoryFieldLength,
		len(nasMessageContainer),
	)
	assert.EqualError(t, err, errorMsg)
	assert.Equal(t, []byte{}, encodedNASMessageContainer)
}

func TestEncodeVariableLengthIE(t *testing.T) {
	// successfully encode VLR name
	vlrName := "www.facebook.com"
	_, err := ie.EncodeVariableLengthIE(decode.IEIVLRName, decode.IELengthVLRNameMin, []byte(vlrName))
	assert.NoError(t, err)

	// value field too short
	vlrName = ""
	encodedVLRName, err := ie.EncodeVariableLengthIE(decode.IEIVLRName, decode.IELengthVLRNameMin, []byte(vlrName))
	mandatoryFieldLength := decode.LengthIEI + decode.LengthLengthIndicator
	errorMsg := fmt.Sprintf(
		"failed to encode VLRName, value field length violation, min length: %d, actual length: %d",
		decode.IELengthVLRNameMin-mandatoryFieldLength,
		len(vlrName),
	)
	assert.EqualError(t, err, errorMsg)
	assert.Equal(t, []byte{}, encodedVLRName)
}
