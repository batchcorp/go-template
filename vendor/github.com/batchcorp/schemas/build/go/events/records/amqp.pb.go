// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/records/amqp.proto

package records

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// See: https://godoc.org/github.com/streadway/amqp#Delivery
type AMQPSinkRecord struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Exchange             string   `protobuf:"bytes,4,opt,name=exchange,proto3" json:"exchange,omitempty"`
	RoutingKey           string   `protobuf:"bytes,5,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty"`
	ContentType          string   `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	ContentEncoding      string   `protobuf:"bytes,7,opt,name=content_encoding,json=contentEncoding,proto3" json:"content_encoding,omitempty"`
	Priority             int32    `protobuf:"varint,8,opt,name=priority,proto3" json:"priority,omitempty"`
	Expiration           string   `protobuf:"bytes,9,opt,name=expiration,proto3" json:"expiration,omitempty"`
	MessageId            string   `protobuf:"bytes,10,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	UserId               string   `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AppId                string   `protobuf:"bytes,12,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ReplyTo              string   `protobuf:"bytes,13,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
	CorrelationId        string   `protobuf:"bytes,14,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AMQPSinkRecord) Reset()         { *m = AMQPSinkRecord{} }
func (m *AMQPSinkRecord) String() string { return proto.CompactTextString(m) }
func (*AMQPSinkRecord) ProtoMessage()    {}
func (*AMQPSinkRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ad538798beb46, []int{0}
}

func (m *AMQPSinkRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AMQPSinkRecord.Unmarshal(m, b)
}
func (m *AMQPSinkRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AMQPSinkRecord.Marshal(b, m, deterministic)
}
func (m *AMQPSinkRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AMQPSinkRecord.Merge(m, src)
}
func (m *AMQPSinkRecord) XXX_Size() int {
	return xxx_messageInfo_AMQPSinkRecord.Size(m)
}
func (m *AMQPSinkRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AMQPSinkRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AMQPSinkRecord proto.InternalMessageInfo

func (m *AMQPSinkRecord) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *AMQPSinkRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AMQPSinkRecord) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AMQPSinkRecord) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *AMQPSinkRecord) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

func (m *AMQPSinkRecord) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *AMQPSinkRecord) GetContentEncoding() string {
	if m != nil {
		return m.ContentEncoding
	}
	return ""
}

func (m *AMQPSinkRecord) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *AMQPSinkRecord) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

func (m *AMQPSinkRecord) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *AMQPSinkRecord) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AMQPSinkRecord) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *AMQPSinkRecord) GetReplyTo() string {
	if m != nil {
		return m.ReplyTo
	}
	return ""
}

func (m *AMQPSinkRecord) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func init() {
	proto.RegisterType((*AMQPSinkRecord)(nil), "records.AMQPSinkRecord")
}

func init() { proto.RegisterFile("events/records/amqp.proto", fileDescriptor_f84ad538798beb46) }

