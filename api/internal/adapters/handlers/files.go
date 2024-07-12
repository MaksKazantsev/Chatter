package handlers

import (
	"fmt"
	"github.com/MaksKazantsev/Chatter/api/internal/clients"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type Files struct {
	cl clients.FilesClient
}

func NewFiles(cl clients.FilesClient) *Files {
	return &Files{cl: cl}
}

// Upload godoc
// @Summary Upload
// @Description Upload file to cloud storage
// @Tags Files
// @Produce json
// @Param Authorization header string true "token"
//
//	@Success        201 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /files/upload [post]
func (f *Files) Upload(c *fiber.Ctx) error {
	token := parseAuthHeader(c)

	file, err := c.FormFile("file")
	if err != nil {
		return utils.NewError(err.Error(), utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	fl, err := file.Open()
	if err != nil {
		return utils.NewError(err.Error(), utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	b, err := io.ReadAll(fl)
	if err != nil {
		return utils.NewError(err.Error(), utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	fileLink, err := f.cl.Upload(c.Context(), uuid.New().String(), token, b)
	if err != nil {
		return fmt.Errorf("client error: %w", err)
	}

	_ = c.SendString(fileLink)
	return nil
}

// UpdateAvatar godoc
// @Summary UpdateAvatar
// @Description Updates user's avatar
// @Tags Files
// @Produce json
// @Param Authorization header string true "token"
//
//	@Success        201 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /files/update [post]
func (f *Files) UpdateAvatar(c *fiber.Ctx) error {
	token := parseAuthHeader(c)

	file, err := c.FormFile("file")
	if err != nil {
		return utils.NewError(err.Error(), utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	fl, err := file.Open()
	if err != nil {
		return utils.NewError(err.Error(), utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	b, err := io.ReadAll(fl)
	if err != nil {
		return utils.NewError(err.Error(), utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	err = f.cl.UpdateAvatar(c.Context(), uuid.New().String(), token, b)
	if err != nil {
		return fmt.Errorf("client error: %w", err)
	}

	c.Status(http.StatusCreated)
	return nil
}
