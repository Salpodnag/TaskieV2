package main

import (
	"AvitoFlats/cfg"
	"log/slog"
	"net/http"
)

func main() {
	config, err := cfg.LoadConfig()
	if err != nil {
		slog.Error("er:", config)
	}
	db, err := cfg.NewClient(config)
	if err != nil {
		slog.Error("Failed to connect to database:", "error", err)
	}
	if err := http.ListenAndServe(config.Server.ServerPort, nil); err != nil {
		slog.Error("Server failed:", "error", err, db)
	}
}
