package hotels

import "strings"

type Hotel struct {
	ID                string    `json:"id"`             // e.g. "iJhz"
	DestinationID     int       `json:"destination_id"` // e.g. 5432
	Name              string    `json:"name"`           // e.g. "Beach Villas Singapore"
	Location          Location  `json:"location"`       //
	Description       string    `json:"description"`    // e.g. "Surrounded by tropical gardens, ..." (truncated)
	Amenities         Amenities `json:"amenities"`
	Images            Images    `json:"images"`
	BookingConditions []string  `json:"booking_conditions"`
}

type Location struct {
	Latitude  float64 `json:"lat"`     // e.g. 1.264751
	Longitude float64 `json:"lng"`     // e.g. 103.824006
	Address   string  `json:"address"` // e.g. "8 Sentosa Gateway, Beach Villas, 098269"
	City      string  `json:"city"`    // e.g. "Singapore"
	Country   string  `json:"country"` // e.g. "Singapore"
}

type Amenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

// TODO: categorize amenities
// WARN: feature incomplete! everythign is "general" now
func NewAmenities(raw []string) Amenities {
	amen := Amenities{
		General: make([]string, 0),
		Room:    make([]string, 0),
	}

	for _, r := range raw {
		s := strings.ToLower(strings.TrimSpace(r))
		if len(s) > 1 {
			amen.General = append(amen.General, s)
		}
	}

	return amen
}

type Images struct {
	Rooms     []Image `json:"rooms"`
	Site      []Image `json:"site"`
	Amenities []Image `json:"amenities"`
}

type Image struct {
	Link        string `json:"link"`        // e.g. "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg"
	Description string `json:"description"` // e.g. "Double room"
}
