package server

import (
	"context"
	"errors"
	"log/slog"
	"reflect"
	"testing"

	"hotelsDataMerge/external"
	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/proto"
)

type mockHotels struct {
	hotels []hotels.Hotel
	err    error
}

func (m *mockHotels) GetHotels(hotelIDs []string, destinationID uint64) ([]hotels.Hotel, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.hotels, nil
}

var (
	testHotel = hotels.Hotel{
		Id:            "SjyX",
		DestinationId: 123,
		Name:          "Test Hotel",
		Location: &hotels.HotelLocation{
			Lat:     float64(40.7128),
			Lng:     float64(-74.0060),
			Address: "123 Test St",
			City:    "Test City",
			Country: "Test Country",
		},
		Description: "A test hotel",
		Amenities: &hotels.HotelAmenities{
			General: []string{"WiFi", "Pool"},
			Room:    []string{"TV", "AC"},
		},
		Images: &hotels.HotelImages{
			Rooms: []hotels.HotelImageDetails{
				{Link: "http://example.com/room1.jpg", Description: "Room 1"},
			},
			Site: []hotels.HotelImageDetails{
				{Link: "http://example.com/site1.jpg", Description: "Site 1"},
			},
			Amenities: []hotels.HotelImageDetails{
				{Link: "http://example.com/amenity1.jpg", Description: "Amenity 1"},
			},
		},
		BookingConditions: []string{"No smoking", "No pets"},
	}

	testHotelWithNilLocation = hotels.Hotel{
		Id:                "NilLoc",
		DestinationId:     456,
		Name:              "Hotel with nil location",
		Location:          nil,
		Description:       "Hotel without location",
		Amenities:         nil,
		Images:            nil,
		BookingConditions: []string{},
	}

	testHotelWithEmptyStrings = hotels.Hotel{
		Id:            "EmptyStr",
		DestinationId: 789,
		Name:          "Hotel with empty strings",
		Location: &hotels.HotelLocation{
			Lat:     float64(0),
			Lng:     float64(0),
			Address: "",
			City:    "",
			Country: "",
		},
		Description: "",
		Amenities: &hotels.HotelAmenities{
			General: []string{},
			Room:    []string{},
		},
		Images: &hotels.HotelImages{
			Rooms:     []hotels.HotelImageDetails{},
			Site:      []hotels.HotelImageDetails{},
			Amenities: []hotels.HotelImageDetails{},
		},
		BookingConditions: []string{},
	}
)

func setupTestMaps() {
	hotels.SaveMaps(map[string]hotels.Hotel{})

	hotels.SaveMaps(map[string]hotels.Hotel{
		"SjyX":     testHotel,
		"NilLoc":   testHotelWithNilLocation,
		"EmptyStr": testHotelWithEmptyStrings,
	})
}

