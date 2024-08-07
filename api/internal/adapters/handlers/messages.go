package handlers

import (
	"context"
	"encoding/json"
	"github.com/MaksKazantsev/Chatter/api/internal/clients"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"sort"
	"strings"
	"sync"
)

type Messages struct {
	cl    clients.MessagesClient
	conns map[string]*websocket.Conn
	mu    sync.RWMutex
}

func NewMessages(cl clients.MessagesClient) *Messages {
	return &Messages{cl: cl, conns: make(map[string]*websocket.Conn)}
}

// Join godoc
// @Summary Join
// @Description Join chat room via WebSocket
// @Tags Chat
// @Produce json
// @Param id query string true "user id"
// @Param input body models.Message true "message"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /chat/ws/join [get]
func (m *Messages) Join(c *websocket.Conn) {
	id := c.Query("id")
	token := parseWSAuthHeader(c)

	m.mu.Lock()
	m.conns[id] = c
	m.mu.Unlock()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			return
		}
		var message models.Message
		if err = json.Unmarshal(msg, &message); err != nil {
			return
		}

		message.Token = token
		message.SenderID = id
		s := strings.Split(id+message.ReceiverID, "")
		sort.Strings(s)
		chatID := strings.Join(s, "")
		message.ChatID = chatID

		var receiverOffline bool
		m.mu.RLock()
		if _, ok := m.conns[message.ReceiverID]; !ok {
			receiverOffline = true
		}
		m.mu.RUnlock()

		if err = m.cl.CreateMessage(context.Background(), &message, receiverOffline); err != nil {
			return
		}

		m.mu.RLock()
		if !receiverOffline {
			_ = m.conns[message.ReceiverID].WriteJSON(message)
		}
		_ = m.conns[id].WriteJSON(message)
		m.mu.RUnlock()
	}
}

// DeleteMessage godoc
// @Summary DeleteMessage
// @Description Delete message
// @Tags Chat
// @Produce json
// @Param id path string true "user id"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /chat/message/{id} [delete]
func (m *Messages) DeleteMessage(c *fiber.Ctx) error {
	token := parseAuthHeader(c)
	id := c.Params("id")

	err := m.cl.DeleteMessage(c.Context(), id, token)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	c.Status(http.StatusOK)
	return nil
}

// GetHistory godoc
// @Summary GetHistory
// @Description Get chat history
// @Tags Chat
// @Produce json
// @Param targetID path string true "user id"
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /chat/history/{targetID} [get]
func (m *Messages) GetHistory(c *fiber.Ctx) error {
	var req models.GetHistoryReq
	req.Token = parseAuthHeader(c)
	req.ChatID = c.Params("targetID")

	messages, err := m.cl.GetHistory(c.Context(), req)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}
	_ = c.JSON(fiber.Map{"messages": messages})
	c.Status(http.StatusOK)
	return nil
}
