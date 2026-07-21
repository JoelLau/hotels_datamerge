package hotels

import (
	"encoding/json"
	"strings"
)

type AcmeHotel struct {
	ID            string   `json:"Id"`            // e.g. "iJhz"
	DestinationID int      `json:"DestinationId"` // e.g. 5432
	Name          string   `json:"Name"`          // e.g. "Beach Villas Singapore"
	Latitude      *float64 `json:"Latitude"`      // e.g. 1.264751
	Longitude     *float64 `json:"Longitude"`     // e.g. 103.824006
	Address       string   `json:"Address"`       // e.g. " 8 Sentosa Gateway, Beach Villas "
	City          string   `json:"City"`          // e.g. "Singapore"
	Country       string   `json:"Country"`       // e.g. "SG"
	PostalCode    string   `json:"PostalCode"`    // e.g. "098269"
	Description   string   `json:"Description"`   // e.g. "  This 5 star hotel is located on the coastline of Singapore."
	Facilities    []string `json:"Facilities"`
}

func (h *AcmeHotel) Hotel() *Hotel {
	if h == nil {
		return nil
	}

	return &Hotel{
		ID:            strings.TrimSpace(h.ID),
		DestinationID: h.DestinationID,
		Name:          strings.TrimSpace(h.Name),
		Location: Location{
			// TODO: writing to lat/lng should be atomic - set both or neither
			Latitude:  h.Latitude,
			Longitude: h.Longitude,

			// TODO: nil on empty string
			Address: strings.TrimSpace(h.Address),
			City:    new(strings.TrimSpace(h.City)),
			Country: new(strings.TrimSpace(h.Country)),
		},
		Description: strings.TrimSpace(h.Description),
		Amenities:   NewAmenities(h.Facilities),

		// NOTE: acme doesn't provide images
		Images: Images{Rooms: []Image{}, Site: []Image{}, Amenities: []Image{}},

		// NOTE: acme doesn't come with booking conditions
		BookingConditions: []string{},
	}
}

// we need to override the default unmarshalling logic because
// lat / lng can arrive in 3 different data types (string, null, float)
func (h *AcmeHotel) UnmarshalJSON(data []byte) error {
	type alias AcmeHotel
	aux := struct {
		Latitude  json.RawMessage `json:"Latitude"`
		Longitude json.RawMessage `json:"Longitude"`
		*alias
	}{
		alias: (*alias)(h),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	lat, err := parseAcmeCoordinate(aux.Latitude)
	if err != nil {
		return err
	}
	h.Latitude = lat

	lng, err := parseAcmeCoordinate(aux.Longitude)
	if err != nil {
		return err
	}
	h.Longitude = lng

	return nil
}

func parseAcmeCoordinate(data json.RawMessage) (*float64, error) {
	if len(data) == 0 || string(data) == "null" || string(data) == `""` {
		return nil, nil
	}

	var v float64
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
