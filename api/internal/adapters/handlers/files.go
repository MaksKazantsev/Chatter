package handlers

import (
	"github.com/MaksKazantsev/Chatter/api/internal/clients"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type Files struct {
	cl clients.Files
}

func NewFiles(cl clients.Files) *Files {
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
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}
	fl, err := file.Open()
	if err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}
	b, err := io.ReadAll(fl)
	if err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}
	fileLink, err := f.cl.Upload(c.Context(), uuid.New().String(), token, b)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.SendString(fileLink)
	return nil
}
