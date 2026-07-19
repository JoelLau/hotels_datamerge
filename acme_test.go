package hotels_test

import (
	"encoding/json"
	"testing"

	hotels "hotels_data_merge"

	"github.com/stretchr/testify/require"
)

func TestAcmeHotelUnmarshal(t *testing.T) {
	t.Parallel()

	sampleJSON := `[
  {
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
  },
  {
    "Id": "SjyX",
    "DestinationId": 5432,
    "Name": "InterContinental Singapore Robertson Quay",
    "Latitude": null,
    "Longitude": null,
    "Address": " 1 Nanson Road",
    "City": "Singapore",
    "Country": "SG",
    "PostalCode": "238909",
    "Description": "Enjoy sophisticated waterfront living at the new InterContinental® Singapore Robertson Quay.",
    "Facilities": ["Pool", "WiFi ", "Aircon", "BusinessCenter", "BathTub", "Breakfast", "DryCleaning", "Bar"]
  },
  {
    "Id": "f8c9",
    "DestinationId": 1122,
    "Name": "Hilton Shinjuku Tokyo",
    "Latitude": "",
    "Longitude": "",
    "Address": "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN",
    "City": "Tokyo",
    "Country": "JP",
    "PostalCode": "160-0023",
    "Description": "Hilton Tokyo is located in Shinjuku, the heart of Tokyo's business, shopping and entertainment district.",
    "Facilities": ["Pool", "WiFi ", "BusinessCenter", "DryCleaning", " Breakfast", "Bar", "BathTub"]
  }
]`

	var got []hotels.AcmeHotel
	require.NoError(t, json.Unmarshal([]byte(sampleJSON), &got))
	require.Len(t, got, 3)

	want := []hotels.AcmeHotel{
		{
			ID:            "iJhz",
			DestinationID: 5432,
			Name:          "Beach Villas Singapore",
			Latitude:      new(1.264751),
			Longitude:     new(103.824006),
			Address:       " 8 Sentosa Gateway, Beach Villas ",
			City:          "Singapore",
			Country:       "SG",
			PostalCode:    "098269",
			Description:   "  This 5 star hotel is located on the coastline of Singapore.",
			Facilities:    []string{"Pool", "BusinessCenter", "WiFi ", "DryCleaning", " Breakfast"},
		},
		{
			ID:            "SjyX",
			DestinationID: 5432,
			Name:          "InterContinental Singapore Robertson Quay",
			Latitude:      nil, // null in source
			Longitude:     nil, // null in source
			Address:       " 1 Nanson Road",
			City:          "Singapore",
			Country:       "SG",
			PostalCode:    "238909",
			Description:   "Enjoy sophisticated waterfront living at the new InterContinental® Singapore Robertson Quay.",
			Facilities:    []string{"Pool", "WiFi ", "Aircon", "BusinessCenter", "BathTub", "Breakfast", "DryCleaning", "Bar"},
		},
		{
			ID:            "f8c9",
			DestinationID: 1122,
			Name:          "Hilton Shinjuku Tokyo",
			Latitude:      nil,
			Longitude:     nil,
			Address:       "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN",
			City:          "Tokyo",
			Country:       "JP",
			PostalCode:    "160-0023",
			Description:   "Hilton Tokyo is located in Shinjuku, the heart of Tokyo's business, shopping and entertainment district.",
			Facilities:    []string{"Pool", "WiFi ", "BusinessCenter", "DryCleaning", " Breakfast", "Bar", "BathTub"},
		},
	}

	require.Equal(t, want, got)
}
