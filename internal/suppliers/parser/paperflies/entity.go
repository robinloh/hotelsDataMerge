package paperflies

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/internal/suppliers/utils"
)

type PaperfliesParser struct {
	Logger       *slog.Logger
	SupplierName utils.Suppliers
	RawData      json.RawMessage
}

type PaperfliesParsedData struct {
	HotelID           string                        `json:"hotel_id"`
	DestinationID     uint64                        `json:"destination_id"`
	HotelName         string                        `json:"hotel_name"`
	Location          PaperfliesParsedDataLocation  `json:"location"`
	Details           string                        `json:"details"`
	Amenities         PaperfliesParsedDataAmenities `json:"amenities"`
	Images            PaperfliesParsedDataImages    `json:"images"`
	BookingConditions []string                      `json:"booking_conditions"`
}

type PaperfliesParsedDataLocation struct {
	Address string `json:"address"`
	Country string `json:"country"`
}

type PaperfliesParsedDataAmenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type PaperfliesParsedDataImages struct {
	Rooms []PaperfliesParsedDataImageDetails `json:"rooms"`
	Site  []PaperfliesParsedDataImageDetails `json:"site"`
}

type PaperfliesParsedDataImageDetails struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}
