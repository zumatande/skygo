package search

import (
	"context"

	"github.com/Zumata/v3-common/data"
)

const (
	partnerSNT = "snt"
)

// Generator interface simulates specific property search behavior
type Generator interface {
	Package(context.Context, *data.SearchRequest) (*data.SearchBasic, error)
}

type generator struct {
	hotelID  string
	packager func(context.Context, *data.SearchRequest) []data.BasicPackage
}

func (g generator) Package(ctx context.Context, req *data.SearchRequest) (*data.SearchBasic, error) {
	search := &data.SearchBasic{
		Packages: data.HotelBasicPackages{
			data.SupplierHotelID(g.hotelID): g.packager(ctx, req),
		},
	}

	return search, nil
}
