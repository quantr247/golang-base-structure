package main

import (
	"context"
	"fmt"
	"golang-base-structure/config"
	"golang-base-structure/internal/helper/log"
	"golang-base-structure/internal/registry"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	// Please read in README to implement Oracle DB
	// _ "github.com/mattn/go-oci8"
)

func main() {
	registry.BuildDIContainer()

	cfg := registry.GetDependency(registry.ConfigDIName).(*config.Config)
	api := registry.GetDependency(registry.APIDIName).(http.Handler)

	err := log.InitZap(cfg.Base.App, cfg.Base.Environment)
	if err != nil {
		zap.S().Panic("Can't init zap logger", zap.Error(err))
	}

	httpGateway := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTPAddress),
		Handler: api,
	}

	go func() {
		if err := httpGateway.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Panic("HTTPGateway: Failed to listen and serve", zap.Error(err))
		}
	}()

	zap.S().Info("*****RUNNING******")

	signals := make(chan os.Signal, 1)
	shutdown := make(chan bool, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Stop http gateway
		if err := httpGateway.Shutdown(ctx); err != nil {
			zap.S().Fatalw("Failed to shutdown HTTP Gateway", zap.Error(err))
		}
		shutdown <- true
	}()
	<-shutdown
	zap.S().Info("*****SHUTDOWN*****")
}
