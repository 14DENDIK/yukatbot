package telegram

// WebhookInfo : Webhook API of Telegram
type WebhookInfo struct {
	Ok     bool `json:"ok"`
	Result struct {
		URL string `json:"url"`
	} `json:"result"`
}
