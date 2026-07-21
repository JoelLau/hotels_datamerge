package hotels

import "strings"

// PatagoniaHotel is the raw shape returned by the patagonia supplier endpoint.
type PatagoniaHotel struct {
	ID          string          `json:"id"`          // e.g. "iJhz"
	Destination int             `json:"destination"` // e.g. 5432
	Name        string          `json:"name"`        // e.g. "Beach Villas Singapore"
	Latitude    float64         `json:"lat"`         // e.g. 1.264751
	Longitude   float64         `json:"lng"`         // e.g. 103.824006
	Address     *string         `json:"address"`     // e.g. "8 Sentosa Gateway, Beach Villas, 098269"
	Info        *string         `json:"info"`        // e.g. "Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters."
	Amenities   []string        `json:"amenities"`
	Images      PatagoniaImages `json:"images"`
}

func (h *PatagoniaHotel) Hotel() *Hotel {
	if h == nil {
		return nil
	}

	address := ""
	if h.Address != nil {
		address = strings.TrimSpace(*h.Address)
	}

	description := ""
	if h.Info != nil {
		description = strings.TrimSpace(*h.Info)
	}

	return &Hotel{
		ID:            strings.TrimSpace(h.ID),
		DestinationID: h.Destination,
		Name:          strings.TrimSpace(h.Name),
		Location: &Location{
			Latitude:  new(h.Latitude),
			Longitude: new(h.Longitude),
			Address:   ToNilIfEmpty(address),
			City:      nil,
			Country:   nil,
		},
		Description: description,
		Amenities:   NewAmenities(h.Amenities),
		Images: Images{
			Rooms: Transform(h.Images.Rooms, func(p PatagoniaImage) Image {
				return Image{Link: p.URL, Description: p.Description}
			}),
			Amenities: Transform(h.Images.Amenities, func(p PatagoniaImage) Image {
				return Image{Link: p.URL, Description: p.Description}
			}),
			// NOTE: patagonia doesn't provide site images
			Site: []Image{},
		},

		// NOTE: patagonia doesn't come with booking conditions
		BookingConditions: []string{},
	}
}

type PatagoniaImages struct {
	Rooms     []PatagoniaImage `json:"rooms"`
	Amenities []PatagoniaImage `json:"amenities"`
}

type PatagoniaImage struct {
	URL         string `json:"url"`         // e.g. "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg"
	Description string `json:"description"` // e.g. "Double room"
}
