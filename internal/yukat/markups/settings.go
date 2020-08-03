package markups

import (
	"github.com/14DENDIK/yukatbot/api/telegram"
	"github.com/14DENDIK/yukatbot/internal/yukat/utils"
)

// SettingsMain ...
func SettingsMain(langCode string) *telegram.InlineKeyboardMarkup {
	markup := &telegram.InlineKeyboardMarkup{}
	p := utils.SetTextPrinter(langCode)
	langButton := []telegram.InlineKeyboardButton{
		{
			Text:         p.Sprintf("🇬🇧 Language"),
			CallbackData: "language",
		},
	}
	doneButton := []telegram.InlineKeyboardButton{
		{
			Text:         p.Sprintf("Done"),
			CallbackData: "done",
		},
	}
	markup.InlineKeyboard = append(markup.InlineKeyboard, langButton)
	markup.InlineKeyboard = append(markup.InlineKeyboard, doneButton)
	return markup
}

// SettingsLanguages ...
func SettingsLanguages(langCode string) *telegram.InlineKeyboardMarkup {
	markup := &telegram.InlineKeyboardMarkup{}
	enButton := []telegram.InlineKeyboardButton{
		{
			Text:         "🇬🇧 English",
			CallbackData: "en",
		},
	}

	ruButton := []telegram.InlineKeyboardButton{
		{
			Text:         "🇷🇺 Русский",
			CallbackData: "ru",
		},
	}

	uzButton := []telegram.InlineKeyboardButton{
		{
			Text:         "🇺🇿 O'zbek",
			CallbackData: "uz",
		},
	}
	markup.InlineKeyboard = append(markup.InlineKeyboard, enButton)
	markup.InlineKeyboard = append(markup.InlineKeyboard, ruButton)
	markup.InlineKeyboard = append(markup.InlineKeyboard, uzButton)
	return markup
}
