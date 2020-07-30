package telegram

// CallbackQuery : CallbackQuery API of Telgram
type CallbackQuery struct {
	ID              string  `json:"id"`
	From            User    `json:"user"`
	Message         Message `json:"message"`
	InlineMessageID string  `json:"inline_message_id"`
}
