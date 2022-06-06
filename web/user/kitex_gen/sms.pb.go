// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.0
// source: sms.proto

package kitex_gen

import (
	context "context"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone   string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	ImgCode string `protobuf:"bytes,2,opt,name=imgCode,proto3" json:"imgCode,omitempty"`
	Uuid    string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Request) GetImgCode() string {
	if x != nil {
		return x.ImgCode
	}
	return ""
}

func (x *Request) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//返回错误码和信息
	Errno  string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *Response) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

var File_sms_proto protoreflect.FileDescriptor

var file_sms_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x4d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x32, 0x64, 0x0a, 0x0a, 0x53, 0x4d,
	0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x4d, 0x53, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x07, 0x53, 0x61, 0x76, 0x65, 0x53, 0x4d, 0x53, 0x12,
	0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_proto_rawDescOnce sync.Once
	file_sms_proto_rawDescData = file_sms_proto_rawDesc
)

func file_sms_proto_rawDescGZIP() []byte {
	file_sms_proto_rawDescOnce.Do(func() {
		file_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_proto_rawDescData)
	})
	return file_sms_proto_rawDescData
}

var file_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sms_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: user.Request
	(*Response)(nil), // 1: user.Response
}
var file_sms_proto_depIdxs = []int32{
	0, // 0: user.SMSService.SendSMS:input_type -> user.Request
	0, // 1: user.SMSService.SaveSMS:input_type -> user.Request
	1, // 2: user.SMSService.SendSMS:output_type -> user.Response
	1, // 3: user.SMSService.SaveSMS:output_type -> user.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_proto_init() }
func file_sms_proto_init() {
	if File_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_proto_goTypes,
		DependencyIndexes: file_sms_proto_depIdxs,
		MessageInfos:      file_sms_proto_msgTypes,
	}.Build()
	File_sms_proto = out.File
	file_sms_proto_rawDesc = nil
	file_sms_proto_goTypes = nil
	file_sms_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.3.2. DO NOT EDIT.

type SMSService interface {
	SendSMS(ctx context.Context, req *Request) (res *Response, err error)
	SaveSMS(ctx context.Context, req *Request) (res *Response, err error)
}