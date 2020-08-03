package telegram

// EditMessageReplyMarkup ...
type EditMessageReplyMarkup struct {
	ChatID          int64                `json:"chat_id"`
	MessageID       int64                `json:"message_id"`
	InlineMessageID string               `json:"inline_message_id"`
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup"`
}

// EditMessageText ...
type EditMessageText struct {
	ChatID                int64                `json:"chat_id"`
	MessageID             int64                `json:"message_id"`
	InlineMessageID       string               `json:"inline_message_id"`
	Text                  string               `json:"text"`
	ParseMode             string               `json:"parse_mode"`
	DisableWebPagePreview bool                 `json:"disable_web_page_preview"`
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`
}

// DeleteMessage ...
type DeleteMessage struct {
	ChatID    int64 `json:"chat_id"`
	MessageID int64 `json:"message_id"`
}
