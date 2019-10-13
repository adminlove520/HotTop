// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

package rpc

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

type Ping struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type Pong struct {
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

type Job struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{2}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Result struct {
	Map                  map[string]*Result_HotList `protobuf:"bytes,1,rep,name=Map,proto3" json:"Map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	T                    string                     `protobuf:"bytes,2,opt,name=T,proto3" json:"T,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{3}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetMap() map[string]*Result_HotList {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *Result) GetT() string {
	if m != nil {
		return m.T
	}
	return ""
}

type Result_HotList struct {
	Item                 []*Result_HotList_Item `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Result_HotList) Reset()         { *m = Result_HotList{} }
func (m *Result_HotList) String() string { return proto.CompactTextString(m) }
func (*Result_HotList) ProtoMessage()    {}
func (*Result_HotList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{3, 0}
}

func (m *Result_HotList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result_HotList.Unmarshal(m, b)
}
func (m *Result_HotList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result_HotList.Marshal(b, m, deterministic)
}
func (m *Result_HotList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result_HotList.Merge(m, src)
}
func (m *Result_HotList) XXX_Size() int {
	return xxx_messageInfo_Result_HotList.Size(m)
}
func (m *Result_HotList) XXX_DiscardUnknown() {
	xxx_messageInfo_Result_HotList.DiscardUnknown(m)
}

var xxx_messageInfo_Result_HotList proto.InternalMessageInfo

func (m *Result_HotList) GetItem() []*Result_HotList_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type Result_HotList_Item struct {
	Title                string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Rank                 float32  `protobuf:"fixed32,2,opt,name=Rank,proto3" json:"Rank,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=Url,proto3" json:"Url,omitempty"`
	Key                  string   `protobuf:"bytes,4,opt,name=Key,proto3" json:"Key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result_HotList_Item) Reset()         { *m = Result_HotList_Item{} }
func (m *Result_HotList_Item) String() string { return proto.CompactTextString(m) }
func (*Result_HotList_Item) ProtoMessage()    {}
func (*Result_HotList_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{3, 0, 0}
}

func (m *Result_HotList_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result_HotList_Item.Unmarshal(m, b)
}
func (m *Result_HotList_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result_HotList_Item.Marshal(b, m, deterministic)
}
func (m *Result_HotList_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result_HotList_Item.Merge(m, src)
}
func (m *Result_HotList_Item) XXX_Size() int {
	return xxx_messageInfo_Result_HotList_Item.Size(m)
}
func (m *Result_HotList_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Result_HotList_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Result_HotList_Item proto.InternalMessageInfo

func (m *Result_HotList_Item) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Result_HotList_Item) GetRank() float32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *Result_HotList_Item) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Result_HotList_Item) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "rpc.Ping")
	proto.RegisterType((*Pong)(nil), "rpc.Pong")
	proto.RegisterType((*Job)(nil), "rpc.Job")
	proto.RegisterType((*Result)(nil), "rpc.Result")
	proto.RegisterMapType((map[string]*Result_HotList)(nil), "rpc.Result.MapEntry")
	proto.RegisterType((*Result_HotList)(nil), "rpc.Result.HotList")
	proto.RegisterType((*Result_HotList_Item)(nil), "rpc.Result.HotList.Item")
}

func init() { proto.RegisterFile("job.proto", fileDescriptor_f32c477d91a04ead) }

