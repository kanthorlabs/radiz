package database

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/require"
)

func TestDatabase(t *testing.T) {
	db, err := New(context.Background())
	require.NoError(t, err)

	rows, err := db.Query("SELECT name FROM sqlite_master")
	require.NoError(t, err)

	var names []string
	for rows.Next() {
		var name string
		require.NoError(t, rows.Scan(&name))
		names = append(names, name)
	}

	require.Contains(t, names, "radiz_string")
}
