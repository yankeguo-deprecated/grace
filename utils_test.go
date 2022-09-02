package grace

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestPtr(t *testing.T) {
	v1 := Ptr("hello")
	require.Equal(t, "hello", *v1)
}

func TestRepeat(t *testing.T) {
	v := Repeat(5, 1)
	require.Equal(t, []int{1, 1, 1, 1, 1}, v)
}

func TestMapKeys(t *testing.T) {
	keys := MapKeys(map[string]string{"a": "b", "c": "d"})
	sort.Strings(keys)
	require.Equal(t, []string{"a", "c"}, keys)
}

func TestMapVals(t *testing.T) {
	vals := MapVals(map[string]string{"a": "b", "c": "d"})
	sort.Strings(vals)
	require.Equal(t, []string{"b", "d"}, vals)
}

func TestSliceToMap(t *testing.T) {
	m := SliceToMap([]string{"aa", "b"}, func(v string) int {
		return len(v)
	})
	require.Equal(t, map[int]string{1: "b", 2: "aa"}, m)
}
