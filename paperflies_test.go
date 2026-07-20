package hotels_test

import (
	"encoding/json"
	"testing"

	hotels "hotels_data_merge"

	"github.com/stretchr/testify/require"
)

func TestPaperfliesHotel(t *testing.T) {
	t.Parallel()

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		sampleJSON := `[
  {
    "hotel_id": "iJhz",
    "destination_id": 5432,
    "hotel_name": "Beach Villas Singapore",
    "location": { "address": "8 Sentosa Gateway, Beach Villas, 098269", "country": "Singapore" },
    "details": "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex.",
    "amenities": {
      "general": ["outdoor pool", "indoor pool", "business center", "childcare"],
      "room": ["tv", "coffee machine", "kettle", "hair dryer", "iron"]
    },
    "images": {
      "rooms": [
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", "caption": "Double room" },
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/3.jpg", "caption": "Double room" }
      ],
      "site": [
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/1.jpg", "caption": "Front" }
      ]
    },
    "booking_conditions": [
      "All children are welcome. One child under 12 years stays free of charge when using existing beds.",
      "Pets are not allowed.",
      "WiFi is available in all areas and is free of charge.",
      "Free private parking is possible on site (reservation is not needed).",
      "Guests are required to show a photo identification and credit card upon check-in."
    ]
  },
  {
    "hotel_id": "SjyX",
    "destination_id": 5432,
    "hotel_name": "InterContinental",
    "location": { "address": "1 Nanson Rd, Singapore 238909", "country": "Singapore" },
    "details": "InterContinental Singapore Robertson Quay is luxury's preferred address offering stylishly cosmopolitan riverside living for discerning travelers to Singapore.",
    "amenities": {
      "general": ["outdoor pool", "business center", "childcare", "parking", "bar", "dry cleaning", "wifi", "breakfast", "concierge"],
      "room": ["aircon", "minibar", "tv", "bathtub", "hair dryer"]
    },
    "images": {
      "rooms": [
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i93_m.jpg", "caption": "Double room" },
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i94_m.jpg", "caption": "Bathroom" }
      ],
      "site": [
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i1_m.jpg", "caption": "Restaurant" },
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i2_m.jpg", "caption": "Hotel Exterior" },
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i5_m.jpg", "caption": "Entrance" },
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i24_m.jpg", "caption": "Bar" }
      ]
    },
    "booking_conditions": []
  },
  {
    "hotel_id": "f8c9",
    "destination_id": 1122,
    "hotel_name": "Hilton Tokyo",
    "location": { "address": "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN", "country": "Japan" },
    "details": "This sleek high-rise property is 10 minutes' walk from Shinjuku train station, 6 minutes' walk from the Tokyo Metropolitan Government Building and 3 km from Yoyogi Park.",
    "amenities": {
      "general": ["indoor pool", "business center", "wifi"],
      "room": ["tv", "aircon", "minibar", "bathtub", "hair dryer"]
    },
    "images": {
      "rooms": [
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i1_m.jpg", "caption": "Suite" },
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i15_m.jpg", "caption": "Double room" }
      ],
      "site": [
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i55_m.jpg", "caption": "Bar" }
      ]
    },
    "booking_conditions": [
      "All children are welcome. One child under 6 years stays free of charge when using existing beds. There is no capacity for extra beds in the room.",
      "Pets are not allowed.",
      "Wired internet is available in the hotel rooms and charges are applicable. WiFi is available in the hotel rooms and charges are applicable.",
      "Private parking is possible on site (reservation is not needed) and costs JPY 1500 per day.",
      "When booking more than 9 rooms, different policies and additional supplements may apply.",
      "The hotel's free shuttle is offered from Bus Stop #21 in front of Keio Department Store at Shinjuku Station."
    ]
  }
]`

		var got []hotels.PaperfliesHotel
		require.NoError(t, json.Unmarshal([]byte(sampleJSON), &got))
		require.Len(t, got, 3)

		want := []hotels.PaperfliesHotel{
			{
				HotelID:       "iJhz",
				DestinationID: 5432,
				HotelName:     "Beach Villas Singapore",
				Location: hotels.PaperfliesLocation{
					Address: "8 Sentosa Gateway, Beach Villas, 098269",
					Country: "Singapore",
				},
				Details: "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex.",
				Amenities: hotels.PaperfliesAmenities{
					General: []string{"outdoor pool", "indoor pool", "business center", "childcare"},
					Room:    []string{"tv", "coffee machine", "kettle", "hair dryer", "iron"},
				},
				Images: hotels.PaperfliesImages{
					Rooms: []hotels.PaperfliesImage{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", Caption: "Double room"},
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/3.jpg", Caption: "Double room"},
					},
					Site: []hotels.PaperfliesImage{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/1.jpg", Caption: "Front"},
					},
				},
				BookingConditions: []string{
					"All children are welcome. One child under 12 years stays free of charge when using existing beds.",
					"Pets are not allowed.",
					"WiFi is available in all areas and is free of charge.",
					"Free private parking is possible on site (reservation is not needed).",
					"Guests are required to show a photo identification and credit card upon check-in.",
				},
			},
			{
				HotelID:       "SjyX",
				DestinationID: 5432,
				HotelName:     "InterContinental",
				Location: hotels.PaperfliesLocation{
					Address: "1 Nanson Rd, Singapore 238909",
					Country: "Singapore",
				},
				Details: "InterContinental Singapore Robertson Quay is luxury's preferred address offering stylishly cosmopolitan riverside living for discerning travelers to Singapore.",
				Amenities: hotels.PaperfliesAmenities{
					General: []string{"outdoor pool", "business center", "childcare", "parking", "bar", "dry cleaning", "wifi", "breakfast", "concierge"},
					Room:    []string{"aircon", "minibar", "tv", "bathtub", "hair dryer"},
				},
				Images: hotels.PaperfliesImages{
					Rooms: []hotels.PaperfliesImage{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i93_m.jpg", Caption: "Double room"},
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i94_m.jpg", Caption: "Bathroom"},
					},
					Site: []hotels.PaperfliesImage{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i1_m.jpg", Caption: "Restaurant"},
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i2_m.jpg", Caption: "Hotel Exterior"},
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i5_m.jpg", Caption: "Entrance"},
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i24_m.jpg", Caption: "Bar"},
					},
				},
				BookingConditions: []string{},
			},
			{
				HotelID:       "f8c9",
				DestinationID: 1122,
				HotelName:     "Hilton Tokyo",
				Location: hotels.PaperfliesLocation{
					Address: "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN",
					Country: "Japan",
				},
				Details: "This sleek high-rise property is 10 minutes' walk from Shinjuku train station, 6 minutes' walk from the Tokyo Metropolitan Government Building and 3 km from Yoyogi Park.",
				Amenities: hotels.PaperfliesAmenities{
					General: []string{"indoor pool", "business center", "wifi"},
					Room:    []string{"tv", "aircon", "minibar", "bathtub", "hair dryer"},
				},
				Images: hotels.PaperfliesImages{
					Rooms: []hotels.PaperfliesImage{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i1_m.jpg", Caption: "Suite"},
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i15_m.jpg", Caption: "Double room"},
					},
					Site: []hotels.PaperfliesImage{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i55_m.jpg", Caption: "Bar"},
					},
				},
				BookingConditions: []string{
					"All children are welcome. One child under 6 years stays free of charge when using existing beds. There is no capacity for extra beds in the room.",
					"Pets are not allowed.",
					"Wired internet is available in the hotel rooms and charges are applicable. WiFi is available in the hotel rooms and charges are applicable.",
					"Private parking is possible on site (reservation is not needed) and costs JPY 1500 per day.",
					"When booking more than 9 rooms, different policies and additional supplements may apply.",
					"The hotel's free shuttle is offered from Bus Stop #21 in front of Keio Department Store at Shinjuku Station.",
				},
			},
		}

		require.Equal(t, want, got)
	})
}
