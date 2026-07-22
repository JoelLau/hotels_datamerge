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

	cfg := hotels.NewConfig()

	repo := hotels.NewInMemoryRepository()

	httpClient := hotels.NewHTTPClient(cfg.HTTPTimeout)
	producers := []hotels.DataProducer{
		hotels.NewAcmeProducer(cfg.AcmeSupplierURL, httpClient),
		hotels.NewPatagoniaProducer(cfg.PatagoniaSupplierURL, httpClient),
		hotels.NewPaperfliesProducer(cfg.PaperfliesSupplierURL, httpClient),
	}

	fetcher := hotels.NewFetcher(repo, producers, cfg.FetchInterval, cfg.ProducerTimeout)
	server := hotels.NewServer(repo, slog.Default(), cfg.ServerAddr)

	apps := []App{fetcher, server}
	for _, app := range apps {
		go func(app App) {
			defer cancel() // kill ALL if any apps exit
			if err := app.Run(ctx); err != nil {
				slog.ErrorContext(ctx, "fatal error on app", slog.Any("error", err))
			}
		}(app)
	}

	<-ctx.Done()
	slog.InfoContext(ctx, "ending")
}

type App interface {
	Run(context.Context) error
}

func NewSlogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stderr,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			}),
	)
}
