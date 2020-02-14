package utils

import (
	"strconv"

	"github.com/Zumata/v3-common/data"
)

// NewCurrencyValue creates CurrencyValue that has been rounded.
func NewCurrencyValue(currency string, value float64) *data.CurrencyValue {
	return &data.CurrencyValue{
		Currency: currency,
		Value:    round(value),
	}
}

// CurrencyRound handle currency round
func round(v float64) float64 {
	s := strconv.FormatFloat(v, 'f', 2, 64)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("this shouldn't error out.")
	}
	return f
}
