// Serve / run http server
package httpserver

import (
	"net"
	"net/http"
	"time"

	"github.com/Zumata/go-common/utils"
	"github.com/gorilla/mux"

	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	serverPort = "SERVER_PORT"
)

// Serve request into server
func Serve(router *mux.Router) error {
	port := viper.GetString(serverPort)
	if port == "" {
		port = "3002"
	}

	handler := cors.Default().Handler(router)

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    600 * time.Second,
		WriteTimeout:   600 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Info().Msgf("API Listening at port = %s", port)

	listener, listenErr := net.Listen("tcp", s.Addr)
	if listenErr != nil {
		return listenErr
	}

	sigTermNetworkListenerCloser, newSigTermErr := utils.NewSigTermNetworkListenerCloser()
	if newSigTermErr != nil {
		return newSigTermErr
	}

	go sigTermNetworkListenerCloser.CloseOnNotify(listener)

	return s.Serve(listener)
}
