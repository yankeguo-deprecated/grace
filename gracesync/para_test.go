package gracesync

import (
	"context"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestParaDo(t *testing.T) {
	vs := []string{"hello", "world"}
	var out []string

	err := DoPara(
		context.Background(),
		1,
		vs,
		func(ctx context.Context, v string) (err error) {
			out = append(out, v)
			return
		},
	)

	sort.Strings(vs)
	sort.Strings(out)

	require.NoError(t, err)
	require.Equal(t, vs, out)
}
