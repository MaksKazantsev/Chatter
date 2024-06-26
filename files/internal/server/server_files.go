package server

import (
	"github.com/MaksKazantsev/Chatter/files/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/files/pkg/grpc"
	"golang.org/x/net/context"
)

func (s *server) UploadToStorage(ctx context.Context, req *pkg.UploadReq) (*pkg.UploadRes, error) {
	uuid, err := s.userCl.UserClient.ParseToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	fileLink, err := s.service.UploadToStorage(ctx, s.c.UploadToStorageToService(req, uuid))
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &pkg.UploadRes{PhotoLink: fileLink}, nil
}
