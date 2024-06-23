package clients

import (
	"context"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
)

func (u userCl) SuggestFriendShip(ctx context.Context, req models.FriendShipReq) error {
	_, err := u.cl.SuggestFriendShip(ctx, u.c.SuggestFsToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u userCl) RefuseFriendShip(ctx context.Context, req models.RefuseFriendShipReq) error {
	_, err := u.cl.RefuseFriendShip(ctx, u.c.RefuseFsToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}
