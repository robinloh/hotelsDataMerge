package merger

import (
	"log/slog"

	"hotelsDataMerge/internal/hotels"
)

type IntMerger interface {
	MergeHotelsData(mappedData []hotels.Hotel) (mergedHotels map[string]hotels.Hotel)
}

type intMerger struct {
	logger *slog.Logger
}

func Initialize(logger *slog.Logger) IntMerger {
	return &intMerger{
		logger: logger,
	}
}
