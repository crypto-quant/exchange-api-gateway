// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pb/ticker.proto

package pb_ticker

import (
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

type Ticker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair string  `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	Last float64 `protobuf:"fixed64,2,opt,name=last,proto3" json:"last,omitempty"`
	Buy  float64 `protobuf:"fixed64,3,opt,name=buy,proto3" json:"buy,omitempty"`
	Sell float64 `protobuf:"fixed64,4,opt,name=sell,proto3" json:"sell,omitempty"`
	High float64 `protobuf:"fixed64,5,opt,name=high,proto3" json:"high,omitempty"`
	Low  float64 `protobuf:"fixed64,6,opt,name=low,proto3" json:"low,omitempty"`
	Vol  float64 `protobuf:"fixed64,7,opt,name=vol,proto3" json:"vol,omitempty"`
	Date int64   `protobuf:"varint,8,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Ticker) Reset() {
	*x = Ticker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticker) ProtoMessage() {}

func (x *Ticker) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticker.ProtoReflect.Descriptor instead.
func (*Ticker) Descriptor() ([]byte, []int) {
	return file_pb_ticker_proto_rawDescGZIP(), []int{0}
}

func (x *Ticker) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *Ticker) GetLast() float64 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *Ticker) GetBuy() float64 {
	if x != nil {
		return x.Buy
	}
	return 0
}

func (x *Ticker) GetSell() float64 {
	if x != nil {
		return x.Sell
	}
	return 0
}

func (x *Ticker) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Ticker) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Ticker) GetVol() float64 {
	if x != nil {
		return x.Vol
	}
	return 0
}

func (x *Ticker) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

var File_pb_ticker_proto protoreflect.FileDescriptor

var file_pb_ticker_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x69, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6c, 0x61, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x75, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x62, 0x75, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x65, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69,
	0x67, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77,
	0x12, 0x10, 0x0a, 0x03, 0x76, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x76,
	0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x6f, 0x75, 0x74, 0x2f, 0x70, 0x62,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_ticker_proto_rawDescOnce sync.Once
	file_pb_ticker_proto_rawDescData = file_pb_ticker_proto_rawDesc
)

func file_pb_ticker_proto_rawDescGZIP() []byte {
	file_pb_ticker_proto_rawDescOnce.Do(func() {
		file_pb_ticker_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_ticker_proto_rawDescData)
	})
	return file_pb_ticker_proto_rawDescData
}

var file_pb_ticker_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_ticker_proto_goTypes = []interface{}{
	(*Ticker)(nil), // 0: Ticker
}
var file_pb_ticker_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_ticker_proto_init() }
func file_pb_ticker_proto_init() {
	if File_pb_ticker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_ticker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticker); i {
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
			RawDescriptor: file_pb_ticker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_ticker_proto_goTypes,
		DependencyIndexes: file_pb_ticker_proto_depIdxs,
		MessageInfos:      file_pb_ticker_proto_msgTypes,
	}.Build()
	File_pb_ticker_proto = out.File
	file_pb_ticker_proto_rawDesc = nil
	file_pb_ticker_proto_goTypes = nil
	file_pb_ticker_proto_depIdxs = nil
}
