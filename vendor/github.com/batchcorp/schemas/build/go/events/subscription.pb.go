// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: subscription.proto

package events

import (
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

type Subscription_Action int32

const (
	Subscription_UNSET_ACTION            Subscription_Action = 0
	Subscription_CREATED                 Subscription_Action = 1
	Subscription_UPDATED                 Subscription_Action = 2
	Subscription_DELETED                 Subscription_Action = 3
	Subscription_PAYMENT_METHOD_ATTACHED Subscription_Action = 4
)

// Enum value maps for Subscription_Action.
var (
	Subscription_Action_name = map[int32]string{
		0: "UNSET_ACTION",
		1: "CREATED",
		2: "UPDATED",
		3: "DELETED",
		4: "PAYMENT_METHOD_ATTACHED",
	}
	Subscription_Action_value = map[string]int32{
		"UNSET_ACTION":            0,
		"CREATED":                 1,
		"UPDATED":                 2,
		"DELETED":                 3,
		"PAYMENT_METHOD_ATTACHED": 4,
	}
)

func (x Subscription_Action) Enum() *Subscription_Action {
	p := new(Subscription_Action)
	*p = x
	return p
}

func (x Subscription_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subscription_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_subscription_proto_enumTypes[0].Descriptor()
}

func (Subscription_Action) Type() protoreflect.EnumType {
	return &file_subscription_proto_enumTypes[0]
}

func (x Subscription_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subscription_Action.Descriptor instead.
func (Subscription_Action) EnumDescriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0, 0}
}

type Subscription_ProductType int32

const (
	Subscription_UNSET_PRODUCT_TYPE Subscription_ProductType = 0
	Subscription_PLAN               Subscription_ProductType = 1
	Subscription_ADDON              Subscription_ProductType = 2
)

// Enum value maps for Subscription_ProductType.
var (
	Subscription_ProductType_name = map[int32]string{
		0: "UNSET_PRODUCT_TYPE",
		1: "PLAN",
		2: "ADDON",
	}
	Subscription_ProductType_value = map[string]int32{
		"UNSET_PRODUCT_TYPE": 0,
		"PLAN":               1,
		"ADDON":              2,
	}
)

func (x Subscription_ProductType) Enum() *Subscription_ProductType {
	p := new(Subscription_ProductType)
	*p = x
	return p
}

func (x Subscription_ProductType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subscription_ProductType) Descriptor() protoreflect.EnumDescriptor {
	return file_subscription_proto_enumTypes[1].Descriptor()
}

func (Subscription_ProductType) Type() protoreflect.EnumType {
	return &file_subscription_proto_enumTypes[1]
}

func (x Subscription_ProductType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subscription_ProductType.Descriptor instead.
func (Subscription_ProductType) EnumDescriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0, 1}
}

// Emitted by billing service as a status update
type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action                   Subscription_Action     `protobuf:"varint,1,opt,name=action,proto3,enum=events.Subscription_Action" json:"action,omitempty"`
	StripeId                 string                  `protobuf:"bytes,2,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	StripeStatus             string                  `protobuf:"bytes,3,opt,name=stripe_status,json=stripeStatus,proto3" json:"stripe_status,omitempty"`
	StripeCreatedAt          int64                   `protobuf:"varint,4,opt,name=stripe_created_at,json=stripeCreatedAt,proto3" json:"stripe_created_at,omitempty"`
	StripeCurrentPeriodStart int64                   `protobuf:"varint,5,opt,name=stripe_current_period_start,json=stripeCurrentPeriodStart,proto3" json:"stripe_current_period_start,omitempty"`
	StripeCurrentPeriodEnd   int64                   `protobuf:"varint,6,opt,name=stripe_current_period_end,json=stripeCurrentPeriodEnd,proto3" json:"stripe_current_period_end,omitempty"`
	StripePaymentMethodId    string                  `protobuf:"bytes,7,opt,name=stripe_payment_method_id,json=stripePaymentMethodId,proto3" json:"stripe_payment_method_id,omitempty"`
	StripeCustomer           *Subscription_Customer  `protobuf:"bytes,8,opt,name=stripe_customer,json=stripeCustomer,proto3" json:"stripe_customer,omitempty"`
	StripeProducts           []*Subscription_Product `protobuf:"bytes,9,rep,name=stripe_products,json=stripeProducts,proto3" json:"stripe_products,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetAction() Subscription_Action {
	if x != nil {
		return x.Action
	}
	return Subscription_UNSET_ACTION
}

