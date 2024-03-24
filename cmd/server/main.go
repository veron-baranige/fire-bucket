package main

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/veron-baranige/fire-bucket/internal/config"
)

func main() {
	setupLogger()

	if err := config.LoadEnv("."); err != nil {
		slog.Error("Failed to load environment variables", "err", err)
	}
	slog.Info("Loaded configurations")
}

func setupLogger() {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)
}