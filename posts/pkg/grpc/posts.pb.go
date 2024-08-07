// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: posts.proto

package pkg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EditCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string        `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	CommentID string        `protobuf:"bytes,2,opt,name=CommentID,proto3" json:"CommentID,omitempty"`
	Value     *CommentValue `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *EditCommentReq) Reset() {
	*x = EditCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditCommentReq) ProtoMessage() {}

func (x *EditCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditCommentReq.ProtoReflect.Descriptor instead.
func (*EditCommentReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{0}
}

func (x *EditCommentReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EditCommentReq) GetCommentID() string {
	if x != nil {
		return x.CommentID
	}
	return ""
}

func (x *EditCommentReq) GetValue() *CommentValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type DeleteCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	CommentID string `protobuf:"bytes,2,opt,name=CommentID,proto3" json:"CommentID,omitempty"`
}

func (x *DeleteCommentReq) Reset() {
	*x = DeleteCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentReq) ProtoMessage() {}

func (x *DeleteCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentReq.ProtoReflect.Descriptor instead.
func (*DeleteCommentReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteCommentReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteCommentReq) GetCommentID() string {
	if x != nil {
		return x.CommentID
	}
	return ""
}

type LikePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID string `protobuf:"bytes,1,opt,name=PostID,proto3" json:"PostID,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *LikePostReq) Reset() {
	*x = LikePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikePostReq) ProtoMessage() {}

func (x *LikePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikePostReq.ProtoReflect.Descriptor instead.
func (*LikePostReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{2}
}

func (x *LikePostReq) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

func (x *LikePostReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LeaveCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string        `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	PostID string        `protobuf:"bytes,2,opt,name=PostID,proto3" json:"PostID,omitempty"`
	Value  *CommentValue `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *LeaveCommentReq) Reset() {
	*x = LeaveCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveCommentReq) ProtoMessage() {}

func (x *LeaveCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveCommentReq.ProtoReflect.Descriptor instead.
func (*LeaveCommentReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveCommentReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LeaveCommentReq) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

func (x *LeaveCommentReq) GetValue() *CommentValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type CreatePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	File        string `protobuf:"bytes,4,opt,name=File,proto3" json:"File,omitempty"`
}

func (x *CreatePostReq) Reset() {
	*x = CreatePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostReq) ProtoMessage() {}

func (x *CreatePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostReq.ProtoReflect.Descriptor instead.
func (*CreatePostReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePostReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreatePostReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePostReq) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type DeletePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	PostID string `protobuf:"bytes,2,opt,name=PostID,proto3" json:"PostID,omitempty"`
}

func (x *DeletePostReq) Reset() {
	*x = DeletePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostReq) ProtoMessage() {}

func (x *DeletePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostReq.ProtoReflect.Descriptor instead.
func (*DeletePostReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePostReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeletePostReq) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

type EditPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	PostID      string `protobuf:"bytes,2,opt,name=PostID,proto3" json:"PostID,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	File        string `protobuf:"bytes,5,opt,name=File,proto3" json:"File,omitempty"`
}

func (x *EditPostReq) Reset() {
	*x = EditPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditPostReq) ProtoMessage() {}

func (x *EditPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditPostReq.ProtoReflect.Descriptor instead.
func (*EditPostReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{6}
}

func (x *EditPostReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EditPostReq) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

func (x *EditPostReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EditPostReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EditPostReq) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type GetUserPostsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetUserPostsReq) Reset() {
	*x = GetUserPostsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPostsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPostsReq) ProtoMessage() {}

