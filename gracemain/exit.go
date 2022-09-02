package gracemain

import (
	"context"
	"log"
	"os"
)

var (
	IgnoredErrors = []error{
		context.Canceled,
	}
)

var (
	onExit = DefaultOnExit
	osExit = os.Exit
)

type ExitCoder interface {
	error
	ExitCode() int
}

func Exit(err *error) {
	if onExit != nil {
		onExit(err)
	}

	for _, ignored := range IgnoredErrors {
		if *err == ignored {
			*err = nil
		}
	}

	if *err == nil {
		osExit(0)
		return
	}

	if ec, ok := (*err).(ExitCoder); ok {
		osExit(ec.ExitCode())
	} else {
		osExit(1)
	}
}

func OnExit(fn func(err *error)) {
	onExit = fn
}

func DefaultOnExit(err *error) {
	if *err == nil {
		return
	}
	log.Println("exited with error:", (*err).Error())
}
