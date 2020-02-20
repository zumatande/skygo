package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/Zumata/v3-common/data"
)

// GenerateCancellationPolicy returns data.CPBands
// params come in pairs and in order ...float64, time.Time
func GenerateCancellationPolicy(remarks string, rateDeadlinePairs ...interface{}) (data.CancellationPolicy, error) {
	var policy data.CancellationPolicy
	count := len(rateDeadlinePairs)

	if count != 0 && count%2 != 0 {
		return policy, fmt.Errorf("wrong number of params: got %v; rate and deadline must be in pairs", count)
	} else if count == 0 {
		return policy, errors.New("must pass at least one rate and deadline pair of params")
	}

	bands := make([]data.CancellationPolicyBand, count/2)
	for i := range bands {
		penalty, ok := rateDeadlinePairs[i*2].(float64)
		if !ok {
			return policy, errors.New("error parsing rate: expecting type float64")
		}
		dateTo, ok := rateDeadlinePairs[i*2+1].(time.Time)
		if !ok {
			return policy, errors.New("error parsing deadline: expecting type time.Time")
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
			DateTo:            dateTo,
		}
	}

	policy.CancellationPolicyBands = bands
	policy.Remarks = remarks

	return policy, nil
}

// CPFunc is function signature creating basic cancellation policy
type CPFunc func(time.Time) data.CancellationPolicy
