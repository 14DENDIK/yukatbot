package telegram

// AnswerInlineQuery : AnswerInlineQuery API of Telegram
type AnswerInlineQuery struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text"`
	ShowAlert       bool   `json:"show_alert"`
	URL             string `json:"url"`
}
