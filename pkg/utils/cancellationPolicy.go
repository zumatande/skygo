package utils

import (
	"fmt"
	"time"

	"github.com/Zumata/v3-common/data"
)

// CPGenerator ...
type CPGenerator func(time.Time) data.CancellationPolicy

// NewCPGenerator returns data.CancellationPolicy; params, if not zero length,
// come in pairs and in order ...float64 rate, deadlineDaysFromCheckIn int;
//
// usage:
// 		cp := NewCancellationPolicyGenerator("be cool, cancel early", 0.0, -60, 50.0, -30, 100.0, -5)
// 		cp := NewCancellationPolicyGenerator("no use cancelling, non-refundable")
//
func NewCPGenerator(remarks string, rateDeadlinePairs ...interface{}) CPGenerator {
	count := len(rateDeadlinePairs)

	if count != 0 && count%2 != 0 {
		panic(fmt.Sprintf("wrong number of params: got %v; rate and deadline must be in pairs", count))
	}

	return func(checkIn time.Time) data.CancellationPolicy {
		bands := make([]data.CancellationPolicyBand, count/2)
		for i := range bands {
			var penalty float64
			switch rateDeadlinePairs[i*2].(type) {
			case int:
				penalty = float64(rateDeadlinePairs[i*2].(int))
			case float64:
				penalty = rateDeadlinePairs[i*2].(float64)
			default:
				panic("expecting penalty rate to be of type float64")
			}

			deadlineDaysFromCheckIn, ok := rateDeadlinePairs[i*2+1].(int)
			if !ok {
				panic("expecting days count to be of type int")
			}

			var dateFrom time.Time
			if i == 0 {
				dateFrom = time.Now().UTC()
			} else {
				dateFrom = bands[i-1].DateTo.Add(time.Second)
			}

			bands[i] = data.CancellationPolicyBand{
				PenaltyPercentage: penalty,
				DateFrom:          dateFrom,
				DateTo:            checkIn.AddDate(0, 0, deadlineDaysFromCheckIn),
			}
		}

		if count == 0 {
			bands = []data.CancellationPolicyBand{
				{
					PenaltyPercentage: 100,
					DateFrom:          time.Now().UTC(),
					DateTo:            checkIn,
				},
			}
		}

		return data.CancellationPolicy{
			Remarks:                 remarks,
			CancellationPolicyBands: bands,
		}
	}
}
