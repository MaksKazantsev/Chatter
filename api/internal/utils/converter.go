package utils

import (
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	filesPkg "github.com/MaksKazantsev/Chatter/files/pkg/grpc"
	messagesPkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	userPkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Converter interface {
	ToPb
	ToService
}

func NewConverter() Converter {
	return &converter{}
}

type converter struct{}

type ToPb interface {
	RegReqToPb(req models.SignupReq) *userPkg.RegisterReq
	LogReqToPb(req models.LoginReq) *userPkg.LoginReq
	SendCodeReqToPb(req string) *userPkg.SendReq
	VerifyCodeReqToPb(req models.VerifyCodeReq) *userPkg.VerifyReq
	RecoveryReqToPb(req models.RecoveryReq) *userPkg.RecoveryReq
	UpdateTokensToPb(req string) *userPkg.UpdateTokenReq
	ParseTokenToPb(req string) *userPkg.ParseTokenReq
	DeleteMsgToPb(messageID, token string) *messagesPkg.DeleteMessageReq
	CreateMsgToPb(req *models.Message, receiverOffline bool) *messagesPkg.CreateMessageReq
	GetHistoryToPb(req models.GetHistoryReq) *messagesPkg.GetHistoryReq
	UploadToPb(req models.UploadReq) *filesPkg.UploadReq
	EditProfileToPb(req models.UserProfileReq) *userPkg.EditProfileReq
}

type ToService interface {
	MessageToService(messages []*messagesPkg.Message) []models.Message
	GetFriendsToService(friends []*userPkg.Friend) []models.Friend
	GetFsToService(reqs []*userPkg.FsReq) []models.FsReq
	GetProfileToService(req *userPkg.GetProfileRes) models.UserProfile
}

// User Microservice

func (c converter) GetProfileToService(req *userPkg.GetProfileRes) models.UserProfile {
	return models.UserProfile{
		Avatar:     req.Avatar,
		Bio:        req.Bio,
		Birthday:   req.Birthday.AsTime(),
		LastOnline: req.LastOnline.AsTime(),
		Username:   req.Username,
	}
}

func (c converter) EditProfileToPb(req models.UserProfileReq) *userPkg.EditProfileReq {
	return &userPkg.EditProfileReq{Token: req.Token, Avatar: req.Avatar, Bio: req.Bio, Birthday: timestamppb.New(*req.Birthday)}
}

func (c converter) GetFsToService(reqs []*userPkg.FsReq) []models.FsReq {
	var res []models.FsReq

	for _, v := range reqs {
		req := models.FsReq{
			ReqID:    v.ReqId,
			Username: v.Username,
			Avatar:   v.Avatar,
		}
		res = append(res, req)
	}
	return res
}

func (c converter) GetFriendsToService(friends []*userPkg.Friend) []models.Friend {
	var res []models.Friend

	for _, v := range friends {
		fr := models.Friend{
			FriendID: v.FriendID,
			Username: v.Username,
			Avatar:   v.Avatar,
		}
		res = append(res, fr)
	}
	return res
}

func (c converter) RecoveryReqToPb(req models.RecoveryReq) *userPkg.RecoveryReq {
	return &userPkg.RecoveryReq{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (c converter) VerifyCodeReqToPb(req models.VerifyCodeReq) *userPkg.VerifyReq {
	return &userPkg.VerifyReq{
		Code:  req.Code,
		Email: req.Email,
		Type:  req.Type,
	}
}

func (c converter) SendCodeReqToPb(email string) *userPkg.SendReq {
	return &userPkg.SendReq{
		Email: email,
	}
}

func (c converter) LogReqToPb(req models.LoginReq) *userPkg.LoginReq {
	return &userPkg.LoginReq{Email: req.Email, Password: req.Password}
}
func (c converter) RegReqToPb(req models.SignupReq) *userPkg.RegisterReq {
	return &userPkg.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
}

func (c converter) UpdateTokensToPb(req string) *userPkg.UpdateTokenReq {
	return &userPkg.UpdateTokenReq{RToken: req}
}
func (c converter) ParseTokenToPb(req string) *userPkg.ParseTokenReq {
	return &userPkg.ParseTokenReq{Token: req}
}
func (c converter) UpdateTokens(req string) *userPkg.UpdateTokenReq {
	return &userPkg.UpdateTokenReq{
		RToken: req,
	}
}

// Messages Microservice

func (c converter) MessageToService(messages []*messagesPkg.Message) []models.Message {
	var msgs []models.Message
	for i := 0; i < len(messages); i++ {
		msg := models.Message{
			MessageID:  messages[i].MessageID,
			SenderID:   messages[i].SenderID,
			ReceiverID: messages[i].ReceiverID,
			SentAt:     messages[i].SentAt.AsTime(),
			ChatID:     messages[i].ChatID,
			Value:      messages[i].Value,
		}
		msgs = append(msgs, msg)
	}
	return msgs
}

func (c converter) GetHistoryToPb(req models.GetHistoryReq) *messagesPkg.GetHistoryReq {
	return &messagesPkg.GetHistoryReq{Token: req.Token, ChatID: req.ChatID}
}

func (c converter) CreateMsgToPb(req *models.Message, receiverOffline bool) *messagesPkg.CreateMessageReq {
	return &messagesPkg.CreateMessageReq{Token: req.Token, SenderID: req.SenderID, ReceiverID: req.ReceiverID, ChatID: req.ChatID, Value: req.Value, ReceiverOffline: receiverOffline}
}

func (c converter) DeleteMsgToPb(messageID, token string) *messagesPkg.DeleteMessageReq {
	return &messagesPkg.DeleteMessageReq{MessageID: messageID, Token: token}
}

// Files microservice

func (c converter) UploadToPb(req models.UploadReq) *filesPkg.UploadReq {
	return &filesPkg.UploadReq{
		Token:   req.Token,
		PhotoID: req.PhotoID,
		Photo:   req.Photo,
	}
}
