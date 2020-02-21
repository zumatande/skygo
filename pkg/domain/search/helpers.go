package search

import (
	"time"

	"github.com/Zumata/v3-common/data"
)

func comparisonRateFromBasicPackage(data.BasicPackage) data.ComparisonRate {
	return data.ComparisonRate{}
}

func generateDailyPrices(basePrice data.CurrencyValue, checkIn time.Time, nights int) []*data.BreakdownPrice {
	prices := make([]*data.BreakdownPrice, nights)

	for i := 0; i < nights; i++ {
		prices[i] = &data.BreakdownPrice{
			Date: checkIn.Add(i * 24 * time.Hour),
			Price: &data.CurrencyValue{
				Currency: basePrice.Currency,
				Value:    basePrice.Value,
			},
		}
	}

	return prices
}

// get room surcharge for common supplier room details,
// if not in common suppliers, returns input surcharge
func getRoomSurcharge(room data.SupplierRoomDetails, surcharge ...float64) float64 {
	switch room.RoomType {
	case "STANDARD":
		return 1.0
	case "DELUXE":
		return 1.5
	default:
		if surcharge == nil {
			return 1.0
		}
		return surcharge[0]
	}
}
