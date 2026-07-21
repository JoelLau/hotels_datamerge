package hotels

import (
	"fmt"
	"maps"
	"slices"
)

type ConflictingDestinationIDsError struct {
	HotelID        string
	DestinationIDs []int
}

func (e *ConflictingDestinationIDsError) Error() string {
	return fmt.Sprintf("destination id conflict: destination ids (%v)", e.DestinationIDs)
}

type Hotels []Hotel

// WARN: assumes all ids are the same because
// of how upstream code is structured at this point on 2026-07-19.
func NewHotels(hs []Hotel) (Hotels, error) {
	destinationIDs := map[int]bool{}
	for _, sh := range hs {
		destinationIDs[sh.DestinationID] = true
	}

	if len(destinationIDs) > 1 {
		return hs, &ConflictingDestinationIDsError{
			HotelID:        hs[0].ID,
			DestinationIDs: slices.Collect(maps.Keys(destinationIDs)),
		}
	}

	return Hotels(hs), nil
}

func (hs Hotels) Merge() Hotel {
	return Hotel{
		ID:                hs.mergeID(),
		DestinationID:     hs.mergeDestinationID(),
		Name:              hs.mergeName(),
		Location:          hs.mergeLocation(),
		Description:       hs.mergeDescription(),
		Amenities:         hs.mergeAmenities(),
		Images:            hs.mergeImages(),
		BookingConditions: hs.mergeBookingConditions(),
	}
}

func (hs Hotels) mergeID() string {
	if len(hs) <= 0 {
		return ""
	}

	return hs[0].ID
}

func (hs Hotels) mergeDestinationID() int {
	if len(hs) <= 0 {
		return 0
	}

	return hs[0].DestinationID
}

func (hs Hotels) getNames() []string {
	if hs == nil {
		return nil
	}

	names := make([]string, len(hs))
	for i, h := range hs {
		names[i] = h.Name
	}

	return names
}

func (hs Hotels) mergeName() string {
	longestName := ""
	for _, n := range hs.getNames() {
		if len(n) > len(longestName) {
			longestName = n
		}
	}
	return longestName
}

func (hs Hotels) mergeLocation() Location {
	return Location{
		Latitude:  1.264751,
		Longitude: 103.824006,
		Address:   "8 Sentosa Gateway, Beach Villas, 098269",
		City:      "Singapore",
		Country:   "Singapore",
	}
}

func (hs Hotels) getDescription() []string {
	if hs == nil {
		return nil
	}

	descriptions := make([]string, len(hs))
	for i, h := range hs {
		descriptions[i] = h.Description
	}

	return descriptions
}

func (hs Hotels) mergeDescription() string {
	longestDescription := ""
	for _, d := range hs.getDescription() {
		if len(d) > len(longestDescription) {
			longestDescription = d
		}
	}
	return longestDescription
}

func (hs Hotels) mergeAmenities() Amenities {
	var amenities Amenities
	for _, h := range hs {
		amenities.General = slices.Concat(amenities.General, h.Amenities.General)
		amenities.Room = slices.Concat(amenities.Room, h.Amenities.Room)
	}
	return amenities
}

func (hs Hotels) mergeImages() Images {
	var images Images
	for _, h := range hs {
		images.Rooms = slices.Concat(images.Rooms, h.Images.Rooms)
		images.Site = slices.Concat(images.Site, h.Images.Site)
		images.Amenities = slices.Concat(images.Amenities, h.Images.Amenities)
	}
	return images
}

func (hs Hotels) mergeBookingConditions() []string {
	merged := make([]string, len(hs))
	for _, h := range hs {
		merged = slices.Concat(merged, h.BookingConditions)
	}
	slices.Sort(merged)
	return merged
}
