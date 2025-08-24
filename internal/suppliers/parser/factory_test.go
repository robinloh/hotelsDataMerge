package parser

import (
	"encoding/json"
	"log/slog"
	"reflect"
	"testing"

	"hotelsDataMerge/internal/suppliers/parser/acme"
	"hotelsDataMerge/internal/suppliers/parser/patagonia"
	"hotelsDataMerge/internal/suppliers/parser/paperflies"
	"hotelsDataMerge/internal/suppliers/utils"
)

func TestDefaultParserFactory_CreateParser(t *testing.T) {
	type fields struct {
		logger *slog.Logger
	}
	type args struct {
		supplierName utils.Suppliers
		rawData      json.RawMessage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ParserFactory
	}{
		{
			name: "Success - Create ACME parser",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierName: utils.Acme,
				rawData:      json.RawMessage(`{"hotels":[]}`),
			},
			want: &acme.AcmeParser{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData:      json.RawMessage(`{"hotels":[]}`),
			},
		},
		{
			name: "Success - Create Patagonia parser",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierName: utils.Patagonia,
				rawData:      json.RawMessage(`{"hotels":[]}`),
			},
			want: &patagonia.PatagoniaParser{
				Logger:       slog.Default(),
				SupplierName: utils.Patagonia,
				RawData:      json.RawMessage(`{"hotels":[]}`),
			},
		},
		{
			name: "Success - Create Paperflies parser",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierName: utils.Paperflies,
				rawData:      json.RawMessage(`{"hotels":[]}`),
			},
			want: &paperflies.PaperfliesParser{
				Logger:       slog.Default(),
				SupplierName: utils.Paperflies,
				RawData:      json.RawMessage(`{"hotels":[]}`),
			},
		},
		{
			name: "Success - Create parser with empty raw data",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierName: utils.Acme,
				rawData:      json.RawMessage{},
			},
			want: &acme.AcmeParser{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData:      json.RawMessage{},
			},
		},
		{
			name: "Success - Create parser with nil raw data",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierName: utils.Patagonia,
				rawData:      nil,
			},
			want: &patagonia.PatagoniaParser{
				Logger:       slog.Default(),
				SupplierName: utils.Patagonia,
				RawData:      nil,
			},
		},
		{
			name: "Success - Create parser with nil logger",
			fields: fields{
				logger: nil,
			},
			args: args{
				supplierName: utils.Paperflies,
				rawData:      json.RawMessage(`{"hotels":[]}`),
			},
			want: &paperflies.PaperfliesParser{
				Logger:       nil,
				SupplierName: utils.Paperflies,
				RawData:      json.RawMessage(`{"hotels":[]}`),
			},
		},
		{
			name: "Success - Create parser with complex raw data",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierName: utils.Acme,
				rawData:      json.RawMessage(`{"hotels":[{"id":"hotel1","name":"Test Hotel"}]}`),
			},
			want: &acme.AcmeParser{
				Logger:       slog.Default(),
				SupplierName: utils.Acme,
				RawData:      json.RawMessage(`{"hotels":[{"id":"hotel1","name":"Test Hotel"}]}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &DefaultParserFactory{
				logger: tt.fields.logger,
			}
			if got := f.CreateParser(tt.args.supplierName, tt.args.rawData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
