package server

import (
	"context"
	"github.com/MaksKazantsev/Chatter/messages/internal/log"
	"github.com/MaksKazantsev/Chatter/messages/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CreateMessage(ctx context.Context, req *pkg.CreateMessageReq) (*emptypb.Empty, error) {
	_, err := s.userCl.UserClient.ParseToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	if err = s.validator.ValidateCreateMessageReq(req); err != nil {
		return nil, err
	}

	serviceReq, receiverOffline := s.converter.CreateMessageToService(req)
	if err = s.service.Messages.CreateMessage(log.WithLogger(ctx, s.log), serviceReq, receiverOffline); err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}
