package gracelog

import (
	"fmt"
	"github.com/guoyk93/grace"
	"io"
	"os"
)

type ProcLoggerOptions struct {
	RotatingFileOptions

	ConsoleOut io.Writer
	ConsoleErr io.Writer

	ConsolePrefix string
	FilePrefix    string
}

type ProcLogger struct {
	out Output
	err Output
}

func NewProcLogger(opts ProcLoggerOptions) (pl *ProcLogger, err error) {
	if opts.ConsoleOut == nil {
		opts.ConsoleOut = os.Stdout
	}
	if opts.ConsoleErr == nil {
		opts.ConsoleErr = os.Stderr
	}
	if opts.MaxFileSize == 0 {
		opts.MaxFileSize = 128 * 1024 * 1024
	}
	if opts.MaxFileCount == 0 {
		opts.MaxFileCount = 5
	}

	var fileOut io.WriteCloser
	if fileOut, err = NewRotatingFile(RotatingFileOptions{
		Dir:          opts.Dir,
		Filename:     opts.Filename + ".out",
		MaxFileSize:  opts.MaxFileSize,
		MaxFileCount: opts.MaxFileCount,
	}); err != nil {
		return
	}

	var fileErr io.WriteCloser
	if fileErr, err = NewRotatingFile(RotatingFileOptions{
		Dir:          opts.Dir,
		Filename:     opts.Filename + ".err",
		MaxFileSize:  opts.MaxFileSize,
		MaxFileCount: opts.MaxFileCount,
	}); err != nil {
		return
	}

	pl = &ProcLogger{
		out: MultiOutput(
			NewWriterOutput(fileOut, []byte(opts.FilePrefix), nil),
			NewWriterOutput(opts.ConsoleOut, []byte(opts.ConsolePrefix), nil),
		),
		err: MultiOutput(
			NewWriterOutput(fileErr, []byte(opts.FilePrefix), nil),
			NewWriterOutput(opts.ConsoleErr, []byte(opts.ConsolePrefix), nil),
		),
	}
	return
}

func (pl *ProcLogger) Close() error {
	eg := grace.NewErrorGroup()
	eg.Add(pl.out.Close())
	eg.Add(pl.err.Close())
	return eg.Unwrap()
}

func (pl *ProcLogger) Print(items ...interface{}) {
	_, _ = pl.out.Write(append([]byte(fmt.Sprint(items...)), '\n'))
}

func (pl *ProcLogger) Error(items ...interface{}) {
	_, _ = pl.err.Write(append([]byte(fmt.Sprint(items...)), '\n'))
}

func (pl *ProcLogger) Printf(pattern string, items ...interface{}) {
	_, _ = pl.out.Write(append([]byte(fmt.Sprintf(pattern, items...)), '\n'))
}

func (pl *ProcLogger) Errorf(pattern string, items ...interface{}) {
	_, _ = pl.err.Write(append([]byte(fmt.Sprintf(pattern, items...)), '\n'))
}

func (pl *ProcLogger) Out() Output {
	return pl.out
}

func (pl *ProcLogger) Err() Output {
	return pl.err
}
