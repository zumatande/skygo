package search

import (
	"context"

	"github.com/Zumata/v3-common/data"
)

// Service is search domain interface
type Service interface {
	Search(context.Context, *data.SearchRequest) (*data.SearchBasic, error)
}

// server implements Search service. It maps properties
// to functions that simulate certain behaviour
type server struct {
	mapper map[string]Generator
}

// Search ...
func (s *server) Search(ctx context.Context, req *data.SearchRequest) (*data.SearchBasic, error) {

	return nil, nil
}
