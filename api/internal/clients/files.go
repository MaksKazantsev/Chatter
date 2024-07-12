package clients

import (
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/files/pkg/grpc"
	"golang.org/x/net/context"
)

type FilesClient interface {
	Upload(ctx context.Context, fileID, token string, val []byte) (string, error)
	UpdateAvatar(ctx context.Context, fileID, token string, val []byte) error
}

func NewFiles(cl pkg.FilesClient) FilesClient {
	return &filesClient{cl: cl, c: utils.NewConverter()}
}

type filesClient struct {
	cl pkg.FilesClient
	c  utils.Converter
}

func (f *filesClient) UpdateAvatar(ctx context.Context, fileID, token string, val []byte) error {
	req := models.UploadReq{
		Token:   token,
		Photo:   val,
		PhotoID: fileID,
	}
	_, err := f.cl.UpdateAvatar(ctx, f.c.UpdateAvatarToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (f *filesClient) Upload(ctx context.Context, fileID, token string, val []byte) (string, error) {
	req := models.UploadReq{
		Token:   token,
		Photo:   val,
		PhotoID: fileID,
	}
	res, err := f.cl.UploadToStorage(ctx, f.c.UploadToPb(req))
	if err != nil {
		return "", utils.GRPCErrorToError(err)
	}
	return res.PhotoLink, nil
}
