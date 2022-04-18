// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.0--rc2
// source: proto/hello_grpc.proto

package hello_grpc

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

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_proto_hello_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_proto_hello_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Res) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Req2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age     uint32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Req2) Reset() {
	*x = Req2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req2) ProtoMessage() {}

func (x *Req2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req2.ProtoReflect.Descriptor instead.
func (*Req2) Descriptor() ([]byte, []int) {
	return file_proto_hello_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *Req2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Req2) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Req2) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Res2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age     uint32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Res2) Reset() {
	*x = Res2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res2) ProtoMessage() {}

func (x *Res2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res2.ProtoReflect.Descriptor instead.
func (*Res2) Descriptor() ([]byte, []int) {
	return file_proto_hello_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *Res2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Res2) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Res2) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_hello_grpc_proto protoreflect.FileDescriptor

var file_proto_hello_grpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x22, 0x1f, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x04, 0x52, 0x65, 0x71, 0x32, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46,
	0x0a, 0x04, 0x52, 0x65, 0x73, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x66, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x47,
	0x52, 0x50, 0x43, 0x12, 0x29, 0x0a, 0x05, 0x53, 0x61, 0x79, 0x48, 0x69, 0x12, 0x0f, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x12, 0x2e,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x10, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x32, 0x1a, 0x10, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x32, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_grpc_proto_rawDescOnce sync.Once
	file_proto_hello_grpc_proto_rawDescData = file_proto_hello_grpc_proto_rawDesc
)

func file_proto_hello_grpc_proto_rawDescGZIP() []byte {
	file_proto_hello_grpc_proto_rawDescOnce.Do(func() {
		file_proto_hello_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_grpc_proto_rawDescData)
	})
	return file_proto_hello_grpc_proto_rawDescData
}

var file_proto_hello_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_hello_grpc_proto_goTypes = []interface{}{
	(*Req)(nil),  // 0: hello_grpc.Req
	(*Res)(nil),  // 1: hello_grpc.Res
	(*Req2)(nil), // 2: hello_grpc.Req2
	(*Res2)(nil), // 3: hello_grpc.Res2
}
var file_proto_hello_grpc_proto_depIdxs = []int32{
	0, // 0: hello_grpc.HelloGRPC.SayHi:input_type -> hello_grpc.Req
	2, // 1: hello_grpc.HelloGRPC.SayHello:input_type -> hello_grpc.Req2
	1, // 2: hello_grpc.HelloGRPC.SayHi:output_type -> hello_grpc.Res
	3, // 3: hello_grpc.HelloGRPC.SayHello:output_type -> hello_grpc.Res2
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_hello_grpc_proto_init() }
func file_proto_hello_grpc_proto_init() {
	if File_proto_hello_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_proto_hello_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
		file_proto_hello_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req2); i {
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
		file_proto_hello_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res2); i {
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
			RawDescriptor: file_proto_hello_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_grpc_proto_goTypes,
		DependencyIndexes: file_proto_hello_grpc_proto_depIdxs,
		MessageInfos:      file_proto_hello_grpc_proto_msgTypes,
	}.Build()
	File_proto_hello_grpc_proto = out.File
	file_proto_hello_grpc_proto_rawDesc = nil
	file_proto_hello_grpc_proto_goTypes = nil
	file_proto_hello_grpc_proto_depIdxs = nil
}
