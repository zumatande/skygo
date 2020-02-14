package cancel

import (
	"context"

	"github.com/Zumata/v3-common/data"
)

// CancelFn is interface for simulating specific booking behaviour
type CancelFn func(context.Context, *data.CancelRequest) (*data.Cancellation, error)

// CancelProps implements Cancel service. It maps property codes
// to functions that simulate certain behaviour
type CancelProps struct {
	mapper map[string]CancelFn
}

// Cancel ...
func (c *CancelProps) Cancel(ctx context.Context, req *data.CancelRequest) (*data.Cancellation, error) {
	return nil, nil
}
