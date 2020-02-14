package partner

import (
	"context"

	"github.com/Zumata/v3-common/data"
)

// StatusProps implements Status service. It maps property codes
// to functions that simulate certain behaviour
type StatusProps struct {
	mapper map[string]StatusFn
}

// StatusFn is interface for simulating specific booking behaviour
type StatusFn func(context.Context, *data.BookingStatusRequest) (*data.Booking, error)
