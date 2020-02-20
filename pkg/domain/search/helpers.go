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
