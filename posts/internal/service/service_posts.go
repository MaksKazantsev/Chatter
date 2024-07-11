package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/posts/internal/log"
	"github.com/MaksKazantsev/Chatter/posts/internal/models"
	"github.com/MaksKazantsev/Chatter/posts/pkg"
	"github.com/google/uuid"
	"sort"
	"strings"
	"time"
)

func (s *Service) CreatePost(ctx context.Context, req models.PostReq) error {
	req.PostID = uuid.New().String()

	if err := s.repo.CreatePost(ctx, req); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	s.broker.Produce(ctx, pkg.KafkaMessage{ID: req.PostAuthorID, LastOnline: time.Now()})

	return nil
}

func (s *Service) DeletePost(ctx context.Context, userID, postID string) error {
	log.GetLogger(ctx).Debug("Service layer success")

	if err := s.repo.DeletePost(ctx, userID, postID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	s.broker.Produce(ctx, pkg.KafkaMessage{ID: userID, LastOnline: time.Now()})
	return nil
}

func (s *Service) EditPost(ctx context.Context, req models.PostReq) error {
	log.GetLogger(ctx).Debug("Service layer success")

	if err := s.repo.EditPost(ctx, req); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	s.broker.Produce(ctx, pkg.KafkaMessage{ID: req.PostAuthorID, LastOnline: time.Now()})

	return nil
}

func (s *Service) LikePost(ctx context.Context, userID, postID string) error {
	st := strings.Split(userID+postID, "")
	sort.Strings(st)
	likeID := strings.Join(st, "")

	log.GetLogger(ctx).Debug("Service layer success")

	if err := s.repo.LikePost(ctx, userID, postID, likeID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	s.broker.Produce(ctx, pkg.KafkaMessage{ID: userID, LastOnline: time.Now()})

	return nil
}
func (s *Service) LeaveComment(ctx context.Context, comment models.Comment) error {
	comment.CommentID = uuid.New().String()
	log.GetLogger(ctx).Debug("Service layer success")
	if err := s.repo.LeaveComment(ctx, comment); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	s.broker.Produce(ctx, pkg.KafkaMessage{ID: comment.UserID, LastOnline: time.Now()})

	return nil
}

func (s *Service) GetUserPosts(ctx context.Context, userID, requesterID string) ([]models.Post, error) {
	log.GetLogger(ctx).Debug("Service layer success")

	res, err := s.repo.GetUserPosts(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}

	s.broker.Produce(ctx, pkg.KafkaMessage{ID: requesterID, LastOnline: time.Now()})

	return res, nil
}

func (s *Service) DeleteComment(ctx context.Context, commentID, userID string) error {
	log.GetLogger(ctx).Debug("Service layer success")

	err := s.repo.DeleteComment(ctx, commentID, userID)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	s.broker.Produce(ctx, pkg.KafkaMessage{ID: userID, LastOnline: time.Now()})

	return nil
}

func (s *Service) UnlikePost(ctx context.Context, userID, postID string) error {
	log.GetLogger(ctx).Debug("Service layer success")

	err := s.repo.UnlikePost(ctx, userID, postID)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	s.broker.Produce(ctx, pkg.KafkaMessage{ID: userID, LastOnline: time.Now()})
	return nil
}

func (s *Service) EditComment(ctx context.Context, req models.EditCommentReq) error {
	log.GetLogger(ctx).Debug("Service layer success")

	dbValue := req.Value.TextValue + ":" + req.Value.File

	err := s.repo.EditComment(ctx, dbValue, req.CommentID, req.UserID)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	s.broker.Produce(ctx, pkg.KafkaMessage{ID: req.UserID, LastOnline: time.Now()})

	return nil
}
