package handlers

import (
	"github.com/14DENDIK/yukatbot/api/telegram"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	_ "golang.org/x/text/message/catalog" // Any
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

func (h *Handler) startCommand(tgmessage *telegram.Message) error {
	user, err := h.store.UserRepo.GetOrCreate(&tgmessage.From)
	if err != nil {
		return err
	}
	p := &message.Printer{}
	switch user.LanguageCode {
	case "ru":
		p = message.NewPrinter(language.Russian)
	case "uz":
		p = message.NewPrinter(language.Uzbek)
	default:
		p = message.NewPrinter(language.English)
	}
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
	if err := h.method.SendMessage(reply); err != nil {
		return err
	}
	return nil
}

func (h *Handler) defaultCommand(tgmessage *telegram.Message) error {
	reply := &telegram.SendMessage{
		ChatID:           tgmessage.Chat.ID,
		Text:             "Unknown command",
		ReplyToMessageID: tgmessage.MessageID,
	}
	if err := h.method.SendMessage(reply); err != nil {
		return err
	}
	return nil
}
