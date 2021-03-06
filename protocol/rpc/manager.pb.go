// Code generated by protoc-gen-go.
// source: manager.proto
// DO NOT EDIT!

package rpc

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

type MGExceptionMsgReq struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *MGExceptionMsgReq) Reset()                    { *m = MGExceptionMsgReq{} }
func (m *MGExceptionMsgReq) String() string            { return proto.CompactTextString(m) }
func (*MGExceptionMsgReq) ProtoMessage()               {}
func (*MGExceptionMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *MGExceptionMsgReq) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *MGExceptionMsgReq) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *MGExceptionMsgReq) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *MGExceptionMsgReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type MGExceptionMsgRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *MGExceptionMsgRes) Reset()                    { *m = MGExceptionMsgRes{} }
func (m *MGExceptionMsgRes) String() string            { return proto.CompactTextString(m) }
func (*MGExceptionMsgRes) ProtoMessage()               {}
func (*MGExceptionMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *MGExceptionMsgRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *MGExceptionMsgRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

type MGOfflineMsgReq struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *MGOfflineMsgReq) Reset()                    { *m = MGOfflineMsgReq{} }
func (m *MGOfflineMsgReq) String() string            { return proto.CompactTextString(m) }
func (*MGOfflineMsgReq) ProtoMessage()               {}
func (*MGOfflineMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *MGOfflineMsgReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type OfflineMsg struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *OfflineMsg) Reset()                    { *m = OfflineMsg{} }
func (m *OfflineMsg) String() string            { return proto.CompactTextString(m) }
func (*OfflineMsg) ProtoMessage()               {}
func (*OfflineMsg) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *OfflineMsg) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *OfflineMsg) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *OfflineMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *OfflineMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type MGOfflineMsgRes struct {
	ErrCode uint32        `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string        `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
	Msgs    []*OfflineMsg `protobuf:"bytes,3,rep,name=msgs" json:"msgs,omitempty"`
}

