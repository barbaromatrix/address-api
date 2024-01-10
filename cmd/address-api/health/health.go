package health

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lockp111/go-easyzap"
)

func NewHealthServer() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	killsignal := <-interrupt
	switch killsignal {
	case os.Interrupt:
		easyzap.Info("got sigint signal... interrupt")

	case syscall.SIGTERM:
		easyzap.Info("got sigterm signal... interrupt")
	}

}
