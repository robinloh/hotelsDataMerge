package parser

import (
	"encoding/json"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/utils"
)

func (i *intParser) ParseSuppliersData(resp map[utils.Suppliers]json.RawMessage) ([]hotels.Hotel, error) {
	allHotels := make([]hotels.Hotel, 0)

	for supplierName, rawData := range resp {
		factory := &DefaultParserFactory{
			logger: i.logger,
		}
		parser := factory.CreateParser(supplierName, rawData)
		if parser != nil {
			parsedHotels, err := parser.ParseAndMapSuppliersData()
			if err != nil {
				return nil, err
			}
			allHotels = append(allHotels, parsedHotels...)
		}
	}
	return allHotels, nil
}
