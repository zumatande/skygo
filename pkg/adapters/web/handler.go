package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"bitbucket.org/zumata_scripts/skynettravel/pkg/errs"
)

var errFmt = `{error:"%v"}`

// HTTPService is generic contract domain services
// are prescribed to implement so that error handling
// and response can be uniformly handled in web layer
type HTTPService interface {
	Serve(http.ResponseWriter, *http.Request) error
}

// Handler implements http.Handler interface
type Handler struct {
	Service HTTPService
}

// ServeHTTP processes http request
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Serve(w, r)
	if err != nil {
		switch err := err.(type) {
		case errs.ErrRequest:
			writeErrorJSON(w, err, http.StatusBadRequest)
		case errs.ErrDomain:
			writeErrorJSON(w, err, http.StatusUnprocessableEntity)
		case errs.ErrInternal:
			writeErrorJSON(w, err, http.StatusInternalServerError)
		default:
			log.Error().Err(err).Str("url", r.URL.Path).Msg("untyped error")
			writeErrorJSON(w, err, http.StatusInternalServerError)
		}
	}
}

func writeErrorJSON(w http.ResponseWriter, err error, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	errStr := fmt.Sprintf(errFmt, err.Error())
	w.Write(json.RawMessage(errStr))
}
