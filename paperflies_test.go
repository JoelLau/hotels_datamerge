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

	t.Run("Hotel", func(t *testing.T) {
		t.Parallel()

		t.Run("Example 01: Beach Villas Singapore", func(t *testing.T) {
			t.Parallel()

			given := hotels.PaperfliesHotel{
				HotelID:       "iJhz",
				DestinationID: 5432,
				HotelName:     "Beach Villas Singapore",
				Location: hotels.PaperfliesLocation{
					Address: "8 Sentosa Gateway, Beach Villas, 098269",
					Country: "Singapore",
				},
				Details: "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex and a 2-minute walk from the Waterfront train station. Featuring sundecks and pool, garden or sea views, the plush 1- to 3-bedroom villas offer free Wi-Fi and flat-screens, as well as free-standing baths, minibars, and tea and coffeemaking facilities. Upgraded villas add private pools, fridges and microwaves; some have wine cellars. A 4-bedroom unit offers a kitchen and a living room. There's 24-hour room and butler service. Amenities include posh restaurant, plus an outdoor pool, a hot tub, and free parking.",
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
					"All children are welcome. One child under 12 years stays free of charge when using existing beds. One child under 2 years stays free of charge in a child's cot/crib. One child under 4 years stays free of charge when using existing beds. One older child or adult is charged SGD 82.39 per person per night in an extra bed. The maximum number of children's cots/cribs in a room is 1. There is no capacity for extra beds in the room.",
					"Pets are not allowed.",
					"WiFi is available in all areas and is free of charge.",
					"Free private parking is possible on site (reservation is not needed).",
					"Guests are required to show a photo identification and credit card upon check-in. Please note that all Special Requests are subject to availability and additional charges may apply. Payment before arrival via bank transfer is required. The property will contact you after you book to provide instructions. Please note that the full amount of the reservation is due before arrival. Resorts World Sentosa will send a confirmation with detailed payment information. After full payment is taken, the property's details, including the address and where to collect keys, will be emailed to you. Bag checks will be conducted prior to entry to Adventure Cove Waterpark. === Upon check-in, guests will be provided with complimentary Sentosa Pass (monorail) to enjoy unlimited transportation between Sentosa Island and Harbour Front (VivoCity). === Prepayment for non refundable bookings will be charged by RWS Call Centre. === All guests can enjoy complimentary parking during their stay, limited to one exit from the hotel per day. === Room reservation charges will be charged upon check-in. Credit card provided upon reservation is for guarantee purpose. === For reservations made with inclusive breakfast, please note that breakfast is applicable only for number of adults paid in the room rate. Any children or additional adults are charged separately for breakfast and are to paid directly to the hotel.",
				},
			}

			want := hotels.Hotel{
				ID:            "iJhz",
				DestinationID: 5432,
				Name:          "Beach Villas Singapore",
				Location: &hotels.Location{
					Latitude:  nil,
					Longitude: nil,
					Address:   new("8 Sentosa Gateway, Beach Villas, 098269"),
					Country:   new("Singapore"),
					City:      nil,
				},
				Description: "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex and a 2-minute walk from the Waterfront train station. Featuring sundecks and pool, garden or sea views, the plush 1- to 3-bedroom villas offer free Wi-Fi and flat-screens, as well as free-standing baths, minibars, and tea and coffeemaking facilities. Upgraded villas add private pools, fridges and microwaves; some have wine cellars. A 4-bedroom unit offers a kitchen and a living room. There's 24-hour room and butler service. Amenities include posh restaurant, plus an outdoor pool, a hot tub, and free parking.",
				Amenities: hotels.Amenities{
					General: []string{"outdoor pool", "indoor pool", "business center", "childcare"},
					Room:    []string{"tv", "coffee machine", "kettle", "hair dryer", "iron"},
				},
				Images: hotels.Images{
					Rooms: []hotels.Image{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", Description: "Double room"},
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/3.jpg", Description: "Double room"},
					},
					Site: []hotels.Image{
						{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/1.jpg", Description: "Front"},
					},
					Amenities: []hotels.Image{},
				},
				BookingConditions: []string{
					"All children are welcome. One child under 12 years stays free of charge when using existing beds. One child under 2 years stays free of charge in a child's cot/crib. One child under 4 years stays free of charge when using existing beds. One older child or adult is charged SGD 82.39 per person per night in an extra bed. The maximum number of children's cots/cribs in a room is 1. There is no capacity for extra beds in the room.",
					"Pets are not allowed.",
					"WiFi is available in all areas and is free of charge.",
					"Free private parking is possible on site (reservation is not needed).",
					"Guests are required to show a photo identification and credit card upon check-in. Please note that all Special Requests are subject to availability and additional charges may apply. Payment before arrival via bank transfer is required. The property will contact you after you book to provide instructions. Please note that the full amount of the reservation is due before arrival. Resorts World Sentosa will send a confirmation with detailed payment information. After full payment is taken, the property's details, including the address and where to collect keys, will be emailed to you. Bag checks will be conducted prior to entry to Adventure Cove Waterpark. === Upon check-in, guests will be provided with complimentary Sentosa Pass (monorail) to enjoy unlimited transportation between Sentosa Island and Harbour Front (VivoCity). === Prepayment for non refundable bookings will be charged by RWS Call Centre. === All guests can enjoy complimentary parking during their stay, limited to one exit from the hotel per day. === Room reservation charges will be charged upon check-in. Credit card provided upon reservation is for guarantee purpose. === For reservations made with inclusive breakfast, please note that breakfast is applicable only for number of adults paid in the room rate. Any children or additional adults are charged separately for breakfast and are to paid directly to the hotel.",
				},
			}

			got := given.Hotel()
			require.Equal(t, want, *got)
		})
	})
}
