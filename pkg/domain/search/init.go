package search

import (
	"context"

	"github.com/Zumata/v3-common/data"
)

var defaultService Service

func init() {
	defaultService = &server{
		mapper: map[string]worker{},
		generic: func(_ context.Context, _ *data.SearchRequest) (*data.SearchBasic, error) {
			return &data.SearchBasic{}, nil
		},
	}
}
