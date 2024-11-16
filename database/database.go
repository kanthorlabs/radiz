package database

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed string.sql
var stringt string

func New(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, stringt); err != nil {
		return nil, err
	}

	return db, nil
}
