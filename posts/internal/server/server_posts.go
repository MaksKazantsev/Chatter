package server

import (
	"context"
	"github.com/MaksKazantsev/Chatter/posts/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/posts/pkg/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CreatePost(ctx context.Context, req *pkg.CreatePostReq) (*emptypb.Empty, error) {
	id, err := s.userCl.UserClient.ParseToken(ctx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}

func (s *server) DeletePost(ctx context.Context, req *pkg.DeletePostReq) (*emptypb.Empty, error) {
	id, err := s.userCl.UserClient.ParseToken(ctx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}

func (s *server) EditPost(ctx context.Context, req *pkg.EditPostReq) (*emptypb.Empty, error) {
	id, err := s.userCl.UserClient.ParseToken(ctx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}

func (s *server) GetUserPosts(ctx context.Context, req *pkg.GetUserPostsReq) (*pkg.GetUserPostsRes, error) {
	id, err := s.userCl.UserClient.ParseToken(ctx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}
