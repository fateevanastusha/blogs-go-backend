package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/fateevanastusha/blogs-go-backend/internal/config"
	"github.com/fateevanastusha/blogs-go-backend/internal/logger"
	"github.com/fateevanastusha/blogs-go-backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	l := logger.New(cfg.LogLevel)
	slog.SetDefault(l)

	if err := server.Start(cfg, l); err != nil {
		l.Error("server stopped", "err", err)
		os.Exit(1)
	}
}
