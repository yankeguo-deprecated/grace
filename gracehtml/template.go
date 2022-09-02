package gracehtml

import (
	"errors"
	"html/template"
	"io/fs"
	"strings"

	"github.com/guoyk93/grace"
)

const DefaultExtension = ".gohtml"

type LoadTemplateOptions struct {
	Filesystems []fs.FS
	Extensions  []string
}

func LoadTemplate(opts LoadTemplateOptions) (tpl *template.Template, err error) {
	if len(opts.Filesystems) == 0 {
		err = errors.New("gracehtml.LoadTemplate: missing opts.Filesystems")
		return
	}
	if len(opts.Extensions) == 0 {
		opts.Extensions = []string{DefaultExtension}
	}
	tpl = template.New("__NOT__USED__")
	for _, filesystem := range opts.Filesystems {
		for i := 0; i < 5; i++ {
			for _, ext := range opts.Extensions {
				pattern := strings.Join(grace.Repeat(i+1, "*"), "/") + ext
				var matches []string
				if matches, err = fs.Glob(filesystem, pattern); err != nil {
					return
				}
				for _, match := range matches {
					name := strings.TrimSuffix(match, ext)
					var buf []byte
					if buf, err = fs.ReadFile(filesystem, match); err != nil {
						return
					}
					if tpl, err = tpl.New(name).Parse(string(buf)); err != nil {
						return
					}
				}
			}
		}
	}
	return
}
