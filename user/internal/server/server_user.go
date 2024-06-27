package server

import (
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) EditProfile(ctx context.Context, req *pkg.EditProfileReq) (*emptypb.Empty, error) {
	s.log.Info("Microservice user successfully received request")
	loggerCtx := log.WithLogger(ctx, s.log)
	id, err := s.service.Auth.ParseToken(loggerCtx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	if err = s.service.User.EditProfile(loggerCtx, s.converter.EditProfileReqToService(req, id)); err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}

func (s *server) SuggestFs(ctx context.Context, req *pkg.SuggestFsReq) (*emptypb.Empty, error) {
	s.log.Info("Microservice user successfully received request")
	loggerCtx := log.WithLogger(ctx, s.log)
	id, err := s.service.Auth.ParseToken(loggerCtx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	if err = s.service.User.SuggestFs(loggerCtx, id, req.ReceiverID); err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}

func (s *server) RefuseFs(ctx context.Context, req *pkg.RefuseFsReq) (*emptypb.Empty, error) {
	s.log.Info("Microservice user successfully received request")
	loggerCtx := log.WithLogger(ctx, s.log)
	id, err := s.service.Auth.ParseToken(loggerCtx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}
	if err = s.service.User.RefuseFs(loggerCtx, req.SenderID, id); err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}

func (s *server) GetFs(ctx context.Context, req *pkg.GetFsReq) (*pkg.GetFsRes, error) {
	s.log.Info("Microservice user successfully received request")
	loggerCtx := log.WithLogger(ctx, s.log)
	id, err := s.service.Auth.ParseToken(loggerCtx, req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	res, err := s.service.User.GetFs(loggerCtx, id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return s.converter.GetFsToPb(res), nil
}
