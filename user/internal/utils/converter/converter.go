package converter

import (
	"fmt"
	"github.com/MaksKazantsev/Chatter/user/internal/models"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Converter interface {
	ToPb
	ToService
}

type ToPb interface {
	RegResToPb(req models.RegRes) *pkg.RegisterRes
	LoginResToPb(access, refresh string) *pkg.LoginRes
	VerifyCodeResToPb(access, refresh string) *pkg.VerifyRes
	UpdateTokensResToPb(access, refresh string) *pkg.UpdateTokenRes
	GetFsToPb([]models.FsReq) *pkg.GetFsRes
	GetFriendsToPb([]models.Friend) *pkg.GetFriendsRes
	GetProfileToPb(profile models.GetUserProfile) *pkg.GetProfileRes
}

type ToService interface {
	RegReqToService(req *pkg.RegisterReq) models.RegReq
	LoginReqToService(req *pkg.LoginReq) models.LogReq
	VerifyCodeReqToService(req *pkg.VerifyReq) models.VerifyCodeReq
	RecoveryReqToService(req *pkg.RecoveryReq) models.Credentials
	EditProfileReqToService(req *pkg.EditProfileReq, id string) models.UserProfile
}

func NewConverter() Converter {
	return &converter{}
}

type converter struct {
}

// To service

func (c converter) RegReqToService(req *pkg.RegisterReq) models.RegReq {
	return models.RegReq{Password: req.Password, Email: req.Email, Username: req.Username}
}

func (c converter) LoginReqToService(req *pkg.LoginReq) models.LogReq {
	return models.LogReq{Email: req.Email, Password: req.Password}
}

func (c converter) EditProfileReqToService(req *pkg.EditProfileReq, id string) models.UserProfile {
	return models.UserProfile{
		UUID:     id,
		Avatar:   req.Avatar,
		Birthday: req.Birthday.AsTime(),
		Bio:      req.Bio,
		Username: req.Username,
	}
}

func (c converter) RecoveryReqToService(req *pkg.RecoveryReq) models.Credentials {
	return models.Credentials{
		Password: req.Password,
		Email:    req.Email,
	}
}

func (c converter) VerifyCodeReqToService(req *pkg.VerifyReq) models.VerifyCodeReq {
	return models.VerifyCodeReq{Code: req.Code, Email: req.Email, Type: req.Type}
}

// To protobuf

func (c converter) GetProfileToPb(profile models.GetUserProfile) *pkg.GetProfileRes {
	return &pkg.GetProfileRes{
		Avatar:     profile.Avatar,
		Birthday:   timestamppb.New(profile.Birthday),
		Bio:        profile.Bio,
		Username:   profile.Username,
		LastOnline: timestamppb.New(profile.LastOnline),
	}
}

func (c converter) GetFriendsToPb(friends []models.Friend) *pkg.GetFriendsRes {
	var res []*pkg.Friend

	for _, v := range friends {
		res = append(res, &pkg.Friend{
			FriendID: v.FriendID,
			Avatar:   v.Avatar,
			Username: v.Username,
		})
	}

	return &pkg.GetFriendsRes{
		Friends: res,
	}
}

func (c converter) GetFsToPb(reqs []models.FsReq) *pkg.GetFsRes {
	var res []*pkg.FsReq

	fmt.Println(reqs)

	for i := 0; i < len(reqs); i++ {
		req := &pkg.FsReq{
			ReqId:    reqs[i].ReqID,
			Avatar:   reqs[i].Avatar,
			Username: reqs[i].Username,
		}
		res = append(res, req)
	}
	return &pkg.GetFsRes{
		FsReqs: res,
	}
}

func (c converter) UpdateTokensResToPb(access, refresh string) *pkg.UpdateTokenRes {
	return &pkg.UpdateTokenRes{AToken: access, RToken: refresh}
}

func (c converter) VerifyCodeResToPb(access, refresh string) *pkg.VerifyRes {
	return &pkg.VerifyRes{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func (c converter) RegResToPb(req models.RegRes) *pkg.RegisterRes {
	return &pkg.RegisterRes{UUID: req.UUID, AccessToken: req.AccessToken, RefreshToken: req.RefreshToken}
}

func (c converter) LoginResToPb(access, refresh string) *pkg.LoginRes {
	return &pkg.LoginRes{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}
