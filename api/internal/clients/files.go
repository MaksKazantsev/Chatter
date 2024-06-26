package clients

import (
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/files/pkg/grpc"
	"golang.org/x/net/context"
)

type Files interface {
	Upload(ctx context.Context, fileID, token string, val []byte) (string, error)
}

func NewFiles(cl pkg.FilesClient) Files {
	return &filesCl{cl: cl, c: utils.NewConverter()}
}

type filesCl struct {
	cl pkg.FilesClient
	c  utils.Converter
}

func (f filesCl) Upload(ctx context.Context, fileID, token string, val []byte) (string, error) {
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
