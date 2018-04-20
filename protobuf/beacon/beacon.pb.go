// Code generated by protoc-gen-go. DO NOT EDIT.
// source: beacon/beacon.proto

/*
Package beacon is a generated protocol buffer package.

It is generated from these files:
	beacon/beacon.proto

It has these top-level messages:
	BeaconPacket
	BeaconResponse
*/
package beacon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// BeaconPacket  holds a link to a previous signature, a timestamp and the
// partial signature for this beacon. All participants send and collects many of
// theses partial beacon packets to recreate locally one beacon
type BeaconPacket struct {
	PreviousSig []byte `protobuf:"bytes,1,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	Timestamp   uint64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	PartialSig  []byte `protobuf:"bytes,3,opt,name=partial_sig,json=partialSig,proto3" json:"partial_sig,omitempty"`
}

func (m *BeaconPacket) Reset()                    { *m = BeaconPacket{} }
func (m *BeaconPacket) String() string            { return proto.CompactTextString(m) }
func (*BeaconPacket) ProtoMessage()               {}
func (*BeaconPacket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BeaconPacket) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

func (m *BeaconPacket) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BeaconPacket) GetPartialSig() []byte {
	if m != nil {
		return m.PartialSig
	}
	return nil
}

type BeaconResponse struct {
}

func (m *BeaconResponse) Reset()                    { *m = BeaconResponse{} }
func (m *BeaconResponse) String() string            { return proto.CompactTextString(m) }
func (*BeaconResponse) ProtoMessage()               {}
func (*BeaconResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*BeaconPacket)(nil), "beacon.BeaconPacket")
	proto.RegisterType((*BeaconResponse)(nil), "beacon.BeaconResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Beacon service

type BeaconClient interface {
	NewBeacon(ctx context.Context, in *BeaconPacket, opts ...grpc.CallOption) (*BeaconResponse, error)
}

type beaconClient struct {
	cc *grpc.ClientConn
}

func NewBeaconClient(cc *grpc.ClientConn) BeaconClient {
	return &beaconClient{cc}
}

func (c *beaconClient) NewBeacon(ctx context.Context, in *BeaconPacket, opts ...grpc.CallOption) (*BeaconResponse, error) {
	out := new(BeaconResponse)
	err := grpc.Invoke(ctx, "/beacon.Beacon/NewBeacon", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Beacon service

type BeaconServer interface {
	NewBeacon(context.Context, *BeaconPacket) (*BeaconResponse, error)
}

func RegisterBeaconServer(s *grpc.Server, srv BeaconServer) {
	s.RegisterService(&_Beacon_serviceDesc, srv)
}

func _Beacon_NewBeacon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeaconPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconServer).NewBeacon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beacon.Beacon/NewBeacon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconServer).NewBeacon(ctx, req.(*BeaconPacket))
	}
	return interceptor(ctx, in, info, handler)
}

var _Beacon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "beacon.Beacon",
	HandlerType: (*BeaconServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBeacon",
			Handler:    _Beacon_NewBeacon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beacon/beacon.proto",
}

func init() { proto.RegisterFile("beacon/beacon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0x4a, 0x61, 0x67, 0x8b, 0x48, 0x14, 0x59, 0x44, 0x70, 0xed, 0x41, 0x7a, 0x6a,
	0x40, 0x4f, 0x5e, 0xeb, 0x5d, 0xa4, 0xbd, 0x79, 0x91, 0xa4, 0x19, 0xe3, 0xa0, 0x6d, 0x42, 0x92,
	0xea, 0xdf, 0x17, 0x92, 0x14, 0xdd, 0xd3, 0xcc, 0xfb, 0xe0, 0xcd, 0xbc, 0x19, 0x38, 0x97, 0x28,
	0x46, 0x33, 0xf3, 0x54, 0x5a, 0xeb, 0x4c, 0x30, 0xac, 0x4c, 0xaa, 0xb6, 0x50, 0x75, 0xb1, 0x7b,
	0x11, 0xe3, 0x27, 0x06, 0x76, 0x0b, 0x95, 0x75, 0xf8, 0x4d, 0x66, 0xf1, 0x6f, 0x9e, 0xf4, 0xae,
	0xd8, 0x17, 0x4d, 0xd5, 0x6f, 0x57, 0x36, 0x90, 0x66, 0xd7, 0xb0, 0x09, 0x34, 0xa1, 0x0f, 0x62,
	0xb2, 0xbb, 0xa3, 0x7d, 0xd1, 0x9c, 0xf4, 0x7f, 0x80, 0xdd, 0xc0, 0xd6, 0x0a, 0x17, 0x48, 0x7c,
	0x45, 0xff, 0x71, 0xf4, 0x43, 0x46, 0x03, 0xe9, 0xfa, 0x0c, 0x4e, 0xd3, 0xc6, 0x1e, 0xbd, 0x35,
	0xb3, 0xc7, 0xfb, 0x27, 0x28, 0x13, 0x61, 0x8f, 0xb0, 0x79, 0xc6, 0x9f, 0x2c, 0x2e, 0xda, 0x9c,
	0xf8, 0x7f, 0xc0, 0xab, 0xcb, 0x43, 0xba, 0x0e, 0xe9, 0x9a, 0xd7, 0x3b, 0x4d, 0xe1, 0x63, 0x91,
	0xed, 0x68, 0x26, 0xae, 0x50, 0x91, 0xe7, 0xca, 0x89, 0x59, 0xf1, 0x78, 0xb0, 0x5c, 0xde, 0xf3,
	0x03, 0x64, 0x19, 0xc1, 0xc3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x6c, 0xf8, 0x0a, 0x18,
	0x01, 0x00, 0x00,
}
