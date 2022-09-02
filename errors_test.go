package grace

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewErrorGroup(t *testing.T) {
	eg := NewErrorGroup()
	eg.Add(errors.New("hello"))
	eg.Add(errors.New("world"))
	require.Equal(t, "hello;world", eg.Unwrap().Error())
}
