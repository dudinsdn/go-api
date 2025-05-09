package main

import (
	"log"
	"my/go-api/config"
	"net/http"
)

func main() {
	cfg := config.Load()

	log.Printf("Server running at %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, nil))
}
