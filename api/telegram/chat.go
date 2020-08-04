package telegram

// Chat : Chat API of Telegram
type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// GetChatAdministrators ...
type GetChatAdministrators struct {
	ChatID int64 `json:"chat_id"`
}

// GetChatMember ...
type GetChatMember struct {
	ChatID int64 `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

// ChatMember ...
type ChatMember struct {
	Ok     bool `json:"ok"`
	Result struct {
		User              User   `json:"user"`
		Status            string `json:"status"`
		CanDeleteMessages bool   `json:"can_delete_messages"`
	} `json:"result"`
}
