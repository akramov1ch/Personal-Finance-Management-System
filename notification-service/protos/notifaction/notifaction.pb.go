// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: protos/notifaction/notifaction.proto

package notifactionpb

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

type NotifactionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifactionRes) Reset() {
	*x = NotifactionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notifaction_notifaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifactionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifactionRes) ProtoMessage() {}

func (x *NotifactionRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notifaction_notifaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifactionRes.ProtoReflect.Descriptor instead.
func (*NotifactionRes) Descriptor() ([]byte, []int) {
	return file_protos_notifaction_notifaction_proto_rawDescGZIP(), []int{0}
}

type NotificationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail string `protobuf:"bytes,1,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	Subject   string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotificationReq) Reset() {
	*x = NotificationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notifaction_notifaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationReq) ProtoMessage() {}

func (x *NotificationReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notifaction_notifaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationReq.ProtoReflect.Descriptor instead.
func (*NotificationReq) Descriptor() ([]byte, []int) {
	return file_protos_notifaction_notifaction_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationReq) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *NotificationReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *NotificationReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_notifaction_notifaction_proto protoreflect.FileDescriptor

var file_protos_notifaction_notifaction_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4c, 0x0a,
	0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_notifaction_notifaction_proto_rawDescOnce sync.Once
	file_protos_notifaction_notifaction_proto_rawDescData = file_protos_notifaction_notifaction_proto_rawDesc
)

func file_protos_notifaction_notifaction_proto_rawDescGZIP() []byte {
	file_protos_notifaction_notifaction_proto_rawDescOnce.Do(func() {
		file_protos_notifaction_notifaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_notifaction_notifaction_proto_rawDescData)
	})
	return file_protos_notifaction_notifaction_proto_rawDescData
}

var file_protos_notifaction_notifaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_notifaction_notifaction_proto_goTypes = []any{
	(*NotifactionRes)(nil),  // 0: NotifactionRes
	(*NotificationReq)(nil), // 1: NotificationReq
}
var file_protos_notifaction_notifaction_proto_depIdxs = []int32{
	1, // 0: NotificationService.SendNotification:input_type -> NotificationReq
	0, // 1: NotificationService.SendNotification:output_type -> NotifactionRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_notifaction_notifaction_proto_init() }
func file_protos_notifaction_notifaction_proto_init() {
	if File_protos_notifaction_notifaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_notifaction_notifaction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NotifactionRes); i {
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
		file_protos_notifaction_notifaction_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NotificationReq); i {
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
			RawDescriptor: file_protos_notifaction_notifaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_notifaction_notifaction_proto_goTypes,
		DependencyIndexes: file_protos_notifaction_notifaction_proto_depIdxs,
		MessageInfos:      file_protos_notifaction_notifaction_proto_msgTypes,
	}.Build()
	File_protos_notifaction_notifaction_proto = out.File
	file_protos_notifaction_notifaction_proto_rawDesc = nil
	file_protos_notifaction_notifaction_proto_goTypes = nil
	file_protos_notifaction_notifaction_proto_depIdxs = nil
}
