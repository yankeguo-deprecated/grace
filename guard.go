package grace

import (
	"context"
	"errors"
	"fmt"
)

func Guard(err *error) {
	if r := recover(); r != nil {
		if re, ok := r.(error); ok {
			*err = re
		} else {
			*err = errors.New(fmt.Sprintf("panic: %v", r))
		}
	}
}

func MustContext(ctx context.Context) {
	if ctx.Err() != nil {
		panic(ctx.Err())
	}
}

func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](v T, err error) T {
	if err == nil {
		return v
	} else {
		panic(err)
	}
}

func Must2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err == nil {
		return v1, v2
	} else {
		panic(err)
	}
}

func Must3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	if err == nil {
		return v1, v2, v3
	} else {
		panic(err)
	}
}

func Must4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	if err == nil {
		return v1, v2, v3, v4
	} else {
		panic(err)
	}
}
