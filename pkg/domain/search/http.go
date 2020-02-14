package search

import (
	"context"

	"github.com/Zumata/v3-common/data"
)

// Service is search.HTTP domain dependency
type Service interface {
	Search(context.Context, *data.SearchRequest) (*data.SearchBasic, error)
}

// HTTP is an HTTPService implementation
type HTTP struct {
	Service Service
}

// Serve satisfies HTTPService
func (h *HTTP) Serve(w http.ResponseWriter, req *http.Request) (err error) {

	return err
}
