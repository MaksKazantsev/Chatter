package server

import (
	"context"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) SuggestFriendShip(ctx context.Context, req *pkg.SuggestFriendShipReq) (*emptypb.Empty, error) {
	if err := s.service.Friendship.SuggestFriendship(log.WithLogger(ctx, s.log), s.converter.SuggestFriendShipToService(req)); err != nil {
		return nil, utils.HandleError(err)
	}
	return nil, nil
}

func (s *server) RefuseFriendShip(ctx context.Context, req *pkg.RefuseFriendShipReq) (*emptypb.Empty, error) {
	if err := s.service.Friendship.RefuseFriendship(log.WithLogger(ctx, s.log), s.converter.RefuseFriendShipToService(req)); err != nil {
		return nil, utils.HandleError(err)
	}
	return nil, nil
}
