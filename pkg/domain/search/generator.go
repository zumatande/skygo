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

// Generator function type simulating specific property
// search behavior
type Generator interface {
	Package(context.Context, *data.SearchRequest) (*data.SearchBasic, error)
}

type serviceFunc func(context.Context, *data.SearchRequest) (*data.SearchBasic, error)

type modifier func(*data.SearchBasic, error) (*data.SearchBasic, error)

type cpFunc func(time.Time) data.CancellationPolicy

type generator struct {
	modifiers []modifier
	hotelID   string
	baseRate  data.CurrencyValue
	cpGen     utils.CPFunc
}

func (g generator) Package(ctx context.Context, req *data.SearchRequest) (*data.SearchBasic, error) {
	// defp generate_package(hotel_id, room_type, price, bed_count, check_in_date, night_count) do
	//     food_code = Enum.random(1..7)
	//     check_in = Timex.parse!(check_in_date, "{YYYY}-{0M}-{0D}")
	//     total_price =
	//       price
	//       |> add_room_type_surcharge(room_type)
	//       |> add_food_code_surcharge(food_code)
	//
	//     # wrap rounding with division in case amount is integer
	//     sell_rate_amount = Float.round(total_price / 1, 2)
	//     ...
	// end

	nights := req.Params.CheckOutDate.Day() - req.Params.CheckInDate.Day()
	breakDownPrices := generateDailyPrices(g.baseRate, req.Params.CheckInDate, nights)

	var totalValue float64
	for _, p := range breakDownPrices {
		totalValue += p.Price.Value
	}
	supplierSellRate := data.CurrencyValue{
		Currency: g.baseRate.Currency,
		Value:    totalValue,
	}
	pkg := data.BasicPackage{
		Partner:                 partnerSNT,
		Supplier:                partnerSNT,
		SupplierSellRate:        supplierSellRate,
		SupplierPublishRate:     &supplierSellRate,
		SupplierBreakdownPrices: breakDownPrices,
		SupplierHotelID:         hotelID,
		SupplierRoomDetails:     utils.GetSupplierRoomDetails(roomType, bedCount),
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

func newBasicGenerator(hotelID string, searcher serviceFunc) Generator {
	return generator{
		searcher: searcher,
		hotelID:  hotelID,
	}
}
