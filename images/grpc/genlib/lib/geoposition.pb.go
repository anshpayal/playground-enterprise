// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: lib/geoposition.proto

package lib

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

type GeoPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GeoPosition) Reset() {
	*x = GeoPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_geoposition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoPosition) ProtoMessage() {}

func (x *GeoPosition) ProtoReflect() protoreflect.Message {
	mi := &file_lib_geoposition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoPosition.ProtoReflect.Descriptor instead.
func (*GeoPosition) Descriptor() ([]byte, []int) {
	return file_lib_geoposition_proto_rawDescGZIP(), []int{0}
}

func (x *GeoPosition) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GeoPosition) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_lib_geoposition_proto protoreflect.FileDescriptor

var file_lib_geoposition_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x65, 0x6f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x72, 0x61, 0x6b, 0x65, 0x6e, 0x64,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x47, 0x0a, 0x0b, 0x47,
	0x65, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x6b, 0x65, 0x6e, 0x64, 0x69, 0x6f, 0x2f, 0x70, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67,
	0x65, 0x6e, 0x6c, 0x69, 0x62, 0x2f, 0x6c, 0x69, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lib_geoposition_proto_rawDescOnce sync.Once
	file_lib_geoposition_proto_rawDescData = file_lib_geoposition_proto_rawDesc
)

func file_lib_geoposition_proto_rawDescGZIP() []byte {
	file_lib_geoposition_proto_rawDescOnce.Do(func() {
		file_lib_geoposition_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_geoposition_proto_rawDescData)
	})
	return file_lib_geoposition_proto_rawDescData
}

var file_lib_geoposition_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lib_geoposition_proto_goTypes = []interface{}{
	(*GeoPosition)(nil), // 0: krakend_playground.GeoPosition
}
var file_lib_geoposition_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lib_geoposition_proto_init() }
func file_lib_geoposition_proto_init() {
	if File_lib_geoposition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_geoposition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoPosition); i {
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
			RawDescriptor: file_lib_geoposition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lib_geoposition_proto_goTypes,
		DependencyIndexes: file_lib_geoposition_proto_depIdxs,
		MessageInfos:      file_lib_geoposition_proto_msgTypes,
	}.Build()
	File_lib_geoposition_proto = out.File
	file_lib_geoposition_proto_rawDesc = nil
	file_lib_geoposition_proto_goTypes = nil
	file_lib_geoposition_proto_depIdxs = nil
}
