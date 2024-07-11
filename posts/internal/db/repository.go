package db

import (
	"context"
	"github.com/MaksKazantsev/Chatter/posts/internal/models"
)

type Repository interface {
	CreatePost(ctx context.Context, req models.PostReq) error
	DeletePost(ctx context.Context, userID, postID string) error
	EditPost(ctx context.Context, req models.PostReq) error
	LikePost(ctx context.Context, userID, postID, likeID string) error
	UnlikePost(ctx context.Context, userID, postID string) error
	LeaveComment(ctx context.Context, comment models.Comment) error
	GetUserPosts(ctx context.Context, userID string) ([]models.Post, error)
	DeleteComment(ctx context.Context, commentID, userID string) error
	EditComment(ctx context.Context, dbValue, commentID, userID string) error
}
