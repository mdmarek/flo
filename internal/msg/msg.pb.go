// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	Keyed
	Progress
	Term
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Keyed struct {
	Key      string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	DataType string `protobuf:"bytes,3,opt,name=DataType" json:"DataType,omitempty"`
	TS       int64  `protobuf:"varint,4,opt,name=TS" json:"TS,omitempty"`
	Graph    string `protobuf:"bytes,6,opt,name=Graph" json:"Graph,omitempty"`
}

func (m *Keyed) Reset()                    { *m = Keyed{} }
func (m *Keyed) String() string            { return proto.CompactTextString(m) }
func (*Keyed) ProtoMessage()               {}
func (*Keyed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Keyed) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Keyed) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Keyed) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *Keyed) GetTS() int64 {
	if m != nil {
		return m.TS
	}
	return 0
}

func (m *Keyed) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

type Progress struct {
	Peer         string   `protobuf:"bytes,1,opt,name=Peer" json:"Peer,omitempty"`
	Graph        string   `protobuf:"bytes,2,opt,name=Graph" json:"Graph,omitempty"`
	Source       []string `protobuf:"bytes,3,rep,name=Source" json:"Source,omitempty"`
	Done         bool     `protobuf:"varint,4,opt,name=Done" json:"Done,omitempty"`
	MinEventTime int64    `protobuf:"varint,5,opt,name=MinEventTime" json:"MinEventTime,omitempty"`
}

func (m *Progress) Reset()                    { *m = Progress{} }
func (m *Progress) String() string            { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()               {}
func (*Progress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Progress) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *Progress) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *Progress) GetSource() []string {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Progress) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Progress) GetMinEventTime() int64 {
	if m != nil {
		return m.MinEventTime
	}
	return 0
}

type Term struct {
	Peers []string `protobuf:"bytes,1,rep,name=Peers" json:"Peers,omitempty"`
}

func (m *Term) Reset()                    { *m = Term{} }
func (m *Term) String() string            { return proto.CompactTextString(m) }
func (*Term) ProtoMessage()               {}
func (*Term) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Term) GetPeers() []string {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*Keyed)(nil), "msg.Keyed")
	proto.RegisterType((*Progress)(nil), "msg.Progress")
	proto.RegisterType((*Term)(nil), "msg.Term")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x49, 0xd2, 0x96, 0x76, 0x58, 0x44, 0x06, 0x91, 0x20, 0x1e, 0x4a, 0x4e, 0x39, 0x79,
	0xf1, 0x15, 0x14, 0x0f, 0x8b, 0xb0, 0xa4, 0x79, 0x81, 0xaa, 0x43, 0xdd, 0x43, 0x9a, 0x92, 0x54,
	0xa1, 0x77, 0x1f, 0x5c, 0x32, 0xbb, 0xf8, 0xe7, 0x94, 0xef, 0x17, 0x66, 0xf2, 0xfb, 0x08, 0x74,
	0x21, 0x4f, 0x77, 0x4b, 0x8a, 0x6b, 0x44, 0x15, 0xf2, 0x64, 0x22, 0xd4, 0x7b, 0xda, 0xe8, 0x0d,
	0x2f, 0x41, 0xed, 0x69, 0xd3, 0xa2, 0x17, 0xb6, 0x73, 0x25, 0x22, 0x42, 0xf5, 0x30, 0xae, 0xa3,
	0x96, 0xbd, 0xb0, 0x3b, 0xc7, 0x19, 0x6f, 0xa0, 0x2d, 0xa7, 0xdf, 0x16, 0xd2, 0x8a, 0x47, 0x7f,
	0x18, 0x2f, 0x40, 0xfa, 0x41, 0x57, 0xbd, 0xb0, 0xca, 0x49, 0x3f, 0xe0, 0x15, 0xd4, 0x4f, 0x69,
	0x5c, 0xde, 0x75, 0xc3, 0x83, 0x27, 0x30, 0x5f, 0x02, 0xda, 0x43, 0x8a, 0x53, 0xa2, 0x9c, 0x8b,
	0xe2, 0x40, 0x94, 0xce, 0x56, 0xce, 0xbf, 0x6b, 0xf2, 0xcf, 0x1a, 0x5e, 0x43, 0x33, 0xc4, 0x8f,
	0xf4, 0x5a, 0xb4, 0xca, 0x76, 0xee, 0x4c, 0x5c, 0x32, 0xce, 0xc4, 0xda, 0xd6, 0x71, 0x46, 0x03,
	0xbb, 0xe7, 0xe3, 0xfc, 0xf8, 0x49, 0xf3, 0xea, 0x8f, 0x81, 0x74, 0xcd, 0x95, 0xfe, 0xdd, 0x99,
	0x5b, 0xa8, 0x3c, 0xa5, 0x50, 0x6c, 0xc5, 0x9a, 0xb5, 0xe0, 0x67, 0x4f, 0xf0, 0xd2, 0xf0, 0x0f,
	0xdd, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x07, 0x61, 0x0a, 0x2e, 0x01, 0x00, 0x00,
}
