package main

import (
	"address-api/cmd/address-api/health"
	"address-api/cmd/address-api/server"
	"address-api/internal/config"
	"address-api/internal/model"
	"address-api/internal/pgk/client"
	"fmt"

	"github.com/lockp111/go-easyzap"
)

func main() {
	cfg := config.GetConfig()

	server := server.NewServer(cfg.AppName)

	aweSomeClient := client.NewAwesomeClient(cfg)

	setupHttpServer(server, cfg)

	easyzap.Info("start application")

	health.NewHealthServer()

	fmt.Println(aweSomeClient)
}

func setupHttpServer(server *server.Server, config model.Config) {
	go func() {
		server.Start(config)
	}()
}
