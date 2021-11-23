// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.14.0
// source: user_question.proto

package gen_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 質問
type UserQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserQuestion) Reset() {
	*x = UserQuestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQuestion) ProtoMessage() {}

func (x *UserQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_user_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQuestion.ProtoReflect.Descriptor instead.
func (*UserQuestion) Descriptor() ([]byte, []int) {
	return file_user_question_proto_rawDescGZIP(), []int{0}
}

// 質問ランダム取得レスポンス
type UserQuestion_GetRandomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	FirstAnswer  string `protobuf:"bytes,2,opt,name=first_answer,json=firstAnswer,proto3" json:"first_answer,omitempty"`
	SecondAnswer string `protobuf:"bytes,3,opt,name=second_answer,json=secondAnswer,proto3" json:"second_answer,omitempty"`
	FirstCount   uint64 `protobuf:"varint,4,opt,name=first_count,json=firstCount,proto3" json:"first_count,omitempty"`
	SecondCount  uint64 `protobuf:"varint,5,opt,name=second_count,json=secondCount,proto3" json:"second_count,omitempty"`
	FirstImgUrl  string `protobuf:"bytes,6,opt,name=first_img_url,json=firstImgUrl,proto3" json:"first_img_url,omitempty"`
	SecondImgUrl string `protobuf:"bytes,7,opt,name=second_img_url,json=secondImgUrl,proto3" json:"second_img_url,omitempty"`
}

func (x *UserQuestion_GetRandomResponse) Reset() {
	*x = UserQuestion_GetRandomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQuestion_GetRandomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQuestion_GetRandomResponse) ProtoMessage() {}

func (x *UserQuestion_GetRandomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQuestion_GetRandomResponse.ProtoReflect.Descriptor instead.
func (*UserQuestion_GetRandomResponse) Descriptor() ([]byte, []int) {
	return file_user_question_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UserQuestion_GetRandomResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UserQuestion_GetRandomResponse) GetFirstAnswer() string {
	if x != nil {
		return x.FirstAnswer
	}
	return ""
}

func (x *UserQuestion_GetRandomResponse) GetSecondAnswer() string {
	if x != nil {
		return x.SecondAnswer
	}
	return ""
}

func (x *UserQuestion_GetRandomResponse) GetFirstCount() uint64 {
	if x != nil {
		return x.FirstCount
	}
	return 0
}

func (x *UserQuestion_GetRandomResponse) GetSecondCount() uint64 {
	if x != nil {
		return x.SecondCount
	}
	return 0
}

func (x *UserQuestion_GetRandomResponse) GetFirstImgUrl() string {
	if x != nil {
		return x.FirstImgUrl
	}
	return ""
}

func (x *UserQuestion_GetRandomResponse) GetSecondImgUrl() string {
	if x != nil {
		return x.SecondImgUrl
	}
	return ""
}

var File_user_question_proto protoreflect.FileDescriptor

var file_user_question_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x77, 0x6f, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x90, 0x02, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0xff, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x49, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a,
	0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x49, 0x6d, 0x67,
	0x55, 0x72, 0x6c, 0x32, 0x63, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x12, 0x2e, 0x74, 0x77, 0x6f, 0x5f, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x74, 0x77,
	0x6f, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x67,
	0x65, 0x6e, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_question_proto_rawDescOnce sync.Once
	file_user_question_proto_rawDescData = file_user_question_proto_rawDesc
)

func file_user_question_proto_rawDescGZIP() []byte {
	file_user_question_proto_rawDescOnce.Do(func() {
		file_user_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_question_proto_rawDescData)
	})
	return file_user_question_proto_rawDescData
}

var file_user_question_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_question_proto_goTypes = []interface{}{
	(*UserQuestion)(nil),                   // 0: two_choices.UserQuestion
	(*UserQuestion_GetRandomResponse)(nil), // 1: two_choices.UserQuestion.GetRandomResponse
	(*Empty)(nil),                          // 2: two_choices.Empty
}
var file_user_question_proto_depIdxs = []int32{
	2, // 0: two_choices.UserQuestionService.GetRandom:input_type -> two_choices.Empty
	1, // 1: two_choices.UserQuestionService.GetRandom:output_type -> two_choices.UserQuestion.GetRandomResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_question_proto_init() }
func file_user_question_proto_init() {
	if File_user_question_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_question_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQuestion); i {
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
		file_user_question_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQuestion_GetRandomResponse); i {
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
			RawDescriptor: file_user_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_question_proto_goTypes,
		DependencyIndexes: file_user_question_proto_depIdxs,
		MessageInfos:      file_user_question_proto_msgTypes,
	}.Build()
	File_user_question_proto = out.File
	file_user_question_proto_rawDesc = nil
	file_user_question_proto_goTypes = nil
	file_user_question_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserQuestionServiceClient is the client API for UserQuestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserQuestionServiceClient interface {
	GetRandom(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserQuestion_GetRandomResponse, error)
}

type userQuestionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserQuestionServiceClient(cc grpc.ClientConnInterface) UserQuestionServiceClient {
	return &userQuestionServiceClient{cc}
}

func (c *userQuestionServiceClient) GetRandom(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserQuestion_GetRandomResponse, error) {
	out := new(UserQuestion_GetRandomResponse)
	err := c.cc.Invoke(ctx, "/two_choices.UserQuestionService/GetRandom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserQuestionServiceServer is the server API for UserQuestionService service.
type UserQuestionServiceServer interface {
	GetRandom(context.Context, *Empty) (*UserQuestion_GetRandomResponse, error)
}

// UnimplementedUserQuestionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserQuestionServiceServer struct {
}

func (*UnimplementedUserQuestionServiceServer) GetRandom(context.Context, *Empty) (*UserQuestion_GetRandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandom not implemented")
}

func RegisterUserQuestionServiceServer(s *grpc.Server, srv UserQuestionServiceServer) {
	s.RegisterService(&_UserQuestionService_serviceDesc, srv)
}

func _UserQuestionService_GetRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserQuestionServiceServer).GetRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/two_choices.UserQuestionService/GetRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserQuestionServiceServer).GetRandom(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserQuestionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "two_choices.UserQuestionService",
	HandlerType: (*UserQuestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandom",
			Handler:    _UserQuestionService_GetRandom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_question.proto",
}
