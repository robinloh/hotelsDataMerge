package hotels

func (i *intHotels) GetHotels(hotelIDs []string, destinationID uint64) ([]Hotel, error) {
	var hotelsByHotelID, hotelsByDestinationId []Hotel
	if len(hotelIDs) > 0 {
		for _, hotelID := range hotelIDs {
			hotel, ok := hotelByHotelIDMap[hotelID]
			if !ok {
				continue
			}
			if destinationID != 0 && hotel.DestinationId != destinationID {
				continue
			}
			hotelsByHotelID = append(hotelsByHotelID, hotel)
		}
	} else if destinationID != 0 {
		hotelsByDestinationId = hotelsByDestinationIdMap[destinationID]
	}
	result := removeDuplicates(hotelsByHotelID, hotelsByDestinationId)
	return result, nil
}

func removeDuplicates(hotelsByHotelID []Hotel, hotelsByDestinationId []Hotel) []Hotel {
	combinedHotels := append(hotelsByHotelID, hotelsByDestinationId...)
	seen := make(map[string]bool)
	var uniqueHotels []Hotel
	for _, hotel := range combinedHotels {
		if _, exists := seen[hotel.Id]; !exists {
			seen[hotel.Id] = true
			uniqueHotels = append(uniqueHotels, hotel)
		}
	}
	return uniqueHotels
}
