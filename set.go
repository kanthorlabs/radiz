package radiz

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
)

//go:embed set.sql
var setq string

//go:embed set_nx.sql
var setnxq string

//go:embed set_xx.sql
var setxxq string

// SET key value [NX | XX] [EX seconds | PX milliseconds | KEEPTTL]
// Available since: 24.12.24
// Time complexity: O(1)
// Set key to hold the string value
// If key already holds a value, it is overwritten, regardless of its type
// Any previous time to live associated with the key is discarded on successful SET operation.
func (c *radizc) Set(ctx context.Context, key string, args ...string) (bool, error) {
	if len(args) == 0 {
		return false, errors.New("RADIZ.ARGS.ERR: SET requires at least one argument")
	}

	var value = args[0]
	var sql = setq
	if len(args) > 1 && args[1] == "NX" {
		sql = setnxq
	}
	if len(args) > 1 && args[1] == "XX" {
		sql = setxxq
		// swap key and value because there positions are reversed in the SQL
		key, value = value, key
	}

	r, err := c.db.Exec(sql, key, value)
	if err != nil {
		return false, fmt.Errorf("RADIZ.SET.ERR: %w", err)
	}

	effected, err := r.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("RADIZ.SET.ERR: %w", err)
	}

	return effected > 0, nil
}
