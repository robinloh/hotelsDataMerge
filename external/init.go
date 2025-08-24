package external

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/internal/suppliers/utils"
)

var suppliersURLMap = map[utils.Suppliers]string{
	utils.Acme:       "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme",
	utils.Patagonia:  "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia",
	utils.Paperflies: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies",
}

type ExtSuppliers interface {
	GetSuppliersRawInfo(supplierURL string) (json.RawMessage, error)
}

type externalHandler struct {
	logger *slog.Logger
}

func Initialize(logger *slog.Logger) ExtSuppliers {
	extHandler := &externalHandler{
		logger: logger,
	}
	return extHandler
}

func GetSuppliersURLMap() map[utils.Suppliers]string {
	return suppliersURLMap
}
