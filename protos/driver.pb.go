// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: driver.proto

package controler

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32  `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Cmd      string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *Command) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32  `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Data     string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *Result) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32  `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Seq      int32  `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Data     string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Ping     string `protobuf:"bytes,4,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *Request) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Request) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32  `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Seq      int32  `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Data     string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Ping     string `protobuf:"bytes,4,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *Response) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Response) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

var File_driver_proto protoreflect.FileDescriptor

var file_driver_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x22, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x60,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x22, 0x61, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x69, 0x6e, 0x67, 0x32, 0x75, 0x0a, 0x05, 0x44, 0x72, 0x69, 0x76, 0x65, 0x12, 0x2f, 0x0a, 0x04,
	0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_driver_proto_rawDescOnce sync.Once
	file_driver_proto_rawDescData = file_driver_proto_rawDesc
)

func file_driver_proto_rawDescGZIP() []byte {
	file_driver_proto_rawDescOnce.Do(func() {
		file_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_driver_proto_rawDescData)
	})
	return file_driver_proto_rawDescData
}

var file_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_driver_proto_goTypes = []interface{}{
	(*Command)(nil),  // 0: controler.Command
	(*Result)(nil),   // 1: controler.Result
	(*Request)(nil),  // 2: controler.Request
	(*Response)(nil), // 3: controler.Response
}
var file_driver_proto_depIdxs = []int32{
	0, // 0: controler.Drive.Call:input_type -> controler.Command
	2, // 1: controler.Drive.StreamCall:input_type -> controler.Request
	1, // 2: controler.Drive.Call:output_type -> controler.Result
	3, // 3: controler.Drive.StreamCall:output_type -> controler.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_driver_proto_init() }
func file_driver_proto_init() {
	if File_driver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_driver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_driver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_driver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_driver_proto_goTypes,
		DependencyIndexes: file_driver_proto_depIdxs,
		MessageInfos:      file_driver_proto_msgTypes,
	}.Build()
	File_driver_proto = out.File
	file_driver_proto_rawDesc = nil
	file_driver_proto_goTypes = nil
	file_driver_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DriveClient is the client API for Drive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriveClient interface {
	Call(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Result, error)
	StreamCall(ctx context.Context, opts ...grpc.CallOption) (Drive_StreamCallClient, error)
}

type driveClient struct {
	cc grpc.ClientConnInterface
}

func NewDriveClient(cc grpc.ClientConnInterface) DriveClient {
	return &driveClient{cc}
}

func (c *driveClient) Call(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/controler.Drive/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driveClient) StreamCall(ctx context.Context, opts ...grpc.CallOption) (Drive_StreamCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Drive_serviceDesc.Streams[0], "/controler.Drive/StreamCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &driveStreamCallClient{stream}
	return x, nil
}

type Drive_StreamCallClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type driveStreamCallClient struct {
	grpc.ClientStream
}

func (x *driveStreamCallClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *driveStreamCallClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DriveServer is the server API for Drive service.
type DriveServer interface {
	Call(context.Context, *Command) (*Result, error)
	StreamCall(Drive_StreamCallServer) error
}

// UnimplementedDriveServer can be embedded to have forward compatible implementations.
type UnimplementedDriveServer struct {
}

func (*UnimplementedDriveServer) Call(context.Context, *Command) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedDriveServer) StreamCall(Drive_StreamCallServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCall not implemented")
}

func RegisterDriveServer(s *grpc.Server, srv DriveServer) {
	s.RegisterService(&_Drive_serviceDesc, srv)
}

func _Drive_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriveServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controler.Drive/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriveServer).Call(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Drive_StreamCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DriveServer).StreamCall(&driveStreamCallServer{stream})
}

type Drive_StreamCallServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type driveStreamCallServer struct {
	grpc.ServerStream
}

func (x *driveStreamCallServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *driveStreamCallServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Drive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "controler.Drive",
	HandlerType: (*DriveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Drive_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCall",
			Handler:       _Drive_StreamCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "driver.proto",
}
