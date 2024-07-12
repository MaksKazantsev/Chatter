package clients

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/posts/pkg/grpc"
)

type PostsClient interface {
	CreatePost(ctx context.Context, req models.CreatePostReq) error
	DeletePost(ctx context.Context, token, postID string) error
	EditPost(ctx context.Context, req models.EditPostReq) error
	LikePost(ctx context.Context, token, postID string) error
	LeaveComment(ctx context.Context, req models.LeaveCommentReq) error
	GetUserPosts(ctx context.Context, token, userID string) ([]models.Post, error)
	DeleteComment(ctx context.Context, token, commentID string) error
	UnlikePost(ctx context.Context, token, postID string) error
	EditComment(ctx context.Context, req models.EditCommentReq) error
}

type postsClient struct {
	cl pkg.PostsClient
	c  utils.Converter
}

func NewPosts(cl pkg.PostsClient) PostsClient {
	return &postsClient{c: utils.NewConverter(), cl: cl}
}

func (p *postsClient) CreatePost(ctx context.Context, req models.CreatePostReq) error {
	_, err := p.cl.CreatePost(ctx, p.c.CreatePostToPb(req))
	if err != nil {

		fmt.Println(err)
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (p *postsClient) DeletePost(ctx context.Context, token, postID string) error {
	_, err := p.cl.DeletePost(ctx, &pkg.DeletePostReq{PostID: postID, Token: token})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (p *postsClient) EditPost(ctx context.Context, req models.EditPostReq) error {
	_, err := p.cl.EditPost(ctx, p.c.EditPostToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (p *postsClient) LikePost(ctx context.Context, token, postID string) error {
	_, err := p.cl.LikePost(ctx, &pkg.LikePostReq{PostID: postID, Token: token})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (p *postsClient) LeaveComment(ctx context.Context, req models.LeaveCommentReq) error {
	_, err := p.cl.LeaveComment(ctx, p.c.LeaveCommentToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (p *postsClient) GetUserPosts(ctx context.Context, token, userID string) ([]models.Post, error) {
	fmt.Println(userID)
	res, err := p.cl.GetUserPosts(ctx, p.c.GetUserPostsToPb(token, userID))
	if err != nil {
		return nil, utils.GRPCErrorToError(err)
	}
	return p.c.GetUserPostsToService(res), nil
}

func (p *postsClient) DeleteComment(ctx context.Context, token, commentID string) error {
	_, err := p.cl.DeleteComment(ctx, &pkg.DeleteCommentReq{CommentID: commentID, Token: token})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (p *postsClient) UnlikePost(ctx context.Context, token, postID string) error {
	_, err := p.cl.UnlikePost(ctx, &pkg.LikePostReq{PostID: postID, Token: token})
	if err != nil {
		return utils.GRPCErrorToError(err)
	}

	return nil
}

func (p *postsClient) EditComment(ctx context.Context, req models.EditCommentReq) error {
	_, err := p.cl.EditComment(ctx, p.c.EditCommentToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}
