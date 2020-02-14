package book

import (
	"context"

	"github.com/Zumata/v3-common/data"
)

// Server implements Book service. It maps property codes
// to functions that simulate certain behaviour
type Server struct {
	mapper map[string]Worker
}

// Worker is function type simulating specific booking behaviour
type Worker func(context.Context, *data.BookRequest) (*data.Booking, error)
