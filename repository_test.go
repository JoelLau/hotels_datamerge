package hotels

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	t.Parallel()

	t.Run("GetHotels", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name        string
			seed        []Hotel
			destination *int
			hotelIDs    []string
			want        []Hotel
		}{
			{name: "no data set returns empty", seed: nil, want: nil},
			{
				name:        "no filter returns all hotels",
				seed:        []Hotel{{ID: "a", DestinationID: 1}, {ID: "b", DestinationID: 1}, {ID: "c", DestinationID: 2}},
				destination: nil,
				want:        []Hotel{{ID: "a", DestinationID: 1}, {ID: "b", DestinationID: 1}, {ID: "c", DestinationID: 2}},
			},
			{
				name:        "filters by destination",
				seed:        []Hotel{{ID: "a", DestinationID: 1}, {ID: "b", DestinationID: 1}, {ID: "c", DestinationID: 2}},
				destination: new(1),
				want:        []Hotel{{ID: "a", DestinationID: 1}, {ID: "b", DestinationID: 1}},
			},
			{
				name:     "filters by hotel ids",
				seed:     []Hotel{{ID: "a", DestinationID: 1}, {ID: "b", DestinationID: 1}, {ID: "c", DestinationID: 2}},
				hotelIDs: []string{"a", "c"},
				want:     []Hotel{{ID: "a", DestinationID: 1}, {ID: "c", DestinationID: 2}},
			},
			{
				name:        "filters by destination and hotel ids",
				seed:        []Hotel{{ID: "a", DestinationID: 1}, {ID: "b", DestinationID: 1}, {ID: "c", DestinationID: 2}},
				destination: new(1),
				hotelIDs:    []string{"a", "c"},
				want:        []Hotel{{ID: "a", DestinationID: 1}},
			},
			{
				name:     "empty hotel ids slice means no filter",
				seed:     []Hotel{{ID: "a", DestinationID: 1}, {ID: "b", DestinationID: 1}, {ID: "c", DestinationID: 2}},
				hotelIDs: []string{},
				want:     []Hotel{{ID: "a", DestinationID: 1}, {ID: "b", DestinationID: 1}, {ID: "c", DestinationID: 2}},
			},
			{
				name:     "no matches returns empty",
				seed:     []Hotel{{ID: "a", DestinationID: 1}, {ID: "b", DestinationID: 1}, {ID: "c", DestinationID: 2}},
				hotelIDs: []string{"nonexistent"},
				want:     nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				repo := NewInMemoryRepository()
				if tt.seed != nil {
					repo.SetHotel(tt.seed)
				}

				hs, err := repo.GetHotels(context.Background(), tt.destination, tt.hotelIDs)

				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.want, hs)
			})
		}
	})

	t.Run("SetHotel", func(t *testing.T) {
		t.Parallel()

		t.Run("replaces previous data", func(t *testing.T) {
			t.Parallel()

			repo := NewInMemoryRepository()
			repo.SetHotel([]Hotel{{ID: "a", DestinationID: 1}})
			repo.SetHotel([]Hotel{{ID: "b", DestinationID: 2}})

			hs, err := repo.GetHotels(context.Background(), nil, nil)

			assert.NoError(t, err)
			assert.ElementsMatch(t, []Hotel{{ID: "b", DestinationID: 2}}, hs)
		})
	})

	t.Run("Ping", func(t *testing.T) {
		t.Parallel()

		t.Run("returns no error", func(t *testing.T) {
			t.Parallel()

			repo := NewInMemoryRepository()

			assert.NoError(t, repo.Ping(context.Background()))
		})
	})
}
