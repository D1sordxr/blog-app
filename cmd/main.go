package main

import (
	"BlogWebApp/internal/config"
	"BlogWebApp/internal/storage/postgres"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting blog app", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := postgres.BuildConnection(&cfg.DBConfig)
	if err != nil {
		log.Error("Failed to connect database:", err)
	}
	_ = storage

	//router := gin.Default()
	//api.Setup(router)
	//_ = router.Run(":" + cfg.HTTPServerConfig.Port)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
