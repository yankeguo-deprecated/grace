package gracesync

import (
	"context"
)

func LoopFeed[T any](ctx context.Context, errs chan error, fnC func() (val T, err error), fnF func(ctx context.Context, val T)) {
	for {
		if ctx.Err() != nil {
			return
		}
		if val, err := fnC(); err != nil {
			select {
			case errs <- err:
			case <-ctx.Done():
			}
			return
		} else {
			go fnF(ctx, val)
		}
	}
}
