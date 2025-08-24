package hotel

import "hotelsDataMerge/internal/hotels"

type HotelBuilder interface {
	HotelBaseBuilder
	Build() hotels.Hotel
}

type HotelBaseBuilder interface {
	WithID(existing, new string) *hotelBuilder
	WithDestinationID(existing, new uint64) *hotelBuilder
	WithName(existing, new string) *hotelBuilder
	WithDescription(existing, new string) *hotelBuilder
	WithLocation(existing, new *hotels.HotelLocation) *hotelBuilder
	WithAmenities(existing, new *hotels.HotelAmenities) *hotelBuilder
	WithImages(existing, new *hotels.HotelImages) *hotelBuilder
	WithBookingConditions(existing, new []string) *hotelBuilder
}

type hotelBuilder struct {
	hotel hotels.Hotel
}

func NewHotelBuilder(existing hotels.Hotel) HotelBuilder {
	return &hotelBuilder{
		hotel: existing,
	}
}

func (b *hotelBuilder) Build() hotels.Hotel {
	return b.hotel
}
