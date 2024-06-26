package clients

import (
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
	
	return fileLink, nil
}
