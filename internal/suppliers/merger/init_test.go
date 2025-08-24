package merger

import (
	"log/slog"
	"reflect"
	"testing"
)

func TestInitialize(t *testing.T) {
	type args struct {
		logger *slog.Logger
	}
	tests := []struct {
		name string
		args args
		want IntMerger
	}{
		{
			name: "Success - Initialize with logger",
			args: args{
				logger: slog.Default(),
			},
			want: &intMerger{
				logger: slog.Default(),
			},
		},
		{
			name: "Success - Initialize with nil logger",
			args: args{
				logger: nil,
			},
			want: &intMerger{
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
