package hotels

import "encoding/json"

// AcmeHotel is the raw shape returned by the acme supplier endpoint.
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

// UnmarshalJSON tolerates acme's inconsistent encoding of missing coordinates,
// which show up as a JSON number, null, or an empty string depending on the record.
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
