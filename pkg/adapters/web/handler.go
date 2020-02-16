package web

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/zumatande/skygo/pkg/errs"
)

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
			WriteErrorJSON(w, err, nil, http.StatusBadRequest)
		case errs.ErrDomain:
			WriteErrorJSON(w, err, nil, http.StatusUnprocessableEntity)
		case errs.ErrInternal:
			WriteErrorJSON(w, err, nil, http.StatusInternalServerError)
		default:
			log.Error().Err(err).Str("url", r.URL.Path).Msg("untyped error")
			WriteErrorJSON(w, err, nil, http.StatusInternalServerError)
		}
	}
}

// WriteErrorJSON writes error response in JSON
func WriteErrorJSON(w http.ResponseWriter, err error, details interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["error"] = err.Error()
	if details != nil {
		resp["details"] = details
	} else {
		resp["details"] = []interface{}{}
	}

	// ignore error here and allow panic; resolve by fixing underlying fault
	// rather than by creating error handler for error handler for error handler
	jm, _ := json.Marshal(resp)
	w.Write(jm)
}
