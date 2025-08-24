package hotel

import (
	"reflect"
	"testing"

	"hotelsDataMerge/internal/hotels"
)

func Test_filterAmenities(t *testing.T) {
	type args struct {
		generalAmenities []string
		roomAmenities    []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "Success - Filter overlapping amenities",
			args: args{
				generalAmenities: []string{"WiFi", "Pool", "Gym"},
				roomAmenities:    []string{"TV", "WiFi", "AC", "Pool"},
			},
			want:  []string{"Gym"},
			want1: []string{"TV", "WiFi", "AC", "Pool"},
		},
		{
			name: "Success - No overlapping amenities",
			args: args{
				generalAmenities: []string{"WiFi", "Pool"},
				roomAmenities:    []string{"TV", "AC"},
			},
			want:  []string{"WiFi", "Pool"},
			want1: []string{"TV", "AC"},
		},
		{
			name: "Success - Empty general amenities",
			args: args{
				generalAmenities: []string{},
				roomAmenities:    []string{"TV", "AC", "WiFi"},
			},
			want:  []string{"TV", "AC", "WiFi"},
			want1: []string{},
		},
		{
			name: "Success - Empty room amenities",
			args: args{
				generalAmenities: []string{"WiFi", "Pool"},
				roomAmenities:    []string{},
			},
			want:  []string{"WiFi", "Pool"},
			want1: []string{},
		},
		{
			name: "Success - Both empty",
			args: args{
				generalAmenities: []string{},
				roomAmenities:    []string{},
			},
			want:  []string{},
			want1: []string{},
		},
		{
			name: "Success - All room amenities overlap",
			args: args{
				generalAmenities: []string{"WiFi", "Pool", "TV", "AC"},
				roomAmenities:    []string{"TV", "AC", "WiFi", "Pool"},
			},
			want:  []string{},
			want1: []string{"TV", "AC", "WiFi", "Pool"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := filterAmenities(tt.args.generalAmenities, tt.args.roomAmenities)
			if len(got) != len(tt.want) {
				t.Errorf("filterAmenities() got %d items, want %d", len(got), len(tt.want))
			} else if len(got) > 0 {
				gotMap := make(map[string]bool)
				wantMap := make(map[string]bool)
				for _, item := range got {
					gotMap[item] = true
				}
				for _, item := range tt.want {
					wantMap[item] = true
				}
				if !reflect.DeepEqual(gotMap, wantMap) {
					t.Errorf("filterAmenities() got = %v, want %v", got, tt.want)
				}
			}

			if len(got1) != len(tt.want1) {
				t.Errorf("filterAmenities() got1 %d items, want %d", len(got1), len(tt.want1))
			} else if len(got1) > 0 {
				got1Map := make(map[string]bool)
				want1Map := make(map[string]bool)
				for _, item := range got1 {
					got1Map[item] = true
				}
				for _, item := range tt.want1 {
					want1Map[item] = true
				}
				if !reflect.DeepEqual(got1Map, want1Map) {
					t.Errorf("filterAmenities() got1 = %v, want %v", got1, tt.want1)
				}
			}
		})
	}
}