func (m *MGOfflineMsgRes) Reset()                    { *m = MGOfflineMsgRes{} }
func (m *MGOfflineMsgRes) String() string            { return proto.CompactTextString(m) }
func (*MGOfflineMsgRes) ProtoMessage()               {}
func (*MGOfflineMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *MGOfflineMsgRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *MGOfflineMsgRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *MGOfflineMsgRes) GetMsgs() []*OfflineMsg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func init() {
	proto.RegisterType((*MGExceptionMsgReq)(nil), "rpc.MGExceptionMsgReq")
	proto.RegisterType((*MGExceptionMsgRes)(nil), "rpc.MGExceptionMsgRes")
	proto.RegisterType((*MGOfflineMsgReq)(nil), "rpc.MGOfflineMsgReq")
	proto.RegisterType((*OfflineMsg)(nil), "rpc.offlineMsg")
	proto.RegisterType((*MGOfflineMsgRes)(nil), "rpc.MGOfflineMsgRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ManagerServerRPC service

type ManagerServerRPCClient interface {
	SetExceptionMsg(ctx context.Context, in *MGExceptionMsgReq, opts ...grpc.CallOption) (*MGExceptionMsgRes, error)
	GetOfflineMsgs(ctx context.Context, in *MGOfflineMsgReq, opts ...grpc.CallOption) (*MGOfflineMsgRes, error)
}

type managerServerRPCClient struct {
	cc *grpc.ClientConn
}

func NewManagerServerRPCClient(cc *grpc.ClientConn) ManagerServerRPCClient {
	return &managerServerRPCClient{cc}
}

func (c *managerServerRPCClient) SetExceptionMsg(ctx context.Context, in *MGExceptionMsgReq, opts ...grpc.CallOption) (*MGExceptionMsgRes, error) {
	out := new(MGExceptionMsgRes)
	err := grpc.Invoke(ctx, "/rpc.ManagerServerRPC/SetExceptionMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServerRPCClient) GetOfflineMsgs(ctx context.Context, in *MGOfflineMsgReq, opts ...grpc.CallOption) (*MGOfflineMsgRes, error) {
	out := new(MGOfflineMsgRes)
	err := grpc.Invoke(ctx, "/rpc.ManagerServerRPC/GetOfflineMsgs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ManagerServerRPC service

type ManagerServerRPCServer interface {
	SetExceptionMsg(context.Context, *MGExceptionMsgReq) (*MGExceptionMsgRes, error)
	GetOfflineMsgs(context.Context, *MGOfflineMsgReq) (*MGOfflineMsgRes, error)
}

func RegisterManagerServerRPCServer(s *grpc.Server, srv ManagerServerRPCServer) {
	s.RegisterService(&_ManagerServerRPC_serviceDesc, srv)
}

func _ManagerServerRPC_SetExceptionMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGExceptionMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServerRPCServer).SetExceptionMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ManagerServerRPC/SetExceptionMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServerRPCServer).SetExceptionMsg(ctx, req.(*MGExceptionMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerServerRPC_GetOfflineMsgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGOfflineMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServerRPCServer).GetOfflineMsgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ManagerServerRPC/GetOfflineMsgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServerRPCServer).GetOfflineMsgs(ctx, req.(*MGOfflineMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagerServerRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ManagerServerRPC",
	HandlerType: (*ManagerServerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetExceptionMsg",
			Handler:    _ManagerServerRPC_SetExceptionMsg_Handler,
		},
		{
			MethodName: "GetOfflineMsgs",
			Handler:    _ManagerServerRPC_GetOfflineMsgs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}

func init() { proto.RegisterFile("manager.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xa5, 0xdf, 0xf2, 0x61, 0x18, 0x83, 0xe0, 0x86, 0x90, 0x86, 0x78, 0x20, 0xcb, 0x85, 0x13,
	0x07, 0xbc, 0x7b, 0x01, 0xd2, 0x70, 0x68, 0x34, 0xdb, 0xf8, 0x03, 0x6a, 0x19, 0xd6, 0x26, 0xb6,
	0x5b, 0x67, 0x17, 0xf5, 0x9f, 0xf8, 0x77, 0x4d, 0xb7, 0x95, 0xaa, 0x70, 0xf2, 0xe0, 0x6d, 0xe6,
	0xbd, 0xd9, 0xbc, 0xf7, 0x66, 0x16, 0x7a, 0x59, 0x9c, 0xc7, 0x0a, 0x69, 0x5e, 0x90, 0xb6, 0x9a,
	0x33, 0x2a, 0x12, 0xf1, 0x0a, 0x97, 0x61, 0xb0, 0x7e, 0x4b, 0xb0, 0xb0, 0xa9, 0xce, 0x43, 0xa3,
	0x24, 0x3e, 0xf3, 0x2b, 0xe8, 0x1a, 0xbd, 0xa7, 0x04, 0xef, 0x37, 0x2b, 0xdf, 0x9b, 0x78, 0x33,
	0x26, 0x1b, 0xa0, 0x64, 0x6d, 0x4c, 0x0a, 0x6d, 0xc9, 0xfe, 0xab, 0xd8, 0x03, 0xc0, 0x87, 0xf0,
	0x3f, 0x33, 0x6a, 0xb3, 0xf2, 0xd9, 0xc4, 0x9b, 0x75, 0x65, 0xd5, 0xf0, 0x01, 0xb0, 0xcc, 0x28,
	0xbf, 0xed, 0xb0, 0xb2, 0x14, 0xeb, 0x63, 0x61, 0xc3, 0x7d, 0x38, 0x43, 0xa2, 0xa5, 0xde, 0xa2,
	0x93, 0xed, 0xc9, 0xcf, 0x96, 0x8f, 0xa0, 0x83, 0x44, 0x91, 0x25, 0xa7, 0xd8, 0x95, 0x75, 0x27,
	0xa6, 0xd0, 0x0f, 0x83, 0xdb, 0xdd, 0xee, 0x29, 0xcd, 0xb1, 0x76, 0x3f, 0x00, 0xb6, 0x4f, 0xb7,
	0xb5, 0xef, 0xb2, 0x14, 0x05, 0x80, 0x3e, 0x8c, 0xfc, 0x49, 0xba, 0xc7, 0x9f, 0xb6, 0x7e, 0x91,
	0x8d, 0x4f, 0xa1, 0x9d, 0x19, 0x65, 0x7c, 0x36, 0x61, 0xb3, 0xf3, 0x45, 0x7f, 0x4e, 0x45, 0x32,
	0x6f, 0x72, 0x48, 0x47, 0x2e, 0xde, 0x3d, 0x18, 0x84, 0xd5, 0x5d, 0x23, 0xa4, 0x17, 0x24, 0x79,
	0xb7, 0xe4, 0x4b, 0xe8, 0x47, 0x68, 0xbf, 0x6e, 0x97, 0x8f, 0xdc, 0xf3, 0xa3, 0x5b, 0x8f, 0x4f,
	0xe3, 0x46, 0xb4, 0xf8, 0x0d, 0x5c, 0x04, 0x68, 0x9b, 0x10, 0x86, 0x0f, 0xeb, 0xd9, 0x6f, 0xfb,
	0x1e, 0x9f, 0x42, 0x8d, 0x68, 0x3d, 0x74, 0xdc, 0x37, 0xbb, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xbb, 0x08, 0xc4, 0x66, 0x77, 0x02, 0x00, 0x00,
}
