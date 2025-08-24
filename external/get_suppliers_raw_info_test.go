package external

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_externalHandler_GetSuppliersRawInfo(t *testing.T) {
	type fields struct {
		logger *slog.Logger
	}
	type args struct {
		supplierURL string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        json.RawMessage
		wantErr     bool
		setupServer func() (*httptest.Server, string)
	}{
		{
			name: "Success - Get valid JSON response",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierURL: "",
			},
			want:    json.RawMessage(`{"hotels":[{"id":"hotel1","name":"Test Hotel"}]}`),
			wantErr: false,
			setupServer: func() (*httptest.Server, string) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{"hotels":[{"id":"hotel1","name":"Test Hotel"}]}`))
				}))
				return server, server.URL
			},
		},
		{
			name: "Success - Get empty JSON response",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierURL: "",
			},
			want:    json.RawMessage(`{}`),
			wantErr: false,
			setupServer: func() (*httptest.Server, string) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{}`))
				}))
				return server, server.URL
			},
		},
		{
			name: "Success - Get array JSON response",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierURL: "",
			},
			want:    json.RawMessage(`[{"id":"hotel1"},{"id":"hotel2"}]`),
			wantErr: false,
			setupServer: func() (*httptest.Server, string) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`[{"id":"hotel1"},{"id":"hotel2"}]`))
				}))
				return server, server.URL
			},
		},
		{
			name: "Error - Invalid JSON response",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierURL: "",
			},
			want:    nil,
			wantErr: true,
			setupServer: func() (*httptest.Server, string) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{"invalid": json`))
				}))
				return server, server.URL
			},
		},
		{
			name: "Success - HTTP error response still returns valid JSON",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierURL: "",
			},
			want:    json.RawMessage(`{"error":"Internal Server Error"}`),
			wantErr: false,
			setupServer: func() (*httptest.Server, string) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = w.Write([]byte(`{"error":"Internal Server Error"}`))
				}))
				return server, server.URL
			},
		},
		{
			name: "Error - Invalid URL",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				supplierURL: "invalid://url",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var server *httptest.Server
			var serverURL string

			if tt.setupServer != nil {
				server, serverURL = tt.setupServer()
				defer server.Close()
				tt.args.supplierURL = serverURL
			}

			e := &externalHandler{
				logger: tt.fields.logger,
			}
			got, err := e.GetSuppliersRawInfo(tt.args.supplierURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSuppliersRawInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSuppliersRawInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
