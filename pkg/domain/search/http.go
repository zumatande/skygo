package search

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Zumata/v3-common/data"
	"github.com/zumatande/skygo/pkg/adapters/web"
	"github.com/zumatande/skygo/pkg/errs"
)

// HTTP is an HTTPService implementation
type HTTP struct {
	service Service
}

// Serve satisfies HTTPService
func (h *HTTP) Serve(w http.ResponseWriter, r *http.Request) (err error) {
	ctx := r.Context()
	req := new(data.SearchRequest)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errs.ErrRequest{err}
	}
	if err = json.Unmarshal(body, req); err != nil {
		return errs.ErrRequest{err}
	}
	search, err := h.service.Search(ctx, req)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(search)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	return nil
}

// NewHTTPService returns ...
func NewHTTPService(s Service) web.HTTPService {
	if s == nil {
		s = defaultService
	}

	return &HTTP{
		service: s,
	}
}
