package graceos

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/guoyk93/grace"
	"github.com/stretchr/testify/require"
)

func TestEnvVal(t *testing.T) {
	var v1 string
	var v2 int
	var v3 []int64
	var v4 map[string]bool
	os.Setenv("TEST_V1", "hello")
	os.Setenv("TEST_V2", "99")
	os.Setenv("TEST_V3", "12,24")
	os.Setenv("TEST_V4", "hello:true;world:false")
	err := grace.Do(
		EnvVal(&v1, "TEST_V1", true),
		EnvVal(&v2, "TEST_V2", true),
		EnvSlice(&v3, "TEST_V3", true),
		EnvMap(&v4, "TEST_V4", true),
	)
	require.NoError(t, err)
	require.Equal(t, "hello", v1)
	require.Equal(t, 99, v2)
	require.Equal(t, []int64{12, 24}, v3)
	require.Equal(t, map[string]bool{"hello": true, "world": false}, v4)
}

func envDecodeValCase[T any](t *testing.T, s string, v T) grace.TaskFunc {
	return func() error {
		t.Run(reflect.TypeOf(v).String(), func(t *testing.T) {
			var out T
			err := DecodeEnvVal(&out, s)
			require.NoError(t, err)
			require.Equal(t, v, out)
		})
		t.Run(reflect.TypeOf(v).String()+"/slice", func(t *testing.T) {
			var out []T
			err := DecodeEnvSlice(&out, s+","+s)
			require.NoError(t, err)
			require.Equal(t, []T{v, v}, out)
		})
		t.Run(reflect.TypeOf(v).String()+"/map", func(t *testing.T) {
			var out map[string]T
			err := DecodeEnvMap(&out, "hello:"+s+";"+"world:"+s)
			require.NoError(t, err)
			require.Equal(t, map[string]T{"hello": v, "world": v}, out)
		})
		return nil
	}
}

func TestEnvDecodeVal(t *testing.T) {
	_ = grace.Do(
		envDecodeValCase(t, "hello", "hello"),
		envDecodeValCase(t, "true", true),
		envDecodeValCase(t, "12", 12),
		envDecodeValCase(t, "12", int64(12)),
		envDecodeValCase(t, "12", uint64(12)),
		envDecodeValCase(t, "1.2", float64(1.2)),
		envDecodeValCase(t, "2s", time.Second*2),
	)
}
