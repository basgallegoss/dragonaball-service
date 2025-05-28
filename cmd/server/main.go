package main

import (
	"log"
	"os"

	"github.com/basgallegoss/dragonball-service/internal/infrastructure/router"
)

func main() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN ")
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		log.Fatal("API_URL")
	}
	r, err := router.SetupRouter(dsn, apiURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
