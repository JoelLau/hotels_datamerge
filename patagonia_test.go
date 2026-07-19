package hotels_test

import (
	"encoding/json"
	"testing"

	hotels "hotels_data_merge"

	"github.com/stretchr/testify/require"
)

func TestPatagoniaHotelUnmarshal(t *testing.T) {
	t.Parallel()

	sampleJSON := `[
  {
    "id": "iJhz",
    "destination": 5432,
    "name": "Beach Villas Singapore",
    "lat": 1.264751,
    "lng": 103.824006,
    "address": "8 Sentosa Gateway, Beach Villas, 098269",
    "info": "Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters.",
    "amenities": ["Aircon", "Tv", "Coffee machine", "Kettle", "Hair dryer", "Iron", "Tub"],
    "images": {
      "rooms": [
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", "description": "Double room" },
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg", "description": "Bathroom" }
      ],
      "amenities": [
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg", "description": "RWS" },
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/6.jpg", "description": "Sentosa Gateway" }
      ]
    }
  },
  {
    "id": "f8c9",
    "destination": 1122,
    "name": "Hilton Tokyo Shinjuku",
    "lat": 35.6926,
    "lng": 139.690965,
    "address": null,
    "info": null,
    "amenities": null,
    "images": {
      "rooms": [
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i10_m.jpg", "description": "Suite" },
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i11_m.jpg", "description": "Suite - Living room" }
      ],
      "amenities": [
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i57_m.jpg", "description": "Bar" }
      ]
    }
  }
]`

	var got []hotels.PatagoniaHotel
	require.NoError(t, json.Unmarshal([]byte(sampleJSON), &got))
	require.Len(t, got, 2)

	want := []hotels.PatagoniaHotel{
		{
			ID:          "iJhz",
			Destination: 5432,
			Name:        "Beach Villas Singapore",
			Latitude:    1.264751,
			Longitude:   103.824006,
			Address:     new("8 Sentosa Gateway, Beach Villas, 098269"),
			Info:        new("Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters."),
			Amenities:   []string{"Aircon", "Tv", "Coffee machine", "Kettle", "Hair dryer", "Iron", "Tub"},
			Images: hotels.PatagoniaImages{
				Rooms: []hotels.PatagoniaImage{
					{URL: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", Description: "Double room"},
					{URL: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg", Description: "Bathroom"},
				},
				Amenities: []hotels.PatagoniaImage{
					{URL: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg", Description: "RWS"},
					{URL: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/6.jpg", Description: "Sentosa Gateway"},
				},
			},
		},
		{
			ID:          "f8c9",
			Destination: 1122,
			Name:        "Hilton Tokyo Shinjuku",
			Latitude:    35.6926,
			Longitude:   139.690965,
			Address:     nil,
			Info:        nil,
			Amenities:   nil,
			Images: hotels.PatagoniaImages{
				Rooms: []hotels.PatagoniaImage{
					{URL: "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i10_m.jpg", Description: "Suite"},
					{URL: "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i11_m.jpg", Description: "Suite - Living room"},
				},
				Amenities: []hotels.PatagoniaImage{
					{URL: "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i57_m.jpg", Description: "Bar"},
				},
			},
		},
	}

	require.Equal(t, want, got)
}
