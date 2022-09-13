package grace

import "context"

func Do(tasks ...Task) (err error) {
	for _, task := range tasks {
		if err = task.Do(); err != nil {
			return
		}
	}
	return
}

func DoContext(ctx context.Context, tasks ...ContextTask) (err error) {
	for _, task := range tasks {
		if err = task.Do(ctx); err != nil {
			return
		}
	}
	return
}

func MustDo(tasks ...Task) {
	Must0(Do(tasks...))
}

func MustDoContext(ctx context.Context, tasks ...ContextTask) {
	Must0(DoContext(ctx, tasks...))
}