func (x *Subscription) GetStripeId() string {
	if x != nil {
		return x.StripeId
	}
	return ""
}

func (x *Subscription) GetStripeStatus() string {
	if x != nil {
		return x.StripeStatus
	}
	return ""
}

func (x *Subscription) GetStripeCreatedAt() int64 {
	if x != nil {
		return x.StripeCreatedAt
	}
	return 0
}

func (x *Subscription) GetStripeCurrentPeriodStart() int64 {
	if x != nil {
		return x.StripeCurrentPeriodStart
	}
	return 0
}

func (x *Subscription) GetStripeCurrentPeriodEnd() int64 {
	if x != nil {
		return x.StripeCurrentPeriodEnd
	}
	return 0
}

func (x *Subscription) GetStripePaymentMethodId() string {
	if x != nil {
		return x.StripePaymentMethodId
	}
	return ""
}

func (x *Subscription) GetStripeCustomer() *Subscription_Customer {
	if x != nil {
		return x.StripeCustomer
	}
	return nil
}

func (x *Subscription) GetStripeProducts() []*Subscription_Product {
	if x != nil {
		return x.StripeProducts
	}
	return nil
}

type Subscription_Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            Subscription_ProductType `protobuf:"varint,1,opt,name=type,proto3,enum=events.Subscription_ProductType" json:"type,omitempty"`
	StripeItemId    string                   `protobuf:"bytes,2,opt,name=stripe_item_id,json=stripeItemId,proto3" json:"stripe_item_id,omitempty"` // subscription item id
	StripeProductId string                   `protobuf:"bytes,3,opt,name=stripe_product_id,json=stripeProductId,proto3" json:"stripe_product_id,omitempty"`
	StripeName      string                   `protobuf:"bytes,4,opt,name=stripe_name,json=stripeName,proto3" json:"stripe_name,omitempty"`
	StripePrice     float64                  `protobuf:"fixed64,5,opt,name=stripe_price,json=stripePrice,proto3" json:"stripe_price,omitempty"` // float64 in Go
	StripeCurrency  string                   `protobuf:"bytes,6,opt,name=stripe_currency,json=stripeCurrency,proto3" json:"stripe_currency,omitempty"`
	StripeQuantity  int64                    `protobuf:"varint,7,opt,name=stripe_quantity,json=stripeQuantity,proto3" json:"stripe_quantity,omitempty"`
	// How many "units" does this product represent? Ie. 10 seats, 5 collections, 100000 bytes, etc.
	BaseValue int64 `protobuf:"varint,8,opt,name=base_value,json=baseValue,proto3" json:"base_value,omitempty"`
}

func (x *Subscription_Product) Reset() {
	*x = Subscription_Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription_Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription_Product) ProtoMessage() {}

func (x *Subscription_Product) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription_Product.ProtoReflect.Descriptor instead.
func (*Subscription_Product) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Subscription_Product) GetType() Subscription_ProductType {
	if x != nil {
		return x.Type
	}
	return Subscription_UNSET_PRODUCT_TYPE
}

func (x *Subscription_Product) GetStripeItemId() string {
	if x != nil {
		return x.StripeItemId
	}
	return ""
}

func (x *Subscription_Product) GetStripeProductId() string {
	if x != nil {
		return x.StripeProductId
	}
	return ""
}

func (x *Subscription_Product) GetStripeName() string {
	if x != nil {
		return x.StripeName
	}
	return ""
}

func (x *Subscription_Product) GetStripePrice() float64 {
	if x != nil {
		return x.StripePrice
	}
	return 0
}

func (x *Subscription_Product) GetStripeCurrency() string {
	if x != nil {
		return x.StripeCurrency
	}
	return ""
}

