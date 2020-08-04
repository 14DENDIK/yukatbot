package handlers

import (
	"github.com/14DENDIK/yukatbot/api/telegram"
	"github.com/14DENDIK/yukatbot/internal/yukat/markups"
	"github.com/14DENDIK/yukatbot/internal/yukat/models"
	"github.com/14DENDIK/yukatbot/internal/yukat/utils"
)

func (h *Handler) callbacksHandler(body *telegram.Update) error {
	user, err := h.store.UserRepo.Get(&body.CallbackQuery.From)
	if err != nil {
		return err
	}
	switch body.CallbackQuery.Data {
	case "language":
		if err := h.languageCallback(&body.CallbackQuery, user); err != nil {
			return err
		}
	case "ru":
		fallthrough
	case "uz":
		fallthrough
	case "en":
		if err := h.changeLangCallback(&body.CallbackQuery, user); err != nil {
			return err
		}
	case "done":
		if err := h.langDoneCallback(&body.CallbackQuery, user); err != nil {
			return err
		}
	default:
		if err := h.defaultCallback(&body.CallbackQuery, user); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) languageCallback(callback *telegram.CallbackQuery, user *models.User) error {
	p := utils.SetTextPrinter(user.LanguageCode)
	reply := &telegram.EditMessageReplyMarkup{
		ChatID:      callback.Message.Chat.ID,
		MessageID:   callback.Message.MessageID,
		ReplyMarkup: *markups.SettingsLanguages(user.LanguageCode),
	}
	answer := &telegram.AnswerCallbackQuery{
		CallbackQueryID: callback.ID,
		Text:            p.Sprintf("Choose language"),
	}
	if err := h.method.RunPostMethod("editMessageReplyMarkup", reply); err != nil {
		return err
	}
	if err := h.method.RunPostMethod("answerCallbackQuery", answer); err != nil {
		return err
	}
	return nil
}

func (h *Handler) changeLangCallback(callback *telegram.CallbackQuery, user *models.User) error {
	user.LanguageCode = callback.Data
	p := utils.SetTextPrinter(user.LanguageCode)
	if err := h.store.UserRepo.Update(user); err != nil {
		return err
	}
	text, err := h.store.CommandsRepo.Get("/settings", user.LanguageCode)
	if err != nil {
		return err
	}
	replyText := &telegram.EditMessageText{
		ChatID:      callback.Message.Chat.ID,
		MessageID:   callback.Message.MessageID,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: *markups.SettingsMain(user.LanguageCode),
	}
	answer := &telegram.AnswerCallbackQuery{
		CallbackQueryID: callback.ID,
		Text:            p.Sprintf("Language changed"),
	}
	if err = h.method.RunPostMethod("editMessageText", replyText); err != nil {
		return err
	}
	if err = h.method.RunPostMethod("answerCallbackQuery", answer); err != nil {
		return err
	}
	return nil
}

func (h *Handler) langDoneCallback(callback *telegram.CallbackQuery, user *models.User) error {
	p := utils.SetTextPrinter(user.LanguageCode)
	reply := *&telegram.AnswerCallbackQuery{
		CallbackQueryID: callback.ID,
		Text:            p.Sprintf("Settings are changed"),
		ShowAlert:       true,
	}
	if err := h.method.RunPostMethod("deleteMessage", &telegram.DeleteMessage{
		ChatID:    callback.Message.Chat.ID,
		MessageID: callback.Message.MessageID,
	}); err != nil {
		return err
	}
	if err := h.method.RunPostMethod("answerCallbackQuery", reply); err != nil {
		return err
	}
	return nil
}

func (h *Handler) defaultCallback(callback *telegram.CallbackQuery, user *models.User) error {
	p := utils.SetTextPrinter(user.LanguageCode)
	reply := &telegram.AnswerCallbackQuery{
		CallbackQueryID: callback.ID,
		Text:            p.Sprintf("Wrong callback input!"),
	}

	if err := h.method.RunPostMethod("answerCallbackQuery", reply); err != nil {
		return err
	}

	return nil
}
