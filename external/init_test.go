package external

import (
	"log/slog"
	"reflect"
	"testing"

	"hotelsDataMerge/internal/suppliers/utils"
)

func TestGetSuppliersURLMap(t *testing.T) {
	tests := []struct {
		name string
		want map[utils.Suppliers]string
	}{
		{
			name: "Success - Get suppliers URL map",
			want: map[utils.Suppliers]string{
				utils.Acme:       "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme",
				utils.Patagonia:  "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia",
				utils.Paperflies: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSuppliersURLMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSuppliersURLMap() = %v, want %v", got, tt.want)
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
		want ExtSuppliers
	}{
		{
			name: "Success - Initialize with logger",
			args: args{
				logger: slog.Default(),
			},
			want: &externalHandler{
				logger: slog.Default(),
			},
		},
		{
			name: "Success - Initialize with nil logger",
			args: args{
				logger: nil,
			},
			want: &externalHandler{
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
