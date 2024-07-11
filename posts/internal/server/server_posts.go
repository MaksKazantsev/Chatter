package server

import (
	"context"
	"github.com/MaksKazantsev/Chatter/posts/internal/log"
	"github.com/MaksKazantsev/Chatter/posts/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/posts/pkg/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CreatePost(ctx context.Context, req *pkg.CreatePostReq) (*emptypb.Empty, error) {
	s.log.Debug("Posts microservice successfully received request")

	id, err := s.cache.Get(ctx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}
	if id == "" {
		id, err = s.userCl.UserClient.ParseToken(ctx, req.Token)
		if err != nil {
			return nil, err
		}
		err = s.cache.Save(ctx, req.Token, id)
		if err != nil {
			return nil, utils.HandleError(err)
		}
	}

	if err = s.srvc.CreatePost(log.WithLogger(ctx, s.log), s.c.CreatePostToService(req, id)); err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}

func (s *server) DeletePost(ctx context.Context, req *pkg.DeletePostReq) (*emptypb.Empty, error) {
	s.log.Debug("Posts microservice successfully received request")

	id, err := s.cache.Get(ctx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}
	if id == "" {
		id, err = s.userCl.UserClient.ParseToken(ctx, req.Token)
		if err != nil {
			return nil, err
		}
		err = s.cache.Save(ctx, req.Token, id)
		if err != nil {
			return nil, utils.HandleError(err)
		}
	}

	if err = s.srvc.DeletePost(log.WithLogger(ctx, s.log), id, req.PostID); err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}

func (s *server) EditPost(ctx context.Context, req *pkg.EditPostReq) (*emptypb.Empty, error) {
	s.log.Debug("Posts microservice successfully received request")

	id, err := s.cache.Get(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if id == "" {
		id, err = s.userCl.UserClient.ParseToken(ctx, req.Token)
		if err != nil {
			return nil, err
		}
		_ = s.cache.Save(ctx, req.Token, id)
	}

	if err = s.srvc.EditPost(log.WithLogger(ctx, s.log), s.c.EditPostToService(req, id)); err != nil {
		return nil, err
	}

	return nil, nil

}

func (s *server) LikePost(ctx context.Context, req *pkg.LikePostReq) (*emptypb.Empty, error) {
	s.log.Debug("Posts microservice successfully received request")

	id, err := s.cache.Get(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if id == "" {
		id, err = s.userCl.UserClient.ParseToken(ctx, req.Token)
		if err != nil {
			return nil, err
		}
		_ = s.cache.Save(ctx, req.Token, id)
	}

	if err = s.srvc.LikePost(log.WithLogger(ctx, s.log), id, req.PostID); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *server) LeaveComment(ctx context.Context, req *pkg.LeaveCommentReq) (*emptypb.Empty, error) {
	s.log.Debug("Posts microservice successfully received request")

	id, err := s.cache.Get(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if id == "" {
		id, err = s.userCl.UserClient.ParseToken(ctx, req.Token)
		if err != nil {
			return nil, err
		}
		_ = s.cache.Save(ctx, req.Token, id)
	}

	if err = s.srvc.LeaveComment(log.WithLogger(ctx, s.log), s.c.LeaveCommentToService(req, id)); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *server) GetUserPosts(ctx context.Context, req *pkg.GetUserPostsReq) (*pkg.GetUserPostsRes, error) {
	s.log.Debug("Posts microservice successfully received request")

	id, err := s.cache.Get(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if id == "" {
		id, err = s.userCl.UserClient.ParseToken(ctx, req.Token)
		if err != nil {
			return nil, err
		}
		err = s.cache.Save(ctx, req.Token, id)
		if err != nil {
			return nil, err
		}
	}

	res, err := s.srvc.GetUserPosts(log.WithLogger(ctx, s.log), req.UserID, id)
	if err != nil {
		return nil, err
	}

	return s.c.GetUserPostsToPb(res), nil
}

func (s *server) DeleteComment(ctx context.Context, req *pkg.DeleteCommentReq) (*emptypb.Empty, error) {
	s.log.Debug("Posts microservice successfully received request")

	id, err := s.cache.Get(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if id == "" {
		id, err = s.userCl.UserClient.ParseToken(ctx, req.Token)
		if err != nil {
			return nil, err
		}
		err = s.cache.Save(ctx, req.Token, id)
		if err != nil {

			return nil, err
		}
	}

	if err = s.srvc.DeleteComment(log.WithLogger(ctx, s.log), req.CommentID, id); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *server) UnlikePost(ctx context.Context, req *pkg.LikePostReq) (*emptypb.Empty, error) {
	s.log.Debug("Posts microservice successfully received request")

	id, err := s.cache.Get(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if id == "" {
		id, err = s.userCl.UserClient.ParseToken(ctx, req.Token)
		if err != nil {
			return nil, err
		}
		_ = s.cache.Save(ctx, req.Token, id)
	}

	if err = s.srvc.UnlikePost(log.WithLogger(ctx, s.log), id, req.PostID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *server) EditComment(ctx context.Context, req *pkg.EditCommentReq) (*emptypb.Empty, error) {
	s.log.Debug("Posts microservice successfully received request")

	id, err := s.cache.Get(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if id == "" {
		id, err = s.userCl.UserClient.ParseToken(ctx, req.Token)
		if err != nil {
			return nil, err
		}
		_ = s.cache.Save(ctx, req.Token, id)
	}

	if err = s.srvc.EditComment(log.WithLogger(ctx, s.log), s.c.EditCommentToService(req, id)); err != nil {
		return nil, err
	}

	return nil, nil
}
