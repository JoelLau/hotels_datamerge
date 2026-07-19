package hotels_test

import (
	"encoding/json"
	"errors"
	hotels "hotels_data_merge"
	"testing"

	Hotels "hotels_data_merge"

	"github.com/stretchr/testify/require"
)

func TestNewHotels(t *testing.T) {
	t.Parallel()

	t.Run("multiple hotels: same id & destination_id", func(t *testing.T) {
		t.Parallel()

		given := hotels.Hotels{
			{ID: "iJhz", DestinationID: 5432},
			{ID: "iJhz", DestinationID: 5432},
			{ID: "iJhz", DestinationID: 5432},
		}

		got, err := Hotels.NewHotels(given)
		require.NoError(t, err)
		require.Equal(t, given, got)
	})

	t.Run("multiple hotels: same id, diff destination_id", func(t *testing.T) {
		t.Parallel()

		given := hotels.Hotels{
			{ID: "iJhz", DestinationID: 1111},
			{ID: "iJhz", DestinationID: 2222},
			{ID: "iJhz", DestinationID: 3333},
		}

		got, err := Hotels.NewHotels(given)

		conflictErr, ok := errors.AsType[*Hotels.ConflictingDestinationIDsError](err)
		require.True(t, ok)
		require.Equal(t, "iJhz", conflictErr.HotelID)
		require.ElementsMatch(t, []int{1111, 2222, 3333}, conflictErr.DestinationIDs)

		want := hotels.Hotels{
			{ID: "iJhz", DestinationID: 1111},
			{ID: "iJhz", DestinationID: 2222},
			{ID: "iJhz", DestinationID: 3333},
		}
		require.Equal(t, want, got)
	})
}

