package gracemain

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App interface {
	Setup() (err error)
	Run() (err error)
	Shutdown() (err error)
}

func Run(app App) (err error) {
	if err = app.Setup(); err != nil {
		_ = app.Shutdown()
		return
	}

	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		chErr <- app.Run()
	}()

	select {
	case err = <-chErr:
		return
	case sig := <-chSig:
		log.Println("signal caught:", sig.String())
	}

	if err = app.Shutdown(); err != nil {
		return
	}
	return
}

func Main(app App) {
	var err error
	defer Exit(&err)

	err = Run(app)
	return
}
