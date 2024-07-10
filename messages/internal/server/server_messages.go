package server

import (
	"context"
	"github.com/MaksKazantsev/Chatter/messages/internal/log"
	"github.com/MaksKazantsev/Chatter/messages/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CreateMessage(ctx context.Context, req *pkg.CreateMessageReq) (*emptypb.Empty, error) {
	s.log.Debug("Message microservice successfully received request")

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

	if err = s.validator.ValidateCreateMessageReq(req); err != nil {
		return nil, err
	}

	serviceReq, receiverOffline := s.converter.CreateMessageToService(req)
	if err = s.service.Messages.CreateMessage(log.WithLogger(ctx, s.log), serviceReq, receiverOffline); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *server) DeleteMessage(ctx context.Context, req *pkg.DeleteMessageReq) (*emptypb.Empty, error) {
	s.log.Debug("Message microservice successfully received request")

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

	if err = s.service.Messages.DeleteMessage(log.WithLogger(ctx, s.log), req.MessageID, id); err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}

func (s *server) GetHistory(ctx context.Context, req *pkg.GetHistoryReq) (*pkg.GetHistoryRes, error) {
	s.log.Debug("Message microservice successfully received request")

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

	res, err := s.service.Messages.GetHistory(log.WithLogger(ctx, s.log), s.converter.GetHistoryToService(req), id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return s.converter.MessageToPb(res), nil
}
