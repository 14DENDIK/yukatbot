package telegram

// AnswerCallbackQuery ...
type AnswerCallbackQuery struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text"`
	ShowAlert       bool   `json:"show_alert"`
	URL             string `json:"url"`
}
