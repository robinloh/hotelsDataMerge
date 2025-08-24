package fetcher

import (
	"encoding/json"
	"log/slog"
	"reflect"
	"testing"

	"hotelsDataMerge/external"
	"hotelsDataMerge/internal/suppliers/utils"
)

type mockExtSuppliers struct {
	responses map[string]json.RawMessage
	errors    map[string]error
}

func (m *mockExtSuppliers) GetSuppliersRawInfo(supplierURL string) (json.RawMessage, error) {
	if err, exists := m.errors[supplierURL]; exists {
		return nil, err
	}
	if resp, exists := m.responses[supplierURL]; exists {
		return resp, nil
	}
	return json.RawMessage{}, nil
}

func Test_intFetcher_GetLatestSupplierData(t *testing.T) {
	type fields struct {
		logger       *slog.Logger
		extSuppliers external.ExtSuppliers
	}
	tests := []struct {
		name            string
		fields          fields
		wantHotelRawMap map[utils.Suppliers]json.RawMessage
		wantErr         bool
	}{
		{
			name: "Success - Get data from all suppliers",
			fields: fields{
				logger: slog.Default(),
				extSuppliers: &mockExtSuppliers{
					responses: map[string]json.RawMessage{
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme":       json.RawMessage(`{"hotels":[]}`),
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia":  json.RawMessage(`{"hotels":[]}`),
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies": json.RawMessage(`{"hotels":[]}`),
					},
					errors: map[string]error{},
				},
			},
			wantHotelRawMap: map[utils.Suppliers]json.RawMessage{
				utils.Acme:       json.RawMessage(`{"hotels":[]}`),
				utils.Patagonia:  json.RawMessage(`{"hotels":[]}`),
				utils.Paperflies: json.RawMessage(`{"hotels":[]}`),
			},
			wantErr: false,
		},
		{
			name: "Success - Get data with different content",
			fields: fields{
				logger: slog.Default(),
				extSuppliers: &mockExtSuppliers{
					responses: map[string]json.RawMessage{
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme":       json.RawMessage(`{"hotels":[{"id":"hotel1"}]}`),
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia":  json.RawMessage(`{"hotels":[{"id":"hotel2"}]}`),
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies": json.RawMessage(`{"hotels":[{"id":"hotel3"}]}`),
					},
					errors: map[string]error{},
				},
			},
			wantHotelRawMap: map[utils.Suppliers]json.RawMessage{
				utils.Acme:       json.RawMessage(`{"hotels":[{"id":"hotel1"}]}`),
				utils.Patagonia:  json.RawMessage(`{"hotels":[{"id":"hotel2"}]}`),
				utils.Paperflies: json.RawMessage(`{"hotels":[{"id":"hotel3"}]}`),
			},
			wantErr: false,
		},
		{
			name: "Error - Any supplier fails (returns partial results)",
			fields: fields{
				logger: slog.Default(),
				extSuppliers: &mockExtSuppliers{
					responses: map[string]json.RawMessage{
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme":       json.RawMessage(`{"hotels":[]}`),
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia":  json.RawMessage(`{"hotels":[]}`),
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies": json.RawMessage(`{"hotels":[]}`),
					},
					errors: map[string]error{
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia": &json.SyntaxError{},
					},
				},
			},
			wantHotelRawMap: map[utils.Suppliers]json.RawMessage{},
			wantErr:         true,
		},
		{
			name: "Success - Empty responses from all suppliers",
			fields: fields{
				logger: slog.Default(),
				extSuppliers: &mockExtSuppliers{
					responses: map[string]json.RawMessage{
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme":       json.RawMessage{},
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia":  json.RawMessage{},
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies": json.RawMessage{},
					},
					errors: map[string]error{},
				},
			},
			wantHotelRawMap: map[utils.Suppliers]json.RawMessage{
				utils.Acme:       json.RawMessage{},
				utils.Patagonia:  json.RawMessage{},
				utils.Paperflies: json.RawMessage{},
			},
			wantErr: false,
		},
		{
			name: "Success - Mixed content and empty responses",
			fields: fields{
				logger: slog.Default(),
				extSuppliers: &mockExtSuppliers{
					responses: map[string]json.RawMessage{
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme":       json.RawMessage(`{"hotels":[{"id":"hotel1"}]}`),
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia":  json.RawMessage{},
						"https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies": json.RawMessage(`{"hotels":[]}`),
					},
					errors: map[string]error{},
				},
			},
			wantHotelRawMap: map[utils.Suppliers]json.RawMessage{
				utils.Acme:       json.RawMessage(`{"hotels":[{"id":"hotel1"}]}`),
				utils.Patagonia:  json.RawMessage{},
				utils.Paperflies: json.RawMessage(`{"hotels":[]}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &intFetcher{
				logger:       tt.fields.logger,
				extSuppliers: tt.fields.extSuppliers,
			}
			gotHotelRawMap, err := i.GetLatestSupplierData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatestSupplierData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetLatestSupplierData() expected error but got none")
				}
			} else {
				if !reflect.DeepEqual(gotHotelRawMap, tt.wantHotelRawMap) {
					t.Errorf("GetLatestSupplierData() gotHotelRawMap = %v, want %v", gotHotelRawMap, tt.wantHotelRawMap)
				}
			}
		})
	}
}
