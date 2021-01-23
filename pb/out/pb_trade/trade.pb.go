// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: trade.proto

package pb_trade

import (
	pb_common "github.com/crypto-quant/exchange-api-gateway/pb/out/pb_common"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Trade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair    string         `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	Tid     int64          `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`
	OrderId int64          `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Side    pb_common.Side `protobuf:"varint,4,opt,name=side,proto3,enum=Side" json:"side,omitempty"`
	Amount  float64        `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Price   float64        `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Date    int64          `protobuf:"varint,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Trade) Reset() {
	*x = Trade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trade) ProtoMessage() {}

func (x *Trade) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trade.ProtoReflect.Descriptor instead.
func (*Trade) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{0}
}

func (x *Trade) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *Trade) GetTid() int64 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *Trade) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Trade) GetSide() pb_common.Side {
	if x != nil {
		return x.Side
	}
	return pb_common.Side_UNKNOWN
}

func (x *Trade) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Trade) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Trade) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

var File_trade_proto protoreflect.FileDescriptor

var file_trade_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x05,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x6f, 0x75, 0x74, 0x2f, 0x70, 0x62, 0x5f, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trade_proto_rawDescOnce sync.Once
	file_trade_proto_rawDescData = file_trade_proto_rawDesc
)

func file_trade_proto_rawDescGZIP() []byte {
	file_trade_proto_rawDescOnce.Do(func() {
		file_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_proto_rawDescData)
	})
	return file_trade_proto_rawDescData
}

var file_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_trade_proto_goTypes = []interface{}{
	(*Trade)(nil),       // 0: Trade
	(pb_common.Side)(0), // 1: Side
}
var file_trade_proto_depIdxs = []int32{
	1, // 0: Trade.side:type_name -> Side
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trade_proto_init() }
func file_trade_proto_init() {
	if File_trade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trade); i {
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
			RawDescriptor: file_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trade_proto_goTypes,
		DependencyIndexes: file_trade_proto_depIdxs,
		MessageInfos:      file_trade_proto_msgTypes,
	}.Build()
	File_trade_proto = out.File
	file_trade_proto_rawDesc = nil
	file_trade_proto_goTypes = nil
	file_trade_proto_depIdxs = nil
}
