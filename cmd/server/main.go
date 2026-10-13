package main

import (
	"log"

	"github.com/fateevanastusha/blogs-go-backend/internal/config"
	"github.com/fateevanastusha/blogs-go-backend/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
