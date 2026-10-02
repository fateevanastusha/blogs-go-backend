package main

import (
	"log"
	"os"

	"github.com/fateevanastusha/blogs-go-backend/internal/server"
)

func main() {
	address := ":3005"
	if port := os.Getenv("PORT"); port != "" {
		address = ":" + port
	}

	if err := server.Start(address); err != nil {
		log.Fatal(err)
	}
}
