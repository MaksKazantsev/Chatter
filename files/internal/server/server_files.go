package server

import (
	"github.com/MaksKazantsev/Chatter/files/internal/log"
	"github.com/MaksKazantsev/Chatter/files/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/files/pkg/grpc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) UploadToStorage(ctx context.Context, req *pkg.UploadReq) (*pkg.UploadRes, error) {
	uuid, err := s.userCl.UserClient.ParseToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	fileLink, err := s.service.UploadToStorage(log.WithLogger(ctx, s.l), s.c.UploadToStorageToService(req, uuid))
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &pkg.UploadRes{PhotoLink: fileLink}, nil
}

func (s *server) UpdateAvatar(ctx context.Context, req *pkg.UpdateAvatarReq) (*emptypb.Empty, error) {
	uuid, err := s.userCl.UserClient.ParseToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	err = s.service.UpdateAvatar(log.WithLogger(ctx, s.l), s.c.UpdateAvatarToService(req, uuid))
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}
