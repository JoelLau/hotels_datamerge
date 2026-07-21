package hotels

import "strings"

type PaperfliesHotel struct {
	HotelID           string              `json:"hotel_id"`       // e.g. "iJhz"
	DestinationID     int                 `json:"destination_id"` // e.g. 5432
	HotelName         string              `json:"hotel_name"`     // e.g. "Beach Villas Singapore"
	Location          PaperfliesLocation  `json:"location"`       //
	Details           string              `json:"details"`        // e.g. "Surrounded by tropical gardens ..." (truncated for brevity)
	Amenities         PaperfliesAmenities `json:"amenities"`
	Images            PaperfliesImages    `json:"images"`
	BookingConditions []string            `json:"booking_conditions"`
}

func (h *PaperfliesHotel) Hotel() *Hotel {
	if h == nil {
		return nil
	}

	return &Hotel{
		ID:            strings.TrimSpace(h.HotelID),
		DestinationID: h.DestinationID,
		Name:          strings.TrimSpace(h.HotelName),
		Location: &Location{
			Latitude:  nil,
			Longitude: nil,
			Address:   ToNilIfEmpty(strings.TrimSpace(h.Location.Address)),
			City:      nil,
			Country:   ToNilIfEmpty(strings.TrimSpace(h.Location.Country)),
		},
		Description: strings.TrimSpace(h.Details),
		Amenities: Amenities{
			General: Filter(
				Transform(h.Amenities.General, func(s string) string { return strings.ToLower(strings.TrimSpace(s)) }),
				func(s string) bool { return len(s) > 0 },
			),
			Room: Filter(
				Transform(h.Amenities.Room, func(s string) string { return strings.ToLower(strings.TrimSpace(s)) }),
				func(s string) bool { return len(s) > 0 },
			),
		},
		Images: Images{
			Rooms: Transform(h.Images.Rooms, func(p PaperfliesImage) Image {
				return Image{
					Link:        p.Link,
					Description: p.Caption,
				}
			}),
			Site: Transform(h.Images.Site, func(p PaperfliesImage) Image {
				return Image{
					Link:        p.Link,
					Description: p.Caption,
				}
			}),
			Amenities: []Image{}, // NOTE: paperflies doesn't provide images of amenities
		},

		BookingConditions: Filter(
			Transform(h.BookingConditions, strings.TrimSpace),
			func(s string) bool { return (len(s) > 0) },
		),
	}
}

type PaperfliesLocation struct {
	Address string `json:"address"` // e.g. "8 Sentosa Gateway, Beach Villas, 098269"
	Country string `json:"country"` // e.g. "Singapore"
}

type PaperfliesAmenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type PaperfliesImages struct {
	Rooms []PaperfliesImage `json:"rooms"`
	Site  []PaperfliesImage `json:"site"`
}

type PaperfliesImage struct {
	Link    string `json:"link"`    // e.g. "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg"
	Caption string `json:"caption"` // e.g. "Double room"
}
