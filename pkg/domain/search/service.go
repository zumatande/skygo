package search

import (
	"context"

	"github.com/Zumata/v3-common/data"
)

// Worker is function type simulating specific search behaviour
type Worker func(context.Context, *data.SearchRequest) (*data.SearchBasic, error)

// Server implements Search service. It maps property codes
// to functions that simulate certain behaviour
type Server struct {
	mapper map[string]Worker
}

// Search ...
func (s *Server) Search(ctx context.Context, req *data.SearchRequest) (*data.SearchBasic, error) {
	
	return nil, nil
}

