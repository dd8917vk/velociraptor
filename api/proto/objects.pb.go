// Code generated by protoc-gen-go. DO NOT EDIT.
// source: objects.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ObjectReference struct {
	// Types that are valid to be assigned to Union:
	//	*ObjectReference_Client
	//	*ObjectReference_Hunt
	//	*ObjectReference_Flow
	//	*ObjectReference_VfsFile
	//	*ObjectReference_ApprovalRequest
	Union                isObjectReference_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ObjectReference) Reset()         { *m = ObjectReference{} }
func (m *ObjectReference) String() string { return proto.CompactTextString(m) }
func (*ObjectReference) ProtoMessage()    {}
func (*ObjectReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_objects_992ff9aeba418b8d, []int{0}
}
func (m *ObjectReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectReference.Unmarshal(m, b)
}
func (m *ObjectReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectReference.Marshal(b, m, deterministic)
}
func (dst *ObjectReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectReference.Merge(dst, src)
}
func (m *ObjectReference) XXX_Size() int {
	return xxx_messageInfo_ObjectReference.Size(m)
}
func (m *ObjectReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectReference.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectReference proto.InternalMessageInfo

type isObjectReference_Union interface {
	isObjectReference_Union()
}

type ObjectReference_Client struct {
	Client *ClientReference `protobuf:"bytes,2,opt,name=client,proto3,oneof"`
}
type ObjectReference_Hunt struct {
	Hunt *HuntReference `protobuf:"bytes,3,opt,name=hunt,proto3,oneof"`
}
type ObjectReference_Flow struct {
	Flow *FlowReference `protobuf:"bytes,4,opt,name=flow,proto3,oneof"`
}
type ObjectReference_VfsFile struct {
	VfsFile *VfsFileReference `protobuf:"bytes,6,opt,name=vfs_file,json=vfsFile,proto3,oneof"`
}
type ObjectReference_ApprovalRequest struct {
	ApprovalRequest *ApprovalRequestReference `protobuf:"bytes,7,opt,name=approval_request,json=approvalRequest,proto3,oneof"`
}

func (*ObjectReference_Client) isObjectReference_Union()          {}
func (*ObjectReference_Hunt) isObjectReference_Union()            {}
func (*ObjectReference_Flow) isObjectReference_Union()            {}
func (*ObjectReference_VfsFile) isObjectReference_Union()         {}
func (*ObjectReference_ApprovalRequest) isObjectReference_Union() {}

func (m *ObjectReference) GetUnion() isObjectReference_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *ObjectReference) GetClient() *ClientReference {
	if x, ok := m.GetUnion().(*ObjectReference_Client); ok {
		return x.Client
	}
	return nil
}

func (m *ObjectReference) GetHunt() *HuntReference {
	if x, ok := m.GetUnion().(*ObjectReference_Hunt); ok {
		return x.Hunt
	}
	return nil
}

func (m *ObjectReference) GetFlow() *FlowReference {
	if x, ok := m.GetUnion().(*ObjectReference_Flow); ok {
		return x.Flow
	}
	return nil
}

func (m *ObjectReference) GetVfsFile() *VfsFileReference {
	if x, ok := m.GetUnion().(*ObjectReference_VfsFile); ok {
		return x.VfsFile
	}
	return nil
}

func (m *ObjectReference) GetApprovalRequest() *ApprovalRequestReference {
	if x, ok := m.GetUnion().(*ObjectReference_ApprovalRequest); ok {
		return x.ApprovalRequest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ObjectReference) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ObjectReference_OneofMarshaler, _ObjectReference_OneofUnmarshaler, _ObjectReference_OneofSizer, []interface{}{
		(*ObjectReference_Client)(nil),
		(*ObjectReference_Hunt)(nil),
		(*ObjectReference_Flow)(nil),
		(*ObjectReference_VfsFile)(nil),
		(*ObjectReference_ApprovalRequest)(nil),
	}
}

func _ObjectReference_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ObjectReference)
	// union
	switch x := m.Union.(type) {
	case *ObjectReference_Client:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Client); err != nil {
			return err
		}
	case *ObjectReference_Hunt:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hunt); err != nil {
			return err
		}
	case *ObjectReference_Flow:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Flow); err != nil {
			return err
		}
	case *ObjectReference_VfsFile:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VfsFile); err != nil {
			return err
		}
	case *ObjectReference_ApprovalRequest:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApprovalRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ObjectReference.Union has unexpected type %T", x)
	}
	return nil
}

