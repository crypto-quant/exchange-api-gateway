// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: depth.proto

package pb_depth

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

type Depth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair string              `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	Date int64               `protobuf:"varint,2,opt,name=date,proto3" json:"date,omitempty"`
	Asks []*Depth_PriceLevel `protobuf:"bytes,3,rep,name=asks,proto3" json:"asks,omitempty"`
	Bids []*Depth_PriceLevel `protobuf:"bytes,4,rep,name=bids,proto3" json:"bids,omitempty"`
}

func (x *Depth) Reset() {
	*x = Depth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Depth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Depth) ProtoMessage() {}

func (x *Depth) ProtoReflect() protoreflect.Message {
	mi := &file_depth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Depth.ProtoReflect.Descriptor instead.
func (*Depth) Descriptor() ([]byte, []int) {
	return file_depth_proto_rawDescGZIP(), []int{0}
}

func (x *Depth) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *Depth) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Depth) GetAsks() []*Depth_PriceLevel {
	if x != nil {
		return x.Asks
	}
	return nil
}

func (x *Depth) GetBids() []*Depth_PriceLevel {
	if x != nil {
		return x.Bids
	}
	return nil
}

type Depth_PriceLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price  float64 `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Depth_PriceLevel) Reset() {
	*x = Depth_PriceLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Depth_PriceLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Depth_PriceLevel) ProtoMessage() {}

func (x *Depth_PriceLevel) ProtoReflect() protoreflect.Message {
	mi := &file_depth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Depth_PriceLevel.ProtoReflect.Descriptor instead.
func (*Depth_PriceLevel) Descriptor() ([]byte, []int) {
	return file_depth_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Depth_PriceLevel) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Depth_PriceLevel) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_depth_proto protoreflect.FileDescriptor

var file_depth_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x65, 0x70, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01,
	0x0a, 0x05, 0x44, 0x65, 0x70, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x44, 0x65, 0x70, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73, 0x1a, 0x3a, 0x0a,
	0x0a, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x6f, 0x75, 0x74,
	0x2f, 0x70, 0x62, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_depth_proto_rawDescOnce sync.Once
	file_depth_proto_rawDescData = file_depth_proto_rawDesc
)

func file_depth_proto_rawDescGZIP() []byte {
	file_depth_proto_rawDescOnce.Do(func() {
		file_depth_proto_rawDescData = protoimpl.X.CompressGZIP(file_depth_proto_rawDescData)
	})
	return file_depth_proto_rawDescData
}

var file_depth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_depth_proto_goTypes = []interface{}{
	(*Depth)(nil),            // 0: Depth
	(*Depth_PriceLevel)(nil), // 1: Depth.PriceLevel
}
var file_depth_proto_depIdxs = []int32{
	1, // 0: Depth.asks:type_name -> Depth.PriceLevel
	1, // 1: Depth.bids:type_name -> Depth.PriceLevel
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_depth_proto_init() }
func file_depth_proto_init() {
	if File_depth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_depth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Depth); i {
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
		file_depth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Depth_PriceLevel); i {
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
			RawDescriptor: file_depth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_depth_proto_goTypes,
		DependencyIndexes: file_depth_proto_depIdxs,
		MessageInfos:      file_depth_proto_msgTypes,
	}.Build()
	File_depth_proto = out.File
	file_depth_proto_rawDesc = nil
	file_depth_proto_goTypes = nil
	file_depth_proto_depIdxs = nil
}