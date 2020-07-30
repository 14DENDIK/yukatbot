package store

import "github.com/jackc/pgx/v4"

// Store ...
type Store struct {
	userRepo *userRepo
}

// New ...
func New(db *pgx.Conn) *Store {
	return &Store{
		userRepo: newUserRepo(db),
	}
}
