// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/streamer.proto

package protos // import "magma/orc8r/cloud/go/protos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// --------------------------------------------------------------------------
// Streamer provides a pipeline for the cloud to push the updates to the
// gateway as and when the update happens.
//
// The Streamer interface defines the semantics and consistency guarantees
// between the cloud and the gateway while abstracting the details of how
// its implemented in the cloud and what the gateway does with the updates.
//
// - The gateways call the GetUpdates() streaming API with a StreamRequest
//   indicating the stream name and the offset to continue streaming from.
// - The cloud sends a stream of DataUpdateBatch containing a batch of updates.
// - If resync is true, then the gateway can cleanup all its data and add
//   all the keys (the batch is guaranteed to contain only unique keys).
// - If resync is false, then the gateway can update the keys, or add new
//   ones if the key is not already present.
// - Key deletions are not yet supported (#15109350)
// --------------------------------------------------------------------------
type StreamRequest struct {
	GatewayId string `protobuf:"bytes,1,opt,name=gatewayId,proto3" json:"gatewayId,omitempty"`
	// Stream name to attach to. (Eg:) subscriberdb, config, etc.
	StreamName string `protobuf:"bytes,2,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	// Any extra data to send up with the stream request. This value will be
	// different per stream provider.
	ExtraArgs            *any.Any `protobuf:"bytes,3,opt,name=extra_args,json=extraArgs,proto3" json:"extra_args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_streamer_875df42a7991f84c, []int{0}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

func (m *StreamRequest) GetStreamName() string {
	if m != nil {
		return m.StreamName
	}
	return ""
}

func (m *StreamRequest) GetExtraArgs() *any.Any {
	if m != nil {
		return m.ExtraArgs
	}
	return nil
}

type DataUpdate struct {
	// Unique key for each item
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value can be file contents, protobuf serialized message, etc.
	// For key deletions, the value field would be absent.
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataUpdate) Reset()         { *m = DataUpdate{} }
func (m *DataUpdate) String() string { return proto.CompactTextString(m) }
func (*DataUpdate) ProtoMessage()    {}
func (*DataUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_streamer_875df42a7991f84c, []int{1}
}
func (m *DataUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataUpdate.Unmarshal(m, b)
}
func (m *DataUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataUpdate.Marshal(b, m, deterministic)
}
func (dst *DataUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataUpdate.Merge(dst, src)
}
func (m *DataUpdate) XXX_Size() int {
	return xxx_messageInfo_DataUpdate.Size(m)
}
func (m *DataUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_DataUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_DataUpdate proto.InternalMessageInfo

func (m *DataUpdate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataUpdate) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type DataUpdateBatch struct {
	Updates []*DataUpdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
	// If resync is true, the updates would be a snapshot of all the
	// contents in the cloud.
	Resync               bool     `protobuf:"varint,2,opt,name=resync,proto3" json:"resync,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataUpdateBatch) Reset()         { *m = DataUpdateBatch{} }
func (m *DataUpdateBatch) String() string { return proto.CompactTextString(m) }
func (*DataUpdateBatch) ProtoMessage()    {}
func (*DataUpdateBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_streamer_875df42a7991f84c, []int{2}
}
func (m *DataUpdateBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataUpdateBatch.Unmarshal(m, b)
}
func (m *DataUpdateBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataUpdateBatch.Marshal(b, m, deterministic)
}
func (dst *DataUpdateBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataUpdateBatch.Merge(dst, src)
}
func (m *DataUpdateBatch) XXX_Size() int {
	return xxx_messageInfo_DataUpdateBatch.Size(m)
}
func (m *DataUpdateBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_DataUpdateBatch.DiscardUnknown(m)
}

var xxx_messageInfo_DataUpdateBatch proto.InternalMessageInfo

func (m *DataUpdateBatch) GetUpdates() []*DataUpdate {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *DataUpdateBatch) GetResync() bool {
	if m != nil {
		return m.Resync
	}
	return false
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "magma.orc8r.StreamRequest")
	proto.RegisterType((*DataUpdate)(nil), "magma.orc8r.DataUpdate")
	proto.RegisterType((*DataUpdateBatch)(nil), "magma.orc8r.DataUpdateBatch")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamerClient is the client API for Streamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamerClient interface {
	// Get the stream of updates from the cloud.
	// The RPC call would be kept open to push new updates as they happen.
	GetUpdates(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Streamer_GetUpdatesClient, error)
}

type streamerClient struct {
	cc *grpc.ClientConn
}

func NewStreamerClient(cc *grpc.ClientConn) StreamerClient {
	return &streamerClient{cc}
}

func (c *streamerClient) GetUpdates(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Streamer_GetUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Streamer_serviceDesc.Streams[0], "/magma.orc8r.Streamer/GetUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerGetUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Streamer_GetUpdatesClient interface {
	Recv() (*DataUpdateBatch, error)
	grpc.ClientStream
}

type streamerGetUpdatesClient struct {
	grpc.ClientStream
}

func (x *streamerGetUpdatesClient) Recv() (*DataUpdateBatch, error) {
	m := new(DataUpdateBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamerServer is the server API for Streamer service.
type StreamerServer interface {
	// Get the stream of updates from the cloud.
	// The RPC call would be kept open to push new updates as they happen.
	GetUpdates(*StreamRequest, Streamer_GetUpdatesServer) error
}

func RegisterStreamerServer(s *grpc.Server, srv StreamerServer) {
	s.RegisterService(&_Streamer_serviceDesc, srv)
}

func _Streamer_GetUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamerServer).GetUpdates(m, &streamerGetUpdatesServer{stream})
}

type Streamer_GetUpdatesServer interface {
	Send(*DataUpdateBatch) error
	grpc.ServerStream
}

type streamerGetUpdatesServer struct {
	grpc.ServerStream
}

func (x *streamerGetUpdatesServer) Send(m *DataUpdateBatch) error {
	return x.ServerStream.SendMsg(m)
}

var _Streamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.Streamer",
	HandlerType: (*StreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUpdates",
			Handler:       _Streamer_GetUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orc8r/protos/streamer.proto",
}

func init() {
	proto.RegisterFile("orc8r/protos/streamer.proto", fileDescriptor_streamer_875df42a7991f84c)
}

var fileDescriptor_streamer_875df42a7991f84c = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4f, 0xf2, 0x40,
	0x10, 0xc5, 0xbf, 0x7e, 0x44, 0x84, 0xa9, 0x46, 0xb3, 0x21, 0x5a, 0x01, 0x23, 0xe9, 0x89, 0xd3,
	0x56, 0xc1, 0x83, 0x57, 0x88, 0x89, 0xd1, 0x83, 0x87, 0x25, 0x78, 0x30, 0x26, 0x64, 0x28, 0xe3,
	0x9a, 0x48, 0xbb, 0xb8, 0xbb, 0x55, 0x7b, 0xf6, 0x1f, 0x37, 0xee, 0x96, 0x20, 0x07, 0x4f, 0xed,
	0x9b, 0xfd, 0xcd, 0xbc, 0xc9, 0x1b, 0xe8, 0x28, 0x9d, 0x5e, 0xe9, 0x64, 0xa5, 0x95, 0x55, 0x26,
	0x31, 0x56, 0x13, 0x66, 0xa4, 0xb9, 0xd3, 0x2c, 0xcc, 0x50, 0x66, 0xc8, 0x1d, 0xd2, 0x3e, 0x91,
	0x4a, 0xc9, 0x25, 0x79, 0x74, 0x5e, 0x3c, 0x27, 0x98, 0x97, 0x9e, 0x8b, 0xbf, 0x02, 0xd8, 0x9f,
	0xb8, 0x56, 0x41, 0x6f, 0x05, 0x19, 0xcb, 0xba, 0xd0, 0x94, 0x68, 0xe9, 0x03, 0xcb, 0xdb, 0x45,
	0x14, 0xf4, 0x82, 0x7e, 0x53, 0x6c, 0x0a, 0xec, 0x0c, 0x42, 0xef, 0x34, 0xcb, 0x31, 0xa3, 0xe8,
	0xbf, 0x7b, 0x07, 0x5f, 0xba, 0xc7, 0x8c, 0xd8, 0x10, 0x80, 0x3e, 0xad, 0xc6, 0x19, 0x6a, 0x69,
	0xa2, 0x5a, 0x2f, 0xe8, 0x87, 0x83, 0x16, 0xf7, 0x0b, 0xf0, 0xf5, 0x02, 0x7c, 0x94, 0x97, 0xa2,
	0xe9, 0xb8, 0x91, 0x96, 0x26, 0xbe, 0x04, 0xb8, 0x46, 0x8b, 0xd3, 0xd5, 0x02, 0x2d, 0xb1, 0x43,
	0xa8, 0xbd, 0x52, 0x59, 0x79, 0xff, 0xfc, 0xb2, 0x16, 0xec, 0xbc, 0xe3, 0xb2, 0xf0, 0x7e, 0x7b,
	0xc2, 0x8b, 0xf8, 0x09, 0x0e, 0x36, 0x5d, 0x63, 0xb4, 0xe9, 0x0b, 0xbb, 0x80, 0xdd, 0xc2, 0x49,
	0x13, 0x05, 0xbd, 0x5a, 0x3f, 0x1c, 0x1c, 0xf3, 0x5f, 0x41, 0xf0, 0x0d, 0x2e, 0xd6, 0x1c, 0x3b,
	0x82, 0xba, 0x26, 0x53, 0xe6, 0xa9, 0x1b, 0xde, 0x10, 0x95, 0x1a, 0x3c, 0x40, 0x63, 0x52, 0x65,
	0xca, 0xee, 0x00, 0x6e, 0xc8, 0x4e, 0xab, 0x8e, 0xf6, 0xd6, 0xcc, 0xad, 0xf4, 0xda, 0xdd, 0x3f,
	0xfc, 0xdc, 0x7a, 0xf1, 0xbf, 0xf3, 0x60, 0x7c, 0xfa, 0xd8, 0x71, 0x48, 0xe2, 0xcf, 0x97, 0x2e,
	0x55, 0xb1, 0x48, 0xa4, 0xaa, 0xee, 0x38, 0xaf, 0xbb, 0xef, 0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x75, 0x62, 0xfb, 0xc6, 0xde, 0x01, 0x00, 0x00,
}
