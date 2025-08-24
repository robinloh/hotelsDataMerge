package parser

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/external"
	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/utils"
)

type IntParser interface {
	ParseSuppliersData(resp map[utils.Suppliers]json.RawMessage) ([]hotels.Hotel, error)
}

type intParser struct {
	logger       *slog.Logger
	extSuppliers external.ExtSuppliers
}

func Initialize(logger *slog.Logger) IntParser {
	return &intParser{
		logger: logger,
	}
}
