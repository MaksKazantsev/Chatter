package handlers

import (
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// GetFriendsSection godoc
// @Summary GetFriendsSection
// @Description Get something from friends section
// @Tags User
// @Produce json
// @Param section query string true "section"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/friends/get [get]
func (u *User) GetFriendsSection(c *fiber.Ctx) error {
	token := parseAuthHeader(c)
	section := c.Query("section")
	switch section {
	case "all":
		friends, err := u.cl.GetFriends(c.Context(), token)
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		_ = c.JSON(fiber.Map{"friends": friends})
		c.Status(http.StatusOK)
		return nil
	case "requests":
		reqs, err := u.cl.GetFs(c.Context(), token)
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		_ = c.JSON(fiber.Map{"fs_reqs": reqs})
		c.Status(http.StatusOK)
		return nil
	default:
		friends, err := u.cl.GetFriends(c.Context(), token)
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		_ = c.JSON(fiber.Map{"friends": friends})
		c.Status(http.StatusOK)
		return nil
	}
}

// DeleteFriend godoc
// @Summary DeleteFriend
// @Description Get all user's friends requests
// @Tags User
// @Produce json
// @Param targetID path string true "UserID to delete from friends"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/friends/delete/{targetID} [delete]
func (u *User) DeleteFriend(c *fiber.Ctx) error {
	token := parseAuthHeader(c)
	targetID := c.Params("targetID")

	err := u.cl.DeleteFriend(c.Context(), token, targetID)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// SuggestFs godoc
// @Summary SuggestFs
// @Description Suggest friendship to user
// @Tags User
// @Produce json
// @Param targetID path string true "UserID to suggest"
// @Param Authorization header string true "token"
//
//	@Success        201 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/friends/suggest/{targetID} [post]
func (u *User) SuggestFs(c *fiber.Ctx) error {
	token := parseAuthHeader(c)
	targetID := c.Params("targetID")

	err := u.cl.SuggestFs(c.Context(), token, targetID)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// RefuseFs godoc
// @Summary RefuseFs
// @Description Refuse friendship request
// @Tags User
// @Produce json
// @Param targetID path string true "SenderID"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/friends/refuse/{targetID} [get]
func (u *User) RefuseFs(c *fiber.Ctx) error {
	token := parseAuthHeader(c)
	targetID := c.Params("targetID")

	err := u.cl.RefuseFs(c.Context(), token, targetID)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// AcceptFs godoc
// @Summary AcceptFs
// @Description Accept friendship request
// @Tags User
// @Produce json
// @Param targetID path string true "SenderID"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/friends/accept/{targetID} [get]
func (u *User) AcceptFs(c *fiber.Ctx) error {
	token := parseAuthHeader(c)
	targetID := c.Params("targetID")

	err := u.cl.AcceptFs(c.Context(), token, targetID)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

// EditProfile godoc
// @Summary EditProfile
// @Description Edit user profile
// @Tags User
// @Produce json
// @Param input body models.UserProfileReq false "edit request"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/profile/edit [put]
func (u *User) EditProfile(c *fiber.Ctx) error {
	var req models.UserProfileReq
	token := parseAuthHeader(c)

	if err := c.BodyParser(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	req.Token = token

	if err := u.cl.EditProfile(c.Context(), req); err != nil {
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
	}

	c.Status(http.StatusOK)
	return nil
}

// GetProfile godoc
// @Summary GetProfile
// @Description Get user's profile
// @Tags User
// @Produce json
// @Param targetID path string true "user's id"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/profile/{targetID} [get]
func (u *User) GetProfile(c *fiber.Ctx) error {
	token := parseAuthHeader(c)
	targetID := c.Params("targetID")

	profile, err := u.cl.GetProfile(c.Context(), token, targetID)
	if err != nil {
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
	}

	_ = c.JSON(fiber.Map{"profile": profile})
	c.Status(http.StatusOK)
	return nil
}

// EditAvatar godoc
// @Summary EditAvatar
// @Description Edit user's avatar
// @Tags User
// @Produce json
// @Param input body models.ProfileAvatar true "user's new avatar"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/profile/avatar/edit [put]
func (u *User) EditAvatar(c *fiber.Ctx) error {
	token := parseAuthHeader(c)
	var req models.ProfileAvatar

	if err := c.BodyParser(&req); err != nil {
		c.Status(http.StatusBadRequest)
	}

	if err := u.cl.EditAvatar(c.Context(), token, req.Avatar); err != nil {
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
	}

	c.Status(http.StatusOK)
	return nil
}

// DeleteAvatar godoc
// @Summary DeleteAvatar
// @Description Delete user's avatar
// @Tags User
// @Produce json
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /user/profile/avatar/delete [delete]
func (u *User) DeleteAvatar(c *fiber.Ctx) error {
	token := parseAuthHeader(c)

	if err := u.cl.DeleteAvatar(c.Context(), token); err != nil {
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
	}

	c.Status(http.StatusOK)
	return nil
}
