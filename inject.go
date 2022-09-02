package grace

import (
	"context"
	"reflect"
)

type injectKey string

func InjectKey[T any](v *T) any {
	typ := reflect.TypeOf(v).Elem()
	return injectKey(typ.PkgPath() + "::" + typ.Name())
}

func Inject[T any](ctx context.Context, v T) context.Context {
	return context.WithValue(ctx, InjectKey(&v), v)
}

func Extract[T any](ctx context.Context) (out T, ok bool) {
	out, ok = ctx.Value(InjectKey(&out)).(T)
	return
}
