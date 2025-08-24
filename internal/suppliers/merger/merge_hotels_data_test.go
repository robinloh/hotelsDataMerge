package merger

import (
	"log/slog"
	"reflect"
	"testing"

	"hotelsDataMerge/internal/hotels"
)

func Test_intMerger_MergeHotelsData(t *testing.T) {
	type fields struct {
		logger *slog.Logger
	}
	type args struct {
		mappedData []hotels.Hotel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]hotels.Hotel
	}{
		{
			name: "Success - Merge single hotel",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				mappedData: []hotels.Hotel{
					{
						Id:            "hotel1",
						DestinationId: 123,
						Name:          "Hotel 1",
						Description:   "Description 1",
					},
				},
			},
			want: map[string]hotels.Hotel{
				"hotel1": {
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
					Description:   "Description 1",
				},
			},
		},
		{
			name: "Success - Merge multiple unique hotels",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				mappedData: []hotels.Hotel{
					{
						Id:            "hotel1",
						DestinationId: 123,
						Name:          "Hotel 1",
						Description:   "Description 1",
					},
					{
						Id:            "hotel2",
						DestinationId: 456,
						Name:          "Hotel 2",
						Description:   "Description 2",
					},
				},
			},
			want: map[string]hotels.Hotel{
				"hotel1": {
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
					Description:   "Description 1",
				},
				"hotel2": {
					Id:            "hotel2",
					DestinationId: 456,
					Name:          "Hotel 2",
					Description:   "Description 2",
				},
			},
		},
		{
			name: "Success - Merge hotels with same ID (new values preferred)",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				mappedData: []hotels.Hotel{
					{
						Id:            "hotel1",
						DestinationId: 123,
						Name:          "Hotel 1",
						Description:   "Description 1",
					},
					{
						Id:            "hotel1",
						DestinationId: 789,
						Name:          "Hotel 1 Updated",
						Description:   "Description 1 Updated",
					},
				},
			},
			want: map[string]hotels.Hotel{
				"hotel1": {
					Id:            "hotel1",
					DestinationId: 789,
					Name:          "Hotel 1 Updated",
					Description:   "Description 1 Updated",
				},
			},
		},
		{
			name: "Success - Merge hotels with same ID (last values preferred)",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				mappedData: []hotels.Hotel{
					{
						Id:            "hotel1",
						DestinationId: 123,
						Name:          "Hotel 1",
						Description:   "Description 1",
					},
					{
						Id:            "hotel1",
						DestinationId: 789,
						Name:          "Hotel 1 Updated",
						Description:   "Description 1 Updated",
					},
					{
						Id:            "hotel1",
						DestinationId: 999,
						Name:          "Hotel 1 Final",
						Description:   "Description 1 Final",
					},
				},
			},
			want: map[string]hotels.Hotel{
				"hotel1": {
					Id:            "hotel1",
					DestinationId: 999,
					Name:          "Hotel 1 Updated",
					Description:   "Description 1 Updated",
				},
			},
		},
		{
			name: "Success - Merge empty hotel list",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				mappedData: []hotels.Hotel{},
			},
			want: map[string]hotels.Hotel{},
		},
		{
			name: "Success - Merge hotels with complex data",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				mappedData: []hotels.Hotel{
					{
						Id:            "hotel1",
						DestinationId: 123,
						Name:          "Hotel 1",
						Description:   "Description 1",
						Location: &hotels.HotelLocation{
							Lat:     float64(40.0),
							Lng:     float64(-74.0),
							Address: "Address 1",
							City:    "City 1",
							Country: "Country 1",
						},
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
			},
			want: map[string]hotels.Hotel{
				"hotel1": {
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
					Description:   "Description 1",
					Location: &hotels.HotelLocation{
						Lat:     float64(40.0),
						Lng:     float64(-74.0),
						Address: "Address 1",
						City:    "City 1",
						Country: "Country 1",
					},
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
		},
		{
			name: "Success - Merge with nil logger",
			fields: fields{
				logger: nil,
			},
			args: args{
				mappedData: []hotels.Hotel{
					{
						Id:            "hotel1",
						DestinationId: 123,
						Name:          "Hotel 1",
						Description:   "Description 1",
					},
				},
			},
			want: map[string]hotels.Hotel{
				"hotel1": {
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
					Description:   "Description 1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &intMerger{
				logger: tt.fields.logger,
			}
			if got := i.MergeHotelsData(tt.args.mappedData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeHotelsData() = %v, want %v", got, tt.want)
			}
		})
	}
}
