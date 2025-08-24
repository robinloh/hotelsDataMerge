package hotels

import (
	"log/slog"
	"reflect"
	"testing"
)

func Test_intHotels_GetHotels(t *testing.T) {
	type fields struct {
		logger *slog.Logger
	}
	type args struct {
		hotelIDs      []string
		destinationID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Hotel
		wantErr bool
	}{
		{
			name: "Success - Get hotels by hotel IDs",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				hotelIDs:      []string{"hotel1", "hotel2"},
				destinationID: 0,
			},
			want: []Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
				},
				{
					Id:            "hotel2",
					DestinationId: 456,
					Name:          "Hotel 2",
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Get hotels by destination ID",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				hotelIDs:      []string{},
				destinationID: 123,
			},
			want: []Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
				},
				{
					Id:            "hotel3",
					DestinationId: 123,
					Name:          "Hotel 3",
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Get hotels by both hotel IDs and destination ID",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				hotelIDs:      []string{"hotel1", "hotel2"},
				destinationID: 123,
			},
			want: []Hotel{
				{
					Id:            "hotel1",
					DestinationId: 123,
					Name:          "Hotel 1",
				},
			},
			wantErr: false,
		},
		{
			name: "Success - No hotels found by hotel IDs",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				hotelIDs:      []string{"nonexistent"},
				destinationID: 0,
			},
			want:    []Hotel{},
			wantErr: false,
		},
		{
			name: "Success - No hotels found by destination ID",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				hotelIDs:      []string{},
				destinationID: 999,
			},
			want:    []Hotel{},
			wantErr: false,
		},
		{
			name: "Success - Empty request returns empty result",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				hotelIDs:      []string{},
				destinationID: 0,
			},
			want:    []Hotel{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveMaps(map[string]Hotel{
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
			})

			i := &intHotels{
				logger: tt.fields.logger,
			}
			got, err := i.GetHotels(tt.args.hotelIDs, tt.args.destinationID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHotels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("GetHotels() got %d hotels, want %d", len(got), len(tt.want))
			} else {
				gotMap := make(map[string]Hotel)
				wantMap := make(map[string]Hotel)
				
				for _, hotel := range got {
					gotMap[hotel.Id] = hotel
				}
				for _, hotel := range tt.want {
					wantMap[hotel.Id] = hotel
				}
				
				for id, wantHotel := range wantMap {
					if gotHotel, exists := gotMap[id]; !exists {
						t.Errorf("GetHotels() missing hotel with ID %s", id)
					} else if !reflect.DeepEqual(gotHotel, wantHotel) {
						t.Errorf("GetHotels() hotel %s = %v, want %v", id, gotHotel, wantHotel)
					}
				}
			}
		})
	}
}
