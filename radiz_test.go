package radiz

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

func TestSqlite(t *testing.T) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	require.NoError(t, err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS radiz_string (key VARCHAR(255) NOT NULL PRIMARY KEY, value TEXT NOT NULL);")
	require.NoError(t, err)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_, err := db.Exec("INSERT INTO radiz_string (key, value) VALUES (?, ?) ON CONFLICT(key) DO UPDATE SET value = excluded.value", strconv.Itoa(i), faker.New().Lorem().Word())
			require.NoError(t, err)
		}()
	}
	wg.Wait()

	var c int
	require.NoError(t, db.QueryRow("SELECT count(*) as total FROM radiz_string").Scan(&c))
	fmt.Println("-------------", c)
}
