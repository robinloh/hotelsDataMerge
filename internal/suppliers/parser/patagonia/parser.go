package patagonia

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/utils"
)

func (p *PatagoniaParser) ParseAndMapSuppliersData() ([]hotels.Hotel, error) {
	var parsedData []PatagoniaParsedData
	if err := json.Unmarshal(p.RawData, &parsedData); err != nil {
		p.Logger.Error("[Patagonia] Failed to unmarshal response", slog.Any("error", err))
		return nil, err
	}

	mappedHotels := make([]hotels.Hotel, 0, len(parsedData))
	for _, data := range parsedData {
		hotel := hotels.Hotel{
			Id:            data.Id,
			DestinationId: data.Destination,
			Name:          utils.TrimSpacesInString(data.Name),
			Location: &hotels.HotelLocation{
				Lat:     data.Lat,
				Lng:     data.Lng,
				Address: utils.TrimSpacesInString(data.Address),
			},
			Description: utils.TrimSpacesInString(data.Info),
			Amenities: &hotels.HotelAmenities{
				General: utils.TrimSpacesInSlices(data.Amenities),
			},
			Images: &hotels.HotelImages{
				Rooms:     mapRoomImages(data.Images.Rooms),
				Amenities: mapAmenitiesImages(data.Images.Amenities),
			},
		}
		mappedHotels = append(mappedHotels, hotel)
	}

	return mappedHotels, nil
}

func mapRoomImages(roomsImageDetails []PatagoniaParsedDataImageDetails) []hotels.HotelImageDetails {
	var rooms []hotels.HotelImageDetails
	for _, room := range roomsImageDetails {
		rooms = append(rooms, hotels.HotelImageDetails{
			Link:        room.Url,
			Description: room.Description,
		})
	}
	return rooms
}

func mapAmenitiesImages(amenitiesImageDetails []PatagoniaParsedDataImageDetails) []hotels.HotelImageDetails {
	var amenities []hotels.HotelImageDetails
	for _, amenity := range amenitiesImageDetails {
		amenities = append(amenities, hotels.HotelImageDetails{
			Link:        amenity.Url,
			Description: amenity.Description,
		})
	}
	return amenities
}
