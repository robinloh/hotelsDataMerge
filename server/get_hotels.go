package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"hotelsDataMerge/external"
	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/proto"

	"google.golang.org/grpc/status"
)

var methodName = "[GetHotels]"

func (h *hotelsDataMergeService) GetHotels(ctx context.Context, req *proto.GetHotelsRequest) (resp *proto.GetHotelsResponse, err error) {
	h.logger.InfoContext(ctx, methodName+fmt.Sprintf(" API request : %+v", req))

	if !external.FetchSuppliersMutex.TryRLock() {
		h.logger.ErrorContext(ctx, fmt.Sprintf("%s Cannot acquire read lock - suppliers data update in progress", methodName))
		return resp, status.Error(http.StatusServiceUnavailable, "Service temporarily unavailable - data update in progress")
	}
	defer external.FetchSuppliersMutex.RUnlock()

	if err = h.validateRequest(req); err != nil {
		h.logger.ErrorContext(ctx, fmt.Sprintf("%s Invalid request. %s", methodName, err))
		return resp, status.Error(http.StatusBadRequest, "Invalid request")
	}
	hotelsList, err := h.hotels.GetHotels(req.HotelIDs, req.DestinationId)
	if err != nil {
		h.logger.ErrorContext(ctx, fmt.Sprintf("%s Error getting hotels: %s", methodName, err))
		return resp, status.Error(http.StatusServiceUnavailable, err.Error())
	}
	resp = h.constructResponse(hotelsList)
	h.logger.InfoContext(ctx, methodName+fmt.Sprintf(" API response : %+v", resp))
	return resp, nil
}

func (h *hotelsDataMergeService) validateRequest(req *proto.GetHotelsRequest) (err error) {
	if len(req.HotelIDs) == 0 && req.DestinationId == 0 {
		return errors.New("no request parameters were specified")
	}
	if len(req.HotelIDs) > 0 {
		hotelIDsMap := hotels.GetHotelIDsMap()
		for _, hotelID := range req.HotelIDs {
			if !hotelIDsMap[hotelID] {
				return fmt.Errorf("hotel ID %s does not exist", hotelID)
			}
		}
	}
	if req.DestinationId != 0 {
		destinationIDsMap := hotels.GetDestinationIDsMap()
		if !destinationIDsMap[req.DestinationId] {
			return fmt.Errorf("destination ID '%d' does not exist", req.DestinationId)
		}
	}
	return nil
}

func (h *hotelsDataMergeService) constructResponse(hotels []hotels.Hotel) (resp *proto.GetHotelsResponse) {
	if len(hotels) == 0 {
		return resp
	}
	hotelsResp := make([]*proto.Hotel, 0)
	for _, hotel := range hotels {
		hotelResp := &proto.Hotel{
			Id:                hotel.Id,
			DestinationId:     int64(hotel.DestinationId),
			Name:              hotel.Name,
			Location:          &proto.Location{},
			Description:       hotel.Description,
			Amenities:         &proto.HotelAmenities{},
			Images:            &proto.Image{},
			BookingConditions: hotel.BookingConditions,
		}
		if hotel.Location != nil {
			if lat, ok := hotel.Location.Lat.(float64); ok {
				hotelResp.Location.Lat = lat
			}
			if lng, ok := hotel.Location.Lng.(float64); ok {
				hotelResp.Location.Lng = lng
			}
			if len(hotel.Location.Address) > 0 {
				hotelResp.Location.Address = hotel.Location.Address
			}
			if len(hotel.Location.City) > 0 {
				hotelResp.Location.City = hotel.Location.City
			}
			if len(hotel.Location.Country) > 0 {
				hotelResp.Location.Country = hotel.Location.Country
			}
		}
		if len(hotel.Description) > 0 {
			hotelResp.Description = hotel.Description
		}
		if hotel.Amenities != nil {
			if len(hotel.Amenities.General) > 0 {
				hotelResp.Amenities.General = hotel.Amenities.General
			}
			if len(hotel.Amenities.Room) > 0 {
				hotelResp.Amenities.Room = hotel.Amenities.Room
			}
		}
		if hotel.Images != nil {
			if len(hotel.Images.Rooms) > 0 {
				hotelResp.Images.Rooms = constructRoomImageDetails(hotel.Images.Rooms)
			}
			if len(hotel.Images.Site) > 0 {
				hotelResp.Images.Site = constructSiteImageDetails(hotel.Images.Site)
			}
			if len(hotel.Images.Amenities) > 0 {
				hotelResp.Images.Amenities = constructAmenitiesImageDetails(hotel.Images.Amenities)
			}
		}
		hotelsResp = append(hotelsResp, hotelResp)
	}
	resp = &proto.GetHotelsResponse{
		Hotels: hotelsResp,
	}
	return resp
}

func constructRoomImageDetails(imageDetails []hotels.HotelImageDetails) []*proto.Room {
	roomImages := make([]*proto.Room, 0, len(imageDetails))
	for _, image := range imageDetails {
		roomImages = append(roomImages, &proto.Room{
			Link:        image.Link,
			Description: image.Description,
		})
	}
	return roomImages
}

func constructSiteImageDetails(imageDetails []hotels.HotelImageDetails) []*proto.Site {
	siteImages := make([]*proto.Site, 0, len(imageDetails))
	for _, image := range imageDetails {
		siteImages = append(siteImages, &proto.Site{
			Link:        image.Link,
			Description: image.Description,
		})
	}
	return siteImages
}

func constructAmenitiesImageDetails(imageDetails []hotels.HotelImageDetails) []*proto.ImageAmenity {
	amenitiesImages := make([]*proto.ImageAmenity, 0, len(imageDetails))
	for _, image := range imageDetails {
		amenitiesImages = append(amenitiesImages, &proto.ImageAmenity{
			Link:        image.Link,
			Description: image.Description,
		})
	}
	return amenitiesImages
}
