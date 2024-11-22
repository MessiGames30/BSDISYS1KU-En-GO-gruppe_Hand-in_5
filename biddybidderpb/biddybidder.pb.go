// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: biddybidder.proto

package biddybidderpb

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

// message containing successtatus
type SuccessStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SuccessStart) Reset() {
	*x = SuccessStart{}
	mi := &file_biddybidder_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuccessStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessStart) ProtoMessage() {}

func (x *SuccessStart) ProtoReflect() protoreflect.Message {
	mi := &file_biddybidder_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessStart.ProtoReflect.Descriptor instead.
func (*SuccessStart) Descriptor() ([]byte, []int) {
	return file_biddybidder_proto_rawDescGZIP(), []int{0}
}

func (x *SuccessStart) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// auctionDetails
type AuctionObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeCreated   int64 `protobuf:"varint,1,opt,name=timeCreated,proto3" json:"timeCreated,omitempty"`
	Duration      int64 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	HighestBid    int64 `protobuf:"varint,3,opt,name=highestBid,proto3" json:"highestBid,omitempty"`       // The current highest bid
	HighestBidder int64 `protobuf:"varint,4,opt,name=highestBidder,proto3" json:"highestBidder,omitempty"` // The current highest bidder
	CurrentTime   int64 `protobuf:"varint,5,opt,name=currentTime,proto3" json:"currentTime,omitempty"`
}

func (x *AuctionObject) Reset() {
	*x = AuctionObject{}
	mi := &file_biddybidder_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuctionObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionObject) ProtoMessage() {}

func (x *AuctionObject) ProtoReflect() protoreflect.Message {
	mi := &file_biddybidder_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionObject.ProtoReflect.Descriptor instead.
func (*AuctionObject) Descriptor() ([]byte, []int) {
	return file_biddybidder_proto_rawDescGZIP(), []int{1}
}

func (x *AuctionObject) GetTimeCreated() int64 {
	if x != nil {
		return x.TimeCreated
	}
	return 0
}

func (x *AuctionObject) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *AuctionObject) GetHighestBid() int64 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *AuctionObject) GetHighestBidder() int64 {
	if x != nil {
		return x.HighestBidder
	}
	return 0
}

func (x *AuctionObject) GetCurrentTime() int64 {
	if x != nil {
		return x.CurrentTime
	}
	return 0
}

// auctionDetails
type AuctionDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeleft      int64 `protobuf:"varint,1,opt,name=timeleft,proto3" json:"timeleft,omitempty"`           // The time that is left on the auction
	HighestBid    int64 `protobuf:"varint,2,opt,name=highestBid,proto3" json:"highestBid,omitempty"`       // The current highest bid
	HighestBidder int64 `protobuf:"varint,3,opt,name=highestBidder,proto3" json:"highestBidder,omitempty"` // The current highest bidder
}

func (x *AuctionDetails) Reset() {
	*x = AuctionDetails{}
	mi := &file_biddybidder_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuctionDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionDetails) ProtoMessage() {}

func (x *AuctionDetails) ProtoReflect() protoreflect.Message {
	mi := &file_biddybidder_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionDetails.ProtoReflect.Descriptor instead.
func (*AuctionDetails) Descriptor() ([]byte, []int) {
	return file_biddybidder_proto_rawDescGZIP(), []int{2}
}

func (x *AuctionDetails) GetTimeleft() int64 {
	if x != nil {
		return x.Timeleft
	}
	return 0
}

func (x *AuctionDetails) GetHighestBid() int64 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *AuctionDetails) GetHighestBidder() int64 {
	if x != nil {
		return x.HighestBidder
	}
	return 0
}

// Message a client sends to the chat
type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidderId int64 `protobuf:"varint,1,opt,name=bidderId,proto3" json:"bidderId,omitempty"`
	Amount   int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"` // Bidding amount for the auction item
}

func (x *Bid) Reset() {
	*x = Bid{}
	mi := &file_biddybidder_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_biddybidder_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_biddybidder_proto_rawDescGZIP(), []int{3}
}

