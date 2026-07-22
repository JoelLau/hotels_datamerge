package hotels

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/render"

	"hotels_data_merge/gen/api"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) SetHotel(hotels []Hotel) {
}

func (repo *Repository) GetHotels(ctx context.Context, destination *int, hotelIDs []string) ([]Hotel, error) {
	return nil, nil
}

func (repo *Repository) Ping(ctx context.Context) error {
	return nil
}

type Repo interface {
	SetHotel(hotels []Hotel)
	GetHotels(ctx context.Context, destination *int, hotelIDs []string) ([]Hotel, error)
	Ping(ctx context.Context) error
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
type Server struct {
	repo   Repo
	server *http.Server
}

func NewServer(r Repo, logger *slog.Logger) *Server {
	s := &Server{repo: r}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(httplog.RequestLogger(&httplog.Logger{Logger: logger}))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))
	api.HandlerFromMux(s, router)

	s.server = &http.Server{Addr: ":8080", Handler: router}

	return s
}

func (s *Server) Run(ctx context.Context) error {
	slog.InfoContext(ctx, "starting server", slog.String("addr", s.server.Addr))

	errCh := make(chan error, 1)
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
	}

	slog.InfoContext(ctx, "exiting server")
	return s.server.Shutdown(context.Background())
}

// GetLivez implements api.ServerInterface.
func (s *Server) GetLivez(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, api.StatusResponse{Status: api.Healthy})
}

// GetReadyz implements api.ServerInterface.
func (s *Server) GetReadyz(w http.ResponseWriter, r *http.Request) {
	if err := s.repo.Ping(r.Context()); err != nil {
		render.Status(r, http.StatusServiceUnavailable)
		render.JSON(w, r, api.StatusResponse{Status: api.Unhealthy})
		return
	}
	render.JSON(w, r, api.StatusResponse{Status: api.Healthy})
}

// GetApiV1Hotels implements api.ServerInterface.
func (s *Server) GetApiV1Hotels(w http.ResponseWriter, r *http.Request, params api.GetApiV1HotelsParams) {
	var destination *int
	if params.Destination != nil {
		var d int
		if _, err := fmt.Sscanf(*params.Destination, "%d", &d); err == nil {
			destination = &d
		}
	}

	var hotelIDs []string
	if params.Hotels != nil {
		hotelIDs = *params.Hotels
	}

	hs, err := s.repo.GetHotels(r.Context(), destination, hotelIDs)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]any{
			"title":  "failed to fetch hotels",
			"status": http.StatusInternalServerError,
		})
		return
	}

	render.JSON(w, r, hs)
}
