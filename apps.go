package hotels

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

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

type Repo interface {
	SetHotel(hotels []Hotel)
	GetHotels(ctx context.Context, destination *int, hotelIDs []string) ([]Hotel, error)
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

func NewServer(r Repo) *Server {
	s := &Server{repo: r}

	router := chi.NewRouter()
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
	writeJSON(w, http.StatusOK, api.StatusResponse{Status: api.Healthy})
}

// GetReadyz implements api.ServerInterface.
func (s *Server) GetReadyz(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, api.StatusResponse{Status: api.Healthy})
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
		writeJSON(w, http.StatusInternalServerError, map[string]any{
			"title":  "failed to fetch hotels",
			"status": http.StatusInternalServerError,
		})
		return
	}

	writeJSON(w, http.StatusOK, hs)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
