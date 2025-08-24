package acme

import (
	"encoding/json"
	"log/slog"

	"hotelsDataMerge/internal/suppliers/utils"
)

type AcmeParser struct {
	Logger       *slog.Logger
	SupplierName utils.Suppliers
	RawData      json.RawMessage
}

type AcmeParsedData struct {
	Id            string   `json:"Id"`
	DestinationId uint64   `json:"DestinationId"`
	Name          string   `json:"Name"`
	Latitude      any      `json:"Latitude"`
	Longitude     any      `json:"Longitude"`
	Address       string   `json:"Address"`
	City          string   `json:"City"`
	Country       string   `json:"Country"`
	PostalCode    string   `json:"PostalCode"`
	Description   string   `json:"Description"`
	Facilities    []string `json:"Facilities"`
}
