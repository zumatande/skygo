// Router definition goes there
// Add path here
package httpserver

import (
	"github.com/Zumata/v3-common/middleware/recovery"
	"github.com/justinas/alice"

	"bitbucket.org/zumata_scripts/skynet2/pkg/adapters"
	"bitbucket.org/zumata_scripts/skynet2/pkg/adapters/gateway/suppliercaller/operation"
	httphandler "bitbucket.org/zumata_scripts/skynet2/pkg/adapters/http"
	"bitbucket.org/zumata_scripts/skynet2/pkg/domain"
	"bitbucket.org/zumata_scripts/skynet2/pkg/usecase"
	"bitbucket.org/zumata_scripts/skynet2/pkg/usecase/hotel"
	"bitbucket.org/zumata_scripts/skynet2/pkg/usecase/search"
	"github.com/gorilla/mux"
)

// InitRoutes Initialise routes
func InitRoutes(router *mux.Router, a *adapters.API) {
	recoveryMiddleware := recovery.MiddlewareWrapper(domain.InternalServerError)
	apiKeyChain := alice.New(recoveryMiddleware)

	// router.Handle("/search", apiKeyChain.Then(insert_handler_here)).Methods(http.MethodPost)
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(http.StatusText(http.StatusOK))) }).Methods(http.MethodGet)
}
