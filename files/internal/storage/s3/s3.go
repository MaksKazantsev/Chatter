package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/files/internal/storage"
	"github.com/MaksKazantsev/Chatter/files/internal/utils"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v6"
	"os"
)

type Strg struct {
	cl *minio.Client
}

func (s *Strg) Upload(ctx context.Context, id string, val []byte) (string, error) {
	bucketName := "b9b14e14-29afa9b5-ceb5-4e04-b798-0b903a19130d"
	exists, err := s.cl.BucketExists(bucketName)
	if err != nil || !exists {
		return "", utils.NewError(err.Error(), utils.ErrInternal)
	}
	_, err = s.cl.PutObject(bucketName, id, bytes.NewReader(val), int64(len(val)), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return "", utils.NewError(err.Error(), utils.ErrInternal)
	}
	return fmt.Sprintf("https://s3.timeweb.cloud/%s/%s", bucketName, id), nil

}

var _ storage.Storage = &Strg{}

func NewStorage() *Strg {
	_ = godotenv.Load(".env")
	cl, err := minio.New("s3.timeweb.cloud", os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), false)

	if err != nil {
		panic(err)
	}
	return &Strg{cl: cl}
}
