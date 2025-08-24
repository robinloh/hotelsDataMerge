package parser

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/parser/acme"
	"hotelsDataMerge/internal/suppliers/parser/paperflies"
	"hotelsDataMerge/internal/suppliers/parser/patagonia"
	"hotelsDataMerge/internal/suppliers/utils"
)

type DefaultParserFactory struct {
	logger *slog.Logger
}

type ParserFactory interface {
	ParseAndMapSuppliersData() ([]hotels.Hotel, error)
}

func (f *DefaultParserFactory) CreateParser(supplierName utils.Suppliers, rawData json.RawMessage) ParserFactory {
	switch supplierName {
	case utils.Acme:
		return &acme.AcmeParser{
			Logger:       f.logger,
			SupplierName: supplierName,
			RawData:      rawData,
		}
	case utils.Patagonia:
		return &patagonia.PatagoniaParser{
			Logger:       f.logger,
			SupplierName: supplierName,
			RawData:      rawData,
		}
	case utils.Paperflies:
		return &paperflies.PaperfliesParser{
			Logger:       f.logger,
			SupplierName: supplierName,
			RawData:      rawData,
		}
	default:
		return nil
	}
}