func (x *GetUserPostsReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPostsReq.ProtoReflect.Descriptor instead.
func (*GetUserPostsReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserPostsReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetUserPostsReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserPostsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=Posts,proto3" json:"Posts,omitempty"`
}

func (x *GetUserPostsRes) Reset() {
	*x = GetUserPostsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPostsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPostsRes) ProtoMessage() {}

func (x *GetUserPostsRes) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPostsRes.ProtoReflect.Descriptor instead.
func (*GetUserPostsRes) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserPostsRes) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      string                 `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	PostID      string                 `protobuf:"bytes,2,opt,name=PostID,proto3" json:"PostID,omitempty"`
	Title       string                 `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	File        string                 `protobuf:"bytes,5,opt,name=File,proto3" json:"File,omitempty"`
	Likes       []*Like                `protobuf:"bytes,6,rep,name=Likes,proto3" json:"Likes,omitempty"`
	Comments    []*Comment             `protobuf:"bytes,7,rep,name=Comments,proto3" json:"Comments,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{9}
}

func (x *Post) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Post) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Post) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Post) GetLikes() []*Like {
	if x != nil {
		return x.Likes
	}
	return nil
}

func (x *Post) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string                 `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	PostID    string                 `protobuf:"bytes,2,opt,name=PostID,proto3" json:"PostID,omitempty"`
	CommentID string                 `protobuf:"bytes,3,opt,name=CommentID,proto3" json:"CommentID,omitempty"`
	Value     *CommentValue          `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{10}
}

func (x *Comment) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Comment) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

func (x *Comment) GetCommentID() string {
	if x != nil {
		return x.CommentID
	}
	return ""
}

func (x *Comment) GetValue() *CommentValue {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CommentValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TextValue string `protobuf:"bytes,1,opt,name=TextValue,proto3" json:"TextValue,omitempty"`
	File      string `protobuf:"bytes,2,opt,name=File,proto3" json:"File,omitempty"`
}

func (x *CommentValue) Reset() {
	*x = CommentValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentValue) ProtoMessage() {}

func (x *CommentValue) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentValue.ProtoReflect.Descriptor instead.
func (*CommentValue) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{11}
}

func (x *CommentValue) GetTextValue() string {
	if x != nil {
		return x.TextValue
	}
	return ""
}

func (x *CommentValue) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	PostID string `protobuf:"bytes,2,opt,name=PostID,proto3" json:"PostID,omitempty"`
}

func (x *Like) Reset() {
	*x = Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Like) ProtoMessage() {}

func (x *Like) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Like.ProtoReflect.Descriptor instead.
func (*Like) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{12}
}

func (x *Like) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Like) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

var File_posts_proto protoreflect.FileDescriptor

var file_posts_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0e, 0x45,
	0x64, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x23, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x46, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x3b,
	0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x64, 0x0a, 0x0f, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x71, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x46, 0x69, 0x6c, 0x65, 0x22, 0x3d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x44, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x3f, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2e,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x12, 0x1b, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0xff,
	0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x73,
	0x12, 0x24, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xb6, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x40, 0x0a, 0x0c, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x78,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65,
	0x78, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x36, 0x0a, 0x04, 0x4c,
	0x69, 0x6b, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x44, 0x32, 0xed, 0x03, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x34, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x0e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x08, 0x45, 0x64, 0x69,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x08, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x4c, 0x69,
	0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x32, 0x0a, 0x0a, 0x55, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x0c, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x3a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x11, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x0b, 0x45,
	0x64, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x45, 0x64, 0x69,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x61, 0x6b, 0x73, 0x4b, 0x61, 0x7a, 0x61, 0x6e, 0x74, 0x73, 0x65, 0x76, 0x2f,
	0x43, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_posts_proto_rawDescOnce sync.Once
	file_posts_proto_rawDescData = file_posts_proto_rawDesc
)

func file_posts_proto_rawDescGZIP() []byte {
	file_posts_proto_rawDescOnce.Do(func() {
		file_posts_proto_rawDescData = protoimpl.X.CompressGZIP(file_posts_proto_rawDescData)
	})
	return file_posts_proto_rawDescData
}

var file_posts_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_posts_proto_goTypes = []interface{}{
	(*EditCommentReq)(nil),        // 0: EditCommentReq
	(*DeleteCommentReq)(nil),      // 1: DeleteCommentReq
	(*LikePostReq)(nil),           // 2: LikePostReq
	(*LeaveCommentReq)(nil),       // 3: LeaveCommentReq
	(*CreatePostReq)(nil),         // 4: CreatePostReq
	(*DeletePostReq)(nil),         // 5: DeletePostReq
	(*EditPostReq)(nil),           // 6: EditPostReq
	(*GetUserPostsReq)(nil),       // 7: GetUserPostsReq
	(*GetUserPostsRes)(nil),       // 8: GetUserPostsRes
	(*Post)(nil),                  // 9: Post
	(*Comment)(nil),               // 10: Comment
	(*CommentValue)(nil),          // 11: CommentValue
	(*Like)(nil),                  // 12: Like
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 14: google.protobuf.Empty
}
var file_posts_proto_depIdxs = []int32{
	11, // 0: EditCommentReq.Value:type_name -> CommentValue
	11, // 1: LeaveCommentReq.Value:type_name -> CommentValue
	9,  // 2: GetUserPostsRes.Posts:type_name -> Post
	12, // 3: Post.Likes:type_name -> Like
	10, // 4: Post.Comments:type_name -> Comment
	13, // 5: Post.CreatedAt:type_name -> google.protobuf.Timestamp
	11, // 6: Comment.Value:type_name -> CommentValue
	13, // 7: Comment.CreatedAt:type_name -> google.protobuf.Timestamp
	4,  // 8: Posts.CreatePost:input_type -> CreatePostReq
	5,  // 9: Posts.DeletePost:input_type -> DeletePostReq
	6,  // 10: Posts.EditPost:input_type -> EditPostReq
	7,  // 11: Posts.GetUserPosts:input_type -> GetUserPostsReq
	2,  // 12: Posts.LikePost:input_type -> LikePostReq
	2,  // 13: Posts.UnlikePost:input_type -> LikePostReq
	3,  // 14: Posts.LeaveComment:input_type -> LeaveCommentReq
	1,  // 15: Posts.DeleteComment:input_type -> DeleteCommentReq
	0,  // 16: Posts.EditComment:input_type -> EditCommentReq
	14, // 17: Posts.CreatePost:output_type -> google.protobuf.Empty
	14, // 18: Posts.DeletePost:output_type -> google.protobuf.Empty
	14, // 19: Posts.EditPost:output_type -> google.protobuf.Empty
	8,  // 20: Posts.GetUserPosts:output_type -> GetUserPostsRes
	14, // 21: Posts.LikePost:output_type -> google.protobuf.Empty
	14, // 22: Posts.UnlikePost:output_type -> google.protobuf.Empty
	14, // 23: Posts.LeaveComment:output_type -> google.protobuf.Empty
	14, // 24: Posts.DeleteComment:output_type -> google.protobuf.Empty
	14, // 25: Posts.EditComment:output_type -> google.protobuf.Empty
	17, // [17:26] is the sub-list for method output_type
	8,  // [8:17] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_posts_proto_init() }
func file_posts_proto_init() {
	if File_posts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_posts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditCommentReq); i {
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
		file_posts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentReq); i {
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
		file_posts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikePostReq); i {
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
		file_posts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveCommentReq); i {
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
		file_posts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostReq); i {
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
		file_posts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostReq); i {
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
		file_posts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditPostReq); i {
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
		file_posts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPostsReq); i {
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
		file_posts_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPostsRes); i {
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
		file_posts_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_posts_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_posts_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentValue); i {
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
		file_posts_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Like); i {
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
			RawDescriptor: file_posts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_posts_proto_goTypes,
		DependencyIndexes: file_posts_proto_depIdxs,
		MessageInfos:      file_posts_proto_msgTypes,
	}.Build()
	File_posts_proto = out.File
	file_posts_proto_rawDesc = nil
	file_posts_proto_goTypes = nil
	file_posts_proto_depIdxs = nil
}
