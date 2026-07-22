package hotels

import (
	"context"
	"fmt"
	"hotels_data_merge/gen/api"
	"hotels_data_merge/internal/swaggerdocs"
	"log/slog"
	"net/http"
	"strings"
	"time"

	_ "embed"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/render"
)

//go:embed openapi.yaml
var openapiSpec []byte

// serves repo data via REST endpoint
type Server struct {
	repo   Repo
	server *http.Server
}

func NewServer(r Repo, logger *slog.Logger, addr string) *Server {
	s := &Server{repo: r}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(httplog.RequestLogger(&httplog.Logger{Logger: logger}))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))
	api.HandlerFromMux(s, router)

	router.Get("/openapi.yaml", swaggerdocs.SpecHandler(openapiSpec, "application/yaml"))
	router.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusMovedPermanently)
	})
	router.Mount("/docs/", swaggerdocs.Handler("/docs/"))
	s.server = &http.Server{Addr: addr, Handler: router}

	return s
}

func (s *Server) Handler() http.Handler {
	return s.server.Handler
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

func (s *Server) GetLivez(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, api.StatusResponse{Status: api.Healthy})
}

func (s *Server) GetReadyz(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, api.StatusResponse{Status: api.Healthy})
}

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
		hotelIDs = strings.Split(*params.Hotels, ",")
		hotelIDs = Transform(hotelIDs, strings.TrimSpace)
	}

	hs, err := s.repo.GetHotels(r.Context(), destination, hotelIDs)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to fetch hotels", slog.Any("error", err))
		title := "failed to fetch hotels"
		status := http.StatusInternalServerError
		render.Status(r, status)
		render.JSON(w, r, api.ErrorResponse{Title: &title, Status: &status})
		return
	}

	render.JSON(w, r, hs)
}
