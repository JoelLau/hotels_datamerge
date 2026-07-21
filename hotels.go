package hotels

import (
	"fmt"
	"maps"
	"slices"
	"strings"
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

// assumes that only 1 unique ID exists in the slice
func (hs Hotels) mergeID() string {
	if len(hs) <= 0 {
		return ""
	}

	return hs[0].ID
}

// assumes that only 1 unique detinationID exists in the slice
func (hs Hotels) mergeDestinationID() int {
	if len(hs) <= 0 {
		return 0
	}

	return hs[0].DestinationID
}

func (hs Hotels) mergeName() string {
	if hs == nil {
		return ""
	}

	return LongestString(
		Filter(
			Transform(hs, func(h Hotel) string { return strings.TrimSpace(h.Name) }),
			func(s string) bool { return len(s) > 0 },
		),
	)
}

// TODO: lat / lng should come from the same source
// TODO: setting lat / lng must be atomic; both or neither should be nil
// TODO: attempt to resolve ISO country / city codes; consider normalizing at supplier level
func (hs Hotels) mergeLocation() *Location {
	return &Location{
		Latitude:  new(1.264751),
		Longitude: new(103.824006),
		Address:   new("8 Sentosa Gateway, Beach Villas, 098269"),
		City:      new("Singapore"),
		Country:   new("Singapore"),
	}
}

func (hs Hotels) mergeDescription() string {
	if hs == nil {
		return ""
	}

	return LongestString(
		Filter(
			Transform(hs, func(h Hotel) string { return strings.TrimSpace(h.Description) }),
			func(s string) bool { return len(s) > 0 },
		),
	)
}

func (hs Hotels) mergeAmenities() Amenities {
	mergeAmenityField := func(strArrArr [][]string) []string {
		result := make([]string, 0)
		for _, strArr := range strArrArr {
			result = slices.Concat(result, strArr)
		}

		result = Transform(result, func(s string) string {
			r := ToLowerCaseWithSpaces(strings.TrimSpace(s))
			if override, ok := amenityOverrides[r]; ok {
				return override
			}
			return r
		})
		result = Filter(result, func(s string) bool { return len(s) > 0 })
		slices.Sort(result)
		result = slices.Compact(result)

		return result
	}

	return Amenities{
		General: mergeAmenityField(Transform(hs, func(h Hotel) []string { return h.Amenities.General })),
		Room:    mergeAmenityField(Transform(hs, func(h Hotel) []string { return h.Amenities.Room })),
	}
}

func (hs Hotels) mergeImages() Images {
	mergeImageField := func(imgArrArr [][]Image) []Image {
		// key: link
		// val: image
		linkToImageMap := make(map[string]Image)
		for _, imgArr := range imgArrArr {
			for _, img := range imgArr {
				k := strings.TrimSpace(img.Link)
				if len(img.Description) > len(linkToImageMap[k].Description) {
					linkToImageMap[k] = img
				}
			}
		}

		result := slices.Collect(maps.Values(linkToImageMap))
		slices.SortStableFunc(result, func(a, b Image) int {
			return strings.Compare(a.Link, b.Link)
		})

		return result
	}

	return Images{
		Rooms:     mergeImageField(Transform(hs, func(h Hotel) []Image { return h.Images.Rooms })),
		Site:      mergeImageField(Transform(hs, func(h Hotel) []Image { return h.Images.Site })),
		Amenities: mergeImageField(Transform(hs, func(h Hotel) []Image { return h.Images.Amenities })),
	}
}

func (hs Hotels) mergeBookingConditions() []string {
	merged := make([]string, len(hs))
	for _, h := range hs {
		merged = slices.Concat(merged, h.BookingConditions)
	}

	merged = Filter(
		Transform(merged, func(s string) string { return strings.TrimSpace(s) }),
		func(s string) bool { return len(s) > 0 },
	)
	// slices.Sort(merged)
	merged = slices.Compact(merged)
	return merged
}
