// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hardware/hardware.proto

package hardware

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PushRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ac56d7fc2e671f, []int{0}
}

func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRequest.Unmarshal(m, b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
}
func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}
func (m *PushRequest) XXX_Size() int {
	return xxx_messageInfo_PushRequest.Size(m)
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ac56d7fc2e671f, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type GetRequest struct {
	MAC                  string   `protobuf:"bytes,1,opt,name=MAC,proto3" json:"MAC,omitempty"`
	IP                   string   `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	ID                   string   `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ac56d7fc2e671f, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetMAC() string {
	if m != nil {
		return m.MAC
	}
	return ""
}

func (m *GetRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *GetRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Hardware struct {
	JSON                 string   `protobuf:"bytes,1,opt,name=JSON,proto3" json:"JSON,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hardware) Reset()         { *m = Hardware{} }
func (m *Hardware) String() string { return proto.CompactTextString(m) }
func (*Hardware) ProtoMessage()    {}
func (*Hardware) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ac56d7fc2e671f, []int{3}
}

func (m *Hardware) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hardware.Unmarshal(m, b)
}
func (m *Hardware) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hardware.Marshal(b, m, deterministic)
}
func (m *Hardware) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hardware.Merge(m, src)
}
func (m *Hardware) XXX_Size() int {
	return xxx_messageInfo_Hardware.Size(m)
}
func (m *Hardware) XXX_DiscardUnknown() {
	xxx_messageInfo_Hardware.DiscardUnknown(m)
}

var xxx_messageInfo_Hardware proto.InternalMessageInfo

func (m *Hardware) GetJSON() string {
	if m != nil {
		return m.JSON
	}
	return ""
}

func init() {
	proto.RegisterType((*PushRequest)(nil), "github.com.tinkerbell.tink.protos.hardware.PushRequest")
	proto.RegisterType((*Empty)(nil), "github.com.tinkerbell.tink.protos.hardware.Empty")
	proto.RegisterType((*GetRequest)(nil), "github.com.tinkerbell.tink.protos.hardware.GetRequest")
	proto.RegisterType((*Hardware)(nil), "github.com.tinkerbell.tink.protos.hardware.Hardware")
}

func init() {
	proto.RegisterFile("hardware/hardware.proto", fileDescriptor_61ac56d7fc2e671f)
}

var fileDescriptor_61ac56d7fc2e671f = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x48, 0x2c, 0x4a,
	0x29, 0x4f, 0x2c, 0x4a, 0xd5, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb4, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x4a, 0x32, 0xf3, 0xb2, 0x53, 0x8b,
	0x92, 0x52, 0x73, 0x72, 0xc0, 0x4c, 0x88, 0x8a, 0x62, 0x3d, 0x98, 0x0e, 0x25, 0x45, 0x2e, 0xee,
	0x80, 0xd2, 0xe2, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0x94,
	0xc4, 0x92, 0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x89, 0x9d, 0x8b, 0xd5,
	0x35, 0xb7, 0xa0, 0xa4, 0x52, 0xc9, 0x8e, 0x8b, 0xcb, 0x3d, 0xb5, 0x04, 0xa6, 0x54, 0x80, 0x8b,
	0xd9, 0xd7, 0xd1, 0x19, 0xaa, 0x12, 0xc4, 0x14, 0xe2, 0xe3, 0x62, 0xf2, 0x0c, 0x90, 0x60, 0x02,
	0x0b, 0x30, 0x79, 0x06, 0x80, 0xf9, 0x2e, 0x12, 0xcc, 0x50, 0xbe, 0x8b, 0x92, 0x1c, 0x17, 0x87,
	0x07, 0xd4, 0x5e, 0x90, 0x45, 0x5e, 0xc1, 0xfe, 0x7e, 0x30, 0x8b, 0x40, 0x6c, 0xa3, 0xc3, 0x6c,
	0x5c, 0xfc, 0x30, 0x05, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x45, 0x5c, 0x2c, 0x20,
	0xf7, 0x09, 0x99, 0xeb, 0x11, 0xef, 0x29, 0x3d, 0x24, 0x1f, 0x49, 0x19, 0x92, 0xa2, 0x11, 0xec,
	0x4f, 0xa1, 0x52, 0x2e, 0x56, 0xa7, 0x4a, 0x90, 0x87, 0xcc, 0x48, 0xd1, 0x8b, 0x08, 0x1a, 0x29,
	0x13, 0x52, 0xf4, 0xc1, 0x83, 0xa4, 0x84, 0x8b, 0xc5, 0xa9, 0xd2, 0x33, 0x60, 0x40, 0x6c, 0x75,
	0xa1, 0xb3, 0xad, 0x05, 0x5c, 0xcc, 0x8e, 0x39, 0x39, 0x42, 0xa4, 0x47, 0x0e, 0x79, 0xf6, 0x19,
	0x30, 0x0a, 0xe5, 0x71, 0xb1, 0x79, 0xe6, 0xa5, 0x83, 0x12, 0x2e, 0x19, 0x96, 0x92, 0x91, 0x88,
	0xca, 0xb9, 0x58, 0xc3, 0x13, 0x4b, 0x92, 0x33, 0xe8, 0x1b, 0xb0, 0x06, 0x8c, 0x4e, 0x5c, 0x51,
	0x1c, 0x30, 0xe1, 0x24, 0x36, 0xb0, 0x3a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x70,
	0x0f, 0x65, 0x2b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HardwareServiceClient is the client API for HardwareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HardwareServiceClient interface {
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*Empty, error)
	ByMAC(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error)
	ByIP(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error)
	ByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error)
	All(ctx context.Context, in *Empty, opts ...grpc.CallOption) (HardwareService_AllClient, error)
	Ingest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Watch(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (HardwareService_WatchClient, error)
}

type hardwareServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHardwareServiceClient(cc grpc.ClientConnInterface) HardwareServiceClient {
	return &hardwareServiceClient{cc}
}

func (c *hardwareServiceClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.hardware.HardwareService/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hardwareServiceClient) ByMAC(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error) {
	out := new(Hardware)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.hardware.HardwareService/ByMAC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hardwareServiceClient) ByIP(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error) {
	out := new(Hardware)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.hardware.HardwareService/ByIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hardwareServiceClient) ByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error) {
	out := new(Hardware)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.hardware.HardwareService/ByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hardwareServiceClient) All(ctx context.Context, in *Empty, opts ...grpc.CallOption) (HardwareService_AllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HardwareService_serviceDesc.Streams[0], "/github.com.tinkerbell.tink.protos.hardware.HardwareService/All", opts...)
	if err != nil {
		return nil, err
	}
	x := &hardwareServiceAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HardwareService_AllClient interface {
	Recv() (*Hardware, error)
	grpc.ClientStream
}

type hardwareServiceAllClient struct {
	grpc.ClientStream
}

func (x *hardwareServiceAllClient) Recv() (*Hardware, error) {
	m := new(Hardware)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hardwareServiceClient) Ingest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.hardware.HardwareService/Ingest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hardwareServiceClient) Watch(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (HardwareService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HardwareService_serviceDesc.Streams[1], "/github.com.tinkerbell.tink.protos.hardware.HardwareService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &hardwareServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HardwareService_WatchClient interface {
	Recv() (*Hardware, error)
	grpc.ClientStream
}

type hardwareServiceWatchClient struct {
	grpc.ClientStream
}

func (x *hardwareServiceWatchClient) Recv() (*Hardware, error) {
	m := new(Hardware)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HardwareServiceServer is the server API for HardwareService service.
type HardwareServiceServer interface {
	Push(context.Context, *PushRequest) (*Empty, error)
	ByMAC(context.Context, *GetRequest) (*Hardware, error)
	ByIP(context.Context, *GetRequest) (*Hardware, error)
	ByID(context.Context, *GetRequest) (*Hardware, error)
	All(*Empty, HardwareService_AllServer) error
	Ingest(context.Context, *Empty) (*Empty, error)
	Watch(*GetRequest, HardwareService_WatchServer) error
}

// UnimplementedHardwareServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHardwareServiceServer struct {
}

func (*UnimplementedHardwareServiceServer) Push(ctx context.Context, req *PushRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (*UnimplementedHardwareServiceServer) ByMAC(ctx context.Context, req *GetRequest) (*Hardware, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByMAC not implemented")
}
func (*UnimplementedHardwareServiceServer) ByIP(ctx context.Context, req *GetRequest) (*Hardware, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByIP not implemented")
}
func (*UnimplementedHardwareServiceServer) ByID(ctx context.Context, req *GetRequest) (*Hardware, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByID not implemented")
}
func (*UnimplementedHardwareServiceServer) All(req *Empty, srv HardwareService_AllServer) error {
	return status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (*UnimplementedHardwareServiceServer) Ingest(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ingest not implemented")
}
func (*UnimplementedHardwareServiceServer) Watch(req *GetRequest, srv HardwareService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterHardwareServiceServer(s *grpc.Server, srv HardwareServiceServer) {
	s.RegisterService(&_HardwareService_serviceDesc, srv)
}

func _HardwareService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HardwareServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.hardware.HardwareService/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HardwareServiceServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HardwareService_ByMAC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HardwareServiceServer).ByMAC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.hardware.HardwareService/ByMAC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HardwareServiceServer).ByMAC(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HardwareService_ByIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HardwareServiceServer).ByIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.hardware.HardwareService/ByIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HardwareServiceServer).ByIP(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HardwareService_ByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HardwareServiceServer).ByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.hardware.HardwareService/ByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HardwareServiceServer).ByID(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HardwareService_All_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HardwareServiceServer).All(m, &hardwareServiceAllServer{stream})
}

type HardwareService_AllServer interface {
	Send(*Hardware) error
	grpc.ServerStream
}

type hardwareServiceAllServer struct {
	grpc.ServerStream
}

func (x *hardwareServiceAllServer) Send(m *Hardware) error {
	return x.ServerStream.SendMsg(m)
}

func _HardwareService_Ingest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HardwareServiceServer).Ingest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.hardware.HardwareService/Ingest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HardwareServiceServer).Ingest(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HardwareService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HardwareServiceServer).Watch(m, &hardwareServiceWatchServer{stream})
}

type HardwareService_WatchServer interface {
	Send(*Hardware) error
	grpc.ServerStream
}

type hardwareServiceWatchServer struct {
	grpc.ServerStream
}

func (x *hardwareServiceWatchServer) Send(m *Hardware) error {
	return x.ServerStream.SendMsg(m)
}

var _HardwareService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.tinkerbell.tink.protos.hardware.HardwareService",
	HandlerType: (*HardwareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _HardwareService_Push_Handler,
		},
		{
			MethodName: "ByMAC",
			Handler:    _HardwareService_ByMAC_Handler,
		},
		{
			MethodName: "ByIP",
			Handler:    _HardwareService_ByIP_Handler,
		},
		{
			MethodName: "ByID",
			Handler:    _HardwareService_ByID_Handler,
		},
		{
			MethodName: "Ingest",
			Handler:    _HardwareService_Ingest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "All",
			Handler:       _HardwareService_All_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Watch",
			Handler:       _HardwareService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hardware/hardware.proto",
}
