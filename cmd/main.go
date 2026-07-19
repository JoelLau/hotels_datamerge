package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	slog.SetDefault(NewSlogger())

	slog.InfoContext(ctx, "starting")

	slog.InfoContext(ctx, "ending")
}

type App interface {
	Run(context.Context) error
}

// TODO: add config options
func NewSlogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stderr,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			}),
	)
}
