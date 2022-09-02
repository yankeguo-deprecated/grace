package gracejson

import (
	"path/filepath"
	"testing"

	"github.com/guoyk93/grace"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalJSON(t *testing.T) {
	type valType struct {
		Hello string `json:"hello"`
	}
	buf := []byte(`{"hello":"world"}`)
	val, err := UnmarshalJSON[valType](buf)
	require.NoError(t, err)
	require.Equal(t, "world", val.Hello)
}

func TestUnmarshalJSONFile(t *testing.T) {
	type valType struct {
		Hello string `json:"hello"`
	}
	val, err := UnmarshalJSONFile[valType](filepath.Join("testdata", "test.json"))
	require.NoError(t, err)
	require.Equal(t, "world", val.Hello)
}

func TestMarshalPretty(t *testing.T) {
	v := grace.Must(MarshalPretty(map[string]interface{}{"hello": "world"}))
	require.Equal(t, `{
  "hello": "world"
}`, string(v))
}
