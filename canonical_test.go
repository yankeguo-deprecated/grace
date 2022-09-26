package grace

import (
	"context"
)

var (
	_ Executor10[int] = Func10[int](func(i1 int) {})
	_ Executor01[int] = Func01[int](func() (o1 int) {
		return
	})
	_ Executor11[int, int] = Func11[int, int](func(i1 int) (o1 int) {
		return
	})
	_ Executor12[int, int, int] = Func12[int, int, int](func(i1 int) (o1 int, o2 int) {
		return
	})
	_ Executor21[int, int, int] = Func21[int, int, int](func(i1 int, i2 int) (o1 int) {
		return
	})
	_ Executor22[int, int, int, int] = Func22[int, int, int, int](func(i1 int, i2 int) (o1 int, o2 int) {
		return
	})
)

var (
	_ Task        = TaskFunc(func() (err error) { return })
	_ ContextTask = ContextTaskFunc(func(ctx context.Context) (err error) { return })
)
