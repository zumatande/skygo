package search

import (
	"context"

	"github.com/Zumata/v3-common/data"
)

// Service is search domain interface
type Service interface {
	Search(context.Context, *data.SearchRequest) (*data.SearchBasic, error)
}

// worker is function type simulating specific search behaviour
type worker func(context.Context, *data.SearchRequest) (*data.SearchBasic, error)

// server implements Search service. It maps property codes
// to functions that simulate certain behaviour
type server struct {
	mapper map[string]worker
	generic worker
}

// Search ...
func (s *server) Search(ctx context.Context, req *data.SearchRequest) (*data.SearchBasic, error) {

	return s.generic(ctx, req)
}

