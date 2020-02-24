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

type modifier func(*data.SearchBasic) error

type generator struct {
	baseRate      data.CurrencyValue
	cpGen         utils.CPGenerator
	hotelID       string
	modifiers     []modifier
	supplierRooms map[string]data.SupplierRoomDetails
}

func (g generator) Package(ctx context.Context, req *data.SearchRequest) (*data.SearchBasic, error) {
	nights := utils.GetNightsCount(req.Params.CheckInDate, req.Params.CheckOutDate)
	bedCount := utils.GetBedCount(req)
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
