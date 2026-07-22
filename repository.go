package hotels

import "context"

type InMemoryRepository struct {
	// NOTE: data can't be denormalized because we'll apply a variable number of filters
	// NOTE: nil means no data set yet (repo not ready)
	hotels []Hotel
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		hotels: nil,
	}
}

func (repo *InMemoryRepository) SetHotel(hotels []Hotel) {
	if len(hotels) <= 0 {
		repo.hotels = make([]Hotel, 0)
		return
	}

	repo.hotels = hotels
}

func (repo *InMemoryRepository) GetHotels(ctx context.Context, destination *int, hotelIDs []string) ([]Hotel, error) {
	idSet := make(map[string]struct{})
	if len(hotelIDs) > 0 {
		idSet = make(map[string]struct{}, len(hotelIDs))
		for _, id := range hotelIDs {
			idSet[id] = struct{}{}
		}
	}

	// NOTE: consider mapping the filters into a chain of functions
	result := repo.hotels

	// filter by destination
	result = Filter(result, func(h Hotel) bool { return destination == nil || h.DestinationID == *destination })

	// filter by hotel ids
	result = Filter(result, func(h Hotel) bool {
		if len(idSet) == 0 {
			return true
		}
		_, ok := idSet[h.ID]
		return ok
	})

	return result, nil
}

func (repo *InMemoryRepository) Ping(ctx context.Context) error {
	return nil
}
