package gracelog

import (
	"bufio"
	"fmt"
	"github.com/guoyk93/grace"
	"io"
	"os"
)

type ProcLogChannel struct {
	FilePrefix    string
	File          io.WriteCloser
	ConsolePrefix string
	Console       io.Writer
}

func (pc *ProcLogChannel) Close() error {
	if pc.File != nil {
		return pc.File.Close()
	}
	return nil
}

func (pc *ProcLogChannel) Write(buf []byte) (n int, err error) {
	if pc.FilePrefix == "" {
		if _, err = pc.File.Write(buf); err != nil {
			return
		}
	} else {
		if _, err = pc.File.Write(append([]byte(pc.FilePrefix), buf...)); err != nil {
			return
		}
	}

	if pc.ConsolePrefix == "" {
		if _, err = pc.Console.Write(buf); err != nil {
			return
		}
	} else {
		if _, err = pc.Console.Write(append([]byte(pc.ConsolePrefix), buf...)); err != nil {
			return
		}
	}

	n = len(buf)

	return
}

func (pc *ProcLogChannel) ReadFrom(r io.Reader) (n int64, err error) {
	br := bufio.NewReader(r)
	for {
		var line []byte
		if line, err = br.ReadBytes('\n'); err == nil {
			_, _ = pc.Write(line)
			n += int64(len(line))
		} else {
			if len(line) != 0 {
				_, _ = pc.Write(append(line, '\n'))
				n += int64(len(line))
			}
			break
		}
	}
	return
}

type ProcLoggerOptions struct {
	RotatingFileOptions

	ConsoleOut io.Writer
	ConsoleErr io.Writer

	ConsolePrefix string
	FilePrefix    string
}

type ProcLogger struct {
	out *ProcLogChannel
	err *ProcLogChannel
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

	pl = &ProcLogger{
		out: &ProcLogChannel{
			ConsolePrefix: opts.ConsolePrefix,
			FilePrefix:    opts.FilePrefix,
		},
		err: &ProcLogChannel{
			ConsolePrefix: opts.ConsolePrefix,
			FilePrefix:    opts.FilePrefix,
		},
	}

	if pl.out.File, err = NewRotatingFile(RotatingFileOptions{
		Dir:          opts.Dir,
		Filename:     opts.Filename + ".out",
		MaxFileSize:  opts.MaxFileSize,
		MaxFileCount: opts.MaxFileCount,
	}); err != nil {
		return
	}

	if pl.err.File, err = NewRotatingFile(RotatingFileOptions{
		Dir:          opts.Dir,
		Filename:     opts.Filename + ".err",
		MaxFileSize:  opts.MaxFileSize,
		MaxFileCount: opts.MaxFileCount,
	}); err != nil {
		return
	}

	pl.out.Console = opts.ConsoleOut
	pl.err.Console = opts.ConsoleErr
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

func (pl *ProcLogger) StreamOut(r io.Reader) {
	_, _ = pl.out.ReadFrom(r)
}

func (pl *ProcLogger) StreamErr(r io.Reader) {
	_, _ = pl.err.ReadFrom(r)
}
