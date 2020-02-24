package utils

import (
	"math"
	"time"
)

// GetNightsCount returns the nights stay between dates
func GetNightsCount(checkIn, checkOut time.Time) int {
	checkIn = time.Date(checkIn.Year(), checkIn.Month(), checkIn.Day(), 0, 0, 0, 0, checkIn.Location())
	checkOut = time.Date(checkOut.Year(), checkOut.Month(), checkOut.Day(), 0, 0, 0, 0, checkOut.Location())
	diff := int(math.Round(checkOut.Sub(checkIn).Hours()))
	nights := diff / 24

	if diff%24 > 0 {
		nights++
	}

	return nights
}
