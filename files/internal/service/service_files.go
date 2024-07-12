package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MaksKazantsev/Chatter/files/internal/models"
	"github.com/MaksKazantsev/Chatter/files/internal/utils"
	"github.com/MaksKazantsev/Chatter/files/pkg"
)

func (s *Service) UploadToStorage(ctx context.Context, req models.UploadToStorageReq) (string, error) {
	link, err := s.s3.Upload(ctx, req.FileID, req.File)
	if err != nil {
		return "", fmt.Errorf("storage error: %w", err)
	}
	return link, nil
}

func (s *Service) UpdateAvatar(ctx context.Context, req models.UploadToStorageReq) error {
	link, err := s.s3.Upload(ctx, req.FileID, req.File)

	b, err := json.Marshal(models.UpdateAvatarMessage{ID: req.UserID, Avatar: link})
	if err != nil {
		return utils.NewError("failed to marshal data: "+err.Error(), utils.ErrInternal)
	}

	s.broker.Publish(ctx, pkg.Message{Type: "UpdateAvatar", Data: b})

	if err != nil {
		return fmt.Errorf("storage error: %w", err)
	}
	return nil
}
