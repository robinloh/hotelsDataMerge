package server

import (
	"log/slog"
	"reflect"
	"testing"

	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/proto"
)

func TestNewHotelsDataMergeService(t *testing.T) {
	type args struct {
		logger *slog.Logger
	}
	tests := []struct {
		name string
		args args
		want proto.HotelDataMergeServer
	}{
		{
			name: "Success",
			args: args{logger: slog.Default()},
			want: &hotelsDataMergeService{
				logger: slog.Default(),
				hotels: hotels.Initialize(slog.Default()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHotelsDataMergeService(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHotelsDataMergeService() = %v, want %v", got, tt.want)
			}
		})
	}
}
