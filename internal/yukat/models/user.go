package models

// User : Model for storing data in database
type User struct {
	ID           int64  `json:"id"`
	TelegramID   int64  `json:"telegram_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
