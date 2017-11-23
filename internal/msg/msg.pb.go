// Code generated by protoc-gen-go.
// source: msg.proto
// DO NOT EDIT!

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	Event
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

type Event struct {
	Key      string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	DataType string `protobuf:"bytes,3,opt,name=DataType" json:"DataType,omitempty"`
	Time     int64  `protobuf:"varint,4,opt,name=Time" json:"Time,omitempty"`
	Graph    string `protobuf:"bytes,6,opt,name=Graph" json:"Graph,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *Event) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Event) GetGraph() string {
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
	proto.RegisterType((*Event)(nil), "msg.Event")
	proto.RegisterType((*Progress)(nil), "msg.Progress")
	proto.RegisterType((*Term)(nil), "msg.Term")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x6a, 0x87, 0x30,
	0x14, 0xc5, 0x89, 0x51, 0xd1, 0x8b, 0x43, 0xb9, 0x94, 0x12, 0x4a, 0x07, 0xc9, 0x94, 0xa9, 0x4b,
	0x5f, 0xa1, 0xa5, 0x43, 0x29, 0x48, 0xea, 0x0b, 0xd8, 0x72, 0xb1, 0x0e, 0x31, 0x92, 0xd8, 0x82,
	0x7b, 0x1f, 0xbc, 0xe4, 0x2a, 0xff, 0x8f, 0xc9, 0x73, 0xe4, 0x9c, 0xfc, 0x4e, 0x02, 0xb5, 0x8b,
	0xe3, 0xe3, 0x12, 0xfc, 0xea, 0x51, 0xba, 0x38, 0xea, 0x08, 0xc5, 0xcb, 0x2f, 0xcd, 0x2b, 0xde,
	0x80, 0x7c, 0xa3, 0x4d, 0x89, 0x56, 0x98, 0xda, 0x26, 0x89, 0x08, 0xf9, 0xf3, 0xb0, 0x0e, 0x2a,
	0x6b, 0x85, 0x69, 0x2c, 0x6b, 0xbc, 0x87, 0x2a, 0x7d, 0xfb, 0x6d, 0x21, 0x25, 0x39, 0x7a, 0xf2,
	0x29, 0xdf, 0x4f, 0x8e, 0x54, 0xde, 0x0a, 0x23, 0x2d, 0x6b, 0xbc, 0x85, 0xe2, 0x35, 0x0c, 0xcb,
	0xb7, 0x2a, 0x39, 0xbc, 0x1b, 0xfd, 0x27, 0xa0, 0xea, 0x82, 0x1f, 0x03, 0xc5, 0x98, 0x6a, 0x1d,
	0x51, 0x38, 0xc8, 0xac, 0xcf, 0xb5, 0xec, 0xa2, 0x86, 0x77, 0x50, 0x7e, 0xf8, 0x9f, 0xf0, 0x95,
	0xd0, 0xd2, 0xd4, 0xf6, 0x70, 0x3c, 0xd4, 0xcf, 0x3b, 0xb8, 0xb2, 0xac, 0x51, 0x43, 0xf3, 0x3e,
	0xcd, 0x7c, 0x35, 0x1e, 0x55, 0xf0, 0xa8, 0xab, 0x7f, 0xfa, 0x01, 0xf2, 0x9e, 0x82, 0x4b, 0xb4,
	0x44, 0x8d, 0x4a, 0xf0, 0xb1, 0xbb, 0xf9, 0x2c, 0xf9, 0x95, 0x9e, 0xfe, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x18, 0x05, 0x21, 0x78, 0x32, 0x01, 0x00, 0x00,
}
