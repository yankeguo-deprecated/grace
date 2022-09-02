package gracetrack

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	ctx := Init(context.Background())

	Group(ctx, "test-a").SetName("TEST A")
	Group(ctx, "test-a").Add("aaa")
	Group(ctx, "test-a").Add("bbb")

	Group(ctx, "test-b").SetName("TEST B")

	Group(ctx, "test-c").SetName("TEST C")
	Group(ctx, "test-c").Add("aaa")
	Group(ctx, "test-c").Add("bbb")

	expected := []string{
		"TEST A",
		"  * aaa",
		"  * bbb",
		"TEST C",
		"  * aaa",
		"  * bbb",
	}

	require.Equal(t, expected, DumpPlain(ctx))
}
