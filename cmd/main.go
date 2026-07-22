package main

import (
	"context"
	hotels "hotels_data_merge"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	slog.SetDefault(NewSlogger())

	slog.InfoContext(ctx, "starting")

	repo := hotels.NewRepository()

	fetcher := hotels.NewFetcher(repo)
	server := hotels.NewServer(repo, slog.Default())

	go func() {
		// TODO: retry on failure
		if err := fetcher.Run(ctx); err != nil {
			slog.ErrorContext(ctx, "fatal error on fetcher", slog.Any("error", err))
			cancel() // attempt exit if 1 "app" fails
			return
		}
	}()

	go func() {
		// TODO: retry on failure
		if err := server.Run(ctx); err != nil {
			slog.ErrorContext(ctx, "fatal error on server", slog.Any("error", err))
			cancel() // attempt exit if 1 "app" fails
			return
		}
	}()

	<-ctx.Done()
	slog.InfoContext(ctx, "ending")
}

type App interface {
	Run(context.Context) error
}

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
