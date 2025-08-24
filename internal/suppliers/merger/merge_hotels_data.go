package merger

import (
	"hotelsDataMerge/internal/hotels"
	mergerHotel "hotelsDataMerge/internal/suppliers/merger/hotel"
)

// MergeHotelsData merges mergerHotel data using the Builder pattern
func (i *intMerger) MergeHotelsData(mappedData []hotels.Hotel) map[string]hotels.Hotel {
	hotelByHotelIDMap := make(map[string]hotels.Hotel)

	for _, hotel := range mappedData {
		if mergedHotel, exists := hotelByHotelIDMap[hotel.Id]; !exists {
			hotelByHotelIDMap[hotel.Id] = hotel
		} else {
			mergedHotel = buildMergedHotel(mergedHotel, hotel)
			hotelByHotelIDMap[hotel.Id] = mergedHotel
		}
	}

	return hotelByHotelIDMap
}

func buildMergedHotel(existing, new hotels.Hotel) hotels.Hotel {
	hotelBuilder := mergerHotel.NewHotelBuilder(existing)

	hotelBuilder.WithID(existing.Id, new.Id)
	hotelBuilder.WithDestinationID(existing.DestinationId, new.DestinationId)
	hotelBuilder.WithName(existing.Name, new.Name)
	hotelBuilder.WithDescription(existing.Description, new.Description)
	hotelBuilder.WithLocation(existing.Location, new.Location)
	hotelBuilder.WithAmenities(existing.Amenities, new.Amenities)
	hotelBuilder.WithImages(existing.Images, new.Images)
	hotelBuilder.WithBookingConditions(existing.BookingConditions, new.BookingConditions)

	return hotelBuilder.Build()
}
