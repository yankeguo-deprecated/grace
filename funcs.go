package grace

import "context"

type Func01[O1 any] func() (o1 O1)

type Func02[O1 any, O2 any] func() (o1 O1, o2 O2)

type Func10[I1 any] func(i1 I1)

type Func11[I1 any, O1 any] func(i1 I1) (o1 O1)

type Func12[I1 any, O1 any, O2 any] func(i1 I1) (o1 O1, o2 O2)

type Func21[I1 any, I2 any, O1 any] func(i1 I1, i2 I2) (o1 O1)

type Func22[I1 any, I2 any, O1 any, O2 any] func(i1 I1, i2 I2) (o1 O1, o2 O2)

type TaskFunc = Func01[error]

type ContextTaskFunc = Func11[context.Context, error]
