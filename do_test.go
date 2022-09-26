package grace

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDo(t *testing.T) {
	var count int
	task := TaskFunc(func() error {
		count++
		return nil
	})
	err := Do(task, task, task)
	require.NoError(t, err)
	require.Equal(t, int(3), count)

	task1 := TaskFunc(func() error {
		count++
		return io.ErrUnexpectedEOF
	})
	err = Do(task, task, task, task1)
	require.Equal(t, io.ErrUnexpectedEOF, err)
	require.Equal(t, int(7), count)

}

func TestDoContext(t *testing.T) {
	type customKey string
	ctx := context.WithValue(context.Background(), customKey("a"), "a")
	var count int
	task := ContextTaskFunc(func(_ctx context.Context) error {
		require.Equal(t, ctx, _ctx)
		count++
		return nil
	})
	err := DoContext(ctx, task, task, task)
	require.NoError(t, err)
	require.Equal(t, int(3), count)

	task1 := ContextTaskFunc(func(_ctx context.Context) error {
		require.Equal(t, ctx, _ctx)
		count++
		return io.ErrUnexpectedEOF
	})
	err = DoContext(ctx, task, task, task, task1)
	require.Equal(t, io.ErrUnexpectedEOF, err)
	require.Equal(t, int(7), count)
}
