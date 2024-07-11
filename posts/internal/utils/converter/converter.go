package converter

import (
	"github.com/MaksKazantsev/Chatter/posts/internal/models"
	pkg "github.com/MaksKazantsev/Chatter/posts/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Converter interface {
	CreatePostToService(req *pkg.CreatePostReq, authorID string) models.PostReq
	EditPostToService(req *pkg.EditPostReq, authorID string) models.PostReq
	LeaveCommentToService(req *pkg.LeaveCommentReq, userID string) models.Comment
	GetUserPostsToPb([]models.Post) *pkg.GetUserPostsRes
	EditCommentToService(req *pkg.EditCommentReq, userID string) models.EditCommentReq
}

type converter struct {
}

func (c converter) EditCommentToService(req *pkg.EditCommentReq, userID string) models.EditCommentReq {
	return models.EditCommentReq{
		CommentID: req.CommentID,
		Value:     models.CommentValue{TextValue: req.Value.TextValue, File: req.Value.File},
		UserID:    userID,
	}
}

func (c converter) GetUserPostsToPb(posts []models.Post) *pkg.GetUserPostsRes {
	var res []*pkg.Post

	for _, p := range posts {
		post := &pkg.Post{}
		for _, l := range p.Likes {
			if p.PostID == l.PostID {
				post.Likes = append(post.Likes, &pkg.Like{UserID: l.UserID, PostID: l.PostID})
			}
		}
		for _, co := range p.Comments {
			if p.PostID == co.PostID {
				post.Comments = append(post.Comments, &pkg.Comment{UserID: co.UserID, PostID: co.PostID, CommentID: co.CommentID, Value: &pkg.CommentValue{TextValue: co.Value.TextValue, File: co.Value.File}, CreatedAt: timestamppb.New(co.CreatedAt)})
			}
		}
		post.PostID = p.PostID
		post.UserID = p.PostAuthorID
		post.Title = p.PostTitle
		post.Description = p.PostDescription
		post.File = p.PostFile
		post.CreatedAt = timestamppb.New(p.CreatedAt)
		res = append(res, post)
	}

	return &pkg.GetUserPostsRes{Posts: res}
}

func NewConverter() Converter {
	return &converter{}
}

// To Service

func (c converter) LeaveCommentToService(req *pkg.LeaveCommentReq, userID string) models.Comment {
	return models.Comment{
		PostID: req.PostID,
		UserID: userID,
		Value:  models.CommentValue{TextValue: req.Value.TextValue, File: req.Value.File},
	}
}

func (c converter) EditPostToService(req *pkg.EditPostReq, authorID string) models.PostReq {
	return models.PostReq{
		PostID:          req.PostID,
		PostAuthorID:    authorID,
		PostTitle:       req.Title,
		PostDescription: req.Description,
		PostFile:        req.File,
	}
}

func (c converter) CreatePostToService(req *pkg.CreatePostReq, authorID string) models.PostReq {
	return models.PostReq{
		PostFile:        req.File,
		PostDescription: req.Description,
		PostTitle:       req.Title,
		PostAuthorID:    authorID,
	}
}
