package main

import (
	"log"
	"net/http"
	"os"

	"github.com/crusty0530/Dumpd/internal/config"
	"github.com/crusty0530/Dumpd/internal/health"
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

	router := http.NewServeMux()
	router.Handle("/health", health.Status{Message: "Server Ok!"})
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
