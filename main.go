package main

import (
	"net/http"
	"strconv"
	"time"

	"bitbucket.org/zumata_scripts/skynet2/pkg/adapters"
	"bitbucket.org/zumata_scripts/skynet2/pkg/adapters/gateway/repositories/postgres"
	"bitbucket.org/zumata_scripts/skynet2/pkg/domain"
	"bitbucket.org/zumata_scripts/skynet2/pkg/domain/logger"
	"bitbucket.org/zumata_scripts/skynet2/pkg/infrastructure/httpserver"
	"bitbucket.org/zumata_scripts/skynet2/pkg/infrastructure/loggerfactory"
	pgconn "bitbucket.org/zumata_scripts/skynet2/pkg/infrastructure/postgres"
	"bitbucket.org/zumata_scripts/skynet2/pkg/usecase"
	"github.com/Zumata/v3-common/httpclient"
	"github.com/Zumata/v3-common/middleware/recovery"
	"github.com/justinas/alice"
	"github.com/spf13/viper"

	"github.com/Zumata/v3-common/handlers"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/plugin/ochttp"
)

func main() {
	viper.AutomaticEnv()

	// build logger
	err := loggerfactory.GetLogFactoryBuilder(viper.GetString("LOGGER_TYPE")).Build()
	if err != nil {
		log.Fatal().Msgf("error get log factory builder, err = %v", err)
	}

	logger.Log.Info("Creating connection")
	conn, err := pgconn.CreateConn()
	if err != nil {
		logger.Log.Fatalf("error create db connection, err = %s", err)
	}

	var repo usecase.Repository = postgres.NewRepo(conn)

	api := adapters.API{}
	timeout := getTimeout()
	var httpClient httpclient.Client = &http.Client{
		Timeout:   timeout,
		Transport: &ochttp.Transport{},
	}

	api.HTTPClient = httpClient

	stdChain := alice.New(recovery.MiddlewareWrapper(domain.InternalServerError))
	router := handlers.StandardRouter(stdChain)

	httpserver.InitRoutes(router, &api, repo)

	if err := httpserver.Serve(router); err != nil {
		logger.Log.Fatalf("error create server connection, err = %s", err)
	}
}

func getTimeout() time.Duration {
	defaultTimeout := time.Minute

	timeoutStr := viper.GetString("HTTP_CLIENT_TIMEOUT")
	if timeoutStr == "" {
		return defaultTimeout
	}

	timeout, err := strconv.ParseInt(timeoutStr, 10, 64)
	if err != nil {
		return defaultTimeout
	}

	return time.Duration(timeout) * time.Second
}
