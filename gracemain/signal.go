package gracemain

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func WithSignalCancel(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	if len(signals) == 0 {
		signals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	ctx, cancel := context.WithCancel(ctx)

	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-chSig
		log.Println("signal caught:", sig.String())
		cancel()
	}()

	return ctx, cancel
}
