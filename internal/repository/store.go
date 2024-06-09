package repository

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	GetUserActivites(ctx context.Context, projects []string, search string, userId int32) ([]UserActivity, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