func (x *Subscription_Product) GetStripeQuantity() int64 {
	if x != nil {
		return x.StripeQuantity
	}
	return 0
}

func (x *Subscription_Product) GetBaseValue() int64 {
	if x != nil {
		return x.BaseValue
	}
	return 0
}

type Subscription_Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StripeId        string `protobuf:"bytes,1,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	StripeName      string `protobuf:"bytes,2,opt,name=stripe_name,json=stripeName,proto3" json:"stripe_name,omitempty"`
	StripeEmail     string `protobuf:"bytes,3,opt,name=stripe_email,json=stripeEmail,proto3" json:"stripe_email,omitempty"`
	StripeCreatedAt int64  `protobuf:"varint,4,opt,name=stripe_created_at,json=stripeCreatedAt,proto3" json:"stripe_created_at,omitempty"`
}

func (x *Subscription_Customer) Reset() {
	*x = Subscription_Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription_Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription_Customer) ProtoMessage() {}

func (x *Subscription_Customer) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription_Customer.ProtoReflect.Descriptor instead.
func (*Subscription_Customer) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Subscription_Customer) GetStripeId() string {
	if x != nil {
		return x.StripeId
	}
	return ""
}

func (x *Subscription_Customer) GetStripeName() string {
	if x != nil {
		return x.StripeName
	}
	return ""
}

func (x *Subscription_Customer) GetStripeEmail() string {
	if x != nil {
		return x.StripeEmail
	}
	return ""
}

func (x *Subscription_Customer) GetStripeCreatedAt() int64 {
	if x != nil {
		return x.StripeCreatedAt
	}
	return 0
}

var File_subscription_proto protoreflect.FileDescriptor

var file_subscription_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xf2, 0x08, 0x0a,
	0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3d, 0x0a, 0x1b, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x39, 0x0a, 0x19, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x16, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x0e, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0f, 0x73,
	0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x1a, 0xc6, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x34,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x70,
	0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x97, 0x01, 0x0a, 0x08,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5e, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x41, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43,
	0x48, 0x45, 0x44, 0x10, 0x04, 0x22, 0x3a, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x5f, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x4c, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x44, 0x4f, 0x4e, 0x10,
	0x02, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscription_proto_rawDescOnce sync.Once
	file_subscription_proto_rawDescData = file_subscription_proto_rawDesc
)

func file_subscription_proto_rawDescGZIP() []byte {
	file_subscription_proto_rawDescOnce.Do(func() {
		file_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscription_proto_rawDescData)
	})
	return file_subscription_proto_rawDescData
}

var file_subscription_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_subscription_proto_goTypes = []interface{}{
	(Subscription_Action)(0),      // 0: events.Subscription.Action
	(Subscription_ProductType)(0), // 1: events.Subscription.ProductType
	(*Subscription)(nil),          // 2: events.Subscription
	(*Subscription_Product)(nil),  // 3: events.Subscription.Product
	(*Subscription_Customer)(nil), // 4: events.Subscription.Customer
}
var file_subscription_proto_depIdxs = []int32{
	0, // 0: events.Subscription.action:type_name -> events.Subscription.Action
	4, // 1: events.Subscription.stripe_customer:type_name -> events.Subscription.Customer
	3, // 2: events.Subscription.stripe_products:type_name -> events.Subscription.Product
	1, // 3: events.Subscription.Product.type:type_name -> events.Subscription.ProductType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_subscription_proto_init() }
func file_subscription_proto_init() {
	if File_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription_Product); i {
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
		file_subscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription_Customer); i {
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
			RawDescriptor: file_subscription_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_subscription_proto_goTypes,
		DependencyIndexes: file_subscription_proto_depIdxs,
		EnumInfos:         file_subscription_proto_enumTypes,
		MessageInfos:      file_subscription_proto_msgTypes,
	}.Build()
	File_subscription_proto = out.File
	file_subscription_proto_rawDesc = nil
	file_subscription_proto_goTypes = nil
	file_subscription_proto_depIdxs = nil
}