var fileDescriptor_f32c477d91a04ead = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x9b, 0xf9, 0xd3, 0xaf, 0xbd, 0xfd, 0x16, 0x72, 0xed, 0x62, 0xcc, 0xa6, 0x65, 0x16,
	0x52, 0x41, 0x66, 0x51, 0x37, 0xe2, 0x4e, 0x8a, 0x60, 0x5b, 0x0b, 0x12, 0xc6, 0x07, 0xc8, 0x94,
	0x30, 0x8e, 0x9d, 0x26, 0x21, 0x4d, 0x95, 0x79, 0x08, 0x9f, 0xc0, 0x97, 0x95, 0x64, 0xa6, 0xd0,
	0x85, 0xbb, 0x93, 0xfb, 0x3b, 0xf7, 0xe4, 0xc0, 0x85, 0xe1, 0x87, 0x2a, 0x32, 0x6d, 0x94, 0x55,
	0x18, 0x1a, 0xbd, 0x4d, 0x29, 0x44, 0xaf, 0x95, 0x2c, 0x11, 0x21, 0xd2, 0x95, 0x2c, 0x13, 0x32,
	0x25, 0xb3, 0x21, 0xf3, 0xda, 0x33, 0xd5, 0x31, 0x75, 0xc6, 0x94, 0x2c, 0xd3, 0x2b, 0x08, 0x57,
	0xaa, 0x70, 0x48, 0xf2, 0xbd, 0x38, 0x21, 0xa7, 0xd3, 0x9f, 0x00, 0xfa, 0x4c, 0x1c, 0x8e, 0xb5,
	0xc5, 0x6b, 0x08, 0x37, 0x5c, 0x27, 0x64, 0x1a, 0xce, 0x46, 0xf3, 0x71, 0x66, 0xf4, 0x36, 0x6b,
	0x49, 0xb6, 0xe1, 0xfa, 0x49, 0x5a, 0xd3, 0x30, 0x67, 0xc0, 0xff, 0x40, 0xf2, 0x24, 0xf0, 0x19,
	0x24, 0xa7, 0xdf, 0x04, 0xfe, 0x3d, 0x2b, 0xfb, 0x52, 0x1d, 0x2c, 0xde, 0x42, 0x54, 0x59, 0xb1,
	0xef, 0x22, 0x92, 0xf3, 0x88, 0xce, 0x92, 0x2d, 0xad, 0xd8, 0x33, 0xef, 0xa2, 0x39, 0x44, 0xee,
	0x85, 0x63, 0x88, 0xf3, 0xca, 0xd6, 0xa7, 0x5e, 0xed, 0xc3, 0x95, 0x65, 0x5c, 0xee, 0xfc, 0x47,
	0x01, 0xf3, 0x1a, 0x2f, 0x20, 0x7c, 0x33, 0x75, 0x12, 0x7a, 0x9f, 0x93, 0x6e, 0xb2, 0x16, 0x4d,
	0x12, 0xb5, 0x93, 0xb5, 0x68, 0xe8, 0x1a, 0x06, 0xa7, 0xba, 0x8e, 0xee, 0x44, 0xd3, 0xe5, 0x3a,
	0x89, 0x37, 0x10, 0x7f, 0xf2, 0xfa, 0x28, 0x7c, 0xec, 0x68, 0x7e, 0xf9, 0x47, 0x45, 0xd6, 0x3a,
	0x1e, 0x82, 0x7b, 0x32, 0x5f, 0x42, 0xfc, 0x58, 0x0a, 0x69, 0x71, 0x02, 0xd1, 0xc2, 0xf0, 0x2f,
	0x1c, 0xf8, 0x85, 0x95, 0x2a, 0xe8, 0xe8, 0x6c, 0x35, 0xed, 0xe1, 0x04, 0xe2, 0xc5, 0xbb, 0xd8,
	0xee, 0x70, 0xe8, 0xe7, 0xee, 0x4c, 0xb4, 0x93, 0xee, 0x02, 0xbd, 0xa2, 0xef, 0xef, 0x78, 0xf7,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xeb, 0x08, 0xf7, 0xd4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	Craw(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Result, error)
	Check(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Craw(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/rpc.Agent/Craw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Check(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/rpc.Agent/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	Craw(context.Context, *Job) (*Result, error)
	Check(context.Context, *Ping) (*Pong, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) Craw(ctx context.Context, req *Job) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Craw not implemented")
}
func (*UnimplementedAgentServer) Check(ctx context.Context, req *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Craw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Craw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Agent/Craw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Craw(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Agent/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Check(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Craw",
			Handler:    _Agent_Craw_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Agent_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}
