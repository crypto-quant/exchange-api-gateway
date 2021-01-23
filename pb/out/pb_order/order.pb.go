// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: order.proto

package pb_order

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair                   string                `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	Price                  float64               `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Amount                 float64               `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	AvgPrice               float64               `protobuf:"fixed64,4,opt,name=avgPrice,proto3" json:"avgPrice,omitempty"`
	DealAmount             float64               `protobuf:"fixed64,5,opt,name=dealAmount,proto3" json:"dealAmount,omitempty"`
	Fee                    float64               `protobuf:"fixed64,6,opt,name=fee,proto3" json:"fee,omitempty"`
	Cid                    string                `protobuf:"bytes,7,opt,name=cid,proto3" json:"cid,omitempty"` // client order id
	OrderId                int64                 `protobuf:"varint,8,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Status                 pb_common.OrderStatus `protobuf:"varint,9,opt,name=status,proto3,enum=OrderStatus" json:"status,omitempty"` // string{"UNFINISH", "PART_FINISH", "FINISH", "CANCEL", "REJECT", "CANCEL_ING", "FAIL"}
	Side                   pb_common.Side        `protobuf:"varint,10,opt,name=side,proto3,enum=Side" json:"side,omitempty"`
	TimeInForce            string                `protobuf:"bytes,11,opt,name=timeInForce,proto3" json:"timeInForce,omitempty"` // GTC, IOC, FOK
	OrderType              string                `protobuf:"bytes,12,opt,name=orderType,proto3" json:"orderType,omitempty"`
	OrderTime              int64                 `protobuf:"varint,13,opt,name=orderTime,proto3" json:"orderTime,omitempty"`
	FinishedTime           int64                 `protobuf:"varint,14,opt,name=finishedTime,proto3" json:"finishedTime,omitempty"`
	TradeId                int64                 `protobuf:"varint,15,opt,name=tradeId,proto3" json:"tradeId,omitempty"`
	CancelledClientOrderId string                `protobuf:"bytes,16,opt,name=cancelledClientOrderId,proto3" json:"cancelledClientOrderId,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *Order) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order) GetAvgPrice() float64 {
	if x != nil {
		return x.AvgPrice
	}
	return 0
}

func (x *Order) GetDealAmount() float64 {
	if x != nil {
		return x.DealAmount
	}
	return 0
}

func (x *Order) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Order) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *Order) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Order) GetStatus() pb_common.OrderStatus {
	if x != nil {
		return x.Status
	}
	return pb_common.OrderStatus_ORDER_UNFINISH
}

func (x *Order) GetSide() pb_common.Side {
	if x != nil {
		return x.Side
	}
	return pb_common.Side_UNKNOWN
}

func (x *Order) GetTimeInForce() string {
	if x != nil {
		return x.TimeInForce
	}
	return ""
}

func (x *Order) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *Order) GetOrderTime() int64 {
	if x != nil {
		return x.OrderTime
	}
	return 0
}

func (x *Order) GetFinishedTime() int64 {
	if x != nil {
		return x.FinishedTime
	}
	return 0
}

func (x *Order) GetTradeId() int64 {
	if x != nil {
		return x.TradeId
	}
	return 0
}

func (x *Order) GetCancelledClientOrderId() string {
	if x != nil {
		return x.CancelledClientOrderId
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x03, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x76, 0x67, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x76, 0x67, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73, 0x69,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x46, 0x6f, 0x72, 0x63,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x46,
	0x6f, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x12, 0x36,
	0x0a, 0x16, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42, 0x0e, 0x5a, 0x0c, 0x6f, 0x75, 0x74, 0x2f, 0x70, 0x62,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_order_proto_goTypes = []interface{}{
	(*Order)(nil),              // 0: Order
	(pb_common.OrderStatus)(0), // 1: OrderStatus
	(pb_common.Side)(0),        // 2: Side
}
var file_order_proto_depIdxs = []int32{
	1, // 0: Order.status:type_name -> OrderStatus
	2, // 1: Order.side:type_name -> Side
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
