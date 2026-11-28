package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crusty0530/Dumpd/internal/config"
	"github.com/crusty0530/Dumpd/internal/health"
	"github.com/crusty0530/Dumpd/internal/platform/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error:", err)
	}

	cfg, err := config.Load(os.Getenv)
	if err != nil {
		log.Fatal("Invalid configuration variables: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pool, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}

	defer pool.Close()
	log.Printf("connected to database")

	router := http.NewServeMux()
	router.Handle("/health", health.Status{Message: "Server Ok!"})
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
