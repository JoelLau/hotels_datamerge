package hotels_test

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	hotels "hotels_data_merge"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	repo := hotels.NewInMemoryRepository()
	repo.SetHotel([]hotels.Hotel{
		{ID: "iJhz", DestinationID: 5432, Name: "Beach Villas Singapore"},
		{ID: "SjyX", DestinationID: 5432, Name: "InterContinental Singapore Robertson Quay"},
		{ID: "f8c9", DestinationID: 1122, Name: "Mystery Hotel"},
	})

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := hotels.NewServer(repo, logger, "127.0.0.1:0")

	ts := httptest.NewServer(server.Handler())
	t.Cleanup(ts.Close)

	t.Run("GET /livez reports healthy", func(t *testing.T) {
		t.Parallel()

		resp, err := http.Get(ts.URL + "/livez")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("GET /readyz reports healthy once repo is set", func(t *testing.T) {
		t.Parallel()

		resp, err := http.Get(ts.URL + "/readyz")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("GET /api/v1/hotels returns all seeded hotels", func(t *testing.T) {
		t.Parallel()

		resp, err := http.Get(ts.URL + "/api/v1/hotels")
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var got []hotels.Hotel
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))
		assert.Len(t, got, 3)
	})

	t.Run("GET /api/v1/hotels filters by destination", func(t *testing.T) {
		t.Parallel()

		resp, err := http.Get(ts.URL + "/api/v1/hotels?destination=5432")
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var got []hotels.Hotel
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))
		assert.Len(t, got, 2)
		for _, h := range got {
			assert.Equal(t, 5432, h.DestinationID)
		}
	})

	t.Run("GET /api/v1/hotels filters by hotel ids", func(t *testing.T) {
		t.Parallel()

		resp, err := http.Get(ts.URL + "/api/v1/hotels?hotels=iJhz,f8c9")
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var got []hotels.Hotel
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))

		gotIDs := make([]string, len(got))
		for i, h := range got {
			gotIDs[i] = h.ID
		}
		assert.ElementsMatch(t, []string{"iJhz", "f8c9"}, gotIDs)
	})
}
