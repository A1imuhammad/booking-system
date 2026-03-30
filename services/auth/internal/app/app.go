package app

import (
	"booking-system/services/auth/internal/app/server"
	"booking-system/services/auth/internal/config"
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
) *App{
	// TODO: implement database

	// TODO: implement app

	// TODO: implement server


	return &App{}
}
