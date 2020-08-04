package telegram

// Message : Message API of Telegram
type Message struct {
	MessageID      int64                `json:"message_id"`
	From           User                 `json:"from"`
	Chat           Chat                 `json:"chat"`
	ReplyToMessage *Message             `json:"reply_to_message"`
	Text           string               `json:"text"`
	Entities       []MessageEntity      `json:"entities"`
	Contact        Contact              `json:"contact"`
	NewChatMembers []User               `json:"new_chat_members"`
	LeftChatMember User                 `json:"left_chat_member"`
	ReplyMarkup    InlineKeyboardMarkup `json:"reply_markup"`
}
