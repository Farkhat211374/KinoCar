// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.1
// source: movie_message.proto

package pb

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

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Rating   float32  `protobuf:"fixed32,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Year     int32    `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
	Director string   `protobuf:"bytes,5,opt,name=director,proto3" json:"director,omitempty"`
	Genre    []string `protobuf:"bytes,6,rep,name=genre,proto3" json:"genre,omitempty"`
	Duration int32    `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_movie_message_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Movie) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Movie) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *Movie) GetGenre() []string {
	if x != nil {
		return x.Genre
	}
	return nil
}

func (x *Movie) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_movie_message_proto protoreflect.FileDescriptor

var file_movie_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xa7, 0x01, 0x0a, 0x05, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x61, 0x72, 0x6b, 0x68, 0x61, 0x74, 0x2f, 0x4b, 0x69, 0x6e, 0x6f, 0x43, 0x61,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_message_proto_rawDescOnce sync.Once
	file_movie_message_proto_rawDescData = file_movie_message_proto_rawDesc
)

func file_movie_message_proto_rawDescGZIP() []byte {
	file_movie_message_proto_rawDescOnce.Do(func() {
		file_movie_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_message_proto_rawDescData)
	})
	return file_movie_message_proto_rawDescData
}

var file_movie_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_movie_message_proto_goTypes = []interface{}{
	(*Movie)(nil), // 0: pb.Movie
}
var file_movie_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_movie_message_proto_init() }
func file_movie_message_proto_init() {
	if File_movie_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
			RawDescriptor: file_movie_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_movie_message_proto_goTypes,
		DependencyIndexes: file_movie_message_proto_depIdxs,
		MessageInfos:      file_movie_message_proto_msgTypes,
	}.Build()
	File_movie_message_proto = out.File
	file_movie_message_proto_rawDesc = nil
	file_movie_message_proto_goTypes = nil
	file_movie_message_proto_depIdxs = nil
}