package butty

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"url/internal/config"
	"url/internal/storage"
	"url/pkg/logger"
)

var buttyService *Service

type Service struct {
	Cfg            *config.Cfg
	Logger         *zap.Logger
	Storage        storage.Storager
	UrlPostCounter prometheus.Counter
	UrlGetCounter  prometheus.Counter
}

func NewButtyService() {
	bs := Service{}
	buttyService = &bs

	bs.Cfg = config.NewConfig()
	bs.Logger = logger.New(bs.Cfg.Service.LogLevel)

	switch bs.Cfg.Data.Database.Driver {
	case "mysql":
		bs.Logger.Info("NewButtyService", zap.String("database", "mysql"))
		bs.Storage = storage.NewMysqlStorage()
		bs.Logger.Info("Create mysql storage", zap.Any("storage", bs.Storage))
	default:
		bs.Logger.Info("NewButtyService", zap.String("database", "inmemory"))
		bs.Storage, _ = storage.NewInMemoryStorage()
		bs.Logger.Info("Create inmemory storage", zap.Any("storage", bs.Storage))
	}
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

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	go func() {
		err := http.ListenAndServe("0.0.0.0:9000", mux)
		if err != nil {
			bs.Logger.Error("Error start promhttp server", zap.Error(err))
		}
	}()
	bs.Logger.Info("Start prom metrics server")
}

func GetService() *Service {
	if buttyService == nil {
		panic("butty service is nil")
	}
	return buttyService
}

func (bs *Service) Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	for {
		s := <-c
		switch s {
		case syscall.SIGINT:
			bs.Logger.Info("Get signal SIGINT")
			bs.Storage.Close()
		}
	}
}
