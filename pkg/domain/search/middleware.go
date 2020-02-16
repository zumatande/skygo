package search

import (
	"context"

	"github.com/go-playground/validator/v10"

	"github.com/Zumata/v3-common/data"

	"github.com/zumatande/skygo/pkg/errs"
)

// Validator is validates search request
type Validator struct {
	service Service
}

// Search ...
func (v Validator) Search(ctx context.Context, req *data.SearchRequest) (*data.SearchBasic, error) {
	// apply `validate` tags and custom validation functions on
	// data.SearchRequest definition
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return nil, errs.ErrRequest{err}
	}

	return v.service.Search(ctx, req)
}