func _ObjectReference_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ObjectReference)
	switch tag {
	case 2: // union.client
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientReference)
		err := b.DecodeMessage(msg)
		m.Union = &ObjectReference_Client{msg}
		return true, err
	case 3: // union.hunt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HuntReference)
		err := b.DecodeMessage(msg)
		m.Union = &ObjectReference_Hunt{msg}
		return true, err
	case 4: // union.flow
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FlowReference)
		err := b.DecodeMessage(msg)
		m.Union = &ObjectReference_Flow{msg}
		return true, err
	case 6: // union.vfs_file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VfsFileReference)
		err := b.DecodeMessage(msg)
		m.Union = &ObjectReference_VfsFile{msg}
		return true, err
	case 7: // union.approval_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApprovalRequestReference)
		err := b.DecodeMessage(msg)
		m.Union = &ObjectReference_ApprovalRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ObjectReference_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ObjectReference)
	// union
	switch x := m.Union.(type) {
	case *ObjectReference_Client:
		s := proto.Size(x.Client)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ObjectReference_Hunt:
		s := proto.Size(x.Hunt)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ObjectReference_Flow:
		s := proto.Size(x.Flow)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ObjectReference_VfsFile:
		s := proto.Size(x.VfsFile)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ObjectReference_ApprovalRequest:
		s := proto.Size(x.ApprovalRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ClientReference struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientReference) Reset()         { *m = ClientReference{} }
func (m *ClientReference) String() string { return proto.CompactTextString(m) }
func (*ClientReference) ProtoMessage()    {}
func (*ClientReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_objects_992ff9aeba418b8d, []int{1}
}
func (m *ClientReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientReference.Unmarshal(m, b)
}
func (m *ClientReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientReference.Marshal(b, m, deterministic)
}
func (dst *ClientReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientReference.Merge(dst, src)
}
func (m *ClientReference) XXX_Size() int {
	return xxx_messageInfo_ClientReference.Size(m)
}
func (m *ClientReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientReference.DiscardUnknown(m)
}

var xxx_messageInfo_ClientReference proto.InternalMessageInfo

func (m *ClientReference) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type HuntReference struct {
	HuntId               string   `protobuf:"bytes,1,opt,name=hunt_id,json=huntId,proto3" json:"hunt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HuntReference) Reset()         { *m = HuntReference{} }
func (m *HuntReference) String() string { return proto.CompactTextString(m) }
func (*HuntReference) ProtoMessage()    {}
func (*HuntReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_objects_992ff9aeba418b8d, []int{2}
}
func (m *HuntReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HuntReference.Unmarshal(m, b)
}
func (m *HuntReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HuntReference.Marshal(b, m, deterministic)
}
func (dst *HuntReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HuntReference.Merge(dst, src)
}
func (m *HuntReference) XXX_Size() int {
	return xxx_messageInfo_HuntReference.Size(m)
}
func (m *HuntReference) XXX_DiscardUnknown() {
	xxx_messageInfo_HuntReference.DiscardUnknown(m)
}

var xxx_messageInfo_HuntReference proto.InternalMessageInfo

func (m *HuntReference) GetHuntId() string {
	if m != nil {
		return m.HuntId
	}
	return ""
}

type FlowReference struct {
	FlowId               string   `protobuf:"bytes,1,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowReference) Reset()         { *m = FlowReference{} }
func (m *FlowReference) String() string { return proto.CompactTextString(m) }
func (*FlowReference) ProtoMessage()    {}
func (*FlowReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_objects_992ff9aeba418b8d, []int{3}
}
func (m *FlowReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowReference.Unmarshal(m, b)
}
func (m *FlowReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowReference.Marshal(b, m, deterministic)
}
func (dst *FlowReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowReference.Merge(dst, src)
}
func (m *FlowReference) XXX_Size() int {
	return xxx_messageInfo_FlowReference.Size(m)
}
func (m *FlowReference) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowReference.DiscardUnknown(m)
}

var xxx_messageInfo_FlowReference proto.InternalMessageInfo

func (m *FlowReference) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *FlowReference) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type VfsFileReference struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	VfsPathComponents    []string `protobuf:"bytes,2,rep,name=vfs_path_components,json=vfsPathComponents,proto3" json:"vfs_path_components,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VfsFileReference) Reset()         { *m = VfsFileReference{} }
func (m *VfsFileReference) String() string { return proto.CompactTextString(m) }
func (*VfsFileReference) ProtoMessage()    {}
func (*VfsFileReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_objects_992ff9aeba418b8d, []int{4}
}
func (m *VfsFileReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VfsFileReference.Unmarshal(m, b)
}
func (m *VfsFileReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VfsFileReference.Marshal(b, m, deterministic)
}
func (dst *VfsFileReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VfsFileReference.Merge(dst, src)
}
func (m *VfsFileReference) XXX_Size() int {
	return xxx_messageInfo_VfsFileReference.Size(m)
}
func (m *VfsFileReference) XXX_DiscardUnknown() {
	xxx_messageInfo_VfsFileReference.DiscardUnknown(m)
}

var xxx_messageInfo_VfsFileReference proto.InternalMessageInfo

func (m *VfsFileReference) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *VfsFileReference) GetVfsPathComponents() []string {
	if m != nil {
		return m.VfsPathComponents
	}
	return nil
}

type ApprovalRequestReference struct {
	//    ApprovalRequest.ApprovalType approval_type = 1;
	ApprovalId           string   `protobuf:"bytes,2,opt,name=approval_id,json=approvalId,proto3" json:"approval_id,omitempty"`
	SubjectId            string   `protobuf:"bytes,3,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApprovalRequestReference) Reset()         { *m = ApprovalRequestReference{} }
