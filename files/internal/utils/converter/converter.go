package converter

import (
	"github.com/MaksKazantsev/Chatter/files/internal/models"
	pkg "github.com/MaksKazantsev/Chatter/files/pkg/grpc"
)

type Converter interface {
	ToService
}

type ToService interface {
	UploadToStorageToService(req *pkg.UploadReq, uuid string) models.UploadToStorageReq
	UpdateAvatarToService(req *pkg.UpdateAvatarReq, uuid string) models.UploadToStorageReq
}

type converter struct {
}

func (c converter) UpdateAvatarToService(req *pkg.UpdateAvatarReq, uuid string) models.UploadToStorageReq {
	return models.UploadToStorageReq{UserID: uuid, File: req.Photo, FileID: req.PhotoID}
}

func (c converter) UploadToStorageToService(req *pkg.UploadReq, uuid string) models.UploadToStorageReq {
	return models.UploadToStorageReq{UserID: uuid, File: req.Photo, FileID: req.PhotoID}
}

func NewConverter() Converter {
	return &converter{}
}