func (x *Bid) GetBidderId() int64 {
	if x != nil {
		return x.BidderId
	}
	return 0
}

func (x *Bid) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	mi := &file_biddybidder_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_biddybidder_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_biddybidder_proto_rawDescGZIP(), []int{4}
}

func (x *Ack) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type Time struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Time) Reset() {
	*x = Time{}
	mi := &file_biddybidder_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Time) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Time) ProtoMessage() {}

func (x *Time) ProtoReflect() protoreflect.Message {
	mi := &file_biddybidder_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Time.ProtoReflect.Descriptor instead.
func (*Time) Descriptor() ([]byte, []int) {
	return file_biddybidder_proto_rawDescGZIP(), []int{5}
}

func (x *Time) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

// Empty message
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_biddybidder_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_biddybidder_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_biddybidder_proto_rawDescGZIP(), []int{6}
}

var File_biddybidder_proto protoreflect.FileDescriptor

var file_biddybidder_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x69, 0x64, 0x64, 0x79, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb5, 0x01,
	0x0a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x72, 0x0a, 0x0e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c,
	0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c,
	0x65, 0x66, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x69, 0x67, 0x68,
	0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x03, 0x62, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1d, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0x07, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x8d, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x04, 0x2e, 0x62, 0x69,
	0x64, 0x1a, 0x04, 0x2e, 0x61, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x06, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x1a, 0x0d, 0x2e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x25, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x1a, 0x06, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x31, 0x5a, 0x2f, 0x42, 0x53, 0x44, 0x49,
	0x53, 0x59, 0x53, 0x31, 0x4b, 0x55, 0x2d, 0x45, 0x6e, 0x2d, 0x47, 0x4f, 0x2d, 0x67, 0x72, 0x75,
	0x70, 0x70, 0x65, 0x5f, 0x48, 0x61, 0x6e, 0x64, 0x2d, 0x69, 0x6e, 0x5f, 0x35, 0x2f, 0x62, 0x69,
	0x64, 0x64, 0x79, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_biddybidder_proto_rawDescOnce sync.Once
	file_biddybidder_proto_rawDescData = file_biddybidder_proto_rawDesc
)

func file_biddybidder_proto_rawDescGZIP() []byte {
	file_biddybidder_proto_rawDescOnce.Do(func() {
		file_biddybidder_proto_rawDescData = protoimpl.X.CompressGZIP(file_biddybidder_proto_rawDescData)
	})
	return file_biddybidder_proto_rawDescData
}

var file_biddybidder_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_biddybidder_proto_goTypes = []any{
	(*SuccessStart)(nil),   // 0: successStart
	(*AuctionObject)(nil),  // 1: auctionObject
	(*AuctionDetails)(nil), // 2: auctionDetails
	(*Bid)(nil),            // 3: bid
	(*Ack)(nil),            // 4: ack
	(*Time)(nil),           // 5: time
	(*Empty)(nil),          // 6: empty
}
var file_biddybidder_proto_depIdxs = []int32{
	3, // 0: Auction.Bid:input_type -> bid
	6, // 1: Auction.Result:input_type -> empty
	5, // 2: Auction.StartFunction:input_type -> time
	1, // 3: Auction.SyncAuction:input_type -> auctionObject
	4, // 4: Auction.Bid:output_type -> ack
	2, // 5: Auction.Result:output_type -> auctionDetails
	0, // 6: Auction.StartFunction:output_type -> successStart
	6, // 7: Auction.SyncAuction:output_type -> empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_biddybidder_proto_init() }
func file_biddybidder_proto_init() {
	if File_biddybidder_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_biddybidder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biddybidder_proto_goTypes,
		DependencyIndexes: file_biddybidder_proto_depIdxs,
		MessageInfos:      file_biddybidder_proto_msgTypes,
	}.Build()
	File_biddybidder_proto = out.File
	file_biddybidder_proto_rawDesc = nil
	file_biddybidder_proto_goTypes = nil
	file_biddybidder_proto_depIdxs = nil
}
