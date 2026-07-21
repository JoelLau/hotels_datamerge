package hotels_test

import (
	"encoding/json"
	"testing"

	hotels "hotels_data_merge"

	"github.com/stretchr/testify/require"
)

func TestPatagoniaHotel(t *testing.T) {
	t.Parallel()

	t.Run("Unmarshal", func(t *testing.T) {
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
	})

	t.Run("Hotel", func(t *testing.T) {
		t.Parallel()

		t.Run("Example 01: Beach Villas Singapore", func(t *testing.T) {
			t.Parallel()

			given := hotels.PatagoniaHotel{
				ID:          "iJhz",
				Destination: 5432,
				Name:        "Beach Villas Singapore",
				Latitude:    1.264751,
				Longitude:   103.824006,
				Address:     new("8 Sentosa Gateway, Beach Villas, 098269"),
				Info:        new("Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters. Guests will find themselves in paradise with this series of exquisite tropical sanctuaries, making it the perfect setting for an idyllic retreat. Within each villa, guests will discover living areas and bedrooms that open out to mini gardens, private timber sundecks and verandahs elegantly framing either lush greenery or an expanse of sea. Guests are assured of a superior slumber with goose feather pillows and luxe mattresses paired with 400 thread count Egyptian cotton bed linen, tastefully paired with a full complement of luxurious in-room amenities and bathrooms boasting rain showers and free-standing tubs coupled with an exclusive array of ESPA amenities and toiletries. Guests also get to enjoy complimentary day access to the facilities at Asia’s flagship spa – the world-renowned ESPA."),
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
			}

			want := hotels.Hotel{
				ID:            "iJhz",
				DestinationID: 5432,
				Name:          "Beach Villas Singapore",
				Location: &hotels.Location{
					Latitude:  new(1.264751),
					Longitude: new(103.824006),
					Address:   new("8 Sentosa Gateway, Beach Villas, 098269"),
					City:      nil,
					Country:   nil,
				},
				Description: "Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters. Guests will find themselves in paradise with this series of exquisite tropical sanctuaries, making it the perfect setting for an idyllic retreat. Within each villa, guests will discover living areas and bedrooms that open out to mini gardens, private timber sundecks and verandahs elegantly framing either lush greenery or an expanse of sea. Guests are assured of a superior slumber with goose feather pillows and luxe mattresses paired with 400 thread count Egyptian cotton bed linen, tastefully paired with a full complement of luxurious in-room amenities and bathrooms boasting rain showers and free-standing tubs coupled with an exclusive array of ESPA amenities and toiletries. Guests also get to enjoy complimentary day access to the facilities at Asia’s flagship spa – the world-renowned ESPA.",
				Amenities: hotels.Amenities{
					General: []string{"aircon", "tv", "coffee machine", "kettle", "hair dryer", "iron", "tub"},
					Room:    []string{},
				},
				Images: hotels.Images{
					Rooms: []hotels.Image{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", Description: "Double room"},
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg", Description: "Bathroom"},
					},
					Site: []hotels.Image{},
					Amenities: []hotels.Image{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg", Description: "RWS"},
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/6.jpg", Description: "Sentosa Gateway"},
					},
				},
				BookingConditions: []string{},
			}

			got := given.Hotel()
			require.Equal(t, want, *got)
		})
	})
}