var fileDescriptor_f84ad538798beb46 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xdf, 0x8a, 0xd5, 0x30,
	0x10, 0xc6, 0xa9, 0x67, 0xcf, 0x9f, 0xce, 0x39, 0x7b, 0x94, 0x80, 0x6c, 0x56, 0xfc, 0x53, 0x05,
	0xa1, 0xde, 0xb4, 0x17, 0xfa, 0x02, 0x0a, 0x5e, 0x14, 0x11, 0xb5, 0xee, 0x95, 0x37, 0x25, 0x4d,
	0x86, 0x36, 0xec, 0x69, 0x12, 0x93, 0x54, 0xb6, 0x2f, 0xe5, 0x33, 0x4a, 0x67, 0xeb, 0xae, 0xde,
	0xe5, 0xfb, 0x7d, 0xdf, 0x4c, 0x06, 0x66, 0xe0, 0x12, 0x7f, 0xa1, 0x89, 0xa1, 0xf4, 0x28, 0xad,
	0x57, 0xa1, 0x14, 0xc3, 0x4f, 0x57, 0x38, 0x6f, 0xa3, 0x65, 0xdb, 0x85, 0xbd, 0xfa, 0xbd, 0x82,
	0xe3, 0xfb, 0xcf, 0xdf, 0xbe, 0x7e, 0xd7, 0xe6, 0xba, 0x26, 0xc6, 0x18, 0x9c, 0xb5, 0x56, 0x4d,
	0x3c, 0xc9, 0x92, 0xfc, 0x50, 0xd3, 0x9b, 0x3d, 0x85, 0x34, 0xea, 0x01, 0x43, 0x14, 0x83, 0xe3,
	0x0f, 0xb2, 0x24, 0x5f, 0xd5, 0xf7, 0x60, 0xae, 0x88, 0x93, 0x43, 0xbe, 0xca, 0x92, 0x3c, 0xad,
	0xe9, 0xcd, 0x9e, 0xc0, 0x0e, 0x6f, 0x64, 0x2f, 0x4c, 0x87, 0xfc, 0x8c, 0xf8, 0x9d, 0x66, 0x2f,
	0x60, 0xef, 0xed, 0x18, 0xb5, 0xe9, 0x9a, 0x6b, 0x9c, 0xf8, 0x9a, 0x6c, 0x58, 0xd0, 0x27, 0x9c,
	0xd8, 0x4b, 0x38, 0x48, 0x6b, 0x22, 0x9a, 0xd8, 0x50, 0xe3, 0x0d, 0x25, 0xf6, 0x0b, 0xbb, 0x9a,
	0xfb, 0xbf, 0x81, 0x47, 0x7f, 0x23, 0x68, 0xa4, 0x55, 0xda, 0x74, 0x7c, 0x4b, 0xb1, 0x87, 0x0b,
	0xff, 0xb8, 0xe0, 0x79, 0x14, 0xe7, 0xb5, 0xf5, 0x3a, 0x4e, 0x7c, 0x97, 0x25, 0xf9, 0xba, 0xbe,
	0xd3, 0xec, 0x39, 0x00, 0xde, 0x38, 0xed, 0x45, 0xd4, 0xd6, 0xf0, 0xf4, 0x76, 0x92, 0x7b, 0xc2,
	0x9e, 0x01, 0x0c, 0x18, 0x82, 0xe8, 0xb0, 0xd1, 0x8a, 0x03, 0xf9, 0xe9, 0x42, 0x2a, 0xc5, 0x2e,
	0x60, 0x3b, 0x06, 0xf4, 0xb3, 0xb7, 0x27, 0x6f, 0x33, 0xcb, 0x4a, 0xb1, 0xc7, 0xb0, 0x11, 0xce,
	0xcd, 0xfc, 0x40, 0x7c, 0x2d, 0x9c, 0xab, 0x14, 0xbb, 0x84, 0x9d, 0x47, 0x77, 0x9a, 0x9a, 0x68,
	0xf9, 0x39, 0x19, 0x5b, 0xd2, 0x57, 0x96, 0xbd, 0x86, 0xa3, 0xb4, 0xde, 0xe3, 0x89, 0x3e, 0x9e,
	0x2b, 0x8f, 0x14, 0x38, 0xff, 0x87, 0x56, 0xea, 0xc3, 0x17, 0xb8, 0x08, 0x7d, 0xd1, 0x8a, 0x28,
	0xfb, 0xe2, 0x76, 0xbf, 0xc5, 0xb2, 0xcb, 0x1f, 0xef, 0x3a, 0x1d, 0xfb, 0xb1, 0x2d, 0xa4, 0x1d,
	0x4a, 0x0a, 0x48, 0xeb, 0x5d, 0x19, 0x64, 0x8f, 0x83, 0x08, 0x65, 0x3b, 0xea, 0x93, 0x2a, 0x3b,
	0x5b, 0xfe, 0x7f, 0x15, 0xed, 0x86, 0x2e, 0xe2, 0xed, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5,
	0xfd, 0x28, 0x58, 0x2e, 0x02, 0x00, 0x00,
}
