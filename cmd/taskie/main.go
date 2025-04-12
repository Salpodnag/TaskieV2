package main

import (
	"AvitoFlats/cfg"
	"log/slog"
)

func main() {
	cfg, err := cfg.LoadConfig()
	if err != nil {
		slog.Error("er:", cfg)
	}
}
