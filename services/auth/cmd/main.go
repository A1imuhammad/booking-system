package main

import (
	"booking-system/services/auth/internal/config"
	"fmt"
)

func main() {
	// TODO: implement config
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: implement logger

	// TODO: implement server

	// TODO: graceful shutdown
}

func setupLogger() {
	// TODO: implement logger
}
