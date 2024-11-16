# something funny

## Sqlite with shared cace

Must use connection string like this `file::memory:?cache=shared`

Otherwise you will get this error

```go
func TestSqlite(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	db.SetMaxOpenConns()
	require.NoError(t, err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS radiz_string (key VARCHAR(255) NOT NULL PRIMARY KEY, value TEXT NOT NULL);")
	require.NoError(t, err)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_, err := db.Exec("INSERT INTO radiz_string (key, value) VALUES (?, ?) ON CONFLICT(key) DO UPDATE SET value = excluded.value", faker.New().Lorem().Word(), faker.New().Lorem().Word())
			require.NoError(t, err)
		}()
	}
	wg.Wait()
}
```

Got error

```
Received unexpected error:
  no such table: radiz_string
Test:       	TestSqlite
```
