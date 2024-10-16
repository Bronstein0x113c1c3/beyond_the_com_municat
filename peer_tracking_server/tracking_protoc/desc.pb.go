// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: desc.proto

package tracking_protoc

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

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_desc_proto_rawDescGZIP(), []int{0}
}

func (x *StopRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IncomingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *IncomingRequest) Reset() {
	*x = IncomingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingRequest) ProtoMessage() {}

func (x *IncomingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingRequest.ProtoReflect.Descriptor instead.
func (*IncomingRequest) Descriptor() ([]byte, []int) {
	return file_desc_proto_rawDescGZIP(), []int{1}
}

func (x *IncomingRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *IncomingRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type SelfNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReceipentList []*IncomingRequest `protobuf:"bytes,2,rep,name=receipent_list,json=receipentList,proto3" json:"receipent_list,omitempty"`
}

func (x *SelfNotify) Reset() {
	*x = SelfNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfNotify) ProtoMessage() {}

func (x *SelfNotify) ProtoReflect() protoreflect.Message {
	mi := &file_desc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfNotify.ProtoReflect.Descriptor instead.
func (*SelfNotify) Descriptor() ([]byte, []int) {
	return file_desc_proto_rawDescGZIP(), []int{2}
}

func (x *SelfNotify) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SelfNotify) GetReceipentList() []*IncomingRequest {
	if x != nil {
		return x.ReceipentList
	}
	return nil
}

type OtherNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Info *IncomingRequest `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *OtherNotify) Reset() {
	*x = OtherNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherNotify) ProtoMessage() {}

func (x *OtherNotify) ProtoReflect() protoreflect.Message {
	mi := &file_desc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherNotify.ProtoReflect.Descriptor instead.
func (*OtherNotify) Descriptor() ([]byte, []int) {
	return file_desc_proto_rawDescGZIP(), []int{3}
}

func (x *OtherNotify) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OtherNotify) GetInfo() *IncomingRequest {
	if x != nil {
		return x.Info
	}
	return nil
}

type Notify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*Notify_Self
	//	*Notify_Other
	Request isNotify_Request `protobuf_oneof:"request"`
	Status  bool             `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Notify) Reset() {
	*x = Notify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notify) ProtoMessage() {}

func (x *Notify) ProtoReflect() protoreflect.Message {
	mi := &file_desc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notify.ProtoReflect.Descriptor instead.
func (*Notify) Descriptor() ([]byte, []int) {
	return file_desc_proto_rawDescGZIP(), []int{4}
}

func (m *Notify) GetRequest() isNotify_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *Notify) GetSelf() *SelfNotify {
	if x, ok := x.GetRequest().(*Notify_Self); ok {
		return x.Self
	}
	return nil
}

func (x *Notify) GetOther() *OtherNotify {
	if x, ok := x.GetRequest().(*Notify_Other); ok {
		return x.Other
	}
	return nil
}

func (x *Notify) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type isNotify_Request interface {
	isNotify_Request()
}

type Notify_Self struct {
	Self *SelfNotify `protobuf:"bytes,1,opt,name=self,proto3,oneof"`
}

type Notify_Other struct {
	Other *OtherNotify `protobuf:"bytes,2,opt,name=other,proto3,oneof"`
}

func (*Notify_Self) isNotify_Request() {}

func (*Notify_Other) isNotify_Request() {}

var File_desc_proto protoreflect.FileDescriptor

var file_desc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b,
	0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x55, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x66, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x37, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x0b, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x74,
	0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x65, 0x6c, 0x66, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x48, 0x00, 0x52, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x12, 0x24, 0x0a, 0x05, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0x4c, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x22, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x10, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x30, 0x01, 0x12, 0x1c, 0x0a, 0x03, 0x4f, 0x75, 0x74, 0x12, 0x0c, 0x2e, 0x53, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_desc_proto_rawDescOnce sync.Once
	file_desc_proto_rawDescData = file_desc_proto_rawDesc
)

func file_desc_proto_rawDescGZIP() []byte {
	file_desc_proto_rawDescOnce.Do(func() {
		file_desc_proto_rawDescData = protoimpl.X.CompressGZIP(file_desc_proto_rawDescData)
	})
	return file_desc_proto_rawDescData
}

var file_desc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_desc_proto_goTypes = []interface{}{
	(*StopRequest)(nil),     // 0: StopRequest
	(*IncomingRequest)(nil), // 1: IncomingRequest
	(*SelfNotify)(nil),      // 2: SelfNotify
	(*OtherNotify)(nil),     // 3: OtherNotify
	(*Notify)(nil),          // 4: Notify
}
var file_desc_proto_depIdxs = []int32{
	1, // 0: SelfNotify.receipent_list:type_name -> IncomingRequest
	1, // 1: OtherNotify.info:type_name -> IncomingRequest
	2, // 2: Notify.self:type_name -> SelfNotify
	3, // 3: Notify.other:type_name -> OtherNotify
	1, // 4: Tracking.Add:input_type -> IncomingRequest
	0, // 5: Tracking.Out:input_type -> StopRequest
	4, // 6: Tracking.Add:output_type -> Notify
	4, // 7: Tracking.Out:output_type -> Notify
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_desc_proto_init() }
func file_desc_proto_init() {
	if File_desc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_desc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_desc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomingRequest); i {
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
		file_desc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfNotify); i {
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
		file_desc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherNotify); i {
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
		file_desc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notify); i {
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
	file_desc_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Notify_Self)(nil),
		(*Notify_Other)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_desc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_desc_proto_goTypes,
		DependencyIndexes: file_desc_proto_depIdxs,
		MessageInfos:      file_desc_proto_msgTypes,
	}.Build()
	File_desc_proto = out.File
	file_desc_proto_rawDesc = nil
	file_desc_proto_goTypes = nil
	file_desc_proto_depIdxs = nil
}
