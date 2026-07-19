package hotels

// PaperfliesHotel is the raw shape returned by the paperflies supplier endpoint.
type PaperfliesHotel struct {
	HotelID           string              `json:"hotel_id"`       // e.g. "iJhz"
	DestinationID     int                 `json:"destination_id"` // e.g. 5432
	HotelName         string              `json:"hotel_name"`     // e.g. "Beach Villas Singapore"
	Location          PaperfliesLocation  `json:"location"`
	Details           string              `json:"details"` // e.g. "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex."
	Amenities         PaperfliesAmenities `json:"amenities"`
	Images            PaperfliesImages    `json:"images"`
	BookingConditions []string            `json:"booking_conditions"`
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
