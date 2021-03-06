// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.5
// source: fields.proto

package service

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

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Признак обязательного поля
	Required *bool `protobuf:"varint,1,opt,name=required,proto3,oneof" json:"required,omitempty"`
	// Признак Примера значения поля
	Example *string `protobuf:"bytes,2,opt,name=example,proto3,oneof" json:"example,omitempty"`
	// Признак что поле передаётся в заголовке
	InHeader *bool `protobuf:"varint,3,opt,name=in_header,json=inHeader,proto3,oneof" json:"in_header,omitempty"`
	// Название заголовка для поля
	HeaderName *string `protobuf:"bytes,4,opt,name=header_name,json=headerName,proto3,oneof" json:"header_name,omitempty"`
	// Признак что поле передаётся в куках
	InCookies *bool `protobuf:"varint,5,opt,name=in_cookies,json=inCookies,proto3,oneof" json:"in_cookies,omitempty"`
	// Название для куки
	CookiesName *string `protobuf:"bytes,6,opt,name=cookies_name,json=cookiesName,proto3,oneof" json:"cookies_name,omitempty"`
	// Признак того что поле в теле передаётся
	InBody *bool `protobuf:"varint,7,opt,name=in_body,json=inBody,proto3,oneof" json:"in_body,omitempty"`
	// Название для json
	JsonName *string `protobuf:"bytes,8,opt,name=json_name,json=jsonName,proto3,oneof" json:"json_name,omitempty"`
	// Признак того что поле передаётся в урле
	InUrl *bool `protobuf:"varint,9,opt,name=in_url,json=inUrl,proto3,oneof" json:"in_url,omitempty"`
	// Название для url
	UrlName *string `protobuf:"bytes,10,opt,name=url_name,json=urlName,proto3,oneof" json:"url_name,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fields_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_fields_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_fields_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetRequired() bool {
	if x != nil && x.Required != nil {
		return *x.Required
	}
	return false
}

func (x *Options) GetExample() string {
	if x != nil && x.Example != nil {
		return *x.Example
	}
	return ""
}

func (x *Options) GetInHeader() bool {
	if x != nil && x.InHeader != nil {
		return *x.InHeader
	}
	return false
}

func (x *Options) GetHeaderName() string {
	if x != nil && x.HeaderName != nil {
		return *x.HeaderName
	}
	return ""
}

func (x *Options) GetInCookies() bool {
	if x != nil && x.InCookies != nil {
		return *x.InCookies
	}
	return false
}

func (x *Options) GetCookiesName() string {
	if x != nil && x.CookiesName != nil {
		return *x.CookiesName
	}
	return ""
}

func (x *Options) GetInBody() bool {
	if x != nil && x.InBody != nil {
		return *x.InBody
	}
	return false
}

func (x *Options) GetJsonName() string {
	if x != nil && x.JsonName != nil {
		return *x.JsonName
	}
	return ""
}

func (x *Options) GetInUrl() bool {
	if x != nil && x.InUrl != nil {
		return *x.InUrl
	}
	return false
}

func (x *Options) GetUrlName() string {
	if x != nil && x.UrlName != nil {
		return *x.UrlName
	}
	return ""
}

var File_fields_proto protoreflect.FileDescriptor

var file_fields_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x65, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x64, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0xe2, 0x03, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x69, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x02, 0x52, 0x08, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x09, 0x69, 0x6e, 0x43, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05,
	0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1c, 0x0a, 0x07, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x06, 0x52, 0x06, 0x69, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x88, 0x01, 0x01, 0x12, 0x20,
	0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x07, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1a, 0x0a, 0x06, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x08, 0x52, 0x05, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08,
	0x75, 0x72, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09,
	0x52, 0x07, 0x75, 0x72, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6e, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69,
	0x65, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x72, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x65, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x64, 0x6f, 0x2e, 0x72, 0x75, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fields_proto_rawDescOnce sync.Once
	file_fields_proto_rawDescData = file_fields_proto_rawDesc
)

func file_fields_proto_rawDescGZIP() []byte {
	file_fields_proto_rawDescOnce.Do(func() {
		file_fields_proto_rawDescData = protoimpl.X.CompressGZIP(file_fields_proto_rawDescData)
	})
	return file_fields_proto_rawDescData
}

var file_fields_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fields_proto_goTypes = []interface{}{
	(*Options)(nil), // 0: eldorado.service.Options
}
var file_fields_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fields_proto_init() }
func file_fields_proto_init() {
	if File_fields_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fields_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
	file_fields_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fields_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fields_proto_goTypes,
		DependencyIndexes: file_fields_proto_depIdxs,
		MessageInfos:      file_fields_proto_msgTypes,
	}.Build()
	File_fields_proto = out.File
	file_fields_proto_rawDesc = nil
	file_fields_proto_goTypes = nil
	file_fields_proto_depIdxs = nil
}
