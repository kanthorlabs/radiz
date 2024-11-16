package radiz

import "context"

// GET key
// Available since: 24.12.24
// Time complexity: O(1)
// Get the value of key.
// If the key does not exist the special value nil is returned.
// An error is returned if the value stored at key is not a string, because GET only handles string values.
func (c *radizc) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}
