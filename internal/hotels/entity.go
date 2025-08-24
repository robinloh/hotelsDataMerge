package hotels

type Hotel struct {
	Id                string          `json:"id"`
	DestinationId     uint64          `json:"destination_id"`
	Name              string          `json:"name"`
	Location          *HotelLocation  `json:"location"`
	Description       string          `json:"description"`
	Amenities         *HotelAmenities `json:"amenities"`
	Images            *HotelImages    `json:"images"`
	BookingConditions []string        `json:"booking_conditions"`
}

type HotelLocation struct {
	Lat     any    `json:"lat"`
	Lng     any    `json:"lng"`
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type HotelAmenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type HotelImages struct {
	Rooms     []HotelImageDetails `json:"rooms"`
	Site      []HotelImageDetails `json:"site"`
	Amenities []HotelImageDetails `json:"amenities"`
}

type HotelImageDetails struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}