func TestMerge(t *testing.T) {
	t.Run("REQUIREMENTS.md example", func(t *testing.T) {
		t.Parallel()

		// The "iJhz" record as reported by each of the 3 suppliers, taken
		// verbatim from the acme/patagonia/paperflies sample fixtures.
		acmeJSON := `{
  "Id": "iJhz",
  "DestinationId": 5432,
  "Name": "Beach Villas Singapore",
  "Latitude": 1.264751,
  "Longitude": 103.824006,
  "Address": " 8 Sentosa Gateway, Beach Villas ",
  "City": "Singapore",
  "Country": "SG",
  "PostalCode": "098269",
  "Description": "  This 5 star hotel is located on the coastline of Singapore.",
  "Facilities": ["Pool", "BusinessCenter", "WiFi ", "DryCleaning", " Breakfast"]
}`

		patagoniaJSON := `{
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
}`

		paperfliesJSON := `{
  "hotel_id": "iJhz",
  "destination_id": 5432,
  "hotel_name": "Beach Villas Singapore",
  "location": { "address": "8 Sentosa Gateway, Beach Villas, 098269", "country": "Singapore" },
  "details": "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex and a 2-minute walk from the Waterfront train station. Featuring sundecks and pool, garden or sea views, the plush 1- to 3-bedroom villas offer free Wi-Fi and flat-screens, as well as free-standing baths, minibars, and tea and coffeemaking facilities. Upgraded villas add private pools, fridges and microwaves; some have wine cellars. A 4-bedroom unit offers a kitchen and a living room. There's 24-hour room and butler service. Amenities include posh restaurant, plus an outdoor pool, a hot tub, and free parking.",
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
    "All children are welcome. One child under 12 years stays free of charge when using existing beds. One child under 2 years stays free of charge in a child's cot/crib. One child under 4 years stays free of charge when using existing beds. One older child or adult is charged SGD 82.39 per person per night in an extra bed. The maximum number of children's cots/cribs in a room is 1. There is no capacity for extra beds in the room.",
    "Pets are not allowed.",
    "WiFi is available in all areas and is free of charge.",
    "Free private parking is possible on site (reservation is not needed).",
    "Guests are required to show a photo identification and credit card upon check-in. Please note that all Special Requests are subject to availability and additional charges may apply. Payment before arrival via bank transfer is required. The property will contact you after you book to provide instructions. Please note that the full amount of the reservation is due before arrival. Resorts World Sentosa will send a confirmation with detailed payment information. After full payment is taken, the property's details, including the address and where to collect keys, will be emailed to you. Bag checks will be conducted prior to entry to Adventure Cove Waterpark. === Upon check-in, guests will be provided with complimentary Sentosa Pass (monorail) to enjoy unlimited transportation between Sentosa Island and Harbour Front (VivoCity). === Prepayment for non refundable bookings will be charged by RWS Call Centre. === All guests can enjoy complimentary parking during their stay, limited to one exit from the hotel per day. === Room reservation charges will be charged upon check-in. Credit card provided upon reservation is for guarantee purpose. === For reservations made with inclusive breakfast, please note that breakfast is applicable only for number of adults paid in the room rate. Any children or additional adults are charged separately for breakfast and are to paid directly to the hotel."
  ]
}`

		var acmeHotel hotels.AcmeHotel
		require.NoError(t, json.Unmarshal([]byte(acmeJSON), &acmeHotel))

		var patagoniaHotel hotels.PatagoniaHotel
		require.NoError(t, json.Unmarshal([]byte(patagoniaJSON), &patagoniaHotel))

		var paperfliesHotel hotels.PaperfliesHotel
		require.NoError(t, json.Unmarshal([]byte(paperfliesJSON), &paperfliesHotel))

		given := hotels.Hotels{
			{
				ID:            acmeHotel.ID,
				DestinationID: acmeHotel.DestinationID,
				Name:          acmeHotel.Name,
				Location: hotels.Location{
					Latitude:  *acmeHotel.Latitude,
					Longitude: *acmeHotel.Longitude,
					Address:   acmeHotel.Address,
					City:      acmeHotel.City,
					Country:   acmeHotel.Country,
				},
				Description: acmeHotel.Description,
				Amenities: hotels.Amenities{
					General: acmeHotel.Facilities,
				},
			},
			{
				ID:            patagoniaHotel.ID,
				DestinationID: patagoniaHotel.Destination,
				Name:          patagoniaHotel.Name,
				Location: hotels.Location{
					Latitude:  patagoniaHotel.Latitude,
					Longitude: patagoniaHotel.Longitude,
					Address:   *patagoniaHotel.Address,
				},
				Description: *patagoniaHotel.Info,
				Amenities: hotels.Amenities{
					Room: patagoniaHotel.Amenities,
				},
				Images: hotels.Images{
					Rooms: []hotels.Image{
						{Link: patagoniaHotel.Images.Rooms[0].URL, Description: patagoniaHotel.Images.Rooms[0].Description},
						{Link: patagoniaHotel.Images.Rooms[1].URL, Description: patagoniaHotel.Images.Rooms[1].Description},
					},
					Amenities: []hotels.Image{
						{Link: patagoniaHotel.Images.Amenities[0].URL, Description: patagoniaHotel.Images.Amenities[0].Description},
						{Link: patagoniaHotel.Images.Amenities[1].URL, Description: patagoniaHotel.Images.Amenities[1].Description},
					},
				},
			},
			{
				ID:            paperfliesHotel.HotelID,
				DestinationID: paperfliesHotel.DestinationID,
				Name:          paperfliesHotel.HotelName,
				Location: hotels.Location{
					Address: paperfliesHotel.Location.Address,
					Country: paperfliesHotel.Location.Country,
				},
				Description: paperfliesHotel.Details,
				Amenities: hotels.Amenities{
					General: paperfliesHotel.Amenities.General,
					Room:    paperfliesHotel.Amenities.Room,
				},
				Images: hotels.Images{
					Rooms: []hotels.Image{
						{Link: paperfliesHotel.Images.Rooms[0].Link, Description: paperfliesHotel.Images.Rooms[0].Caption},
						{Link: paperfliesHotel.Images.Rooms[1].Link, Description: paperfliesHotel.Images.Rooms[1].Caption},
					},
					Site: []hotels.Image{
						{Link: paperfliesHotel.Images.Site[0].Link, Description: paperfliesHotel.Images.Site[0].Caption},
					},
				},
				BookingConditions: paperfliesHotel.BookingConditions,
			},
		}

		want := hotels.Hotel{
			ID:            "iJhz",
			DestinationID: 5432,
			Name:          "Beach Villas Singapore",
			Location: hotels.Location{
				Latitude:  1.264751,
				Longitude: 103.824006,
				Address:   "8 Sentosa Gateway, Beach Villas, 098269",
				City:      "Singapore",
				Country:   "Singapore",
			},
			Description: "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex and a 2-minute walk from the Waterfront train station. Featuring sundecks and pool, garden or sea views, the plush 1- to 3-bedroom villas offer free Wi-Fi and flat-screens, as well as free-standing baths, minibars, and tea and coffeemaking facilities. Upgraded villas add private pools, fridges and microwaves; some have wine cellars. A 4-bedroom unit offers a kitchen and a living room. There's 24-hour room and butler service. Amenities include posh restaurant, plus an outdoor pool, a hot tub, and free parking.",
			Amenities: hotels.Amenities{
				General: []string{"outdoor pool", "indoor pool", "business center", "childcare", "wifi", "dry cleaning", "breakfast"},
				Room:    []string{"aircon", "tv", "coffee machine", "kettle", "hair dryer", "iron", "bathtub"},
			},
			Images: hotels.Images{
				Rooms: []hotels.Image{
					{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", Description: "Double room"},
					{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/3.jpg", Description: "Double room"},
					{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg", Description: "Bathroom"},
				},
				Site: []hotels.Image{
					{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/1.jpg", Description: "Front"},
				},
				Amenities: []hotels.Image{
					{Link: "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg", Description: "RWS"},
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

		got := given.Merge()
		require.Equal(t, want, got)
	})
}
