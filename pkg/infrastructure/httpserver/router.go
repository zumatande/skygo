// Package httpserver is a convenience package for wiring up
// http servers and dependencies
package httpserver

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"

	"github.com/Zumata/v3-common/middleware/recovery"

	"github.com/zumatande/skygo/pkg/adapters/web"
	"github.com/zumatande/skygo/pkg/domain/search"
	"github.com/zumatande/skygo/pkg/errs"
)

// InitRoutes Initialise routes
func InitRoutes(router *mux.Router) {
	recoveryMiddleware := recovery.MiddlewareWrapper(func(_ context.Context, w http.ResponseWriter) {
		web.WriteErrorJSON(w, errs.ErrPanic, nil, http.StatusInternalServerError)
	})
	apiKeyChain := alice.New(recoveryMiddleware)

	// initialize handlers
	shsrv := search.NewHTTPService(nil)
	router.Handle("/search", apiKeyChain.Then(&web.Handler{shsrv})).Methods(http.MethodPost)

	// router.Handle("/search", apiKeyChain.Then(insert_handler_here)).Methods(http.MethodPost)
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(http.StatusText(http.StatusOK))) }).Methods(http.MethodGet)
}
