package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/barbaromatrix/address-api/cmd/address-api/router"
	"github.com/barbaromatrix/address-api/internal/config"

	"github.com/lockp111/go-easyzap"
)

func main() {
	ctx := context.Background()
	config.Init()
	cfg := config.GetConfig()

	go router.Start(router.NewMetaApp(), cfg.HealthPort)

	if err := router.RunServer(cfg.ServerHost); err != nil {
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
