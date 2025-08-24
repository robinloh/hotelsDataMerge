package hotel

import (
	"reflect"
	"testing"

	"hotelsDataMerge/internal/hotels"
)

func TestNewHotelBuilder(t *testing.T) {
	type args struct {
		existing hotels.Hotel
	}
	tests := []struct {
		name string
		args args
		want HotelBuilder
	}{
		{
			name: "Success - Create builder with empty hotel",
			args: args{
				existing: hotels.Hotel{},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{},
			},
		},
		{
			name: "Success - Create builder with populated hotel",
			args: args{
				existing: hotels.Hotel{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Test Hotel",
					Description:   "A test hotel",
					Location: &hotels.HotelLocation{
						Address: "123 Test St",
						City:    "Test City",
						Country: "Test Country",
					},
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi", "Pool"},
						Room:    []string{"TV", "AC"},
					},
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{
							{Link: "room1.jpg", Description: "Room 1"},
						},
						Site: []hotels.HotelImageDetails{
							{Link: "site1.jpg", Description: "Site 1"},
						},
					},
					BookingConditions: []string{"No smoking", "No pets"},
				},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Test Hotel",
					Description:   "A test hotel",
					Location: &hotels.HotelLocation{
						Address: "123 Test St",
						City:    "Test City",
						Country: "Test Country",
					},
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi", "Pool"},
						Room:    []string{"TV", "AC"},
					},
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{
							{Link: "room1.jpg", Description: "Room 1"},
						},
						Site: []hotels.HotelImageDetails{
							{Link: "site1.jpg", Description: "Site 1"},
						},
					},
					BookingConditions: []string{"No smoking", "No pets"},
				},
			},
		},
		{
			name: "Success - Create builder with hotel having nil pointers",
			args: args{
				existing: hotels.Hotel{
					Id:            "hotel2",
					DestinationId: 456,
					Name:          "Hotel with nil pointers",
					Description:   "A hotel with nil location, amenities, and images",
				},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Id:            "hotel2",
					DestinationId: 456,
					Name:          "Hotel with nil pointers",
					Description:   "A hotel with nil location, amenities, and images",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHotelBuilder(tt.args.existing); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHotelBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hotelBuilder_Build(t *testing.T) {
	type fields struct {
		hotel hotels.Hotel
	}
	tests := []struct {
		name   string
		fields fields
		want   hotels.Hotel
	}{
		{
			name: "Success - Build empty hotel",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			want: hotels.Hotel{},
		},
		{
			name: "Success - Build populated hotel",
			fields: fields{
				hotel: hotels.Hotel{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Test Hotel",
					Description:   "A test hotel",
					Location: &hotels.HotelLocation{
						Address: "123 Test St",
						City:    "Test City",
						Country: "Test Country",
					},
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi", "Pool"},
						Room:    []string{"TV", "AC"},
					},
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{
							{Link: "room1.jpg", Description: "Room 1"},
						},
						Site: []hotels.HotelImageDetails{
							{Link: "site1.jpg", Description: "Site 1"},
						},
					},
					BookingConditions: []string{"No smoking", "No pets"},
				},
			},
			want: hotels.Hotel{
				Id:            "hotel1",
				DestinationId: 123,
				Name:          "Test Hotel",
				Description:   "A test hotel",
				Location: &hotels.HotelLocation{
					Address: "123 Test St",
					City:    "Test City",
					Country: "Test Country",
				},
				Amenities: &hotels.HotelAmenities{
					General: []string{"WiFi", "Pool"},
					Room:    []string{"TV", "AC"},
				},
				Images: &hotels.HotelImages{
					Rooms: []hotels.HotelImageDetails{
						{Link: "room1.jpg", Description: "Room 1"},
					},
					Site: []hotels.HotelImageDetails{
						{Link: "site1.jpg", Description: "Site 1"},
					},
				},
				BookingConditions: []string{"No smoking", "No pets"},
			},
		},
		{
			name: "Success - Build hotel with minimal data",
			fields: fields{
				hotel: hotels.Hotel{
					Id:            "hotel3",
					DestinationId: 789,
					Name:          "Minimal Hotel",
				},
			},
			want: hotels.Hotel{
				Id:            "hotel3",
				DestinationId: 789,
				Name:          "Minimal Hotel",
			},
		},
		{
			name: "Success - Build hotel with nil pointers",
			fields: fields{
				hotel: hotels.Hotel{
					Id:            "hotel4",
					DestinationId: 999,
					Name:          "Hotel with nil pointers",
					Description:   "A hotel with nil location, amenities, and images",
				},
			},
			want: hotels.Hotel{
				Id:            "hotel4",
				DestinationId: 999,
				Name:          "Hotel with nil pointers",
				Description:   "A hotel with nil location, amenities, and images",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hotelBuilder{
				hotel: tt.fields.hotel,
			}
			if got := b.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
