// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package pb

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

type Employee struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BadgeNumber          int32    `protobuf:"varint,2,opt,name=badgeNumber,proto3" json:"badgeNumber,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	VacationAccrualRate  float32  `protobuf:"fixed32,5,opt,name=vacationAccrualRate,proto3" json:"vacationAccrualRate,omitempty"`
	VacationAccrued      float32  `protobuf:"fixed32,6,opt,name=vacationAccrued,proto3" json:"vacationAccrued,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Employee) Reset()         { *m = Employee{} }
func (m *Employee) String() string { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()    {}
func (*Employee) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_7bc54bf07212d4bf, []int{0}
}
func (m *Employee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Employee.Unmarshal(m, b)
}
func (m *Employee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Employee.Marshal(b, m, deterministic)
}
func (dst *Employee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Employee.Merge(dst, src)
}
func (m *Employee) XXX_Size() int {
	return xxx_messageInfo_Employee.Size(m)
}
func (m *Employee) XXX_DiscardUnknown() {
	xxx_messageInfo_Employee.DiscardUnknown(m)
}

var xxx_messageInfo_Employee proto.InternalMessageInfo

func (m *Employee) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Employee) GetBadgeNumber() int32 {
	if m != nil {
		return m.BadgeNumber
	}
	return 0
}

func (m *Employee) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Employee) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Employee) GetVacationAccrualRate() float32 {
	if m != nil {
		return m.VacationAccrualRate
	}
	return 0
}

func (m *Employee) GetVacationAccrued() float32 {
	if m != nil {
		return m.VacationAccrued
	}
	return 0
}

type GetAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_7bc54bf07212d4bf, []int{1}
}
func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (dst *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(dst, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

type GetByBadgeNumberRequest struct {
	BadgeNumber          int32    `protobuf:"varint,1,opt,name=badgeNumber,proto3" json:"badgeNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByBadgeNumberRequest) Reset()         { *m = GetByBadgeNumberRequest{} }
func (m *GetByBadgeNumberRequest) String() string { return proto.CompactTextString(m) }
func (*GetByBadgeNumberRequest) ProtoMessage()    {}
func (*GetByBadgeNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_7bc54bf07212d4bf, []int{2}
}
func (m *GetByBadgeNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByBadgeNumberRequest.Unmarshal(m, b)
}
func (m *GetByBadgeNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByBadgeNumberRequest.Marshal(b, m, deterministic)
}
func (dst *GetByBadgeNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByBadgeNumberRequest.Merge(dst, src)
}
func (m *GetByBadgeNumberRequest) XXX_Size() int {
	return xxx_messageInfo_GetByBadgeNumberRequest.Size(m)
}
func (m *GetByBadgeNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByBadgeNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByBadgeNumberRequest proto.InternalMessageInfo

func (m *GetByBadgeNumberRequest) GetBadgeNumber() int32 {
	if m != nil {
		return m.BadgeNumber
	}
	return 0
}

type EmployeeRequest struct {
	Employee             *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EmployeeRequest) Reset()         { *m = EmployeeRequest{} }
func (m *EmployeeRequest) String() string { return proto.CompactTextString(m) }
func (*EmployeeRequest) ProtoMessage()    {}
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_7bc54bf07212d4bf, []int{3}
}
func (m *EmployeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeRequest.Unmarshal(m, b)
}
func (m *EmployeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeRequest.Marshal(b, m, deterministic)
}
func (dst *EmployeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeRequest.Merge(dst, src)
}
func (m *EmployeeRequest) XXX_Size() int {
	return xxx_messageInfo_EmployeeRequest.Size(m)
}
func (m *EmployeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeRequest proto.InternalMessageInfo

func (m *EmployeeRequest) GetEmployee() *Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

type EmployeeResponse struct {
	Employee             *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EmployeeResponse) Reset()         { *m = EmployeeResponse{} }
func (m *EmployeeResponse) String() string { return proto.CompactTextString(m) }
func (*EmployeeResponse) ProtoMessage()    {}
func (*EmployeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_7bc54bf07212d4bf, []int{4}
}
func (m *EmployeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeResponse.Unmarshal(m, b)
}
func (m *EmployeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeResponse.Marshal(b, m, deterministic)
}
func (dst *EmployeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeResponse.Merge(dst, src)
}
func (m *EmployeeResponse) XXX_Size() int {
	return xxx_messageInfo_EmployeeResponse.Size(m)
}
func (m *EmployeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeResponse proto.InternalMessageInfo

func (m *EmployeeResponse) GetEmployee() *Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

type AddPhotoRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPhotoRequest) Reset()         { *m = AddPhotoRequest{} }
func (m *AddPhotoRequest) String() string { return proto.CompactTextString(m) }
func (*AddPhotoRequest) ProtoMessage()    {}
func (*AddPhotoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_7bc54bf07212d4bf, []int{5}
}
func (m *AddPhotoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPhotoRequest.Unmarshal(m, b)
}
func (m *AddPhotoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPhotoRequest.Marshal(b, m, deterministic)
}
func (dst *AddPhotoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPhotoRequest.Merge(dst, src)
}
func (m *AddPhotoRequest) XXX_Size() int {
	return xxx_messageInfo_AddPhotoRequest.Size(m)
}
func (m *AddPhotoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPhotoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPhotoRequest proto.InternalMessageInfo

func (m *AddPhotoRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type AddPhotoResponse struct {
	IsOk                 bool     `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPhotoResponse) Reset()         { *m = AddPhotoResponse{} }
func (m *AddPhotoResponse) String() string { return proto.CompactTextString(m) }
func (*AddPhotoResponse) ProtoMessage()    {}
func (*AddPhotoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_7bc54bf07212d4bf, []int{6}
}
func (m *AddPhotoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPhotoResponse.Unmarshal(m, b)
}
func (m *AddPhotoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPhotoResponse.Marshal(b, m, deterministic)
}
func (dst *AddPhotoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPhotoResponse.Merge(dst, src)
}
func (m *AddPhotoResponse) XXX_Size() int {
	return xxx_messageInfo_AddPhotoResponse.Size(m)
}
func (m *AddPhotoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPhotoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddPhotoResponse proto.InternalMessageInfo

func (m *AddPhotoResponse) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func init() {
	proto.RegisterType((*Employee)(nil), "Employee")
	proto.RegisterType((*GetAllRequest)(nil), "GetAllRequest")
	proto.RegisterType((*GetByBadgeNumberRequest)(nil), "GetByBadgeNumberRequest")
	proto.RegisterType((*EmployeeRequest)(nil), "EmployeeRequest")
	proto.RegisterType((*EmployeeResponse)(nil), "EmployeeResponse")
	proto.RegisterType((*AddPhotoRequest)(nil), "AddPhotoRequest")
	proto.RegisterType((*AddPhotoResponse)(nil), "AddPhotoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	GetByBadgeNumber(ctx context.Context, in *GetByBadgeNumberRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (EmployeeService_GetAllClient, error)
	Save(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	SaveAll(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_SaveAllClient, error)
	AddPhoto(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddPhotoClient, error)
}

type employeeServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmployeeServiceClient(cc *grpc.ClientConn) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetByBadgeNumber(ctx context.Context, in *GetByBadgeNumberRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/EmployeeService/GetByBadgeNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (EmployeeService_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EmployeeService_serviceDesc.Streams[0], "/EmployeeService/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmployeeService_GetAllClient interface {
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *employeeServiceGetAllClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) Save(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/EmployeeService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) SaveAll(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_SaveAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EmployeeService_serviceDesc.Streams[1], "/EmployeeService/SaveAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceSaveAllClient{stream}
	return x, nil
}

type EmployeeService_SaveAllClient interface {
	Send(*EmployeeRequest) error
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceSaveAllClient struct {
	grpc.ClientStream
}

func (x *employeeServiceSaveAllClient) Send(m *EmployeeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceSaveAllClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) AddPhoto(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddPhotoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EmployeeService_serviceDesc.Streams[2], "/EmployeeService/AddPhoto", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceAddPhotoClient{stream}
	return x, nil
}

type EmployeeService_AddPhotoClient interface {
	Send(*AddPhotoRequest) error
	Recv() (*AddPhotoResponse, error)
	grpc.ClientStream
}

type employeeServiceAddPhotoClient struct {
	grpc.ClientStream
}

func (x *employeeServiceAddPhotoClient) Send(m *AddPhotoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceAddPhotoClient) Recv() (*AddPhotoResponse, error) {
	m := new(AddPhotoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
type EmployeeServiceServer interface {
	GetByBadgeNumber(context.Context, *GetByBadgeNumberRequest) (*EmployeeResponse, error)
	GetAll(*GetAllRequest, EmployeeService_GetAllServer) error
	Save(context.Context, *EmployeeRequest) (*EmployeeResponse, error)
	SaveAll(EmployeeService_SaveAllServer) error
	AddPhoto(EmployeeService_AddPhotoServer) error
}

func RegisterEmployeeServiceServer(s *grpc.Server, srv EmployeeServiceServer) {
	s.RegisterService(&_EmployeeService_serviceDesc, srv)
}

func _EmployeeService_GetByBadgeNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByBadgeNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetByBadgeNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/GetByBadgeNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetByBadgeNumber(ctx, req.(*GetByBadgeNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeeServiceServer).GetAll(m, &employeeServiceGetAllServer{stream})
}

type EmployeeService_GetAllServer interface {
	Send(*EmployeeResponse) error
	grpc.ServerStream
}

type employeeServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *employeeServiceGetAllServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EmployeeService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Save(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_SaveAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).SaveAll(&employeeServiceSaveAllServer{stream})
}

type EmployeeService_SaveAllServer interface {
	Send(*EmployeeResponse) error
	Recv() (*EmployeeRequest, error)
	grpc.ServerStream
}

type employeeServiceSaveAllServer struct {
	grpc.ServerStream
}

func (x *employeeServiceSaveAllServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceSaveAllServer) Recv() (*EmployeeRequest, error) {
	m := new(EmployeeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EmployeeService_AddPhoto_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).AddPhoto(&employeeServiceAddPhotoServer{stream})
}

type EmployeeService_AddPhotoServer interface {
	Send(*AddPhotoResponse) error
	Recv() (*AddPhotoRequest, error)
	grpc.ServerStream
}

type employeeServiceAddPhotoServer struct {
	grpc.ServerStream
}

func (x *employeeServiceAddPhotoServer) Send(m *AddPhotoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceAddPhotoServer) Recv() (*AddPhotoRequest, error) {
	m := new(AddPhotoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EmployeeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByBadgeNumber",
			Handler:    _EmployeeService_GetByBadgeNumber_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _EmployeeService_Save_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _EmployeeService_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SaveAll",
			Handler:       _EmployeeService_SaveAll_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AddPhoto",
			Handler:       _EmployeeService_AddPhoto_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messages.proto",
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_messages_7bc54bf07212d4bf) }

var fileDescriptor_messages_7bc54bf07212d4bf = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0xef, 0xd2, 0x30,
	0x18, 0xc7, 0xd3, 0xb9, 0x1f, 0x8e, 0x07, 0x7f, 0x6c, 0xd4, 0x83, 0x0b, 0xf1, 0xb0, 0x2c, 0xc1,
	0x2c, 0x31, 0x56, 0x02, 0x17, 0x8d, 0x07, 0x03, 0x89, 0xe1, 0x86, 0x66, 0xdc, 0xbc, 0x75, 0xeb,
	0x23, 0x2e, 0x6e, 0x74, 0xae, 0x85, 0x84, 0x57, 0xe2, 0xfb, 0xf2, 0x15, 0x19, 0xea, 0xc6, 0x64,
	0x60, 0xc2, 0xad, 0xfb, 0xfe, 0xe9, 0xb3, 0xe7, 0x93, 0x14, 0x86, 0x05, 0x2a, 0xc5, 0xb7, 0xa8,
	0x58, 0x59, 0x49, 0x2d, 0xc3, 0xdf, 0x04, 0x9c, 0x4f, 0x45, 0x99, 0xcb, 0x23, 0x22, 0x1d, 0x82,
	0x95, 0x09, 0x9f, 0x04, 0x24, 0x7a, 0x88, 0xad, 0x4c, 0xd0, 0x00, 0x06, 0x09, 0x17, 0x5b, 0x5c,
	0xef, 0x8b, 0x04, 0x2b, 0xdf, 0x32, 0xc6, 0xbf, 0x12, 0x7d, 0x09, 0xfd, 0x6f, 0x59, 0xa5, 0xf4,
	0x9a, 0x17, 0xe8, 0x3f, 0x09, 0x48, 0xd4, 0x8f, 0x5b, 0x81, 0x8e, 0xc1, 0xc9, 0x79, 0x6d, 0xda,
	0xc6, 0x3c, 0x7f, 0xd3, 0x29, 0x3c, 0x3f, 0xf0, 0x94, 0xeb, 0x4c, 0xee, 0x16, 0x69, 0x5a, 0xed,
	0x79, 0x1e, 0x73, 0x8d, 0xfe, 0x43, 0x40, 0x22, 0x2b, 0xbe, 0x65, 0xd1, 0x08, 0xdc, 0x0b, 0x19,
	0x85, 0xdf, 0x33, 0xe9, 0xae, 0x1c, 0xba, 0xf0, 0xb8, 0x42, 0xbd, 0xc8, 0xf3, 0x18, 0x7f, 0xee,
	0x51, 0xe9, 0xf0, 0x03, 0xbc, 0x58, 0xa1, 0x5e, 0x1e, 0x97, 0xed, 0xaf, 0xd7, 0x56, 0x77, 0x47,
	0x72, 0xb5, 0x63, 0xf8, 0x0e, 0xdc, 0x86, 0x50, 0x53, 0x9a, 0x80, 0x83, 0xb5, 0x64, 0x1a, 0x83,
	0x59, 0x9f, 0x9d, 0x33, 0x67, 0x2b, 0x7c, 0x0f, 0x5e, 0xdb, 0x54, 0xa5, 0xdc, 0x29, 0xbc, 0xb7,
	0x3a, 0x01, 0x77, 0x21, 0xc4, 0x97, 0xef, 0x52, 0xcb, 0x66, 0x28, 0x05, 0x5b, 0x70, 0xcd, 0x4d,
	0xeb, 0x59, 0x6c, 0xce, 0xe1, 0x2b, 0xf0, 0xda, 0x58, 0x3d, 0x81, 0x82, 0x9d, 0xa9, 0xcf, 0x3f,
	0x4c, 0xce, 0x89, 0xcd, 0x79, 0xf6, 0xcb, 0x6a, 0x97, 0xd8, 0x60, 0x75, 0xc8, 0x52, 0xa4, 0x1f,
	0xc1, 0xeb, 0x42, 0xa1, 0x3e, 0xfb, 0x0f, 0xa7, 0xf1, 0x88, 0x5d, 0xad, 0xf2, 0x06, 0x7a, 0x7f,
	0x31, 0xd3, 0x21, 0xbb, 0xe0, 0x7d, 0x23, 0x3c, 0x25, 0xf4, 0x35, 0xd8, 0x1b, 0x7e, 0x40, 0xea,
	0xb1, 0x0e, 0xce, 0x5b, 0x77, 0xcf, 0xe0, 0xe9, 0x29, 0x7c, 0xba, 0xfc, 0x9e, 0x7c, 0x44, 0xa6,
	0x84, 0xce, 0xc1, 0x69, 0x60, 0x50, 0x8f, 0x75, 0xf0, 0x8d, 0x47, 0xac, 0x4b, 0xea, 0x54, 0x5a,
	0x3e, 0x7e, 0x1d, 0x6c, 0xab, 0x32, 0x15, 0x58, 0xc8, 0xb7, 0x65, 0x92, 0xf4, 0xcc, 0xb3, 0x98,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xde, 0xc5, 0x25, 0x28, 0x03, 0x00, 0x00,
}
