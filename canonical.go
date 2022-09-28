package grace

import "context"

type Executor10[I1 any] interface {
	Do(i1 I1)
}

type Executor01[O1 any] interface {
	Do() (o1 O1)
}

type Executor02[O1 any, O2 any] interface {
	Do() (o1 O1, o2 O2)
}

type Executor11[I1 any, O1 any] interface {
	Do(i1 I1) (o1 O1)
}

type Executor12[I1 any, O1 any, O2 any] interface {
	Do(i1 I1) (o1 O1, o2 O2)
}

type Executor21[I1 any, I2 any, O1 any] interface {
	Do(i1 I1, i2 I2) (o1 O1)
}

type Executor22[I1 any, I2 any, O1 any, O2 any] interface {
	Do(i1 I1, i2 I2) (o1 O1, o2 O2)
}

type Func10[I1 any] func(i1 I1)

func (fn Func10[I1]) Do(i1 I1) {
	fn(i1)
}

type Func01[O1 any] func() (o1 O1)

func (fn Func01[O1]) Do() O1 {
	return fn()
}

type Func02[O1 any, O2 any] func() (o1 O1, o2 O2)

func (fn Func02[O1, O2]) Do() (O1, O2) {
	return fn()
}

type Func11[I1 any, O1 any] func(i1 I1) (o1 O1)

func (fn Func11[I1, O1]) Do(i1 I1) O1 {
	return fn(i1)
}

type Func12[I1 any, O1 any, O2 any] func(i1 I1) (o1 O1, o2 O2)

func (fn Func12[I1, O1, O2]) Do(i1 I1) (O1, O2) {
	return fn(i1)
}

type Func21[I1 any, I2 any, O1 any] func(i1 I1, i2 I2) (o1 O1)

func (fn Func21[I1, I2, O1]) Do(i1 I1, i2 I2) O1 {
	return fn(i1, i2)
}

type Func22[I1 any, I2 any, O1 any, O2 any] func(i1 I1, i2 I2) (o1 O1, o2 O2)

func (fn Func22[I1, I2, O1, O2]) Do(i1 I1, i2 I2) (O1, O2) {
	return fn(i1, i2)
}

type Task = Executor01[error]

type TaskFunc = Func01[error]

type ContextTask = Executor11[context.Context, error]

type ContextTaskFunc = Func11[context.Context, error]
