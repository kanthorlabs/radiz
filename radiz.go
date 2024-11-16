package radiz

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/kanthorlabs/radiz/database"
	_ "github.com/mattn/go-sqlite3"
)

var _ Radiz = (*radizc)(nil)

//go:embed version
var version string

func New(ctx context.Context) (Radiz, error) {
	db, err := database.New(ctx)
	if err != nil {
		return nil, err
	}
	return &radizc{db: db}, nil
}

type Radiz interface {
	Version() string
	Set(ctx context.Context, key string, args ...string) (bool, error)
	Get(ctx context.Context, key string) (string, error)
}

type radizc struct {
	db *sql.DB
}

func (c *radizc) Version() string {
	return version
}

func Sqlite() {

}
