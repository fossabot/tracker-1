// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.proto

package schema

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Source struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}
func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Manages struct {
	Language             string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	System               string   `protobuf:"bytes,2,opt,name=system,proto3" json:"system,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manages) Reset()         { *m = Manages{} }
func (m *Manages) String() string { return proto.CompactTextString(m) }
func (*Manages) ProtoMessage()    {}
func (*Manages) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}
func (m *Manages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manages.Unmarshal(m, b)
}
func (m *Manages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manages.Marshal(b, m, deterministic)
}
func (m *Manages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manages.Merge(m, src)
}
func (m *Manages) XXX_Size() int {
	return xxx_messageInfo_Manages.Size(m)
}
func (m *Manages) XXX_DiscardUnknown() {
	xxx_messageInfo_Manages.DiscardUnknown(m)
}

var xxx_messageInfo_Manages proto.InternalMessageInfo

func (m *Manages) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Manages) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *Manages) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Module struct {
	Language             string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Organization         string   `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
	Module               string   `protobuf:"bytes,3,opt,name=module,proto3" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{2}
}
func (m *Module) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Module.Unmarshal(m, b)
}
func (m *Module) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Module.Marshal(b, m, deterministic)
}
func (m *Module) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Module.Merge(m, src)
}
func (m *Module) XXX_Size() int {
	return xxx_messageInfo_Module.Size(m)
}
func (m *Module) XXX_DiscardUnknown() {
	xxx_messageInfo_Module.DiscardUnknown(m)
}

var xxx_messageInfo_Module proto.InternalMessageInfo

func (m *Module) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Module) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *Module) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

type Depends struct {
	Language             string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	VersionConstraint    string   `protobuf:"bytes,2,opt,name=version_constraint,json=versionConstraint,proto3" json:"version_constraint,omitempty"`
	Scopes               []string `protobuf:"bytes,3,rep,name=scopes,proto3" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Depends) Reset()         { *m = Depends{} }
func (m *Depends) String() string { return proto.CompactTextString(m) }
func (*Depends) ProtoMessage()    {}
func (*Depends) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{3}
}
func (m *Depends) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Depends.Unmarshal(m, b)
}
func (m *Depends) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Depends.Marshal(b, m, deterministic)
}
func (m *Depends) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Depends.Merge(m, src)
}
func (m *Depends) XXX_Size() int {
	return xxx_messageInfo_Depends.Size(m)
}
func (m *Depends) XXX_DiscardUnknown() {
	xxx_messageInfo_Depends.DiscardUnknown(m)
}

var xxx_messageInfo_Depends proto.InternalMessageInfo

func (m *Depends) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Depends) GetVersionConstraint() string {
	if m != nil {
		return m.VersionConstraint
	}
	return ""
}

func (m *Depends) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func init() {
	proto.RegisterType((*Source)(nil), "cloud.deps.dts.api.v1alpha.schema.Source")
	proto.RegisterType((*Manages)(nil), "cloud.deps.dts.api.v1alpha.schema.Manages")
	proto.RegisterType((*Module)(nil), "cloud.deps.dts.api.v1alpha.schema.Module")
	proto.RegisterType((*Depends)(nil), "cloud.deps.dts.api.v1alpha.schema.Depends")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4b, 0x3c, 0x31,
	0x10, 0xc5, 0xb9, 0xef, 0x42, 0xf6, 0xbe, 0xc3, 0x15, 0x9a, 0x42, 0xc2, 0x55, 0x67, 0x2a, 0x1b,
	0x03, 0xe2, 0x7f, 0xa0, 0xb6, 0xd7, 0x68, 0x21, 0xd8, 0xe8, 0x98, 0x0c, 0x7b, 0x0b, 0xd9, 0x24,
	0xe4, 0xc7, 0x81, 0xfe, 0xf5, 0xb2, 0xd9, 0x78, 0x60, 0x73, 0x5d, 0x3e, 0x79, 0xc3, 0x7b, 0x33,
	0x0f, 0x36, 0x49, 0x1f, 0x68, 0x42, 0x15, 0xa2, 0xcf, 0x9e, 0x5f, 0x6b, 0xeb, 0x8b, 0x51, 0x86,
	0x42, 0x52, 0x26, 0x27, 0x85, 0x61, 0x54, 0xc7, 0x3b, 0xb4, 0xe1, 0x80, 0x6a, 0x19, 0x94, 0x5b,
	0x60, 0x2f, 0xbe, 0x44, 0x4d, 0xfc, 0x02, 0xba, 0x12, 0xad, 0x58, 0xed, 0x56, 0x37, 0xff, 0x9f,
	0xe7, 0xa7, 0x7c, 0x85, 0x7e, 0x8f, 0x0e, 0x07, 0x4a, 0x7c, 0x0b, 0x6b, 0x8b, 0x6e, 0x28, 0x38,
	0x50, 0x9b, 0x38, 0x31, 0xbf, 0x02, 0x96, 0xbe, 0x52, 0xa6, 0x49, 0xfc, 0xab, 0x4a, 0x23, 0x2e,
	0xa0, 0x3f, 0x52, 0x4c, 0xa3, 0x77, 0xa2, 0xab, 0xc2, 0x2f, 0xca, 0x0f, 0x60, 0x7b, 0x6f, 0x8a,
	0xa5, 0xb3, 0xbe, 0x12, 0x36, 0x3e, 0x0e, 0xe8, 0xc6, 0x6f, 0xcc, 0xb3, 0xc9, 0xe2, 0xfe, 0xe7,
	0x6f, 0xce, 0x9e, 0xaa, 0x53, 0x8b, 0x68, 0x24, 0x2d, 0xf4, 0x4f, 0x14, 0xc8, 0x99, 0xf3, 0xab,
	0xdf, 0x02, 0x6f, 0x3b, 0xbd, 0x6b, 0xef, 0x52, 0x8e, 0x38, 0xba, 0xdc, 0x82, 0x2e, 0x9b, 0xf2,
	0x78, 0x12, 0xea, 0xa5, 0xda, 0x07, 0x4a, 0xa2, 0xdb, 0x75, 0xf5, 0xd2, 0x4a, 0x0f, 0xeb, 0x37,
	0xb6, 0xd4, 0xf9, 0xc9, 0x6a, 0xf1, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97, 0x1b, 0xed,
	0x69, 0x88, 0x01, 0x00, 0x00,
}