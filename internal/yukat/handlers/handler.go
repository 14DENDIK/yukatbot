package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	if strings.HasPrefix(body.Message.Text, "/") {
		fmt.Println("Command")
	} else if body.CallbackQuery.ID != "" {
		fmt.Println("Callback")
	} else {
		fmt.Print("Just Text")
	}
}
