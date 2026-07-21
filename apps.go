package hotels

import (
	"context"
	"log/slog"
)

type App interface {
	Run(context.Context) error
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) SetHotel(hotels []Hotel) {
}

func (repo *Repository) GetHotelByID(id string) (Hotel, bool) {
	return Hotel{}, false
}

func (repo *Repository) GetHotelByDestinationID(destination int) (Hotel, bool) {
	return Hotel{}, false
}

type Repo interface {
	SetHotel(hotels []Hotel)
	GetHotelByID(id string) (Hotel, bool)
	GetHotelByDestinationID(destination int) (Hotel, bool)
}

// writes data to the repository
//
// TODO: fetch data from suppliers
// TODO: fetch data at intervals
// TODO: write data to repository
type Fetcher struct {
	repo *Repo
}

func NewFetcher(r Repo) *Fetcher {
	return &Fetcher{}
}

func (f *Fetcher) Run(ctx context.Context) error {
	slog.InfoContext(ctx, "starting fetcher")
	<-ctx.Done()
	slog.InfoContext(ctx, "exiting fetcher")
	return nil
}

// serves repo data via REST endpoint
//
// TODO: generate from OpenAPI spec
// TODO: single REST endpoint
// TODO: get / validate query param
type Server struct {
	repo *Repo
}

func NewServer(r Repo) *Server {
	return &Server{}
}

func (f *Server) Run(ctx context.Context) error {
	slog.InfoContext(ctx, "starting server")
	<-ctx.Done()
	slog.InfoContext(ctx, "exiting server")
	return nil
}
