package fetcher

import (
	"log/slog"
	"reflect"
	"testing"

	"hotelsDataMerge/external"
)

func TestInitialize(t *testing.T) {
	type args struct {
		logger       *slog.Logger
		extSuppliers external.ExtSuppliers
	}
	tests := []struct {
		name string
		args args
		want IntFetcher
	}{
		{
			name: "Success - Initialize with logger and external suppliers",
			args: args{
				logger:       slog.Default(),
				extSuppliers: external.Initialize(slog.Default()),
			},
			want: &intFetcher{
				logger:       slog.Default(),
				extSuppliers: external.Initialize(slog.Default()),
			},
		},
		{
			name: "Success - Initialize with nil logger",
			args: args{
				logger:       nil,
				extSuppliers: external.Initialize(slog.Default()),
			},
			want: &intFetcher{
				logger:       nil,
				extSuppliers: external.Initialize(slog.Default()),
			},
		},
		{
			name: "Success - Initialize with nil external suppliers",
			args: args{
				logger:       slog.Default(),
				extSuppliers: nil,
			},
			want: &intFetcher{
				logger:       slog.Default(),
				extSuppliers: nil,
			},
		},
		{
			name: "Success - Initialize with both nil values",
			args: args{
				logger:       nil,
				extSuppliers: nil,
			},
			want: &intFetcher{
				logger:       nil,
				extSuppliers: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Initialize(tt.args.logger, tt.args.extSuppliers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
