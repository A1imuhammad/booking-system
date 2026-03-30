package main

import (
	"booking-system/services/auth/internal/config"
	"log/slog"

	"go.uber.org/zap"
	slogzap "github.com/samber/slog-zap"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	
	log := setupLogger(cfg.Env)
	log.Info("starting application", slog.String("env", cfg.Env))	
	// TODO: implement server

	// TODO: graceful shutdown
}

func setupLogger(env string) *slog.Logger {
	var zapLogger *zap.Logger
	var err error

	switch env {
	case EnvLocal:
		zapLogger, err = zap.NewDevelopment()
	case EnvDev:
		zapLogger, err = zap.NewDevelopment()
	case EnvProd:
		zapLogger, err = zap.NewProduction()

	default:
		zapLogger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	handler := slogzap.Option{
		Logger: zapLogger,
		Level:  slog.LevelDebug,
	}.NewZapHandler()

	log := slog.New(handler).With(
		slog.String("service", "auth"),
	)

	return log
}
