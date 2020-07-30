package telegram

// SendMessage : SendMessage API of Telegram
type SendMessage struct {
	ChatID           int64       `json:"chat_id"`
	Text             string      `json:"text"`
	ParseMode        string      `json:"parse_mode"`
	ReplyToMessageID int64       `json:"reply_to_message_id"`
	ReplyMarkup      interface{} `json:"reply_markup,omitempty"`
}