func Test_hotelBuilder_WithAmenities(t *testing.T) {
	type fields struct {
		hotel hotels.Hotel
	}
	type args struct {
		existing *hotels.HotelAmenities
		new      *hotels.HotelAmenities
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *hotelBuilder
	}{
		{
			name: "Success - Set new amenities",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: &hotels.HotelAmenities{
					General: []string{"WiFi", "Pool"},
					Room:    []string{"TV", "AC"},
				},
				new: &hotels.HotelAmenities{
					General: []string{"Gym", "Spa"},
					Room:    []string{"Mini Bar", "Safe"},
				},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Amenities: &hotels.HotelAmenities{
						General: []string{"wifi", "pool", "gym", "spa"},
						Room:    []string{"tv", "ac", "mini bar", "safe"},
					},
				},
			},
		},
		{
			name: "Success - Set nil new amenities",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: &hotels.HotelAmenities{
					General: []string{"WiFi", "Pool"},
					Room:    []string{"TV", "AC"},
				},
				new: nil,
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Amenities: &hotels.HotelAmenities{
						General: []string{"WiFi", "Pool"},
						Room:    []string{"TV", "AC"},
					},
				},
			},
		},
		{
			name: "Success - Set nil existing amenities",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: nil,
				new: &hotels.HotelAmenities{
					General: []string{"Gym", "Spa"},
					Room:    []string{"Mini Bar", "Safe"},
				},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Amenities: &hotels.HotelAmenities{
						General: []string{"Gym", "Spa"},
						Room:    []string{"Mini Bar", "Safe"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hotelBuilder{
				hotel: tt.fields.hotel,
			}
			got := b.WithAmenities(tt.args.existing, tt.args.new)
			if got.hotel.Amenities == nil && tt.want.hotel.Amenities == nil {
			} else if got.hotel.Amenities == nil || tt.want.hotel.Amenities == nil {
				t.Errorf("WithAmenities() amenities mismatch: got %v, want %v", got.hotel.Amenities, tt.want.hotel.Amenities)
			} else {
				if len(got.hotel.Amenities.General) != len(tt.want.hotel.Amenities.General) {
					t.Errorf("WithAmenities() General amenities count mismatch: got %d, want %d", len(got.hotel.Amenities.General), len(tt.want.hotel.Amenities.General))
				} else {
					gotGeneralMap := make(map[string]bool)
					wantGeneralMap := make(map[string]bool)
					for _, item := range got.hotel.Amenities.General {
						gotGeneralMap[item] = true
					}
					for _, item := range tt.want.hotel.Amenities.General {
						wantGeneralMap[item] = true
					}
					if !reflect.DeepEqual(gotGeneralMap, wantGeneralMap) {
						t.Errorf("WithAmenities() General amenities content mismatch: got = %v, want %v", got.hotel.Amenities.General, tt.want.hotel.Amenities.General)
					}
				}

				if len(got.hotel.Amenities.Room) != len(tt.want.hotel.Amenities.Room) {
					t.Errorf("WithAmenities() Room amenities count mismatch: got %d, want %d", len(got.hotel.Amenities.Room), len(tt.want.hotel.Amenities.Room))
				} else {
					gotRoomMap := make(map[string]bool)
					wantRoomMap := make(map[string]bool)
					for _, item := range got.hotel.Amenities.Room {
						gotRoomMap[item] = true
					}
					for _, item := range tt.want.hotel.Amenities.Room {
						wantRoomMap[item] = true
					}
					if !reflect.DeepEqual(gotRoomMap, wantRoomMap) {
						t.Errorf("WithAmenities() Room amenities content mismatch: got = %v, want %v", got.hotel.Amenities.Room, tt.want.hotel.Amenities.Room)
					}
				}
			}
		})
	}
}

