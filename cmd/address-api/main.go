package main

import (
	"address-api/cmd/address-api/router"
	"address-api/internal/config"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/lockp111/go-easyzap"
)

func main() {
	ctx := context.Background()

	go router.Start(router.NewMetaApp(), config.GetConfig().MetaHost)

	if err := router.RunServer(); err != nil {
		easyzap.Error(ctx, err, "Error while initializing gRPC server")
	}

	// gracefully shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	switch <-signalChannel {
	case os.Interrupt:
		easyzap.Infof("Received SIGINT, stopping...")
	case syscall.SIGTERM:
		easyzap.Infof("Received SIGTERM, stopping...")
	}
}
