package graceyaml

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestUnmarshalYAMLFile(t *testing.T) {
	type Data struct {
		A string `yaml:"a"`
		C struct {
			D string `yaml:"d"`
		} `yaml:"c"`
	}

	data, err := UnmarshalYAMLFile[Data](filepath.Join("testdata", "test.yaml"))
	require.NoError(t, err)
	require.Equal(t, "b", data.A)
	require.Equal(t, "e", data.C.D)
}
