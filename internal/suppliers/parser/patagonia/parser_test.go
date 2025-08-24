package patagonia

import (
	"encoding/json"
	"log/slog"
	"testing"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/utils"
)

func TestPatagoniaParser_ParseAndMapSuppliersData(t *testing.T) {
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
			name: "Success - Parse valid Patagonia data",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Patagonia,
				RawData: json.RawMessage(`[
					{
						"id": "hotel1",
						"destination": 123,
						"name": "  Test Hotel  ",
						"lat": 40.7128,
						"lng": -74.0060,
						"address": "  123 Test St  ",
						"info": "  A test hotel  ",
						"amenities": ["WiFi", "Pool", "Gym"],
						"images": {
							"rooms": [
								{
									"url": "http://example.com/room1.jpg",
									"description": "Room 1"
								}
							],
							"amenities": [
								{
									"url": "http://example.com/amenity1.jpg",
									"description": "Amenity 1"
								}
							]
						}
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
					},
					Description: "A test hotel",
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi", "Pool", "Gym"},
					},
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{
							{
								Link:        "http://example.com/room1.jpg",
								Description: "Room 1",
							},
						},
						Amenities: []hotels.HotelImageDetails{
							{
								Link:        "http://example.com/amenity1.jpg",
								Description: "Amenity 1",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Parse multiple hotels",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Patagonia,
				RawData: json.RawMessage(`[
					{
						"id": "hotel1",
						"destination": 123,
						"name": "Hotel 1",
						"lat": 10.0,
						"lng": 20.0,
						"address": "Address 1",
						"info": "Description 1",
						"amenities": ["WiFi"],
						"images": {
							"rooms": [],
							"amenities": []
						}
					},
					{
						"id": "hotel2",
						"destination": 456,
						"name": "Hotel 2",
						"lat": 30.0,
						"lng": 40.0,
						"address": "Address 2",
						"info": "Description 2",
						"amenities": ["Pool", "Gym"],
						"images": {
							"rooms": [],
							"amenities": []
						}
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
					},
					Description: "Description 1",
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi"},
					},
					Images: &hotels.HotelImages{
						Rooms:     []hotels.HotelImageDetails{},
						Amenities: []hotels.HotelImageDetails{},
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
					},
					Description: "Description 2",
					Amenities: &hotels.HotelAmenities{
						General: []string{"Pool", "Gym"},
					},
					Images: &hotels.HotelImages{
						Rooms:     []hotels.HotelImageDetails{},
						Amenities: []hotels.HotelImageDetails{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Parse empty array",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Patagonia,
				RawData:      json.RawMessage(`[]`),
			},
			want:    []hotels.Hotel{},
			wantErr: false,
		},
		{
			name: "Error - Invalid JSON",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Patagonia,
				RawData:      json.RawMessage(`{"invalid": json`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error - Malformed JSON array",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Patagonia,
				RawData:      json.RawMessage(`[{"id": "hotel1", "name": "Test Hotel"`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success - Parse hotel with minimal data",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Patagonia,
				RawData: json.RawMessage(`[
					{
						"id": "minimal",
						"destination": 0,
						"name": "",
						"lat": 0,
						"lng": 0,
						"address": "",
						"info": "",
						"amenities": [],
						"images": {
							"rooms": [],
							"amenities": []
						}
					}
				]`),
			},
			want: []hotels.Hotel{
				{
					Id:            "minimal",
					DestinationId: 0,
					Name:          "",
					Location: &hotels.HotelLocation{
						Lat:     float64(0),
						Lng:     float64(0),
						Address: "",
					},
					Description: "",
					Amenities: &hotels.HotelAmenities{
						General: []string{},
					},
					Images: &hotels.HotelImages{
						Rooms:     []hotels.HotelImageDetails{},
						Amenities: []hotels.HotelImageDetails{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Error - Parse hotel with string coordinates (should fail)",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Patagonia,
				RawData: json.RawMessage(`[
					{
						"id": "hotel1",
						"destination": 123,
						"name": "Test Hotel",
						"lat": "40.7128",
						"lng": "-74.0060",
						"address": "123 Test St",
						"info": "A test hotel",
						"amenities": ["WiFi"],
						"images": {
							"rooms": [],
							"amenities": []
						}
					}
				]`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PatagoniaParser{
				Logger:       tt.fields.Logger,
				SupplierName: tt.fields.SupplierName,
				RawData:      tt.fields.RawData,
			}
			got, err := p.ParseAndMapSuppliersData()
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
					}

					if gotHotel.Amenities != nil && wantHotel.Amenities != nil {
						if len(gotHotel.Amenities.General) != len(wantHotel.Amenities.General) {
							t.Errorf("Hotel[%d].Amenities.General count = %d, want %d", i, len(gotHotel.Amenities.General), len(wantHotel.Amenities.General))
						}
					}

					if gotHotel.Images != nil && wantHotel.Images != nil {
						if len(gotHotel.Images.Rooms) != len(wantHotel.Images.Rooms) {
							t.Errorf("Hotel[%d].Images.Rooms count = %d, want %d", i, len(gotHotel.Images.Rooms), len(wantHotel.Images.Rooms))
						}
						if len(gotHotel.Images.Amenities) != len(wantHotel.Images.Amenities) {
							t.Errorf("Hotel[%d].Images.Amenities count = %d, want %d", i, len(gotHotel.Images.Amenities), len(wantHotel.Images.Amenities))
						}
					}
				}
			}
		})
	}
}
