package grace

import "context"

func Do(fns ...TaskFunc) (err error) {
	for _, task := range fns {
		if err = task(); err != nil {
			return
		}
	}
	return
}

func DoContext(ctx context.Context, fns ...ContextTaskFunc) (err error) {
	for _, fn := range fns {
		if err = fn(ctx); err != nil {
			return
		}
	}
	return
}

func MustDo(fns ...TaskFunc) {
	Must0(Do(fns...))
}

func MustDoContext(ctx context.Context, fns ...ContextTaskFunc) {
	Must0(DoContext(ctx, fns...))
}
