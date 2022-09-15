// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema/chat.proto

package schema

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{2}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{3}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

// for client
type StreamRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{4}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// for server
type StreamResponse struct {
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*StreamResponse_ClientMessage
	//	*StreamResponse_ServerShutdown
	//	*StreamResponse_ClientLogin
	//	*StreamResponse_ClientLogout
	Event                isStreamResponse_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{5}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type isStreamResponse_Event interface {
	isStreamResponse_Event()
}

type StreamResponse_ClientMessage struct {
	ClientMessage *StreamResponse_Message `protobuf:"bytes,2,opt,name=client_message,json=clientMessage,proto3,oneof"`
}

type StreamResponse_ServerShutdown struct {
	ServerShutdown *StreamResponse_Shutdown `protobuf:"bytes,3,opt,name=server_shutdown,json=serverShutdown,proto3,oneof"`
}

type StreamResponse_ClientLogin struct {
	ClientLogin *StreamResponse_Login `protobuf:"bytes,4,opt,name=client_login,json=clientLogin,proto3,oneof"`
}

type StreamResponse_ClientLogout struct {
	ClientLogout *StreamResponse_Logout `protobuf:"bytes,5,opt,name=client_logout,json=clientLogout,proto3,oneof"`
}

func (*StreamResponse_ClientMessage) isStreamResponse_Event() {}

func (*StreamResponse_ServerShutdown) isStreamResponse_Event() {}

func (*StreamResponse_ClientLogin) isStreamResponse_Event() {}

func (*StreamResponse_ClientLogout) isStreamResponse_Event() {}

func (m *StreamResponse) GetEvent() isStreamResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StreamResponse) GetClientMessage() *StreamResponse_Message {
	if x, ok := m.GetEvent().(*StreamResponse_ClientMessage); ok {
		return x.ClientMessage
	}
	return nil
}

func (m *StreamResponse) GetServerShutdown() *StreamResponse_Shutdown {
	if x, ok := m.GetEvent().(*StreamResponse_ServerShutdown); ok {
		return x.ServerShutdown
	}
	return nil
}

func (m *StreamResponse) GetClientLogin() *StreamResponse_Login {
	if x, ok := m.GetEvent().(*StreamResponse_ClientLogin); ok {
		return x.ClientLogin
	}
	return nil
}

func (m *StreamResponse) GetClientLogout() *StreamResponse_Logout {
	if x, ok := m.GetEvent().(*StreamResponse_ClientLogout); ok {
		return x.ClientLogout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamResponse_ClientMessage)(nil),
		(*StreamResponse_ServerShutdown)(nil),
		(*StreamResponse_ClientLogin)(nil),
		(*StreamResponse_ClientLogout)(nil),
	}
}

type StreamResponse_Login struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Login) Reset()         { *m = StreamResponse_Login{} }
func (m *StreamResponse_Login) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Login) ProtoMessage()    {}
func (*StreamResponse_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{5, 0}
}

func (m *StreamResponse_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Login.Unmarshal(m, b)
}
func (m *StreamResponse_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Login.Marshal(b, m, deterministic)
}
func (m *StreamResponse_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Login.Merge(m, src)
}
func (m *StreamResponse_Login) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Login.Size(m)
}
func (m *StreamResponse_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Login.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Login proto.InternalMessageInfo

func (m *StreamResponse_Login) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Logout struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Logout) Reset()         { *m = StreamResponse_Logout{} }
func (m *StreamResponse_Logout) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Logout) ProtoMessage()    {}
func (*StreamResponse_Logout) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{5, 1}
}

func (m *StreamResponse_Logout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Logout.Unmarshal(m, b)
}
func (m *StreamResponse_Logout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Logout.Marshal(b, m, deterministic)
}
func (m *StreamResponse_Logout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Logout.Merge(m, src)
}
func (m *StreamResponse_Logout) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Logout.Size(m)
}
func (m *StreamResponse_Logout) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Logout.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Logout proto.InternalMessageInfo

