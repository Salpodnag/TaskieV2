package utils

import (
	"log/slog"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuid, err := uuid.NewV7()
	if err != nil {
		slog.Error("Failed to generate UUID:", "error", err)
		return ""
	}
	return uuid.String()
}
