package telegram

// Message : Message API of Telegram
type Message struct {
	MessageID      int64                `json:"message_id"`
	From           User                 `json:"from"`
	Chat           Chat                 `json:"chat"`
	ReplyToMessage *Message             `json:"reply_to_message"`
	Text           string               `json:"text"`
	Contact        Contact              `json:"contact"`
	ReplyMarkup    InlineKeyboardMarkup `json:"reply_markup"`
}
