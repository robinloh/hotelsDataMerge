package parser

import (
	"encoding/json"
	"log/slog"
	"testing"

	"hotelsDataMerge/external"
	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers/utils"
)

func Test_intParser_ParseSuppliersData(t *testing.T) {
	type fields struct {
		logger       *slog.Logger
		extSuppliers external.ExtSuppliers
	}
	type args struct {
		resp map[utils.Suppliers]json.RawMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []hotels.Hotel
		wantErr bool
	}{
		{
			name: "Success - Parse data from all suppliers",
			fields: fields{
				logger:       slog.Default(),
				extSuppliers: nil,
			},
			args: args{
				resp: map[utils.Suppliers]json.RawMessage{
					utils.Acme:       json.RawMessage(`[{"Id":"hotel1","DestinationId":123,"Name":"Hotel 1","Latitude":40.0,"Longitude":-74.0,"Address":"Address 1","City":"City 1","Country":"Country 1","Description":"Desc 1","Facilities":["WiFi"]}]`),
					utils.Patagonia:  json.RawMessage(`[{"id":"hotel2","destination":456,"name":"Hotel 2","lat":30.0,"lng":-80.0,"address":"Address 2","info":"Desc 2","amenities":["Pool"]}]`),
					utils.Paperflies: json.RawMessage(`[{"hotel_id":"hotel3","destination_id":789,"hotel_name":"Hotel 3","location":{"address":"Address 3","country":"Country 3"},"details":"Desc 3","amenities":{"general":["Gym"]},"images":{"rooms":[],"site":[]},"booking_conditions":[]}]`),
				},
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
					Location: &hotels.HotelLocation{
						Lat:     float64(40.0),
						Lng:     float64(-74.0),
						Address: "Address 1",
						City:    "City 1",
						Country: "Country 1",
					},
					Description: "Desc 1",
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
						Lng:     float64(-80.0),
						Address: "Address 2",
					},
					Description: "Desc 2",
					Amenities: &hotels.HotelAmenities{
						General: []string{"Pool"},
					},
				},
				{
					Id:            "hotel3",
					DestinationId: 789,
					Name:          "Hotel 3",
					Location: &hotels.HotelLocation{
						Address: "Address 3",
						Country: "Country 3",
					},
					Description: "Desc 3",
					Amenities: &hotels.HotelAmenities{
						General: []string{"Gym"},
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
		{
			name: "Success - Parse data from single supplier",
			fields: fields{
				logger:       slog.Default(),
				extSuppliers: nil,
			},
			args: args{
				resp: map[utils.Suppliers]json.RawMessage{
					utils.Acme: json.RawMessage(`[{"Id":"hotel1","DestinationId":123,"Name":"Hotel 1","Latitude":40.0,"Longitude":-74.0,"Address":"Address 1","City":"City 1","Country":"Country 1","Description":"Desc 1","Facilities":["WiFi"]}]`),
				},
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
					Location: &hotels.HotelLocation{
						Lat:     float64(40.0),
						Lng:     float64(-74.0),
						Address: "Address 1",
						City:    "City 1",
						Country: "Country 1",
					},
					Description: "Desc 1",
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Parse empty data from suppliers",
			fields: fields{
				logger:       slog.Default(),
				extSuppliers: nil,
			},
			args: args{
				resp: map[utils.Suppliers]json.RawMessage{
					utils.Acme:       json.RawMessage(`[]`),
					utils.Patagonia:  json.RawMessage(`[]`),
					utils.Paperflies: json.RawMessage(`[]`),
				},
			},
			want:    []hotels.Hotel{},
			wantErr: false,
		},
		{
			name: "Success - Parse mixed empty and populated data",
			fields: fields{
				logger:       slog.Default(),
				extSuppliers: nil,
			},
			args: args{
				resp: map[utils.Suppliers]json.RawMessage{
					utils.Acme:       json.RawMessage(`[]`),
					utils.Patagonia:  json.RawMessage(`[{"id":"hotel2","destination":456,"name":"Hotel 2","lat":30.0,"lng":-80.0,"address":"Address 2","info":"Desc 2","amenities":["Pool"]}]`),
					utils.Paperflies: json.RawMessage(`[]`),
				},
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel2",
					DestinationId: 456,
					Name:          "Hotel 2",
					Location: &hotels.HotelLocation{
						Lat:     float64(30.0),
						Lng:     float64(-80.0),
						Address: "Address 2",
					},
					Description: "Desc 2",
					Amenities: &hotels.HotelAmenities{
						General: []string{"Pool"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Error - Invalid JSON from one supplier",
			fields: fields{
				logger:       slog.Default(),
				extSuppliers: nil,
			},
			args: args{
				resp: map[utils.Suppliers]json.RawMessage{
					utils.Acme:       json.RawMessage(`[{"Id":"hotel1","DestinationId":123,"Name":"Hotel 1"}]`),
					utils.Patagonia:  json.RawMessage(`{"invalid": json`),
					utils.Paperflies: json.RawMessage(`[]`),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success - Empty response map",
			fields: fields{
				logger:       slog.Default(),
				extSuppliers: nil,
			},
			args: args{
				resp: map[utils.Suppliers]json.RawMessage{},
			},
			want:    []hotels.Hotel{},
			wantErr: false,
		},
		{
			name: "Success - Parse with nil logger",
			fields: fields{
				logger:       nil,
				extSuppliers: nil,
			},
			args: args{
				resp: map[utils.Suppliers]json.RawMessage{
					utils.Acme: json.RawMessage(`[{"Id":"hotel1","DestinationId":123,"Name":"Hotel 1","Latitude":40.0,"Longitude":-74.0,"Address":"Address 1","City":"City 1","Country":"Country 1","Description":"Desc 1","Facilities":["WiFi"]}]`),
				},
			},
			want: []hotels.Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
					Location: &hotels.HotelLocation{
						Lat:     float64(40.0),
						Lng:     float64(-74.0),
						Address: "Address 1",
						City:    "City 1",
						Country: "Country 1",
					},
					Description: "Desc 1",
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
			i := &intParser{
				logger:       tt.fields.logger,
				extSuppliers: tt.fields.extSuppliers,
			}
			got, err := i.ParseSuppliersData(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSuppliersData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("ParseSuppliersData() got %d hotels, want %d", len(got), len(tt.want))
			} else {
				gotMap := make(map[string]hotels.Hotel)
				wantMap := make(map[string]hotels.Hotel)

				for _, hotel := range got {
					gotMap[hotel.Id] = hotel
				}
				for _, hotel := range tt.want {
					wantMap[hotel.Id] = hotel
				}

				for id, wantHotel := range wantMap {
					if gotHotel, exists := gotMap[id]; !exists {
						t.Errorf("ParseSuppliersData() missing hotel with ID %s", id)
					} else {
						if gotHotel.Id != wantHotel.Id {
							t.Errorf("Hotel %s.Id = %s, want %s", id, gotHotel.Id, wantHotel.Id)
						}
						if gotHotel.DestinationId != wantHotel.DestinationId {
							t.Errorf("Hotel %s.DestinationId = %d, want %d", id, gotHotel.DestinationId, wantHotel.DestinationId)
						}
						if gotHotel.Name != wantHotel.Name {
							t.Errorf("Hotel %s.Name = %s, want %s", id, gotHotel.Name, wantHotel.Name)
						}
						if gotHotel.Description != wantHotel.Description {
							t.Errorf("Hotel %s.Description = %s, want %s", id, gotHotel.Description, wantHotel.Description)
						}
					}
				}
			}
		})
	}
}
