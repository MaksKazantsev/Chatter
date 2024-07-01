package clients

import (
	"context"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
)

func (u userCl) GetFriends(ctx context.Context, token string) ([]models.Friend, error) {
	res, err := u.cl.GetFriends(ctx, &pkg.GetFsAction{Token: token})
	if err != nil {
		return nil, utils.GRPCErrorToError(err)
	}
	return u.c.GetFriendsToService(res.Friends), nil
}

func (u userCl) GetFs(ctx context.Context, token string) ([]models.FsReq, error) {
	res, err := u.cl.GetFs(ctx, &pkg.GetFsAction{Token: token})
	if err != nil {
		return nil, utils.GRPCErrorToError(err)
	}

	return u.c.GetFsToService(res.FsReqs), nil
}

func (u userCl) DeleteFriend(ctx context.Context, token, targetID string) error {
	_, err := u.cl.DeleteFriend(ctx, &pkg.FsAction{Token: token, TargetID: targetID})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}
func (u userCl) SuggestFs(ctx context.Context, token string, targetID string) error {
	_, err := u.cl.SuggestFs(ctx, &pkg.SuggestFsReq{Token: token, ReceiverID: targetID})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u userCl) RefuseFs(ctx context.Context, token string, targetID string) error {
	_, err := u.cl.RefuseFs(ctx, &pkg.RefuseFsReq{Token: token, SenderID: targetID})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u userCl) AcceptFs(ctx context.Context, token string, targetID string) error {
	_, err := u.cl.AcceptFs(ctx, &pkg.FsAction{Token: token, TargetID: targetID})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u userCl) EditProfile(ctx context.Context, req models.UserProfileReq) error {
	_, err := u.cl.EditProfile(ctx, u.c.EditProfileToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u userCl) GetProfile(ctx context.Context, token, targetID string) (models.UserProfile, error) {
	profile, err := u.cl.GetProfile(ctx, &pkg.SimpleReq{Token: token, TargetID: targetID})
	if err != nil {
		return models.UserProfile{}, utils.GRPCErrorToError(err)
	}
	return u.c.GetProfileToService(profile), nil
}

func (u userCl) EditAvatar(ctx context.Context, token, avatar string) error {
	_, err := u.cl.EditAvatar(ctx, &pkg.EditAvatarReq{Token: token, Avatar: avatar})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u userCl) DeleteAvatar(ctx context.Context, token string) error {
	_, err := u.cl.DeleteAvatar(ctx, &pkg.DeleteAvatarReq{Token: token})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}

	return nil
}
