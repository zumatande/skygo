package search

import (
	"context"
	"time"

	"github.com/Zumata/v3-common/data"
	"github.com/zumatande/skygo/pkg/utils"
)

const (
	partnerSNT = "snt"
)

var (
	defaultGenerator = generator{
		hotelID: "",
		baseRate: data.CurrencyValue{
			"USD",
			112.3,
		},
	}
)

// Generator interface simulates specific property search behavior
type Generator interface {
	Package(context.Context, *data.SearchRequest) (*data.SearchBasic, error)
}

type serviceFunc func(context.Context, *data.SearchRequest) (*data.SearchBasic, error)

type modifier func(*data.SearchBasic, error) (*data.SearchBasic, error)

type generator struct {
	modifiers     []modifier
	hotelID       string
	baseRate      data.CurrencyValue
	supplierRooms map[string]data.SupplierRoomDetails
	cpGen         utils.CPGenerator
}

func (g generator) Package(ctx context.Context, req *data.SearchRequest) (*data.SearchBasic, error) {
	// parse request
	// get bed count
	// get room count

	nights := req.Params.CheckOutDate.Day() - req.Params.CheckInDate.Day()
	breakDownPrices := generateDailyPrices(g.baseRate, req.Params.CheckInDate, nights)

	var sumValue float64
	for _, p := range breakDownPrices {
		sumValue += p.Price.Value
	}
	supplierSellRate := data.CurrencyValue{
		Currency: g.baseRate.Currency,
		Value:    sumValue * roomCount,
	}
	pkg := data.BasicPackage{
		Partner:                 partnerSNT,
		Supplier:                partnerSNT,
		SupplierSellRate:        supplierSellRate,
		SupplierPublishRate:     &supplierSellRate,
		SupplierBreakdownPrices: breakDownPrices,
		SupplierHotelID:         hotelID,
		SupplierRoomDetails:     utils.NewSupplierRoomDetails(roomType, bedCount),
		SupplierDetails: map[string]string{
			"room_type": roomType,
		},
		CreatedAt: time.Now(),
		RefreshAt: time.Now(),
	}

	search := &data.SearchBasic{
		Packages: data.HotelBasicPackages{
			data.SupplierHotelID(g.hotelID): []data.BasicPackage{pkg},
		},
	}

	return search, nil
}
