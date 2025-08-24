package fetcher

import (
	"encoding/json"

	"hotelsDataMerge/external"
	"hotelsDataMerge/internal/suppliers/utils"
)

type GetSuppliersResponse struct {
	SupplierName utils.Suppliers
	Error        error
	RawResp      []byte
}

func (i *intFetcher) GetLatestSupplierData() (hotelRawMap map[utils.Suppliers]json.RawMessage, err error) {
	hotelRawMap = make(map[utils.Suppliers]json.RawMessage)
	for supplierName, supplierURL := range external.GetSuppliersURLMap() {
		rawResp, err := i.extSuppliers.GetSuppliersRawInfo(supplierURL)
		if err != nil {
			return hotelRawMap, err
		}
		hotelRawMap[supplierName] = rawResp
	}
	return hotelRawMap, err
}
