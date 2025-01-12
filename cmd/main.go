package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"videoverse/internal"
	"videoverse/pkg/config"
	"videoverse/pkg/logbox"
	"videoverse/repository"
	"videoverse/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	internal.RunMigrations()
	
	var logger = logbox.NewLogBox()
	if config.ENV == config.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.Debug().Str("LOG_LEVEL", gin.Mode()).Msg("log level set to DEBUG")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.APP_PORT),
		Handler: routes.NewRouter().SetRoutes(repository.NewRepository()),
	}

	logger.Info().Str("address", server.Addr).Str("event", "STARING_SERVER").Msg("")

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error().Err(err).Msg("failed to start server")
			logger.Fatal().Msgf("failed to start %s service\n", config.SERVICE_NAME)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown
	logger.Info().Msg("SHUTTING_DOWN_SERVER")
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal().Str("event", "SERVER_FORCED_TO_SHUTDOWN").Str("address", server.Addr).Err(err).Msg("server forced shutdonw")
	}
}
