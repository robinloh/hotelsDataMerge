package acme

import (
	"encoding/json"
	"log/slog"
	"testing"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/utils"
)

func TestAcmeParser_ParseAndMapSuppliersData(t *testing.T) {
	type fields struct {
		Logger       *slog.Logger
		SupplierName utils.Suppliers
		RawData      json.RawMessage
	}
	tests := []struct {
		name    string
		fields  fields
		want    []hotels.Hotel
		wantErr bool
	}{
		{
			name: "Success - Parse valid ACME data",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData: json.RawMessage(`[
					{
						"Id": "hotel1",
						"DestinationId": 123,
						"Name": "  Test Hotel  ",
						"Latitude": 40.7128,
						"Longitude": -74.0060,
						"Address": "  123 Test St  ",
						"City": "  Test City  ",
						"Country": "  Test Country  ",
						"PostalCode": "12345",
						"Description": "  A test hotel  ",
						"Facilities": ["WiFi", "Pool", "Gym"]
					}
				]`),
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Test Hotel",
					Location: &hotels.HotelLocation{
						Lat:     float64(40.7128),
						Lng:     float64(-74.0060),
						Address: "123 Test St",
						City:    "Test City",
						Country: "Test Country",
					},
					Description: "A test hotel",
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi", "Pool", "Gym"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Parse multiple hotels",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData: json.RawMessage(`[
					{
						"Id": "hotel1",
						"DestinationId": 123,
						"Name": "Hotel 1",
						"Latitude": 10.0,
						"Longitude": 20.0,
						"Address": "Address 1",
						"City": "City 1",
						"Country": "Country 1",
						"PostalCode": "11111",
						"Description": "Description 1",
						"Facilities": ["WiFi"]
					},
					{
						"Id": "hotel2",
						"DestinationId": 456,
						"Name": "Hotel 2",
						"Latitude": 30.0,
						"Longitude": 40.0,
						"Address": "Address 2",
						"City": "City 2",
						"Country": "Country 2",
						"PostalCode": "22222",
						"Description": "Description 2",
						"Facilities": ["Pool", "Gym"]
					}
				]`),
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
					Location: &hotels.HotelLocation{
						Lat:     float64(10.0),
						Lng:     float64(20.0),
						Address: "Address 1",
						City:    "City 1",
						Country: "Country 1",
					},
					Description: "Description 1",
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi"},
					},
				},
				{
					Id:            "hotel2",
					DestinationId: 456,
					Name:          "Hotel 2",
					Location: &hotels.HotelLocation{
						Lat:     float64(30.0),
						Lng:     float64(40.0),
						Address: "Address 2",
						City:    "City 2",
						Country: "Country 2",
					},
					Description: "Description 2",
					Amenities: &hotels.HotelAmenities{
						General: []string{"Pool", "Gym"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Parse hotel with empty facilities",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData: json.RawMessage(`[
					{
						"Id": "hotel1",
						"DestinationId": 123,
						"Name": "Test Hotel",
						"Latitude": 0,
						"Longitude": 0,
						"Address": "",
						"City": "",
						"Country": "",
						"PostalCode": "",
						"Description": "",
						"Facilities": []
					}
				]`),
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Test Hotel",
					Location: &hotels.HotelLocation{
						Lat:     float64(0),
						Lng:     float64(0),
						Address: "",
						City:    "",
						Country: "",
					},
					Description: "",
					Amenities: &hotels.HotelAmenities{
						General: []string{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Parse empty array",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData:      json.RawMessage(`[]`),
			},
			want:    []hotels.Hotel{},
			wantErr: false,
		},
		{
			name: "Error - Invalid JSON",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData:      json.RawMessage(`{"invalid": json`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error - Malformed JSON array",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData:      json.RawMessage(`[{"Id": "hotel1", "Name": "Test Hotel"`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success - Parse with string coordinates",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData: json.RawMessage(`[
					{
						"Id": "hotel1",
						"DestinationId": 123,
						"Name": "Test Hotel",
						"Latitude": "40.7128",
						"Longitude": "-74.0060",
						"Address": "123 Test St",
						"City": "Test City",
						"Country": "Test Country",
						"PostalCode": "12345",
						"Description": "A test hotel",
						"Facilities": ["WiFi"]
					}
				]`),
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Test Hotel",
					Location: &hotels.HotelLocation{
						Lat:     "40.7128",
						Lng:     "-74.0060",
						Address: "123 Test St",
						City:    "Test City",
						Country: "Test Country",
					},
					Description: "A test hotel",
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AcmeParser{
				Logger:       tt.fields.Logger,
				SupplierName: tt.fields.SupplierName,
				RawData:      tt.fields.RawData,
			}
			got, err := a.ParseAndMapSuppliersData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAndMapSuppliersData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("ParseAndMapSuppliersData() got %d hotels, want %d", len(got), len(tt.want))
			} else {
				for i, gotHotel := range got {
					wantHotel := tt.want[i]
					
					if gotHotel.Id != wantHotel.Id {
						t.Errorf("Hotel[%d].Id = %s, want %s", i, gotHotel.Id, wantHotel.Id)
					}
					if gotHotel.DestinationId != wantHotel.DestinationId {
						t.Errorf("Hotel[%d].DestinationId = %d, want %d", i, gotHotel.DestinationId, wantHotel.DestinationId)
					}
					if gotHotel.Name != wantHotel.Name {
						t.Errorf("Hotel[%d].Name = %s, want %s", i, gotHotel.Name, wantHotel.Name)
					}
					if gotHotel.Description != wantHotel.Description {
						t.Errorf("Hotel[%d].Description = %s, want %s", i, gotHotel.Description, wantHotel.Description)
					}
					
					if gotHotel.Location != nil && wantHotel.Location != nil {
						if gotHotel.Location.Lat != wantHotel.Location.Lat {
							t.Errorf("Hotel[%d].Location.Lat = %v, want %v", i, gotHotel.Location.Lat, wantHotel.Location.Lat)
						}
						if gotHotel.Location.Lng != wantHotel.Location.Lng {
							t.Errorf("Hotel[%d].Location.Lng = %v, want %v", i, gotHotel.Location.Lng, wantHotel.Location.Lng)
						}
						if gotHotel.Location.Address != wantHotel.Location.Address {
							t.Errorf("Hotel[%d].Location.Address = %s, want %s", i, gotHotel.Location.Address, wantHotel.Location.Address)
						}
						if gotHotel.Location.City != wantHotel.Location.City {
							t.Errorf("Hotel[%d].Location.City = %s, want %s", i, gotHotel.Location.City, wantHotel.Location.City)
						}
						if gotHotel.Location.Country != wantHotel.Location.Country {
							t.Errorf("Hotel[%d].Location.Country = %s, want %s", i, gotHotel.Location.Country, wantHotel.Location.Country)
						}
					}
					
					if gotHotel.Amenities != nil && wantHotel.Amenities != nil {
						if len(gotHotel.Amenities.General) != len(wantHotel.Amenities.General) {
							t.Errorf("Hotel[%d].Amenities.General count = %d, want %d", i, len(gotHotel.Amenities.General), len(wantHotel.Amenities.General))
						}
						for j, facility := range gotHotel.Amenities.General {
							if j < len(wantHotel.Amenities.General) && facility != wantHotel.Amenities.General[j] {
								t.Errorf("Hotel[%d].Amenities.General[%d] = %s, want %s", i, j, facility, wantHotel.Amenities.General[j])
							}
						}
					}
				}
			}
		})
	}
}
