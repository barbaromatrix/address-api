// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: api/proto/v1/address-api.proto

package v1

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

type IpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *IpRequest) Reset() {
	*x = IpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_address_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpRequest) ProtoMessage() {}

func (x *IpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_address_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpRequest.ProtoReflect.Descriptor instead.
func (*IpRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_address_api_proto_rawDescGZIP(), []int{0}
}

func (x *IpRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type IpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message     string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Country     string  `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	CountryCode string  `protobuf:"bytes,4,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	Region      string  `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	RegionName  string  `protobuf:"bytes,6,opt,name=regionName,proto3" json:"regionName,omitempty"`
	City        string  `protobuf:"bytes,7,opt,name=city,proto3" json:"city,omitempty"`
	Zip         string  `protobuf:"bytes,8,opt,name=zip,proto3" json:"zip,omitempty"`
	Lat         float64 `protobuf:"fixed64,9,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon         float64 `protobuf:"fixed64,10,opt,name=lon,proto3" json:"lon,omitempty"`
	Timezone    string  `protobuf:"bytes,11,opt,name=timezone,proto3" json:"timezone,omitempty"`
}

func (x *IpResponse) Reset() {
	*x = IpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_address_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpResponse) ProtoMessage() {}

func (x *IpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_address_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpResponse.ProtoReflect.Descriptor instead.
func (*IpResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_address_api_proto_rawDescGZIP(), []int{1}
}

func (x *IpResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *IpResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *IpResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *IpResponse) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *IpResponse) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *IpResponse) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *IpResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *IpResponse) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *IpResponse) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *IpResponse) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *IpResponse) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

var File_api_proto_v1_address_api_proto protoreflect.FileDescriptor

var file_api_proto_v1_address_api_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x1b, 0x0a, 0x09,
	0x49, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x98, 0x02, 0x0a, 0x0a, 0x49, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x7a, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x32, 0x4c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49,
	0x70, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x56, 0x61, 0x6e, 0x65, 0x73, 0x73, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x6c, 0x61, 0x72, 0x69,
	0x6e, 0x69, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_address_api_proto_rawDescOnce sync.Once
	file_api_proto_v1_address_api_proto_rawDescData = file_api_proto_v1_address_api_proto_rawDesc
)

func file_api_proto_v1_address_api_proto_rawDescGZIP() []byte {
	file_api_proto_v1_address_api_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_address_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_address_api_proto_rawDescData)
	})
	return file_api_proto_v1_address_api_proto_rawDescData
}

var file_api_proto_v1_address_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_v1_address_api_proto_goTypes = []interface{}{
	(*IpRequest)(nil),  // 0: address.v1.IpRequest
	(*IpResponse)(nil), // 1: address.v1.IpResponse
}
var file_api_proto_v1_address_api_proto_depIdxs = []int32{
	0, // 0: address.v1.Address.GetAddressByIp:input_type -> address.v1.IpRequest
	1, // 1: address.v1.Address.GetAddressByIp:output_type -> address.v1.IpResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_v1_address_api_proto_init() }
func file_api_proto_v1_address_api_proto_init() {
	if File_api_proto_v1_address_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_address_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpRequest); i {
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
		file_api_proto_v1_address_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpResponse); i {
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
			RawDescriptor: file_api_proto_v1_address_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_address_api_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_address_api_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_address_api_proto_msgTypes,
	}.Build()
	File_api_proto_v1_address_api_proto = out.File
	file_api_proto_v1_address_api_proto_rawDesc = nil
	file_api_proto_v1_address_api_proto_goTypes = nil
	file_api_proto_v1_address_api_proto_depIdxs = nil
}
