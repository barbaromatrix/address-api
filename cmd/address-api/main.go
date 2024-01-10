package main

import (
	"address-api/internal/config"
	"fmt"
)

func main() {
	cfg := config.GetConfig()

	fmt.Println(cfg)
}