func (m *ApprovalRequestReference) String() string { return proto.CompactTextString(m) }
func (*ApprovalRequestReference) ProtoMessage()    {}
func (*ApprovalRequestReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_objects_992ff9aeba418b8d, []int{5}
}
func (m *ApprovalRequestReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApprovalRequestReference.Unmarshal(m, b)
}
func (m *ApprovalRequestReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApprovalRequestReference.Marshal(b, m, deterministic)
}
func (dst *ApprovalRequestReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovalRequestReference.Merge(dst, src)
}
func (m *ApprovalRequestReference) XXX_Size() int {
	return xxx_messageInfo_ApprovalRequestReference.Size(m)
}
func (m *ApprovalRequestReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovalRequestReference.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovalRequestReference proto.InternalMessageInfo

func (m *ApprovalRequestReference) GetApprovalId() string {
	if m != nil {
		return m.ApprovalId
	}
	return ""
}

func (m *ApprovalRequestReference) GetSubjectId() string {
	if m != nil {
		return m.SubjectId
	}
	return ""
}

func init() {
	proto.RegisterType((*ObjectReference)(nil), "proto.ObjectReference")
	proto.RegisterType((*ClientReference)(nil), "proto.ClientReference")
	proto.RegisterType((*HuntReference)(nil), "proto.HuntReference")
	proto.RegisterType((*FlowReference)(nil), "proto.FlowReference")
	proto.RegisterType((*VfsFileReference)(nil), "proto.VfsFileReference")
	proto.RegisterType((*ApprovalRequestReference)(nil), "proto.ApprovalRequestReference")
}

func init() { proto.RegisterFile("objects.proto", fileDescriptor_objects_992ff9aeba418b8d) }

var fileDescriptor_objects_992ff9aeba418b8d = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6e, 0xd4, 0x30,
	0x18, 0x85, 0x99, 0x99, 0x36, 0x69, 0xfe, 0x6a, 0x34, 0xc5, 0x20, 0x1a, 0x81, 0x50, 0xad, 0x2c,
	0x50, 0x60, 0x91, 0x41, 0x80, 0x58, 0x20, 0x36, 0xb4, 0x52, 0x69, 0x24, 0x04, 0x28, 0x02, 0x24,
	0xd8, 0x44, 0x69, 0xec, 0x34, 0x46, 0x19, 0x3b, 0xc4, 0x4e, 0xc2, 0x91, 0xd8, 0x73, 0x05, 0x4e,
	0x02, 0xd7, 0x60, 0x81, 0x6c, 0x67, 0x86, 0x49, 0x05, 0x5d, 0x59, 0x79, 0xef, 0xfb, 0x9f, 0xe2,
	0xf7, 0x1b, 0xe6, 0xe2, 0xfc, 0x33, 0xcd, 0x95, 0x8c, 0xea, 0x46, 0x28, 0x81, 0x76, 0xcd, 0x71,
	0xfb, 0x59, 0xdf, 0xf7, 0x51, 0x47, 0x2b, 0x91, 0x33, 0x42, 0xbf, 0x46, 0xb9, 0x58, 0x2d, 0x2f,
	0x44, 0x95, 0xf1, 0x8b, 0xa5, 0x15, 0x9b, 0xac, 0x56, 0xa2, 0x59, 0x1a, 0x78, 0x29, 0xe9, 0x2a,
	0xe3, 0x8a, 0xe5, 0x36, 0x22, 0xf8, 0x36, 0x85, 0xc5, 0x1b, 0x13, 0x9a, 0xd0, 0x82, 0x36, 0x94,
	0xe7, 0x14, 0x3d, 0x04, 0x27, 0xaf, 0x18, 0xe5, 0xca, 0x9f, 0xe2, 0x49, 0xb8, 0xff, 0xe8, 0x96,
	0x65, 0xa3, 0x13, 0x23, 0x6e, 0xb8, 0xb3, 0x6b, 0xc9, 0xc0, 0xa1, 0x07, 0xb0, 0x53, 0xb6, 0x5c,
	0xf9, 0x33, 0xc3, 0xdf, 0x1c, 0xf8, 0xb3, 0x76, 0x4c, 0x1b, 0x46, 0xb3, 0x45, 0x25, 0x7a, 0x7f,
	0x67, 0xc4, 0x9e, 0x56, 0xa2, 0x1f, 0xb1, 0x9a, 0x41, 0x4f, 0x60, 0xaf, 0x2b, 0x64, 0x5a, 0xb0,
	0x8a, 0xfa, 0x8e, 0xe1, 0x0f, 0x07, 0xfe, 0x43, 0x21, 0x4f, 0x59, 0x45, 0xb7, 0x47, 0xdc, 0xce,
	0x6a, 0xe8, 0x15, 0x1c, 0x64, 0x75, 0xdd, 0x88, 0x2e, 0xab, 0xd2, 0x86, 0x7e, 0x69, 0xa9, 0x54,
	0xbe, 0x6b, 0xa6, 0x8f, 0x86, 0xe9, 0x17, 0x83, 0x9d, 0x58, 0x77, 0x3b, 0x65, 0x91, 0x8d, 0xbd,
	0x63, 0x17, 0x76, 0x5b, 0xce, 0x04, 0x0f, 0x22, 0x58, 0x5c, 0x6a, 0x00, 0xdd, 0x01, 0xcf, 0x36,
	0x90, 0x32, 0xe2, 0x4f, 0xf0, 0x24, 0xf4, 0x92, 0x3d, 0x2b, 0xc4, 0x24, 0x08, 0x61, 0x3e, 0x6a,
	0x00, 0x1d, 0x82, 0xab, 0x1b, 0xf8, 0xcb, 0x3a, 0xfa, 0x33, 0x26, 0xc1, 0xf7, 0x09, 0xcc, 0x47,
	0x05, 0xa0, 0x97, 0xe0, 0xea, 0x02, 0x36, 0xe8, 0x71, 0xf4, 0xf3, 0xf7, 0xaf, 0x1f, 0x93, 0x10,
	0xdd, 0x7b, 0x57, 0x52, 0x2c, 0xa9, 0x94, 0x4c, 0x70, 0xcc, 0x08, 0x16, 0x05, 0x56, 0x25, 0xc5,
	0xcd, 0x7a, 0x92, 0x60, 0x3d, 0x17, 0x25, 0x8e, 0x3e, 0x62, 0x82, 0x3e, 0x6e, 0xff, 0xe1, 0xd4,
	0x44, 0x3d, 0x37, 0x51, 0x4f, 0xc1, 0xb3, 0xb7, 0x79, 0x9f, 0xbc, 0x46, 0xf7, 0x75, 0xaa, 0xe5,
	0x74, 0xe8, 0x3f, 0x12, 0x71, 0x93, 0x71, 0x2c, 0x78, 0xb4, 0x75, 0xbf, 0x14, 0x0e, 0x2e, 0x6f,
	0xe1, 0xca, 0x42, 0x50, 0x04, 0x37, 0xf4, 0x36, 0xeb, 0x4c, 0x95, 0x69, 0x2e, 0x56, 0xb5, 0xe0,
	0x94, 0x2b, 0xe9, 0x4f, 0xf1, 0x2c, 0xf4, 0x92, 0xeb, 0x5d, 0x21, 0xdf, 0x66, 0xaa, 0x3c, 0xd9,
	0x18, 0xc1, 0x27, 0xf0, 0xff, 0xb7, 0x28, 0x74, 0x04, 0xfb, 0x9b, 0x1d, 0xaf, 0x6f, 0x96, 0xc0,
	0x5a, 0x8a, 0x09, 0xba, 0x0b, 0x20, 0x5b, 0xf3, 0xb0, 0xb5, 0x3f, 0x33, 0xbe, 0x37, 0x28, 0x31,
	0x39, 0x77, 0xcc, 0x43, 0x78, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x29, 0x47, 0x64, 0x52,
	0x03, 0x00, 0x00,
}
