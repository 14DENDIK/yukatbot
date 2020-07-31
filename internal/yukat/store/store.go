package store

import "github.com/jackc/pgx/v4"

// Store ...
type Store struct {
	UserRepo     *userRepo
	CommandsRepo *commandsRepo
}

// New ...
func New(db *pgx.Conn) *Store {
	return &Store{
		UserRepo:     newUserRepo(db),
		CommandsRepo: newCommandsRepo(db),
	}
}
