package app

import (
	"booking-system/services/auth/internal/app/server"
	"booking-system/services/auth/internal/config"
	storagepg "booking-system/services/auth/internal/storage/postgres"
	"context"
	"fmt"
	"log/slog"
)

type App struct {
	GRPCSrv *server.Server
}

func New(
	log *slog.Logger,
	postgres config.Postgres,
	jwt config.JWT,
	gRPC config.GRPCConfig,
) *App {
	// TODO: implement database
	storage, err := storagepg .New(context.Background(),postgres)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage)
	// TODO: implement app

	// TODO: implement server

	return &App{}
}
