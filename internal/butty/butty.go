// Copyright 2022. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package butty is main app service for butty url
// creation and using.
package butty

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url/internal/config"
	"url/internal/storage"
	"url/pkg/logger"
)

// buttyService is a global variable for access to butty service
var buttyService *Service

// Service is a struct contains main services components
type Service struct {
	// Cfg is a struct for external configuration butty service
	Cfg *config.Cfg

	// Logger main logger services
	Logger *zap.Logger

	// Storage is a common Storage for butty and full links
	Storage storage.Storager

	// UrlPostCounter is counter for created butty links
	UrlPostCounter prometheus.Counter
	PostCounter    uint32

	// UrlGetCounter is counter for used butty links
	UrlGetCounter prometheus.Counter
	GetCounter    uint32
	// Router is an implementation for openapi specs
	Router *gin.Engine

	// ServerButty is a http server for main service
	ServerButty http.Server

	// ServerProm is http server for prometheus metrics
	ServerProm http.Server
}

// NewButtyService is a constructor for butty service
// here create global variable buttyService with
// selected configuration, logger and storage
func NewButtyService() {
	bs := Service{}
	buttyService = &bs

	bs.Cfg = config.NewConfig()
	bs.Logger = logger.New(bs.Cfg.Service.LogLevel)

	switch bs.Cfg.Data.Database.Driver {
	case "mysql":
		bs.Logger.Info("NewButtyService", zap.String("database", "mysql"))
		bs.Storage = storage.NewMysqlStorage()
		bs.Logger.Info("Created mysql storage", zap.Any("storage", bs.Storage))
	default:
		bs.Logger.Info("NewButtyService", zap.String("database", "inmemory"))
		bs.Storage, _ = storage.NewInMemoryStorage()
		bs.Logger.Info("Created inmemory storage", zap.Any("storage", bs.Storage))
	}
	bs.Logger.Debug("service", zap.String("message", "Created new butty service"))
}

// GetService return global service variable
func GetService() *Service {
	if buttyService == nil {
		panic("butty service is nil")
	}
	return buttyService
}

// RunPrometheus start Prometheus http server
func (bs *Service) RunPrometheus() {
	bs.UrlPostCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "butty",
		Name:      "urlPostCounter",
		Help:      "Total butty post url counter",
	})
	prometheus.MustRegister(bs.UrlPostCounter)

	bs.UrlGetCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "butty",
		Name:      "urlGetCounter",
		Help:      "Total butty get url counter",
	})
	prometheus.MustRegister(bs.UrlGetCounter)

	Server := http.NewServeMux()
	Server.Handle("/metrics", promhttp.Handler())

	bs.ServerProm = http.Server{
		Addr:    ":9000",
		Handler: Server,
	}

	err := bs.ServerProm.ListenAndServe()

	if err != nil {
		bs.Logger.Fatal("Error start prom http server", zap.Error(err))
	}

	bs.Logger.Info("Start prom metrics server")
}

// RunButty start butty http server
func (bs *Service) RunButty() {
	bs.Router.StaticFS("/website", http.Dir("./website"))
	bs.Logger.Debug("gin", zap.String("message", "gin created"))
	bs.Logger.Fatal("Error start gin http api", zap.Error(bs.Router.Run("0.0.0.0:"+bs.Cfg.Server.Http.Port)))
}

// Run start app with main services and wait SIGINT for gracefully shutdown
func (bs *Service) Run() {
	go bs.RunPrometheus()
	go bs.RunButty()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	<-c

	bs.Logger.Info("Get signal SIGINT. Shutdown service...")
	bs.Storage.Close()
	bs.Logger.Info("Storage closed.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := bs.ServerProm.Shutdown(ctx); err != nil {
		bs.Logger.Fatal("Prometheus forced to shutdown: ", zap.Error(err))
	}
	bs.Logger.Info("Prometheus server shutdown.")

	ctx1, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel1()

	if err := bs.ServerButty.Shutdown(ctx1); err != nil {
		bs.Logger.Fatal("Server forced to shutdown: ", zap.Error(err))
	}
	bs.Logger.Info("Server shutdown.")
	logger.CloseLogger()
}
