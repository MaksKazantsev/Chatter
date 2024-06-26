package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/files/internal/models"
)

func (s *Service) UploadToStorage(ctx context.Context, req models.UploadToStorageReq) (string, error) {
	link, err := s.s3.Upload(ctx, req.FileID, req.File)
	if err != nil {
		return "", fmt.Errorf("storage error: %w", err)
	}
	if err = s.repo.UploadAvatar(ctx, req.UserID, req.FileID, link); err != nil {
		return "", fmt.Errorf("repo error: %w", err)
	}
	return link, nil
}
