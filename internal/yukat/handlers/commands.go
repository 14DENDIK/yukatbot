package handlers

import (
	"github.com/14DENDIK/yukatbot/api/telegram"
)

func (h *Handler) commandsHandler(body *telegram.Update) error {
	switch body.Message.Text {
	case "/start":
		if err := h.startCommand(&body.Message); err != nil {
			return err
		}
	default:
		if err := h.defaultCommand(&body.Message); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) startCommand(message *telegram.Message) error {
	user, err := h.store.UserRepo.GetOrCreate(&message.From)
	if err != nil {
		return err
	}

	text := "Hello <b>" + user.FirstName + " " + user.LastName + ".</b>\n\n"

	textBody, err := h.store.CommandsRepo.Get(message.Text)
	if err != nil {
		return err
	}

	reply := &telegram.SendMessage{
		ChatID:    message.Chat.ID,
		Text:      text + textBody,
		ParseMode: "HTML",
	}
	if err := h.method.SendMessage(reply); err != nil {
		return err
	}
	return nil
}

func (h *Handler) defaultCommand(message *telegram.Message) error {
	reply := &telegram.SendMessage{
		ChatID:           message.Chat.ID,
		Text:             "Unknown command",
		ReplyToMessageID: message.MessageID,
	}
	if err := h.method.SendMessage(reply); err != nil {
		return err
	}
	return nil
}
