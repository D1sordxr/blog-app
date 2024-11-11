package main

import (
	"BlogWebApp/internal/config"
	api "BlogWebApp/internal/http-server"
	logger "BlogWebApp/internal/lib/logger/sl"
	db "BlogWebApp/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func main() {

	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log.Info("starting blog app", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := db.BuildConnection(&cfg.DBConfig)
	if err != nil {
		log.Error("Failed to connect database:", err)
	}
	_ = storage

	router := gin.Default()
	api.Setup(router)
	// TODO: middleware

	port := cfg.HTTPServerConfig.Port
	if err = router.Run(":" + port); err != nil {
		log.Error("failed to run http server")
	}

}
