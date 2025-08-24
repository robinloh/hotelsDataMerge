package external

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

var (
	FetchSuppliersMutex sync.RWMutex
)

func (e *externalHandler) GetSuppliersRawInfo(supplierURL string) (json.RawMessage, error) {
	resp, err := http.Get(supplierURL)
	if err != nil {
		e.logger.Error("[suppliers] Error in getting the suppliers info", "error", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			e.logger.Error("[suppliers] Error in closing body", "error", err)
		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		e.logger.Error("[suppliers] Error in reading the response body", "error", err)
		return nil, err
	}

	var respRawData json.RawMessage
	if err := json.Unmarshal(respBody, &respRawData); err != nil {
		e.logger.Error("[suppliers] Error in unmarshalling the response body", "error", err)
		return nil, err
	}

	return respRawData, nil
}
