package handlers

import (
	"github.com/MaksKazantsev/Chatter/api/internal/clients"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Posts struct {
	cl clients.PostsClient
}

func NewPosts(cl clients.PostsClient) *Posts {
	return &Posts{cl: cl}
}

// CreatePost godoc
// @Summary CreatePost
// @Description Create new post
// @Tags Post
// @Produce json
// @Param input body models.CreatePostReq true "post request"
// @Param Authorization header string true "token"
//
//	@Success        201 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/posts/create [post]
func (p *Posts) CreatePost(c *fiber.Ctx) error {
	var req models.CreatePostReq
	token := parseAuthHeader(c)

	if err := c.BodyParser(&req); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	req.Token = token

	if err := p.cl.CreatePost(c.Context(), req); err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// DeletePost godoc
// @Summary DeletePost
// @Description Delete your post
// @Tags Post
// @Produce json
// @Param postID path string true "post ID"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/posts/delete/{postID} [delete]
func (p *Posts) DeletePost(c *fiber.Ctx) error {
	postID := c.Params("postID")
	token := parseAuthHeader(c)

	if err := p.cl.DeletePost(c.Context(), token, postID); err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// EditPost godoc
// @Summary EditPost
// @Description Edit your post
// @Tags Post
// @Produce json
// @Param postID path string true "post ID"
// @Param input body models.EditPostReq true "post request"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/posts/edit [put]
func (p *Posts) EditPost(c *fiber.Ctx) error {
	var req models.EditPostReq
	postID := c.Params("postID")
	token := parseAuthHeader(c)

	if err := c.BodyParser(&req); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	req.Token = token
	req.PostID = postID

	if err := p.cl.EditPost(c.Context(), req); err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// LikePost godoc
// @Summary LikePost
// @Description Like post
// @Tags Post
// @Produce json
// @Param input body models.LikePostReq true "like post request"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/posts/like [put]
func (p *Posts) LikePost(c *fiber.Ctx) error {
	var req models.LikePostReq

	token := parseAuthHeader(c)

	if err := c.BodyParser(&req); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	req.Token = token

	if err := p.cl.LikePost(c.Context(), req.Token, req.PostID); err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// DeleteComment godoc
// @Summary DeleteComment
// @Description Delete comment
// @Tags Post
// @Produce json
// @Param commentID path string true "comment ID"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/comment/delete [delete]
func (p *Posts) DeleteComment(c *fiber.Ctx) error {
	commentID := c.Params("commentID")
	token := parseAuthHeader(c)

	if err := p.cl.DeleteComment(c.Context(), token, commentID); err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// LeaveComment godoc
// @Summary LeaveComment
// @Description Leave new comment
// @Tags Post
// @Produce json
// @Param input body models.LeaveCommentReq true "leave comment req"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router        /user/comment/new [post]
func (p *Posts) LeaveComment(c *fiber.Ctx) error {
	var req models.LeaveCommentReq

	token := parseAuthHeader(c)

	if err := c.BodyParser(&req); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	req.Token = token

	if err := p.cl.LeaveComment(c.Context(), req); err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// GetUserPosts godoc
// @Summary GetUserPosts
// @Description Get users post
// @Tags Post
// @Produce json
// @Param userID path string true "user ID"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/posts/get/{userID} [get]
func (p *Posts) GetUserPosts(c *fiber.Ctx) error {

	token := parseAuthHeader(c)
	userID := c.Params("userID")

	res, err := p.cl.GetUserPosts(c.Context(), token, userID)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Status(http.StatusOK).JSON(fiber.Map{"posts": res})
	return nil
}

// UnlikePost godoc
// @Summary  UnlikePost
// @Description Unlike user's post
// @Tags Post
// @Produce json
// @Param input body models.LikePostReq true "dislike post request"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/posts/dislike [put]
func (p *Posts) UnlikePost(c *fiber.Ctx) error {
	var req models.LikePostReq

	token := parseAuthHeader(c)

	if err := c.BodyParser(&req); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	req.Token = token

	err := p.cl.UnlikePost(c.Context(), req.Token, req.PostID)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Status(http.StatusOK)
	return nil
}

// EditComment godoc
// @Summary  EditComment
// @Description Edit your comment
// @Tags Post
// @Produce json
// @Param input body models.EditCommentReq true "edit comment req"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/comment/edit [put]
func (p *Posts) EditComment(c *fiber.Ctx) error {
	var req models.EditCommentReq

	token := parseAuthHeader(c)

	if err := c.BodyParser(&req); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	req.Token = token

	err := p.cl.EditComment(c.Context(), req)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Status(http.StatusOK)
	return nil
}
