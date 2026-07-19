package hotels

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

type PatagoniaImages struct {
	Rooms     []PatagoniaImage `json:"rooms"`
	Amenities []PatagoniaImage `json:"amenities"`
}

type PatagoniaImage struct {
	URL         string `json:"url"`         // e.g. "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg"
	Description string `json:"description"` // e.g. "Double room"
}
