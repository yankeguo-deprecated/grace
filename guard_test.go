package grace

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGuard(t *testing.T) {
	var err error
	defer func() {
		require.Error(t, err)
		require.Equal(t, "hello", err.Error())
	}()
	defer Guard(&err)
	panic(errors.New("hello"))
}

func TestMustContext(t *testing.T) {
	var err error
	defer func() {
		require.Error(t, err)
		require.Equal(t, context.Canceled, err)
	}()
	defer Guard(&err)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	MustContext(ctx)
}
