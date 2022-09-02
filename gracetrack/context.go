package gracetrack

import "context"

type contextKeyType int

const contextKey = contextKeyType(0)

func Init(ctx context.Context) context.Context {
	if v := ctx.Value(contextKey); v != nil {
		if _, ok := v.(*Track); ok {
			return ctx
		}
	}
	return context.WithValue(ctx, contextKey, New())
}

func Extract(ctx context.Context) *Track {
	return ctx.Value(contextKey).(*Track)
}

func Group(ctx context.Context, key string) *TrackGroup {
	return Extract(ctx).Group(key)
}

func DumpPlain(ctx context.Context) []string {
	return Extract(ctx).DumpPlain()
}
