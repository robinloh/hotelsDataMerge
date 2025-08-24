package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"hotelsDataMerge/external"
	"hotelsDataMerge/internal/hotels"
	"hotelsDataMerge/internal/suppliers"
	"hotelsDataMerge/proto"
	"hotelsDataMerge/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

const fetchSuppliersPeriod = time.Second * 5

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	extSuppliers := external.Initialize(logger)
	intSuppliers := suppliers.Initialize(logger, extSuppliers)

	go func() {
		processSuppliersData(intSuppliers, logger)
		time.Sleep(fetchSuppliersPeriod)
	}()

	svc := server.NewHotelsDataMergeService(logger)
	setupServer(svc, logger)
	setupGrpcGateway(logger)
}

func processSuppliersData(intSuppliers *suppliers.IntSuppliers, logger *slog.Logger) {
	if !external.FetchSuppliersMutex.TryLock() {
		logger.Info("Skipping suppliers data fetch - GetHotels operation in progress")
		return
	}
	defer external.FetchSuppliersMutex.Unlock()

	logger.Info("Starting suppliers data fetch and processing")

	rawResp, err := intSuppliers.Fetcher.GetLatestSupplierData()
	if err != nil {
		logger.Error("Failed to fetch suppliers data", "error", err)
		return
	}

	mappedData, err := intSuppliers.Parser.ParseSuppliersData(rawResp)
	if err != nil {
		logger.Error("Failed to parse and map suppliers data", "error", err)
		return
	}

	mergedHotels := intSuppliers.Merger.MergeHotelsData(mappedData)
	hotels.SaveMaps(mergedHotels)
	logger.Info("Suppliers data fetched and processed successfully")
}

func setupServer(svc proto.HotelDataMergeServer, logger *slog.Logger) {
	svr := grpc.NewServer()
	proto.RegisterHotelDataMergeServer(svr, svc)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "8080"))
	if err != nil {
		log.Panicln("Failed to listen to tcp port", err)
	}
	logger.Info(fmt.Sprintf("gRPC server listening at: %s", lis.Addr().String()))
	go func() {
		if err := svr.Serve(lis); err != nil {
			log.Panicln("Failed to serve", err)
		}
	}()
}

func setupGrpcGateway(logger *slog.Logger) {
	conn, err := grpc.NewClient(
		"0.0.0.0:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	mux := runtime.NewServeMux()
	err = proto.RegisterHotelDataMergeHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8090"),
		Handler: mux,
	}
	logger.Info(fmt.Sprintf("Serving gRPC-Gateway on: %s", gwServer.Addr))
	if err = gwServer.ListenAndServe(); err != nil {
		log.Fatalln("Failed to serve gateway:", err)
	}
}
