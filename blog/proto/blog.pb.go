// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: blog.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Blog struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId      string                 `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blog) Reset() {
	*x = Blog{}
	mi := &file_blog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Blog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Blog) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type BlogId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BlogId) Reset() {
	*x = BlogId{}
	mi := &file_blog_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlogId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogId) ProtoMessage() {}

func (x *BlogId) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogId.ProtoReflect.Descriptor instead.
func (*BlogId) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{1}
}

func (x *BlogId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BlogFilter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Condition     string                 `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BlogFilter) Reset() {
	*x = BlogFilter{}
	mi := &file_blog_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogFilter) ProtoMessage() {}

func (x *BlogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogFilter.ProtoReflect.Descriptor instead.
func (*BlogFilter) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{2}
}

func (x *BlogFilter) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

var File_blog_proto protoreflect.FileDescriptor

const file_blog_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"blog.proto\x12\x04blog\x1a\x1bgoogle/protobuf/empty.proto\"c\n" +
	"\x04Blog\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tauthor_id\x18\x02 \x01(\tR\bauthorId\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\"\x18\n" +
	"\x06BlogId\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"*\n" +
	"\n" +
	"BlogFilter\x12\x1c\n" +
	"\tcondition\x18\x01 \x01(\tR\tcondition2\xa7\x02\n" +
	"\vBlogService\x12&\n" +
	"\n" +
	"CreateBlog\x12\n" +
	".blog.Blog\x1a\f.blog.BlogId\x12$\n" +
	"\bReadBlog\x12\f.blog.BlogId\x1a\n" +
	".blog.Blog\x120\n" +
	"\n" +
	"UpdateBlog\x12\n" +
	".blog.Blog\x1a\x16.google.protobuf.Empty\x122\n" +
	"\n" +
	"DeleteBlog\x12\f.blog.BlogId\x1a\x16.google.protobuf.Empty\x121\n" +
	"\tListBlogs\x12\x16.google.protobuf.Empty\x1a\n" +
	".blog.Blog0\x01\x121\n" +
	"\x0fListBlogsFilter\x12\x10.blog.BlogFilter\x1a\n" +
	".blog.Blog0\x01B*Z(github.com/roster90/grp_greet/blog/protob\x06proto3"

var (
	file_blog_proto_rawDescOnce sync.Once
	file_blog_proto_rawDescData []byte
)

func file_blog_proto_rawDescGZIP() []byte {
	file_blog_proto_rawDescOnce.Do(func() {
		file_blog_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_blog_proto_rawDesc), len(file_blog_proto_rawDesc)))
	})
	return file_blog_proto_rawDescData
}

var file_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blog_proto_goTypes = []any{
	(*Blog)(nil),          // 0: blog.Blog
	(*BlogId)(nil),        // 1: blog.BlogId
	(*BlogFilter)(nil),    // 2: blog.BlogFilter
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_blog_proto_depIdxs = []int32{
	0, // 0: blog.BlogService.CreateBlog:input_type -> blog.Blog
	1, // 1: blog.BlogService.ReadBlog:input_type -> blog.BlogId
	0, // 2: blog.BlogService.UpdateBlog:input_type -> blog.Blog
	1, // 3: blog.BlogService.DeleteBlog:input_type -> blog.BlogId
	3, // 4: blog.BlogService.ListBlogs:input_type -> google.protobuf.Empty
	2, // 5: blog.BlogService.ListBlogsFilter:input_type -> blog.BlogFilter
	1, // 6: blog.BlogService.CreateBlog:output_type -> blog.BlogId
	0, // 7: blog.BlogService.ReadBlog:output_type -> blog.Blog
	3, // 8: blog.BlogService.UpdateBlog:output_type -> google.protobuf.Empty
	3, // 9: blog.BlogService.DeleteBlog:output_type -> google.protobuf.Empty
	0, // 10: blog.BlogService.ListBlogs:output_type -> blog.Blog
	0, // 11: blog.BlogService.ListBlogsFilter:output_type -> blog.Blog
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blog_proto_init() }
func file_blog_proto_init() {
	if File_blog_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_blog_proto_rawDesc), len(file_blog_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_proto_goTypes,
		DependencyIndexes: file_blog_proto_depIdxs,
		MessageInfos:      file_blog_proto_msgTypes,
	}.Build()
	File_blog_proto = out.File
	file_blog_proto_goTypes = nil
	file_blog_proto_depIdxs = nil
}
