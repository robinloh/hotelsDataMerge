package server

import (
	"log/slog"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/proto"
)

type hotelsDataMergeService struct {
	logger *slog.Logger
	hotels hotels.IntHotels
	proto.UnimplementedHotelDataMergeServer
}

func NewHotelsDataMergeService(logger *slog.Logger) proto.HotelDataMergeServer {
	return &hotelsDataMergeService{
		logger: logger,
		hotels: hotels.Initialize(logger),
	}
}
