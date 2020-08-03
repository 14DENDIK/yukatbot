package handlers

import (
	"github.com/14DENDIK/yukatbot/api/telegram"
	"github.com/14DENDIK/yukatbot/internal/yukat/markups"
	"github.com/14DENDIK/yukatbot/internal/yukat/utils"
)

func (h *Handler) commandsHandler(body *telegram.Update) error {
	switch body.Message.Text {
	case "/start":
		if err := h.startCommand(&body.Message); err != nil {
			return err
		}
	case "/help":
		if err := h.helpCommand(&body.Message); err != nil {
			return err
		}
	case "/settings":
		if err := h.settingsCommand(&body.Message); err != nil {
			return err
		}
	default:
		if err := h.defaultCommand(&body.Message); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) startCommand(tgmessage *telegram.Message) error {
	user, err := h.store.UserRepo.GetOrCreate(&tgmessage.From)
	if err != nil {
		return err
	}
	p := utils.SetTextPrinter(user.LanguageCode)
	text := p.Sprintf("Hello <b>%s %s.</b>\n\n", user.FirstName, user.LastName)
	textBody, err := h.store.CommandsRepo.Get(tgmessage.Text, user.LanguageCode)
	if err != nil {
		return err
	}

	reply := &telegram.SendMessage{
		ChatID:    tgmessage.Chat.ID,
		Text:      text + textBody,
		ParseMode: "HTML",
	}
	if err := h.method.RunMethod("sendMessage", reply); err != nil {
		return err
	}
	return nil
}

func (h *Handler) helpCommand(tgmessage *telegram.Message) error {
	user, err := h.store.UserRepo.GetOrCreate(&tgmessage.From)
	if err != nil {
		return err
	}
	text, err := h.store.CommandsRepo.Get(tgmessage.Text, user.LanguageCode)
	if err != nil {
		return err
	}
	reply := &telegram.SendMessage{
		ChatID:    tgmessage.Chat.ID,
		Text:      text,
		ParseMode: "HTML",
	}
	if err = h.method.RunMethod("sendMessage", reply); err != nil {
		return err
	}
	return nil
}

func (h *Handler) settingsCommand(tgmessage *telegram.Message) error {
	user, err := h.store.UserRepo.GetOrCreate(&tgmessage.From)
	if err != nil {
		return err
	}
	text, err := h.store.CommandsRepo.Get(tgmessage.Text, user.LanguageCode)
	if err != nil {
		return err
	}
	reply := &telegram.SendMessage{
		ChatID:      tgmessage.Chat.ID,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: markups.SettingsMain(user.LanguageCode),
	}

	// Should work on error handling of goroutines
	go h.method.RunMethod("deleteMessage", &telegram.DeleteMessage{
		ChatID:    tgmessage.Chat.ID,
		MessageID: tgmessage.MessageID,
	})
	if err = h.method.RunMethod("sendMessage", reply); err != nil {
		return err
	}
	return nil
}

func (h *Handler) defaultCommand(tgmessage *telegram.Message) error {
	user, err := h.store.UserRepo.GetOrCreate(&tgmessage.From)
	if err != nil {
		return err
	}
	p := utils.SetTextPrinter(user.LanguageCode)
	reply := &telegram.SendMessage{
		ChatID:           tgmessage.Chat.ID,
		Text:             p.Sprintf("Unknown command"),
		ReplyToMessageID: tgmessage.MessageID,
	}
	if err := h.method.RunMethod("sendMessage", reply); err != nil {
		return err
	}
	return nil
}
