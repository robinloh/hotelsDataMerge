package patagonia

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/internal/suppliers/utils"
)

type PatagoniaParser struct {
	Logger       *slog.Logger
	SupplierName utils.Suppliers
	RawData      json.RawMessage
}

type PatagoniaParsedData struct {
	Id          string                   `json:"id"`
	Destination uint64                   `json:"destination"`
	Name        string                   `json:"name"`
	Lat         float64                  `json:"lat"`
	Lng         float64                  `json:"lng"`
	Address     string                   `json:"address"`
	Info        string                   `json:"info"`
	Amenities   []string                 `json:"amenities"`
	Images      PatagoniaParsedDataImage `json:"images"`
}

type PatagoniaParsedDataImage struct {
	Rooms     []PatagoniaParsedDataImageDetails `json:"rooms"`
	Amenities []PatagoniaParsedDataImageDetails `json:"amenities"`
}
type PatagoniaParsedDataImageDetails struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}
