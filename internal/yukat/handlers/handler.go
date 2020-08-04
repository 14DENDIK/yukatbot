package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/14DENDIK/yukatbot/api/telegram"
	"github.com/14DENDIK/yukatbot/internal/yukat/methods"
	"github.com/14DENDIK/yukatbot/internal/yukat/store"
	"go.uber.org/zap"
)

// Handler ...
type Handler struct {
	store  *store.Store
	logger *zap.Logger
	method *methods.Method
}

// New ...
func New(store *store.Store, logger *zap.Logger, token string) *Handler {
	return &Handler{
		store:  store,
		logger: logger,
		method: methods.New(token),
	}
}

// MainEntry ...
func (h *Handler) MainEntry(res http.ResponseWriter, req *http.Request) {
	body := &telegram.Update{}

	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		h.logger.Sugar().Error("Could not decode req body")
		return
	}

	if len(body.Message.Entities) > 0 {
		if body.Message.Entities[0].Type == "bot_command" {
			h.logger.Sugar().Infof("[%s] Command: %v", req.Method, body.Message.Text)
			if err := h.commandsHandler(body); err != nil {
				h.logger.Sugar().Error(err)
			}
		}
	} else if body.CallbackQuery.ID != "" {
		h.logger.Sugar().Infof("[%s] Callback Data: %v", req.Method, body.CallbackQuery.Data)
		if err := h.callbacksHandler(body); err != nil {
			h.logger.Sugar().Error(err)
		}
	} else if len(body.Message.NewChatMembers) > 0 {
		h.logger.Sugar().Infof("[%s] %v - User(s) joined. Chat: %v", req.Method, len(body.Message.NewChatMembers), body.Message.Chat.Title)
		if err := h.newChatMembersJoined(body); err != nil {
			h.logger.Sugar().Error(err)
		}
	} else if body.Message.LeftChatMember.ID != 0 {
		h.logger.Sugar().Infof("[%s] User(s) left. Chat: %v", req.Method, body.Message.Chat.Title)
		if err := h.leftChatMember(body); err != nil {
			h.logger.Sugar().Error(err)
		}
	} else {
		h.logger.Sugar().Infof("[%s] Message: %v", req.Method, body.Message.Text)
		fmt.Print("Just Text\n")
	}
}