func (m *StreamResponse_Logout) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Message struct {
	Nama                 string   `protobuf:"bytes,1,opt,name=nama,proto3" json:"nama,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Message) Reset()         { *m = StreamResponse_Message{} }
func (m *StreamResponse_Message) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Message) ProtoMessage()    {}
func (*StreamResponse_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{5, 2}
}

func (m *StreamResponse_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Message.Unmarshal(m, b)
}
func (m *StreamResponse_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Message.Marshal(b, m, deterministic)
}
func (m *StreamResponse_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Message.Merge(m, src)
}
func (m *StreamResponse_Message) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Message.Size(m)
}
func (m *StreamResponse_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Message.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Message proto.InternalMessageInfo

func (m *StreamResponse_Message) GetNama() string {
	if m != nil {
		return m.Nama
	}
	return ""
}

type StreamResponse_Shutdown struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Shutdown) Reset()         { *m = StreamResponse_Shutdown{} }
func (m *StreamResponse_Shutdown) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Shutdown) ProtoMessage()    {}
func (*StreamResponse_Shutdown) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fa9e2790b774bb, []int{5, 3}
}

func (m *StreamResponse_Shutdown) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Shutdown.Unmarshal(m, b)
}
func (m *StreamResponse_Shutdown) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Shutdown.Marshal(b, m, deterministic)
}
func (m *StreamResponse_Shutdown) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Shutdown.Merge(m, src)
}
func (m *StreamResponse_Shutdown) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Shutdown.Size(m)
}
func (m *StreamResponse_Shutdown) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Shutdown.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Shutdown proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LoginRequest)(nil), "chat.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "chat.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "chat.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "chat.LogoutResponse")
	proto.RegisterType((*StreamRequest)(nil), "chat.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "chat.StreamResponse")
	proto.RegisterType((*StreamResponse_Login)(nil), "chat.StreamResponse.Login")
	proto.RegisterType((*StreamResponse_Logout)(nil), "chat.StreamResponse.Logout")
	proto.RegisterType((*StreamResponse_Message)(nil), "chat.StreamResponse.Message")
	proto.RegisterType((*StreamResponse_Shutdown)(nil), "chat.StreamResponse.Shutdown")
}

func init() {
	proto.RegisterFile("schema/chat.proto", fileDescriptor_b8fa9e2790b774bb)
}

var fileDescriptor_b8fa9e2790b774bb = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xb5, 0x5b, 0xdb, 0x49, 0xa7, 0xb1, 0x81, 0xa5, 0x07, 0x6b, 0xdb, 0x0a, 0x64, 0x09, 0xa9,
	0xe2, 0xe0, 0xa0, 0x20, 0x24, 0x38, 0x20, 0xa4, 0x20, 0x24, 0x1f, 0xe8, 0xc5, 0xe5, 0xc4, 0xa5,
	0xda, 0x86, 0xc1, 0x8e, 0x88, 0xbd, 0xc1, 0xbb, 0x5b, 0xfe, 0x14, 0xff, 0x85, 0xbf, 0x84, 0xba,
	0x1f, 0x76, 0x2c, 0xb9, 0x37, 0xcf, 0xec, 0x7b, 0x6f, 0xdf, 0xbe, 0x19, 0xc3, 0x33, 0xb1, 0xa9,
	0xb1, 0x61, 0xcb, 0x4d, 0xcd, 0x64, 0xbe, 0xef, 0xb8, 0xe4, 0x24, 0x78, 0xf8, 0xa6, 0x2f, 0x2a,
	0xce, 0xab, 0x1d, 0x2e, 0x75, 0xef, 0x4e, 0xfd, 0x5c, 0xca, 0x6d, 0x83, 0x42, 0xb2, 0x66, 0x6f,
	0x60, 0xd9, 0x6b, 0x58, 0x7c, 0xe5, 0xd5, 0xb6, 0x2d, 0xf1, 0xb7, 0x42, 0x21, 0x09, 0x85, 0xb9,
	0x12, 0xd8, 0xb5, 0xac, 0xc1, 0xd4, 0x7f, 0xe9, 0x5f, 0x9d, 0x94, 0x7d, 0x9d, 0xbd, 0x82, 0xd8,
	0x62, 0xc5, 0x9e, 0xb7, 0x02, 0xc9, 0x19, 0x84, 0x92, 0xff, 0xc2, 0xd6, 0x22, 0x4d, 0x61, 0x61,
	0x5c, 0x49, 0xa7, 0x39, 0x0d, 0x7b, 0x0a, 0x89, 0x83, 0x19, 0xb9, 0xec, 0x23, 0xc4, 0x37, 0xb2,
	0x43, 0xd6, 0x38, 0x62, 0x0a, 0xb3, 0x06, 0x85, 0x60, 0x95, 0xf3, 0xe2, 0x4a, 0x42, 0x20, 0xd0,
	0x16, 0x8f, 0x74, 0x5b, 0x7f, 0x67, 0xff, 0x8e, 0x21, 0x71, 0x7c, 0x6b, 0xf0, 0x3d, 0x9c, 0xf4,
	0x0f, 0xd6, 0x12, 0xa7, 0x2b, 0x9a, 0x9b, 0x48, 0x72, 0x17, 0x49, 0xfe, 0xcd, 0x21, 0xca, 0x01,
	0x4c, 0xbe, 0x40, 0xb2, 0xd9, 0x6d, 0xb1, 0x95, 0xb7, 0xce, 0xc1, 0x91, 0xa6, 0x5f, 0xe4, 0x3a,
	0xe3, 0xf1, 0x3d, 0xf9, 0xb5, 0xc1, 0x14, 0x5e, 0x19, 0x1b, 0x96, 0x6d, 0x90, 0x02, 0x9e, 0x08,
	0xec, 0xee, 0xb1, 0xbb, 0x15, 0xb5, 0x92, 0x3f, 0xf8, 0x9f, 0x36, 0x3d, 0xd6, 0x3a, 0x97, 0x93,
	0x3a, 0x37, 0x16, 0x54, 0x78, 0x65, 0x62, 0x78, 0xae, 0x43, 0x3e, 0xc1, 0xc2, 0x1a, 0xda, 0x3d,
	0xcc, 0x20, 0x0d, 0xec, 0x6b, 0xa6, 0x64, 0xf4, 0x94, 0x0a, 0xaf, 0x3c, 0x35, 0x0c, 0x5d, 0x92,
	0x35, 0xc4, 0x83, 0x00, 0x57, 0x32, 0x0d, 0xb5, 0xc2, 0xf9, 0x63, 0x0a, 0x5c, 0xc9, 0xc2, 0x2b,
	0x17, 0xbd, 0x04, 0x57, 0x92, 0x9e, 0x43, 0x68, 0xc4, 0x5c, 0xfe, 0xfe, 0x90, 0x3f, 0xbd, 0x80,
	0xc8, 0xc0, 0x26, 0x4f, 0x2f, 0x61, 0x76, 0x3d, 0x1a, 0x1e, 0x3b, 0x38, 0x66, 0x14, 0x60, 0xee,
	0x9e, 0xba, 0x9e, 0x41, 0x88, 0xf7, 0xd8, 0xca, 0xd5, 0x5f, 0x1f, 0x82, 0xcf, 0x35, 0x93, 0x64,
	0xd5, 0xdf, 0x6b, 0xdc, 0x1e, 0xae, 0x2c, 0x7d, 0x3e, 0xea, 0xd9, 0x5d, 0xf2, 0xc8, 0xbb, 0xde,
	0xce, 0x00, 0x18, 0x96, 0x92, 0x9e, 0x8d, 0x9b, 0x3d, 0xed, 0x03, 0x44, 0x26, 0x0b, 0x47, 0x1b,
	0xad, 0xa4, 0xa3, 0x8d, 0xe3, 0xca, 0xbc, 0x2b, 0xff, 0x8d, 0xbf, 0x9e, 0x7f, 0x8f, 0xcc, 0x7f,
	0x78, 0x17, 0xe9, 0xe5, 0x7a, 0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xf8, 0x2c, 0xb5, 0x98,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/chat.Chat/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/chat.Chat/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/chat.Chat/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatStreamClient{stream}
	return x, nil
}

type Chat_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type chatStreamClient struct {
	grpc.ClientStream
}

func (x *chatStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Stream(Chat_StreamServer) error
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedChatServer) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedChatServer) Stream(srv Chat_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Stream(&chatStreamServer{stream})
}

type Chat_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type chatStreamServer struct {
	grpc.ServerStream
}

func (x *chatStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Chat_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Chat_Logout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Chat_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "schema/chat.proto",
}
