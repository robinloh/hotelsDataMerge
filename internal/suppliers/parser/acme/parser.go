package acme

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/utils"
)

func (a *AcmeParser) ParseAndMapSuppliersData() ([]hotels.Hotel, error) {
	var parsedData []AcmeParsedData
	if err := json.Unmarshal(a.RawData, &parsedData); err != nil {
		a.Logger.Error("[Acme] Failed to unmarshal response", slog.Any("error", err))
		return nil, err
	}

	mappedHotels := make([]hotels.Hotel, 0, len(parsedData))
	for _, data := range parsedData {
		hotel := hotels.Hotel{
			Id:            data.Id,
			DestinationId: data.DestinationId,
			Name:          utils.TrimSpacesInString(data.Name),
			Location: &hotels.HotelLocation{
				Lat:     data.Latitude,
				Lng:     data.Longitude,
				Address: utils.TrimSpacesInString(data.Address),
				City:    utils.TrimSpacesInString(data.City),
				Country: utils.TrimSpacesInString(data.Country),
			},
			Description: utils.TrimSpacesInString(data.Description),
			Amenities: &hotels.HotelAmenities{
				General: utils.TrimSpacesInSlices(data.Facilities),
			},
		}
		mappedHotels = append(mappedHotels, hotel)
	}
	return mappedHotels, nil
}
