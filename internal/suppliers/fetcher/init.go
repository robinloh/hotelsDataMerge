package fetcher

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/external"
	"hotelsDataMerge/internal/suppliers/utils"
)

type IntFetcher interface {
	GetLatestSupplierData() (hotelRawMap map[utils.Suppliers]json.RawMessage, err error)
}

type intFetcher struct {
	logger       *slog.Logger
	extSuppliers external.ExtSuppliers
}

func Initialize(logger *slog.Logger, extSuppliers external.ExtSuppliers) IntFetcher {
	return &intFetcher{
		logger:       logger,
		extSuppliers: extSuppliers,
	}
}
