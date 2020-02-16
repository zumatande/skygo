package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/zumatande/skygo/pkg/infrastructure/httpserver"
)

func main() {
	// TODO: init routes
	router := new(mux.Router)
	httpserver.InitRoutes(router)

	if err := httpserver.Serve(router); err != nil {
		log.Fatal().Msgf("error create server connection, err = %s", err)
	}
}
