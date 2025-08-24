package hotels

import (
	"log/slog"
)

var (
	hotelIDsMap       = make(map[string]bool)
	hotelByHotelIDMap = make(map[string]Hotel)

	destinationIDsMap        = make(map[uint64]bool)
	hotelsByDestinationIdMap = make(map[uint64][]Hotel)
)

type IntHotels interface {
	GetHotels(hotelIDs []string, destinationID uint64) (hotels []Hotel, err error)
}

type intHotels struct {
	logger *slog.Logger
}

func Initialize(logger *slog.Logger) IntHotels {
	return &intHotels{
		logger: logger,
	}
}

func GetHotelIDsMap() map[string]bool {
	return hotelIDsMap
}

func GetDestinationIDsMap() map[uint64]bool {
	return destinationIDsMap
}

func SaveMaps(hotels map[string]Hotel) {
	// Clear existing maps first
	ClearMaps()
	
	for hotelID, hotel := range hotels {
		hotelIDsMap[hotelID] = true
		hotelByHotelIDMap[hotelID] = hotel
		if hotel.DestinationId != 0 {
			if _, exists := destinationIDsMap[hotel.DestinationId]; !exists {
				destinationIDsMap[hotel.DestinationId] = true
				hotelsByDestinationIdMap[hotel.DestinationId] = make([]Hotel, 0)
			}
			hotelsByDestinationIdMap[hotel.DestinationId] = append(hotelsByDestinationIdMap[hotel.DestinationId], hotel)
		}
	}
}

// ClearMaps clears all the hotel maps
func ClearMaps() {
	hotelIDsMap = make(map[string]bool)
	hotelByHotelIDMap = make(map[string]Hotel)
	destinationIDsMap = make(map[uint64]bool)
	hotelsByDestinationIdMap = make(map[uint64][]Hotel)
}
