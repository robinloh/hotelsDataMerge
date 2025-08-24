package paperflies

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/utils"
)

func (p *PaperfliesParser) ParseAndMapSuppliersData() ([]hotels.Hotel, error) {
	var parsedData []PaperfliesParsedData
	if err := json.Unmarshal(p.RawData, &parsedData); err != nil {
		p.Logger.Error("[Paperflies] Failed to unmarshal response", slog.Any("error", err))
		return nil, err
	}

	mappedHotels := make([]hotels.Hotel, 0, len(parsedData))
	for _, data := range parsedData {
		hotel := hotels.Hotel{
			Id:            data.HotelID,
			DestinationId: data.DestinationID,
			Name:          utils.TrimSpacesInString(data.HotelName),
			Location: &hotels.HotelLocation{
				Address: utils.TrimSpacesInString(data.Location.Address),
				Country: utils.TrimSpacesInString(data.Location.Country),
			},
			Description: utils.TrimSpacesInString(data.Details),
			Amenities: &hotels.HotelAmenities{
				General: utils.TrimSpacesInSlices(data.Amenities.General),
				Room:    utils.TrimSpacesInSlices(data.Amenities.Room),
			},
			Images: &hotels.HotelImages{
				Rooms: mapRoomImages(data.Images.Rooms),
				Site:  mapSiteImages(data.Images.Site),
			},
			BookingConditions: utils.TrimSpacesInSlices(data.BookingConditions),
		}
		mappedHotels = append(mappedHotels, hotel)
	}
	return mappedHotels, nil
}

func mapRoomImages(roomsImageDetails []PaperfliesParsedDataImageDetails) []hotels.HotelImageDetails {
	var rooms []hotels.HotelImageDetails
	for _, room := range roomsImageDetails {
		rooms = append(rooms, hotels.HotelImageDetails{
			Link:        room.Link,
			Description: room.Caption,
		})
	}
	return rooms
}

func mapSiteImages(siteImageDetails []PaperfliesParsedDataImageDetails) []hotels.HotelImageDetails {
	var sites []hotels.HotelImageDetails
	for _, site := range siteImageDetails {
		sites = append(sites, hotels.HotelImageDetails{
			Link:        site.Link,
			Description: site.Caption,
		})
	}
	return sites
}
