package hotel

import (
	"strings"

	"hotelsDataMerge/internal/hotels"
)

func (b *hotelBuilder) WithID(existing, new string) *hotelBuilder {
	if len(new) > 0 {
		b.hotel.Id = new
	} else {
		b.hotel.Id = existing
	}
	return b
}

func (b *hotelBuilder) WithDestinationID(existing, new uint64) *hotelBuilder {
	if new > 0 {
		b.hotel.DestinationId = new
	} else {
		b.hotel.DestinationId = existing
	}
	return b
}

func (b *hotelBuilder) WithName(existing, new string) *hotelBuilder {
	if len(new) > len(existing) {
		b.hotel.Name = new
	} else {
		b.hotel.Name = existing
	}
	return b
}

func (b *hotelBuilder) WithLocation(existing, new *hotels.HotelLocation) *hotelBuilder {
	if existing == nil {
		b.hotel.Location = new
		return b
	}

	if new == nil {
		b.hotel.Location = existing
		return b
	}

	merged := &hotels.HotelLocation{}

	if new.Lat != nil {
		if lat, ok := new.Lat.(float64); ok {
			merged.Lat = lat
		}
	} else {
		merged.Lat = existing.Lat
	}

	if new.Lng != nil {
		if lng, ok := new.Lng.(float64); ok {
			merged.Lng = lng
		}
	} else {
		merged.Lng = existing.Lng
	}

	if len(new.Address) < len(existing.Address) {
		merged.Address = existing.Address
	} else {
		merged.Address = new.Address
	}

	if len(new.City) == 0 {
		merged.City = existing.City
	} else {
		merged.City = new.City
	}

	// If the new country is empty or a 2-letter code, prefer the existing one
	if len(new.Country) == 0 || len(existing.Country) == 2 {
		merged.Country = existing.Country
	} else {
		merged.Country = new.Country
	}

	b.hotel.Location = merged
	return b
}

func (b *hotelBuilder) WithDescription(existing, new string) *hotelBuilder {
	if len(new) > len(existing) {
		b.hotel.Description = new
	} else {
		b.hotel.Description = existing
	}
	return b
}

func (b *hotelBuilder) WithAmenities(existing, new *hotels.HotelAmenities) *hotelBuilder {
	if existing == nil {
		b.hotel.Amenities = new
		return b
	}
	if new == nil {
		b.hotel.Amenities = existing
		return b
	}

	merged := &hotels.HotelAmenities{}
	merged.General = mergeStrings(existing.General, new.General)
	merged.Room = mergeStrings(existing.Room, new.Room)

	merged.General, merged.Room = filterAmenities(merged.General, merged.Room)

	b.hotel.Amenities = merged
	return b
}

func (b *hotelBuilder) WithImages(existing, new *hotels.HotelImages) *hotelBuilder {
	if existing == nil {
		b.hotel.Images = new
		return b
	}
	if new == nil {
		b.hotel.Images = existing
		return b
	}

	merged := &hotels.HotelImages{}
	merged.Rooms = mergeHotelImageDetails(existing.Rooms, new.Rooms)
	merged.Site = mergeHotelImageDetails(existing.Site, new.Site)
	merged.Amenities = mergeHotelImageDetails(existing.Amenities, new.Amenities)

	b.hotel.Images = merged
	return b
}

func (b *hotelBuilder) WithBookingConditions(existing, new []string) *hotelBuilder {
	if len(new) > 0 {
		b.hotel.BookingConditions = new
	} else {
		b.hotel.BookingConditions = existing
	}
	return b
}

func mergeStrings(existing, new []string) []string {
	merged := make([]string, 0)
	mergedMap := make(map[string]bool)
	for _, str := range existing {
		trimmed := strings.TrimSpace(str)
		mergedMap[strings.ToLower(trimmed)] = true
	}
	for _, str := range new {
		trimmed := strings.TrimSpace(str)
		mergedMap[strings.ToLower(trimmed)] = true
	}
	for str := range mergedMap {
		merged = append(merged, str)
	}
	return merged
}

func mergeHotelImageDetails(existing, new []hotels.HotelImageDetails) []hotels.HotelImageDetails {
	if len(new) > 0 {
		return new
	} else {
		return existing
	}
}

// filterRoomAmenities removes any amenities from "room" that are also present in "general"
func filterAmenities(generalAmenities, roomAmenities []string) ([]string, []string) {
	if len(generalAmenities) == 0 {
		return roomAmenities, generalAmenities
	}

	commonMap := make(map[string]bool)
	for _, amenity := range roomAmenities {
		commonMap[amenity] = true
	}

	var filteredGeneralAmenities []string
	for _, amenity := range generalAmenities {
		if !commonMap[amenity] {
			filteredGeneralAmenities = append(filteredGeneralAmenities, amenity)
		}
	}

	return filteredGeneralAmenities, roomAmenities
}
