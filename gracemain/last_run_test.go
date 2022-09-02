package gracemain

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWriteLastRun(t *testing.T) {
	defer os.RemoveAll(filepath.Join("testdata", "last-run.txt"))
	err := WriteLastRun("testdata")
	require.NoError(t, err)
	buf, err := os.ReadFile(filepath.Join("testdata", "last-run.txt"))
	require.NoError(t, err)
	d, err := time.Parse(time.RFC3339, string(buf))
	require.NoError(t, err)
	require.Less(t, time.Now().Sub(d), time.Second*5)
}
