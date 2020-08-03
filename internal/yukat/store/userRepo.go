package store

import (
	"context"

	"github.com/14DENDIK/yukatbot/api/telegram"
	"github.com/14DENDIK/yukatbot/internal/yukat/models"
	"github.com/jackc/pgx/v4"
)

type userRepo struct {
	db *pgx.Conn
}

// NewUserRepo ...
func newUserRepo(db *pgx.Conn) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetOrCreate(t *telegram.User) (*models.User, error) {
	user, err := u.Get(t)
	if err != nil {
		if err != pgx.ErrNoRows {
			return nil, err
		}
		user, err = u.Create(t)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (u *userRepo) Get(t *telegram.User) (*models.User, error) {
	user := &models.User{
		TelegramID: t.ID,
		FirstName:  t.FirstName,
		LastName:   t.LastName,
		Username:   t.Username,
	}
	if err := u.db.QueryRow(
		context.Background(),
		"SELECT id, language_code, current_step FROM users WHERE telegram_id=$1;",
		user.TelegramID,
	).Scan(&user.ID, &user.LanguageCode, &user.CurrentStep); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) Create(t *telegram.User) (*models.User, error) {
	user := &models.User{
		TelegramID:   t.ID,
		FirstName:    t.FirstName,
		LastName:     t.LastName,
		Username:     t.Username,
		LanguageCode: "en",
		CurrentStep:  "main",
	}
	if contains([]string{"en", "ru", "uz"}, t.LanguageCode) {
		user.LanguageCode = t.LanguageCode
	}
	if err := u.db.QueryRow(
		context.Background(),
		"INSERT INTO users(telegram_id, first_name, last_name, username, language_code, current_step) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;",
		user.TelegramID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.LanguageCode,
		user.CurrentStep,
	).Scan(&user.ID); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) Update(user *models.User) error {
	if _, err := u.db.Exec(
		context.Background(),
		"UPDATE users SET first_name=$1, last_name=$2, username=$3, language_code=$4, current_step=$5 WHERE id=$6 AND telegram_id=$7;",
		user.FirstName,
		user.LastName,
		user.Username,
		user.LanguageCode,
		user.CurrentStep,
		user.ID,
		user.TelegramID,
	); err != nil {
		return err
	}
	return nil
}

func contains(codes []string, langCode string) bool {
	for _, c := range codes {
		if c == langCode {
			return true
		}
	}
	return false
}
