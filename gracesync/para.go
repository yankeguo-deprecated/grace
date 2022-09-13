package gracesync

import (
	"context"
	"errors"
	"sync"

	"github.com/guoyk93/grace"
)

func DoPara[T any](ctx context.Context, concurrency int, vs []T, fn grace.Func21[context.Context, T, error]) (err error) {
	if concurrency < 1 {
		panic(errors.New("DoPara: invalid argument 'concurrency', must > 0"))
	}

	ch := make(chan struct{}, concurrency)
	for i := 0; i < concurrency; i++ {
		ch <- struct{}{}
	}
	wg := &sync.WaitGroup{}

	eg := grace.NewErrorGroup()
	for _, _v := range vs {
		v := _v
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-ch
			defer func() {
				ch <- struct{}{}
			}()
			eg.Add(fn(ctx, v))
		}()
	}
	wg.Wait()
	err = eg.Unwrap()
	return
}
