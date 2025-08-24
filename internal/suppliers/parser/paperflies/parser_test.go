package paperflies

import (
	"encoding/json"
	"log/slog"
	"testing"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/utils"
)

func TestPaperfliesParser_ParseAndMapSuppliersData(t *testing.T) {
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
			name: "Success - Parse valid Paperflies data",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Paperflies,
				RawData: json.RawMessage(`[
					{
						"hotel_id": "hotel1",
						"destination_id": 123,
						"hotel_name": "  Test Hotel  ",
						"location": {
							"address": "  123 Test St  ",
							"country": "  Test Country  "
						},
						"details": "  A test hotel  ",
						"amenities": {
							"general": ["WiFi", "Pool"],
							"room": ["TV", "AC"]
						},
						"images": {
							"rooms": [
								{
									"link": "http://example.com/room1.jpg",
									"caption": "Room 1"
								}
							],
							"site": [
								{
									"link": "http://example.com/site1.jpg",
									"caption": "Site 1"
								}
							]
						},
						"booking_conditions": ["No smoking", "No pets"]
					}
				]`),
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Test Hotel",
					Location: &hotels.HotelLocation{
						Address: "123 Test St",
						Country: "Test Country",
					},
					Description: "A test hotel",
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi", "Pool"},
						Room:    []string{"TV", "AC"},
					},
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{
							{
								Link:        "http://example.com/room1.jpg",
								Description: "Room 1",
							},
						},
						Site: []hotels.HotelImageDetails{
							{
								Link:        "http://example.com/site1.jpg",
								Description: "Site 1",
							},
						},
					},
					BookingConditions: []string{"No smoking", "No pets"},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Parse multiple hotels",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Paperflies,
				RawData: json.RawMessage(`[
					{
						"hotel_id": "hotel1",
						"destination_id": 123,
						"hotel_name": "Hotel 1",
						"location": {
							"address": "Address 1",
							"country": "Country 1"
						},
						"details": "Description 1",
						"amenities": {
							"general": ["WiFi"],
							"room": []
						},
						"images": {
							"rooms": [],
							"site": []
						},
						"booking_conditions": []
					},
					{
						"hotel_id": "hotel2",
						"destination_id": 456,
						"hotel_name": "Hotel 2",
						"location": {
							"address": "Address 2",
							"country": "Country 2"
						},
						"details": "Description 2",
						"amenities": {
							"general": [],
							"room": ["TV"]
						},
						"images": {
							"rooms": [],
							"site": []
						},
						"booking_conditions": ["No pets"]
					}
				]`),
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
					Location: &hotels.HotelLocation{
						Address: "Address 1",
						Country: "Country 1",
					},
					Description: "Description 1",
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi"},
						Room:    []string{},
					},
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{},
						Site:  []hotels.HotelImageDetails{},
					},
					BookingConditions: []string{},
				},
				{
					Id:            "hotel2",
					DestinationId: 456,
					Name:          "Hotel 2",
					Location: &hotels.HotelLocation{
						Address: "Address 2",
						Country: "Country 2",
					},
					Description: "Description 2",
					Amenities: &hotels.HotelAmenities{
						General: []string{},
						Room:    []string{"TV"},
					},
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{},
						Site:  []hotels.HotelImageDetails{},
					},
					BookingConditions: []string{"No pets"},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Parse empty array",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Paperflies,
				RawData:      json.RawMessage(`[]`),
			},
			want:    []hotels.Hotel{},
			wantErr: false,
		},
		{
			name: "Error - Invalid JSON",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Paperflies,
				RawData:      json.RawMessage(`{"invalid": json`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error - Malformed JSON array",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Paperflies,
				RawData:      json.RawMessage(`[{"hotel_id": "hotel1"`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success - Parse hotel with minimal data",
			fields: fields{
				Logger:       slog.Default(),
				SupplierName: utils.Paperflies,
				RawData: json.RawMessage(`[
					{
						"hotel_id": "minimal",
						"destination_id": 0,
						"hotel_name": "",
						"location": {
							"address": "",
							"country": ""
						},
						"details": "",
						"amenities": {
							"general": [],
							"room": []
						},
						"images": {
							"rooms": [],
							"site": []
						},
						"booking_conditions": []
					}
				]`),
			},
			want: []hotels.Hotel{
				{
					Id:            "minimal",
					DestinationId: 0,
					Name:          "",
					Location: &hotels.HotelLocation{
						Address: "",
						Country: "",
					},
					Description: "",
					Amenities: &hotels.HotelAmenities{
						General: []string{},
						Room:    []string{},
					},
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{},
						Site:  []hotels.HotelImageDetails{},
					},
					BookingConditions: []string{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaperfliesParser{
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
						if gotHotel.Location.Address != wantHotel.Location.Address {
							t.Errorf("Hotel[%d].Location.Address = %s, want %s", i, gotHotel.Location.Address, wantHotel.Location.Address)
						}
						if gotHotel.Location.Country != wantHotel.Location.Country {
							t.Errorf("Hotel[%d].Location.Country = %s, want %s", i, gotHotel.Location.Country, wantHotel.Location.Country)
						}
					}

					if gotHotel.Amenities != nil && wantHotel.Amenities != nil {
						if len(gotHotel.Amenities.General) != len(wantHotel.Amenities.General) {
							t.Errorf("Hotel[%d].Amenities.General count = %d, want %d", i, len(gotHotel.Amenities.General), len(wantHotel.Amenities.General))
						}
						if len(gotHotel.Amenities.Room) != len(wantHotel.Amenities.Room) {
							t.Errorf("Hotel[%d].Amenities.Room count = %d, want %d", i, len(gotHotel.Amenities.Room), len(wantHotel.Amenities.Room))
						}
					}

					if gotHotel.Images != nil && wantHotel.Images != nil {
						if len(gotHotel.Images.Rooms) != len(wantHotel.Images.Rooms) {
							t.Errorf("Hotel[%d].Images.Rooms count = %d, want %d", i, len(gotHotel.Images.Rooms), len(wantHotel.Images.Rooms))
						}
						if len(gotHotel.Images.Site) != len(wantHotel.Images.Site) {
							t.Errorf("Hotel[%d].Images.Site count = %d, want %d", i, len(gotHotel.Images.Site), len(wantHotel.Images.Site))
						}
					}

					if len(gotHotel.BookingConditions) != len(wantHotel.BookingConditions) {
						t.Errorf("Hotel[%d].BookingConditions count = %d, want %d", i, len(gotHotel.BookingConditions), len(wantHotel.BookingConditions))
					}
				}
			}
		})
	}
}
