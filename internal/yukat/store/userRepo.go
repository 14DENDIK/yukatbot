package store

import (
	"context"
	"database/sql"

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
	user := &models.User{
		TelegramID:   t.ID,
		FirstName:    t.FirstName,
		LastName:     t.LastName,
		Username:     t.Username,
		LanguageCode: t.LanguageCode,
	}
	if err := u.db.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE telegram_id=$1;",
	).Scan(user); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		if err := u.db.QueryRow(
			context.Background(),
			"INSERT INTO users(telegram_id, first_name, last_name, username, language_code VALUES($1, $2, $3, $4, $5 RETURNING id;",
			user.TelegramID,
			user.FirstName,
			user.LastName,
			user.Username,
			user.LanguageCode,
		).Scan(&user.ID); err != nil {
			return nil, err
		}
	}
	return user, nil
}
