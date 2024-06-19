// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: realworld/v1/realworld.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RealWorld_Login_FullMethodName              = "/realworld.v1.RealWorld/Login"
	RealWorld_Register_FullMethodName           = "/realworld.v1.RealWorld/Register"
	RealWorld_GetCurrentUser_FullMethodName     = "/realworld.v1.RealWorld/GetCurrentUser"
	RealWorld_UpdateUser_FullMethodName         = "/realworld.v1.RealWorld/UpdateUser"
	RealWorld_GetProfile_FullMethodName         = "/realworld.v1.RealWorld/GetProfile"
	RealWorld_FollowUser_FullMethodName         = "/realworld.v1.RealWorld/FollowUser"
	RealWorld_UnFollowUser_FullMethodName       = "/realworld.v1.RealWorld/UnFollowUser"
	RealWorld_ListArticles_FullMethodName       = "/realworld.v1.RealWorld/ListArticles"
	RealWorld_FeedArticles_FullMethodName       = "/realworld.v1.RealWorld/FeedArticles"
	RealWorld_GetArticles_FullMethodName        = "/realworld.v1.RealWorld/GetArticles"
	RealWorld_CreateArticles_FullMethodName     = "/realworld.v1.RealWorld/CreateArticles"
	RealWorld_UpdateArticles_FullMethodName     = "/realworld.v1.RealWorld/UpdateArticles"
	RealWorld_DeleteArticles_FullMethodName     = "/realworld.v1.RealWorld/DeleteArticles"
	RealWorld_AddComment_FullMethodName         = "/realworld.v1.RealWorld/AddComment"
	RealWorld_GetComment_FullMethodName         = "/realworld.v1.RealWorld/GetComment"
	RealWorld_DeleteComment_FullMethodName      = "/realworld.v1.RealWorld/DeleteComment"
	RealWorld_FavouriteArticle_FullMethodName   = "/realworld.v1.RealWorld/FavouriteArticle"
	RealWorld_UnFavouriteArticle_FullMethodName = "/realworld.v1.RealWorld/UnFavouriteArticle"
	RealWorld_GetTags_FullMethodName            = "/realworld.v1.RealWorld/GetTags"
)

// RealWorldClient is the client API for RealWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealWorldClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	CreateArticles(ctx context.Context, in *CreateArticlesRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UpdateArticles(ctx context.Context, in *UpdateArticlesRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	DeleteArticles(ctx context.Context, in *DeleteArticlesRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error)
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error)
	FavouriteArticle(ctx context.Context, in *FavouriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UnFavouriteArticle(ctx context.Context, in *UnFavouriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*TagListReply, error)
}

type realWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewRealWorldClient(cc grpc.ClientConnInterface) RealWorldClient {
	return &realWorldClient{cc}
}

func (c *realWorldClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, RealWorld_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, RealWorld_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, RealWorld_GetCurrentUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, RealWorld_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, RealWorld_GetProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, RealWorld_FollowUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, RealWorld_UnFollowUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, RealWorld_ListArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, RealWorld_FeedArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_GetArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) CreateArticles(ctx context.Context, in *CreateArticlesRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_CreateArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UpdateArticles(ctx context.Context, in *UpdateArticlesRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_UpdateArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) DeleteArticles(ctx context.Context, in *DeleteArticlesRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_DeleteArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error) {
	out := new(SingleCommentReply)
	err := c.cc.Invoke(ctx, RealWorld_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error) {
	out := new(MultipleCommentsReply)
	err := c.cc.Invoke(ctx, RealWorld_GetComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error) {
	out := new(SingleCommentReply)
	err := c.cc.Invoke(ctx, RealWorld_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) FavouriteArticle(ctx context.Context, in *FavouriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_FavouriteArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UnFavouriteArticle(ctx context.Context, in *UnFavouriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, RealWorld_UnFavouriteArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*TagListReply, error) {
	out := new(TagListReply)
	err := c.cc.Invoke(ctx, RealWorld_GetTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealWorldServer is the server API for RealWorld service.
// All implementations must embed UnimplementedRealWorldServer
// for forward compatibility
type RealWorldServer interface {
	Login(context.Context, *LoginRequest) (*UserReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error)
	FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error)
	UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error)
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error)
	GetArticles(context.Context, *GetArticlesRequest) (*SingleArticleReply, error)
	CreateArticles(context.Context, *CreateArticlesRequest) (*SingleArticleReply, error)
	UpdateArticles(context.Context, *UpdateArticlesRequest) (*SingleArticleReply, error)
	DeleteArticles(context.Context, *DeleteArticlesRequest) (*SingleArticleReply, error)
	AddComment(context.Context, *AddCommentRequest) (*SingleCommentReply, error)
	GetComment(context.Context, *GetCommentRequest) (*MultipleCommentsReply, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*SingleCommentReply, error)
	FavouriteArticle(context.Context, *FavouriteArticleRequest) (*SingleArticleReply, error)
	UnFavouriteArticle(context.Context, *UnFavouriteArticleRequest) (*SingleArticleReply, error)
	GetTags(context.Context, *GetTagsRequest) (*TagListReply, error)
	mustEmbedUnimplementedRealWorldServer()
}

// UnimplementedRealWorldServer must be embedded to have forward compatible implementations.
type UnimplementedRealWorldServer struct {
}

func (UnimplementedRealWorldServer) Login(context.Context, *LoginRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRealWorldServer) Register(context.Context, *RegisterRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRealWorldServer) GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedRealWorldServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedRealWorldServer) GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedRealWorldServer) FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedRealWorldServer) UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowUser not implemented")
}
func (UnimplementedRealWorldServer) ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedRealWorldServer) FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}
func (UnimplementedRealWorldServer) GetArticles(context.Context, *GetArticlesRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedRealWorldServer) CreateArticles(context.Context, *CreateArticlesRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticles not implemented")
}
func (UnimplementedRealWorldServer) UpdateArticles(context.Context, *UpdateArticlesRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticles not implemented")
}
func (UnimplementedRealWorldServer) DeleteArticles(context.Context, *DeleteArticlesRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticles not implemented")
}
func (UnimplementedRealWorldServer) AddComment(context.Context, *AddCommentRequest) (*SingleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedRealWorldServer) GetComment(context.Context, *GetCommentRequest) (*MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedRealWorldServer) DeleteComment(context.Context, *DeleteCommentRequest) (*SingleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedRealWorldServer) FavouriteArticle(context.Context, *FavouriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavouriteArticle not implemented")
}
func (UnimplementedRealWorldServer) UnFavouriteArticle(context.Context, *UnFavouriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavouriteArticle not implemented")
}
func (UnimplementedRealWorldServer) GetTags(context.Context, *GetTagsRequest) (*TagListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedRealWorldServer) mustEmbedUnimplementedRealWorldServer() {}

// UnsafeRealWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealWorldServer will
// result in compilation errors.
type UnsafeRealWorldServer interface {
	mustEmbedUnimplementedRealWorldServer()
}

func RegisterRealWorldServer(s grpc.ServiceRegistrar, srv RealWorldServer) {
	s.RegisterService(&RealWorld_ServiceDesc, srv)
}

func _RealWorld_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetCurrentUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_FollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UnFollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UnFollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_UnFollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UnFollowUser(ctx, req.(*UnFollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_ListArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_FeedArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).FeedArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_FeedArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).FeedArticles(ctx, req.(*FeedArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetArticles(ctx, req.(*GetArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_CreateArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).CreateArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_CreateArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).CreateArticles(ctx, req.(*CreateArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UpdateArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UpdateArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_UpdateArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UpdateArticles(ctx, req.(*UpdateArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_DeleteArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).DeleteArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_DeleteArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).DeleteArticles(ctx, req.(*DeleteArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_FavouriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavouriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).FavouriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_FavouriteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).FavouriteArticle(ctx, req.(*FavouriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UnFavouriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFavouriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UnFavouriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_UnFavouriteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UnFavouriteArticle(ctx, req.(*UnFavouriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealWorld_GetTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetTags(ctx, req.(*GetTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RealWorld_ServiceDesc is the grpc.ServiceDesc for RealWorld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealWorld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realworld.v1.RealWorld",
	HandlerType: (*RealWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _RealWorld_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _RealWorld_Register_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _RealWorld_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _RealWorld_UpdateUser_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _RealWorld_GetProfile_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _RealWorld_FollowUser_Handler,
		},
		{
			MethodName: "UnFollowUser",
			Handler:    _RealWorld_UnFollowUser_Handler,
		},
		{
			MethodName: "ListArticles",
			Handler:    _RealWorld_ListArticles_Handler,
		},
		{
			MethodName: "FeedArticles",
			Handler:    _RealWorld_FeedArticles_Handler,
		},
		{
			MethodName: "GetArticles",
			Handler:    _RealWorld_GetArticles_Handler,
		},
		{
			MethodName: "CreateArticles",
			Handler:    _RealWorld_CreateArticles_Handler,
		},
		{
			MethodName: "UpdateArticles",
			Handler:    _RealWorld_UpdateArticles_Handler,
		},
		{
			MethodName: "DeleteArticles",
			Handler:    _RealWorld_DeleteArticles_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _RealWorld_AddComment_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _RealWorld_GetComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _RealWorld_DeleteComment_Handler,
		},
		{
			MethodName: "FavouriteArticle",
			Handler:    _RealWorld_FavouriteArticle_Handler,
		},
		{
			MethodName: "UnFavouriteArticle",
			Handler:    _RealWorld_UnFavouriteArticle_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _RealWorld_GetTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "realworld/v1/realworld.proto",
}
