package gracehtml

import (
	"bytes"
	"io/fs"
	"os"
	"testing"

	"github.com/guoyk93/grace"
	"github.com/stretchr/testify/require"
)

func TestNewTemplate(t *testing.T) {
	tpl, err := LoadTemplate(LoadTemplateOptions{
		Filesystems: []fs.FS{
			os.DirFS("testdata"),
		},
		Extensions: []string{
			".gohtml",
			".html",
		},
	})
	require.NoError(t, err)
	buf := &bytes.Buffer{}
	err = tpl.ExecuteTemplate(buf, "bbb/ccc", grace.M{"A": "B"})
	require.NoError(t, err)
	require.Equal(t, "<div>cB</div>", buf.String())
}
