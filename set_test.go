package radiz

import (
	"context"
	"sync"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	c, err := New(context.Background())
	require.NoError(t, err)

	var f = faker.New()
	var key = f.UUID().V4()

	set, err := c.Set(context.Background(), key, f.Lorem().Word())
	require.NoError(t, err)
	require.True(t, set)

	concurrent := faker.New().IntBetween(100, 500)
	var wg sync.WaitGroup
	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			set, err := c.Set(context.Background(), key, f.Lorem().Word())
			require.NoError(t, err)
			require.True(t, set)
		}()
	}
	wg.Wait()
}

func TestSet_NX(t *testing.T) {
	c, err := New(context.Background())
	require.NoError(t, err)

	var f = faker.New()
	var key = f.UUID().V4()

	first, err := c.Set(context.Background(), key, f.Lorem().Word(), "NX")
	require.NoError(t, err)
	require.True(t, first)

	second, err := c.Set(context.Background(), key, f.Lorem().Word(), "NX")
	require.NoError(t, err)
	require.False(t, second)
}

func TestSet_XX(t *testing.T) {
	c, err := New(context.Background())
	require.NoError(t, err)

	var f = faker.New()
	var key = f.UUID().V4()

	// not exist yet so that we cannot set it
	first, err := c.Set(context.Background(), key, f.Lorem().Word(), "XX")
	require.NoError(t, err)
	require.False(t, first)

	// initial
	second, err := c.Set(context.Background(), key, f.Lorem().Word())
	require.NoError(t, err)
	require.True(t, second)

	third, err := c.Set(context.Background(), key, f.Lorem().Word(), "XX")
	require.NoError(t, err)
	require.True(t, third)
}

func TestSet_ErrNotEnoughArgs(t *testing.T) {
	c, err := New(context.Background())
	require.NoError(t, err)

	var f = faker.New()
	var key = f.UUID().V4()

	set, err := c.Set(context.Background(), key)
	require.ErrorContains(t, err, "RADIZ.ARGS.ERR")
	require.False(t, set)
}
