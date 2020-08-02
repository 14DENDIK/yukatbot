package store

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"
)

type commandsRepo struct {
	db *pgx.Conn
}

func newCommandsRepo(db *pgx.Conn) *commandsRepo {
	return &commandsRepo{
		db: db,
	}
}

func (r *commandsRepo) Get(command, langCode string) (string, error) {
	var textBody string
	if err := r.db.QueryRow(
		context.Background(),
		"SELECT text ->> $2 FROM commands WHERE name=$1",
		command,
		langCode,
	).Scan(&textBody); err != nil {
		return "", err
	}
	textBody = strings.Replace(textBody, `\n`, "\n", -1)
	return textBody, nil
}
