package hotels

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

type Repo interface {
	SetHotel(hotels []Hotel)
	GetHotels(ctx context.Context, destination *int, hotelIDs []string) ([]Hotel, error)
	Ping(ctx context.Context) error
}

type Fetcher struct {
	repo            Repo
	producers       []DataProducer
	interval        time.Duration
	producerTimeout time.Duration
}

func NewFetcher(r Repo, producers []DataProducer, interval, producerTimeout time.Duration) *Fetcher {
	return &Fetcher{
		repo:            r,
		producers:       producers,
		interval:        interval,
		producerTimeout: producerTimeout,
	}
}

func (f *Fetcher) Ping(ctx context.Context) error {
	return nil
}

func (f *Fetcher) Run(ctx context.Context) error {
	slog.InfoContext(ctx, "starting fetcher")

	ticker := time.NewTicker(f.interval)
	defer ticker.Stop()

	for {
		f.fetchAll(ctx)

		select {
		case <-ctx.Done():
			slog.InfoContext(ctx, "exiting fetcher")
			return nil
		case <-ticker.C:
		}
	}
}

func (f *Fetcher) fetchAll(ctx context.Context) {
	var mu sync.Mutex
	var all []Hotel

	var wg sync.WaitGroup
	for _, p := range f.producers {
		wg.Add(1)
		go func(p DataProducer) {
			defer wg.Done()

			pctx, cancel := context.WithTimeout(ctx, f.producerTimeout)
			defer cancel()

			hs, err := p.Fetch(pctx)
			if err != nil {
				slog.ErrorContext(ctx, "producer fetch failed", slog.String("producer", p.Name()), slog.Any("error", err))
				return
			}

			mu.Lock()
			all = append(all, hs...)
			mu.Unlock()
		}(p)
	}
	wg.Wait()

	f.repo.SetHotel(GroupAndMerge(all))
}