func Test_hotelsDataMergeService_GetHotels(t *testing.T) {
	setupTestMaps()

	type fields struct {
		logger                            *slog.Logger
		hotels                            hotels.IntHotels
		UnimplementedHotelDataMergeServer proto.UnimplementedHotelDataMergeServer
	}
	type args struct {
		ctx context.Context
		req *proto.GetHotelsRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResp     *proto.GetHotelsResponse
		wantErr      bool
		setupMutex   func()
		cleanupMutex func()
	}{
		{
			name: "Success - Get hotel by ID",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: []hotels.Hotel{testHotel},
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{"SjyX"},
					DestinationId: 0,
				},
			},
			wantResp: &proto.GetHotelsResponse{
				Hotels: []*proto.Hotel{
					{
						Id:            "SjyX",
						DestinationId: 123,
						Name:          "Test Hotel",
						Location: &proto.Location{
							Lat:     40.7128,
							Lng:     -74.0060,
							Address: "123 Test St",
							City:    "Test City",
							Country: "Test Country",
						},
						Description: "A test hotel",
						Amenities: &proto.HotelAmenities{
							General: []string{"WiFi", "Pool"},
							Room:    []string{"TV", "AC"},
						},
						Images: &proto.Image{
							Rooms: []*proto.Room{
								{Link: "http://example.com/room1.jpg", Description: "Room 1"},
							},
							Site: []*proto.Site{
								{Link: "http://example.com/site1.jpg", Description: "Site 1"},
							},
							Amenities: []*proto.ImageAmenity{
								{Link: "http://example.com/amenity1.jpg", Description: "Amenity 1"},
							},
						},
						BookingConditions: []string{"No smoking", "No pets"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Get hotel by destination ID",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: []hotels.Hotel{testHotel},
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{},
					DestinationId: 123,
				},
			},
			wantResp: &proto.GetHotelsResponse{
				Hotels: []*proto.Hotel{
					{
						Id:            "SjyX",
						DestinationId: 123,
						Name:          "Test Hotel",
						Location: &proto.Location{
							Lat:     40.7128,
							Lng:     -74.0060,
							Address: "123 Test St",
							City:    "Test City",
							Country: "Test Country",
						},
						Description: "A test hotel",
						Amenities: &proto.HotelAmenities{
							General: []string{"WiFi", "Pool"},
							Room:    []string{"TV", "AC"},
						},
						Images: &proto.Image{
							Rooms: []*proto.Room{
								{Link: "http://example.com/room1.jpg", Description: "Room 1"},
							},
							Site: []*proto.Site{
								{Link: "http://example.com/site1.jpg", Description: "Site 1"},
							},
							Amenities: []*proto.ImageAmenity{
								{Link: "http://example.com/amenity1.jpg", Description: "Amenity 1"},
							},
						},
						BookingConditions: []string{"No smoking", "No pets"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Hotel with nil location",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: []hotels.Hotel{testHotelWithNilLocation},
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{"NilLoc"},
					DestinationId: 0,
				},
			},
			wantResp: &proto.GetHotelsResponse{
				Hotels: []*proto.Hotel{
					{
						Id:                "NilLoc",
						DestinationId:     456,
						Name:              "Hotel with nil location",
						Location:          &proto.Location{},
						Description:       "Hotel without location",
						Amenities:         &proto.HotelAmenities{},
						Images:            &proto.Image{},
						BookingConditions: []string{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Hotel with empty strings",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: []hotels.Hotel{testHotelWithEmptyStrings},
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{"EmptyStr"},
					DestinationId: 0,
				},
			},
			wantResp: &proto.GetHotelsResponse{
				Hotels: []*proto.Hotel{
					{
						Id:            "EmptyStr",
						DestinationId: 789,
						Name:          "Hotel with empty strings",
						Location: &proto.Location{
							Lat:     0,
							Lng:     0,
							Address: "",
							City:    "",
							Country: "",
						},
						Description: "",
						Amenities: &proto.HotelAmenities{
							General: []string{},
							Room:    []string{},
						},
						Images: &proto.Image{
							Rooms:     []*proto.Room{},
							Site:      []*proto.Site{},
							Amenities: []*proto.ImageAmenity{},
						},
						BookingConditions: []string{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success - Empty hotels list",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: []hotels.Hotel{},
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{},
					DestinationId: 123,
				},
			},
			wantResp: nil,
			wantErr:  false,
		},
		{
			name: "Error - Mutex lock failure",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: []hotels.Hotel{testHotel},
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{"SjyX"},
					DestinationId: 0,
				},
			},
			wantResp: nil,
			wantErr:  true,
			setupMutex: func() {
				external.FetchSuppliersMutex.Lock()
			},
			cleanupMutex: func() {
				external.FetchSuppliersMutex.Unlock()
			},
		},
		{
			name: "Error - Validation: no request parameters",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: []hotels.Hotel{testHotel},
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{},
					DestinationId: 0,
				},
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "Error - Validation: invalid hotel ID",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: []hotels.Hotel{testHotel},
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{"InvalidID"},
					DestinationId: 0,
				},
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "Error - Validation: invalid destination ID",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: []hotels.Hotel{testHotel},
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{},
					DestinationId: 999,
				},
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "Error - Hotels service error",
			fields: fields{
				logger: slog.Default(),
				hotels: &mockHotels{
					hotels: nil,
					err:    errors.New("database connection failed"),
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetHotelsRequest{
					HotelIDs:      []string{"SjyX"},
					DestinationId: 0,
				},
			},
			wantResp: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupMutex != nil {
				tt.setupMutex()
			}
			if tt.cleanupMutex != nil {
				defer tt.cleanupMutex()
			}

			h := &hotelsDataMergeService{
				logger:                            tt.fields.logger,
				hotels:                            tt.fields.hotels,
				UnimplementedHotelDataMergeServer: tt.fields.UnimplementedHotelDataMergeServer,
			}
			gotResp, err := h.GetHotels(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHotels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				if gotResp == nil && tt.wantResp == nil {
					return
				}
				if gotResp == nil || tt.wantResp == nil {
					t.Errorf("GetHotels() gotResp = %v, want %v", gotResp, tt.wantResp)
					return
				}

				if len(gotResp.Hotels) != len(tt.wantResp.Hotels) {
					t.Errorf("GetHotels() got %d hotels, want %d", len(gotResp.Hotels), len(tt.wantResp.Hotels))
					return
				}

				for i, gotHotel := range gotResp.Hotels {
					wantHotel := tt.wantResp.Hotels[i]
					if gotHotel.Id != wantHotel.Id {
						t.Errorf("Hotel[%d].Id = %s, want %s", i, gotHotel.Id, wantHotel.Id)
					}
					if gotHotel.DestinationId != wantHotel.DestinationId {
						t.Errorf("Hotel[%d].DestinationId = %d, want %d", i, gotHotel.DestinationId, wantHotel.DestinationId)
					}
					if gotHotel.Name != wantHotel.Name {
						t.Errorf("Hotel[%d].Name = %s, want %s", i, gotHotel.Name, wantHotel.Name)
					}
					if gotHotel.Description != wantHotel.Description {
						t.Errorf("Hotel[%d].Description = %s, want %s", i, gotHotel.Description, wantHotel.Description)
					}

					if gotHotel.Location != nil && wantHotel.Location != nil {
						if gotHotel.Location.Lat != wantHotel.Location.Lat {
							t.Errorf("Hotel[%d].Location.Lat = %f, want %f", i, gotHotel.Location.Lat, wantHotel.Location.Lat)
						}
						if gotHotel.Location.Lng != wantHotel.Location.Lng {
							t.Errorf("Hotel[%d].Location.Lng = %f, want %f", i, gotHotel.Location.Lng, wantHotel.Location.Lng)
						}
						if gotHotel.Location.Address != wantHotel.Location.Address {
							t.Errorf("Hotel[%d].Location.Address = %s, want %s", i, gotHotel.Location.Address, wantHotel.Location.Address)
						}
						if gotHotel.Location.City != wantHotel.Location.City {
							t.Errorf("Hotel[%d].Location.City = %s, want %s", i, gotHotel.Location.City, wantHotel.Location.City)
						}
						if gotHotel.Location.Country != wantHotel.Location.Country {
							t.Errorf("Hotel[%d].Location.Country = %s, want %s", i, gotHotel.Location.Country, wantHotel.Location.Country)
						}
					}

					if gotHotel.Amenities != nil && wantHotel.Amenities != nil {
						if len(gotHotel.Amenities.General) != len(wantHotel.Amenities.General) {
							t.Errorf("Hotel[%d].Amenities.General count = %d, want %d", i, len(gotHotel.Amenities.General), len(wantHotel.Amenities.Room))
						}
						if len(gotHotel.Amenities.Room) != len(wantHotel.Amenities.Room) {
							t.Errorf("Hotel[%d].Amenities.Room count = %d, want %d", i, len(gotHotel.Amenities.Room), len(wantHotel.Amenities.Room))
						}
					}

					if gotHotel.Images != nil && wantHotel.Images != nil {
						if len(gotHotel.Images.Rooms) != len(wantHotel.Images.Rooms) {
							t.Errorf("Hotel[%d].Images.Rooms count = %d, want %d", i, len(gotHotel.Images.Rooms), len(wantHotel.Images.Rooms))
						}
						if len(gotHotel.Images.Site) != len(wantHotel.Images.Site) {
							t.Errorf("Hotel[%d].Images.Site count = %d, want %d", i, len(gotHotel.Images.Site), len(wantHotel.Images.Site))
						}
						if len(gotHotel.Images.Amenities) != len(wantHotel.Images.Amenities) {
							t.Errorf("Hotel[%d].Images.Amenities count = %d, want %d", i, len(gotHotel.Images.Amenities), len(wantHotel.Images.Amenities))
						}
					}

					if !reflect.DeepEqual(gotHotel.BookingConditions, wantHotel.BookingConditions) {
						t.Errorf("Hotel[%d].BookingConditions = %v, want %v", i, gotHotel.BookingConditions, wantHotel.BookingConditions)
					}
				}
			}
		})
	}
}

func Test_constructRoomImageDetails(t *testing.T) {
	imageDetails := []hotels.HotelImageDetails{
		{Link: "http://example.com/room1.jpg", Description: "Room 1"},
		{Link: "http://example.com/room2.jpg", Description: "Room 2"},
	}

	result := constructRoomImageDetails(imageDetails)

	expected := []*proto.Room{
		{Link: "http://example.com/room1.jpg", Description: "Room 1"},
		{Link: "http://example.com/room2.jpg", Description: "Room 2"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("constructRoomImageDetails() = %v, want %v", result, expected)
	}
}

func Test_constructSiteImageDetails(t *testing.T) {
	imageDetails := []hotels.HotelImageDetails{
		{Link: "http://example.com/site1.jpg", Description: "Site 1"},
		{Link: "http://example.com/site2.jpg", Description: "Site 2"},
	}

	result := constructSiteImageDetails(imageDetails)

	expected := []*proto.Site{
		{Link: "http://example.com/site1.jpg", Description: "Site 1"},
		{Link: "http://example.com/site2.jpg", Description: "Site 2"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("constructSiteImageDetails() = %v, want %v", result, expected)
	}
}

func Test_constructAmenitiesImageDetails(t *testing.T) {
	imageDetails := []hotels.HotelImageDetails{
		{Link: "http://example.com/amenity1.jpg", Description: "Amenity 1"},
		{Link: "http://example.com/amenity2.jpg", Description: "Amenity 2"},
	}

	result := constructAmenitiesImageDetails(imageDetails)

	expected := []*proto.ImageAmenity{
		{Link: "http://example.com/amenity1.jpg", Description: "Amenity 1"},
		{Link: "http://example.com/amenity2.jpg", Description: "Amenity 2"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("constructAmenitiesImageDetails() = %v, want %v", result, expected)
	}
}

func Test_constructResponse_EmptyHotels(t *testing.T) {
	h := &hotelsDataMergeService{
		logger: slog.Default(),
	}

	result := h.constructResponse([]hotels.Hotel{})

	if result != nil {
		t.Errorf("constructResponse() with empty hotels should return nil, got %v", result)
	}
}

func Test_constructResponse_MultipleHotels(t *testing.T) {
	h := &hotelsDataMergeService{
		logger: slog.Default(),
	}

	hotels := []hotels.Hotel{
		{
			Id:            "Hotel1",
			DestinationId: 1,
			Name:          "First Hotel",
			Location: &hotels.HotelLocation{
				Lat: float64(10.0),
				Lng: float64(20.0),
			},
		},
		{
			Id:            "Hotel2",
			DestinationId: 2,
			Name:          "Second Hotel",
			Location:      nil,
		},
	}

	result := h.constructResponse(hotels)

	//nolint:staticcheck
	if result == nil {
		t.Fatal("constructResponse() should not return nil")
	}

	//nolint:staticcheck
	hotelsResult := result.Hotels
	if hotelsResult == nil {
		t.Fatal("constructResponse() should not return nil Hotels")
	}

	if len(hotelsResult) != 2 {
		t.Errorf("Expected 2 hotels, got %d", len(hotelsResult))
	}

	if hotelsResult[0].Id != "Hotel1" {
		t.Errorf("Expected first hotel ID to be 'Hotel1', got %s", hotelsResult[0].Id)
	}

	if hotelsResult[1].Id != "Hotel2" {
		t.Errorf("Expected second hotel ID to be 'Hotel2', got %s", hotelsResult[1].Id)
	}
}
