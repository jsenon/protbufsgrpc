// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

/*
Package mailing is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Client
*/
package mailing

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

type Client struct {
	Id      int32          `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email   string         `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Country string         `protobuf:"bytes,4,opt,name=country" json:"country,omitempty"`
	Inbox   []*Client_Mail `protobuf:"bytes,5,rep,name=inbox" json:"inbox,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Client) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Client) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Client) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Client) GetInbox() []*Client_Mail {
	if m != nil {
		return m.Inbox
	}
	return nil
}

type Client_Mail struct {
	RemoteEmail string `protobuf:"bytes,1,opt,name=remoteEmail" json:"remoteEmail,omitempty"`
	Body        string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Client_Mail) Reset()                    { *m = Client_Mail{} }
func (m *Client_Mail) String() string            { return proto.CompactTextString(m) }
func (*Client_Mail) ProtoMessage()               {}
func (*Client_Mail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Client_Mail) GetRemoteEmail() string {
	if m != nil {
		return m.RemoteEmail
	}
	return ""
}

func (m *Client_Mail) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Client)(nil), "mailing.Client")
	proto.RegisterType((*Client_Mail)(nil), "mailing.Client.Mail")
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x6a, 0xc4, 0x30,
	0x10, 0x44, 0x91, 0x6d, 0xd9, 0x64, 0x0d, 0x81, 0x2c, 0x2e, 0x44, 0x2a, 0x91, 0xca, 0xa4, 0x50,
	0x91, 0xb4, 0xe9, 0x42, 0xca, 0x34, 0xfe, 0x03, 0x29, 0x12, 0x41, 0xc4, 0xd6, 0x06, 0x59, 0x07,
	0xe7, 0x1f, 0xbc, 0xef, 0x3a, 0x24, 0x9f, 0xe1, 0xba, 0x99, 0x61, 0x77, 0xdf, 0x2c, 0x3c, 0x69,
	0x6b, 0xa3, 0x5b, 0x57, 0x43, 0xf4, 0xa7, 0xfe, 0x23, 0x25, 0xc2, 0x6e, 0xd1, 0x7e, 0xf6, 0xe1,
	0xf7, 0xe5, 0xc2, 0xa0, 0xfd, 0x9c, 0xbd, 0x0b, 0x09, 0x1f, 0xa1, 0xf2, 0x56, 0x30, 0xc9, 0x46,
	0x3e, 0x55, 0xde, 0x22, 0x42, 0x13, 0xf4, 0xe2, 0x44, 0x25, 0xd9, 0xf8, 0x30, 0x15, 0x8d, 0x03,
	0x70, 0x97, 0x57, 0x45, 0x5d, 0xc2, 0xdd, 0xa0, 0x80, 0xee, 0x87, 0x4e, 0x21, 0xc5, 0x4d, 0x34,
	0x25, 0x3f, 0x2c, 0xbe, 0x02, 0xf7, 0xc1, 0xd0, 0x59, 0x70, 0x59, 0x8f, 0xfd, 0xdb, 0xa0, 0x6e,
	0x5c, 0xb5, 0x33, 0xd5, 0xb7, 0xf6, 0xf3, 0xb4, 0x8f, 0x3c, 0x7f, 0x40, 0x93, 0x2d, 0x4a, 0xe8,
	0xa3, 0x5b, 0x28, 0xb9, 0xaf, 0x42, 0x62, 0xe5, 0xe2, 0x7d, 0x94, 0x9b, 0x19, 0xb2, 0xdb, 0xd1,
	0x2c, 0x6b, 0xd3, 0x96, 0xc7, 0xde, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x7f, 0x31, 0x36,
	0xed, 0x00, 0x00, 0x00,
}