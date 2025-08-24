package suppliers

import (
	"log/slog"

	"hotelsDataMerge/external"
	"hotelsDataMerge/internal/suppliers/fetcher"
	"hotelsDataMerge/internal/suppliers/merger"
	"hotelsDataMerge/internal/suppliers/parser"
)

type IntSuppliers struct {
	Fetcher fetcher.IntFetcher
	Parser  parser.IntParser
	Merger  merger.IntMerger
}

func Initialize(logger *slog.Logger, extSuppliers external.ExtSuppliers) *IntSuppliers {
	return &IntSuppliers{
		Fetcher: fetcher.Initialize(logger, extSuppliers),
		Parser:  parser.Initialize(logger),
		Merger:  merger.Initialize(logger),
	}
}
