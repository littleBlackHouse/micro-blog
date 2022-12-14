// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: proto/article/article.proto

package article

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

type ArticleListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	LastId  string `protobuf:"bytes,2,opt,name=lastId,proto3" json:"lastId,omitempty"`
	Page    int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit   int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ArticleListRequest) Reset() {
	*x = ArticleListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleListRequest) ProtoMessage() {}

func (x *ArticleListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleListRequest.ProtoReflect.Descriptor instead.
func (*ArticleListRequest) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{0}
}

func (x *ArticleListRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ArticleListRequest) GetLastId() string {
	if x != nil {
		return x.LastId
	}
	return ""
}

func (x *ArticleListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ArticleListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ArticleListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*ArticleListResponse_Data `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ArticleListResponse) Reset() {
	*x = ArticleListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleListResponse) ProtoMessage() {}

func (x *ArticleListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleListResponse.ProtoReflect.Descriptor instead.
func (*ArticleListResponse) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ArticleListResponse) GetData() []*ArticleListResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type ArticleCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Image   string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ArticleCreateRequest) Reset() {
	*x = ArticleCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCreateRequest) ProtoMessage() {}

func (x *ArticleCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCreateRequest.ProtoReflect.Descriptor instead.
func (*ArticleCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleCreateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ArticleCreateRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type ArticleCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ArticleCreateResponse) Reset() {
	*x = ArticleCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCreateResponse) ProtoMessage() {}

func (x *ArticleCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCreateResponse.ProtoReflect.Descriptor instead.
func (*ArticleCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{3}
}

type ArticleListResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                           `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string                           `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Image     string                           `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	CreatedAt string                           `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Author    *ArticleListResponse_Data_Author `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *ArticleListResponse_Data) Reset() {
	*x = ArticleListResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleListResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleListResponse_Data) ProtoMessage() {}

func (x *ArticleListResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleListResponse_Data.ProtoReflect.Descriptor instead.
func (*ArticleListResponse_Data) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ArticleListResponse_Data) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArticleListResponse_Data) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleListResponse_Data) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ArticleListResponse_Data) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ArticleListResponse_Data) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ArticleListResponse_Data) GetAuthor() *ArticleListResponse_Data_Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type ArticleListResponse_Data_Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ArticleListResponse_Data_Author) Reset() {
	*x = ArticleListResponse_Data_Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleListResponse_Data_Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleListResponse_Data_Author) ProtoMessage() {}

func (x *ArticleListResponse_Data_Author) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleListResponse_Data_Author.ProtoReflect.Descriptor instead.
func (*ArticleListResponse_Data_Author) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *ArticleListResponse_Data_Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArticleListResponse_Data_Author) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_proto_article_article_proto protoreflect.FileDescriptor

var file_proto_article_article_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x70, 0x0a, 0x12, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xd7, 0x02, 0x0a, 0x13, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xf2, 0x01,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x34, 0x0a, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x14, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x17, 0x0a, 0x15, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x99, 0x01, 0x0a, 0x07, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_article_article_proto_rawDescOnce sync.Once
	file_proto_article_article_proto_rawDescData = file_proto_article_article_proto_rawDesc
)

func file_proto_article_article_proto_rawDescGZIP() []byte {
	file_proto_article_article_proto_rawDescOnce.Do(func() {
		file_proto_article_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_article_article_proto_rawDescData)
	})
	return file_proto_article_article_proto_rawDescData
}

var file_proto_article_article_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_article_article_proto_goTypes = []interface{}{
	(*ArticleListRequest)(nil),              // 0: article.ArticleListRequest
	(*ArticleListResponse)(nil),             // 1: article.ArticleListResponse
	(*ArticleCreateRequest)(nil),            // 2: article.ArticleCreateRequest
	(*ArticleCreateResponse)(nil),           // 3: article.ArticleCreateResponse
	(*ArticleListResponse_Data)(nil),        // 4: article.ArticleListResponse.Data
	(*ArticleListResponse_Data_Author)(nil), // 5: article.ArticleListResponse.Data.Author
}
var file_proto_article_article_proto_depIdxs = []int32{
	4, // 0: article.ArticleListResponse.data:type_name -> article.ArticleListResponse.Data
	5, // 1: article.ArticleListResponse.Data.author:type_name -> article.ArticleListResponse.Data.Author
	0, // 2: article.Article.List:input_type -> article.ArticleListRequest
	2, // 3: article.Article.Create:input_type -> article.ArticleCreateRequest
	1, // 4: article.Article.List:output_type -> article.ArticleListResponse
	3, // 5: article.Article.Create:output_type -> article.ArticleCreateResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_article_article_proto_init() }
func file_proto_article_article_proto_init() {
	if File_proto_article_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_article_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleListRequest); i {
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
		file_proto_article_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleListResponse); i {
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
		file_proto_article_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleCreateRequest); i {
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
		file_proto_article_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleCreateResponse); i {
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
		file_proto_article_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleListResponse_Data); i {
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
		file_proto_article_article_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleListResponse_Data_Author); i {
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
			RawDescriptor: file_proto_article_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_article_article_proto_goTypes,
		DependencyIndexes: file_proto_article_article_proto_depIdxs,
		MessageInfos:      file_proto_article_article_proto_msgTypes,
	}.Build()
	File_proto_article_article_proto = out.File
	file_proto_article_article_proto_rawDesc = nil
	file_proto_article_article_proto_goTypes = nil
	file_proto_article_article_proto_depIdxs = nil
}
