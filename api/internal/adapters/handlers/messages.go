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
	cl    clients.Messages
	conns map[string]*websocket.Conn
	mu    sync.RWMutex
}

func NewMessages(cl clients.Messages) *Messages {
	return &Messages{cl: cl, conns: make(map[string]*websocket.Conn)}
}

func (m *Messages) Join(c *websocket.Conn) {
	id := c.Query("id")
	token := c.Query("token")

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
