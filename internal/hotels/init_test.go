package hotels

import (
	"log/slog"
	"reflect"
	"testing"
)

func TestGetDestinationIDsMap(t *testing.T) {
	tests := []struct {
		name string
		want map[uint64]bool
	}{
		{
			name: "Success - Get destination IDs map",
			want: map[uint64]bool{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveMaps(map[string]Hotel{})
			
			if got := GetDestinationIDsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDestinationIDsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHotelIDsMap(t *testing.T) {
	tests := []struct {
		name string
		want map[string]bool
	}{
		{
			name: "Success - Get hotel IDs map",
			want: map[string]bool{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveMaps(map[string]Hotel{})
			
			if got := GetHotelIDsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHotelIDsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitialize(t *testing.T) {
	type args struct {
		logger *slog.Logger
	}
	tests := []struct {
		name string
		args args
		want IntHotels
	}{
		{
			name: "Success - Initialize with logger",
			args: args{
				logger: slog.Default(),
			},
			want: &intHotels{
				logger: slog.Default(),
			},
		},
		{
			name: "Success - Initialize with nil logger",
			args: args{
				logger: nil,
			},
			want: &intHotels{
				logger: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Initialize(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveMaps(t *testing.T) {
	type args struct {
		hotels map[string]Hotel
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success - Save empty hotels map",
			args: args{
				hotels: map[string]Hotel{},
			},
		},
		{
			name: "Success - Save hotels with destinations",
			args: args{
				hotels: map[string]Hotel{
					"hotel1": {
						Id:            "hotel1",
						DestinationId: 123,
						Name:          "Hotel 1",
					},
					"hotel2": {
						Id:            "hotel2",
						DestinationId: 456,
						Name:          "Hotel 2",
					},
					"hotel3": {
						Id:            "hotel3",
						DestinationId: 123,
						Name:          "Hotel 3",
					},
				},
			},
		},
		{
			name: "Success - Save hotels without destinations",
			args: args{
				hotels: map[string]Hotel{
					"hotel1": {
						Id:            "hotel1",
						DestinationId: 0,
						Name:          "Hotel 1",
					},
					"hotel2": {
						Id:            "hotel2",
						DestinationId: 0,
						Name:          "Hotel 2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveMaps(map[string]Hotel{})
			
			SaveMaps(tt.args.hotels)
			
			hotelIDsMap := GetHotelIDsMap()
			for hotelID := range tt.args.hotels {
				if !hotelIDsMap[hotelID] {
					t.Errorf("Hotel ID %s not found in hotelIDsMap", hotelID)
				}
			}
			
			destinationIDsMap := GetDestinationIDsMap()
			for _, hotel := range tt.args.hotels {
				if hotel.DestinationId != 0 {
					if !destinationIDsMap[hotel.DestinationId] {
						t.Errorf("Destination ID %d not found in destinationIDsMap", hotel.DestinationId)
					}
				}
			}
		})
	}
}