func Test_hotelBuilder_WithBookingConditions(t *testing.T) {
	type fields struct {
		hotel hotels.Hotel
	}
	type args struct {
		existing []string
		new      []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *hotelBuilder
	}{
		{
			name: "Success - Set new booking conditions",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: []string{"No smoking", "No pets"},
				new:      []string{"No parties", "Quiet hours"},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					BookingConditions: []string{"No parties", "Quiet hours"},
				},
			},
		},
		{
			name: "Success - Set empty new booking conditions",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: []string{"No smoking", "No pets"},
				new:      []string{},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					BookingConditions: []string{"No smoking", "No pets"},
				},
			},
		},
		{
			name: "Success - Set nil new booking conditions",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: []string{"No smoking", "No pets"},
				new:      nil,
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					BookingConditions: []string{"No smoking", "No pets"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hotelBuilder{
				hotel: tt.fields.hotel,
			}
			if got := b.WithBookingConditions(tt.args.existing, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithBookingConditions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hotelBuilder_WithDescription(t *testing.T) {
	type fields struct {
		hotel hotels.Hotel
	}
	type args struct {
		existing string
		new      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *hotelBuilder
	}{
		{
			name: "Success - Set longer new description",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "A nice hotel",
				new:      "A beautiful luxury hotel with amazing amenities",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Description: "A beautiful luxury hotel with amazing amenities",
				},
			},
		},
		{
			name: "Success - Keep existing description (longer)",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "A beautiful luxury hotel with amazing amenities",
				new:      "A nice hotel",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Description: "A beautiful luxury hotel with amazing amenities",
				},
			},
		},
		{
			name: "Success - Set empty new description",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "A nice hotel",
				new:      "",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Description: "A nice hotel",
				},
			},
		},
		{
			name: "Success - Set empty existing description",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "",
				new:      "A nice hotel",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Description: "A nice hotel",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hotelBuilder{
				hotel: tt.fields.hotel,
			}
			if got := b.WithDescription(tt.args.existing, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hotelBuilder_WithDestinationID(t *testing.T) {
	type fields struct {
		hotel hotels.Hotel
	}
	type args struct {
		existing uint64
		new      uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *hotelBuilder
	}{
		{
			name: "Success - Set new destination ID",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: 123,
				new:      456,
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					DestinationId: 456,
				},
			},
		},
		{
			name: "Success - Keep existing destination ID (new is 0)",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: 123,
				new:      0,
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					DestinationId: 123,
				},
			},
		},
		{
			name: "Success - Set new destination ID (existing is 0)",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: 0,
				new:      456,
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					DestinationId: 456,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hotelBuilder{
				hotel: tt.fields.hotel,
			}
			if got := b.WithDestinationID(tt.args.existing, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDestinationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hotelBuilder_WithID(t *testing.T) {
	type fields struct {
		hotel hotels.Hotel
	}
	type args struct {
		existing string
		new      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *hotelBuilder
	}{
		{
			name: "Success - Set new ID",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "hotel1",
				new:      "hotel2",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Id: "hotel2",
				},
			},
		},
		{
			name: "Success - Keep existing ID (new is empty)",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "hotel1",
				new:      "",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Id: "hotel1",
				},
			},
		},
		{
			name: "Success - Set new ID (existing is empty)",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "",
				new:      "hotel2",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Id: "hotel2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hotelBuilder{
				hotel: tt.fields.hotel,
			}
			if got := b.WithID(tt.args.existing, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hotelBuilder_WithImages(t *testing.T) {
	type fields struct {
		hotel hotels.Hotel
	}
	type args struct {
		existing *hotels.HotelImages
		new      *hotels.HotelImages
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *hotelBuilder
	}{
		{
			name: "Success - Set new images",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: &hotels.HotelImages{
					Rooms: []hotels.HotelImageDetails{
						{Link: "room1.jpg", Description: "Room 1"},
					},
				},
				new: &hotels.HotelImages{
					Rooms: []hotels.HotelImageDetails{
						{Link: "room2.jpg", Description: "Room 2"},
					},
				},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{
							{Link: "room2.jpg", Description: "Room 2"},
						},
					},
				},
			},
		},
		{
			name: "Success - Set nil new images",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: &hotels.HotelImages{
					Rooms: []hotels.HotelImageDetails{
						{Link: "room1.jpg", Description: "Room 1"},
					},
				},
				new: nil,
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{
							{Link: "room1.jpg", Description: "Room 1"},
						},
					},
				},
			},
		},
		{
			name: "Success - Set nil existing images",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: nil,
				new: &hotels.HotelImages{
					Rooms: []hotels.HotelImageDetails{
						{Link: "room2.jpg", Description: "Room 2"},
					},
				},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Images: &hotels.HotelImages{
						Rooms: []hotels.HotelImageDetails{
							{Link: "room2.jpg", Description: "Room 2"},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hotelBuilder{
				hotel: tt.fields.hotel,
			}
			if got := b.WithImages(tt.args.existing, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hotelBuilder_WithLocation(t *testing.T) {
	type fields struct {
		hotel hotels.Hotel
	}
	type args struct {
		existing *hotels.HotelLocation
		new      *hotels.HotelLocation
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *hotelBuilder
	}{
		{
			name: "Success - Set new location",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: &hotels.HotelLocation{
					Address: "Old Address",
					City:    "Old City",
				},
				new: &hotels.HotelLocation{
					Address: "New Address",
					City:    "New City",
				},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Location: &hotels.HotelLocation{
						Address: "New Address",
						City:    "New City",
					},
				},
			},
		},
		{
			name: "Success - Set nil new location",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: &hotels.HotelLocation{
					Address: "Old Address",
					City:    "Old City",
				},
				new: nil,
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Location: &hotels.HotelLocation{
						Address: "Old Address",
						City:    "Old City",
					},
				},
			},
		},
		{
			name: "Success - Set nil existing location",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: nil,
				new: &hotels.HotelLocation{
					Address: "New Address",
					City:    "New City",
				},
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Location: &hotels.HotelLocation{
						Address: "New Address",
						City:    "New City",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hotelBuilder{
				hotel: tt.fields.hotel,
			}
			if got := b.WithLocation(tt.args.existing, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hotelBuilder_WithName(t *testing.T) {
	type fields struct {
		hotel hotels.Hotel
	}
	type args struct {
		existing string
		new      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *hotelBuilder
	}{
		{
			name: "Success - Set longer new name",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "Hotel A",
				new:      "Luxury Hotel A with Spa",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Name: "Luxury Hotel A with Spa",
				},
			},
		},
		{
			name: "Success - Keep existing name (longer)",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "Luxury Hotel A with Spa",
				new:      "Hotel A",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Name: "Luxury Hotel A with Spa",
				},
			},
		},
		{
			name: "Success - Set empty new name",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "Hotel A",
				new:      "",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Name: "Hotel A",
				},
			},
		},
		{
			name: "Success - Set empty existing name",
			fields: fields{
				hotel: hotels.Hotel{},
			},
			args: args{
				existing: "",
				new:      "Hotel A",
			},
			want: &hotelBuilder{
				hotel: hotels.Hotel{
					Name: "Hotel A",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hotelBuilder{
				hotel: tt.fields.hotel,
			}
			if got := b.WithName(tt.args.existing, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeHotelImageDetails(t *testing.T) {
	type args struct {
		existing []hotels.HotelImageDetails
		new      []hotels.HotelImageDetails
	}
	tests := []struct {
		name string
		args args
		want []hotels.HotelImageDetails
	}{
		{
			name: "Success - Return new images when available",
			args: args{
				existing: []hotels.HotelImageDetails{
					{Link: "old1.jpg", Description: "Old 1"},
				},
				new: []hotels.HotelImageDetails{
					{Link: "new1.jpg", Description: "New 1"},
					{Link: "new2.jpg", Description: "New 2"},
				},
			},
			want: []hotels.HotelImageDetails{
				{Link: "new1.jpg", Description: "New 1"},
				{Link: "new2.jpg", Description: "New 2"},
			},
		},
		{
			name: "Success - Return existing images when new is empty",
			args: args{
				existing: []hotels.HotelImageDetails{
					{Link: "old1.jpg", Description: "Old 1"},
					{Link: "old2.jpg", Description: "Old 2"},
				},
				new: []hotels.HotelImageDetails{},
			},
			want: []hotels.HotelImageDetails{
				{Link: "old1.jpg", Description: "Old 1"},
				{Link: "old2.jpg", Description: "Old 2"},
			},
		},
		{
			name: "Success - Return existing images when new is nil",
			args: args{
				existing: []hotels.HotelImageDetails{
					{Link: "old1.jpg", Description: "Old 1"},
				},
				new: nil,
			},
			want: []hotels.HotelImageDetails{
				{Link: "old1.jpg", Description: "Old 1"},
			},
		},
		{
			name: "Success - Return empty when both are empty",
			args: args{
				existing: []hotels.HotelImageDetails{},
				new:      []hotels.HotelImageDetails{},
			},
			want: []hotels.HotelImageDetails{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeHotelImageDetails(tt.args.existing, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeHotelImageDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeStrings(t *testing.T) {
	type args struct {
		existing []string
		new      []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Success - Merge unique strings",
			args: args{
				existing: []string{"WiFi", "Pool"},
				new:      []string{"Gym", "Spa"},
			},
			want: []string{"wifi", "pool", "gym", "spa"},
		},
		{
			name: "Success - Merge with duplicates (case insensitive)",
			args: args{
				existing: []string{"WiFi", "Pool"},
				new:      []string{"wifi", "Gym", "pool"},
			},
			want: []string{"wifi", "pool", "gym"},
		},
		{
			name: "Success - Merge with spaces",
			args: args{
				existing: []string{"  WiFi  ", "Pool"},
				new:      []string{"Gym", "  wifi  "},
			},
			want: []string{"wifi", "pool", "gym"},
		},
		{
			name: "Success - Empty existing",
			args: args{
				existing: []string{},
				new:      []string{"Gym", "Spa"},
			},
			want: []string{"gym", "spa"},
		},
		{
			name: "Success - Empty new",
			args: args{
				existing: []string{"WiFi", "Pool"},
				new:      []string{},
			},
			want: []string{"wifi", "pool"},
		},
		{
			name: "Success - Both empty",
			args: args{
				existing: []string{},
				new:      []string{},
			},
			want: []string{},
		},
		{
			name: "Success - Mixed case and spaces",
			args: args{
				existing: []string{"  WiFi  ", "Pool"},
				new:      []string{"gym", "  SPA  ", "wifi"},
			},
			want: []string{"wifi", "pool", "gym", "spa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeStrings(tt.args.existing, tt.args.new)
			if len(got) != len(tt.want) {
				t.Errorf("mergeStrings() got %d items, want %d", len(got), len(tt.want))
			} else {
				gotMap := make(map[string]bool)
				wantMap := make(map[string]bool)
				for _, item := range got {
					gotMap[item] = true
				}
				for _, item := range tt.want {
					wantMap[item] = true
				}
				if !reflect.DeepEqual(gotMap, wantMap) {
					t.Errorf("mergeStrings() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
