package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/zumatande/skygo/pkg/infrastructure/httpserver"
)

func main() {
	zerolog.LevelFieldName = "severity"
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	router := new(mux.Router)
	httpserver.InitRoutes(router)

	if err := httpserver.Serve(router); err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}
