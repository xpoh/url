package butty

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"url/cmd/url"
	"url/internal/config"
	"url/internal/storage"
)

type Service struct {
	cfg    *config.Cfg
	logger *zap.Logger

	urlIn  chan string
	urlOut chan string
	rwmux  sync.RWMutex
	rmux   sync.Mutex

	storage storage.Storager
}

func NewButtyService() *Service {
	bs := Service{}
	cfg := config.NewConfig()
	bs.cfg = cfg
	bs.urlOut = make(chan string, 1)
	bs.urlIn = make(chan string, 1)
	bs.storage, _ = storage.NewInMemoryStorage()
	return &bs
}

func (bs *Service) InitLogger() {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)
	// TODO make loglevel
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	var err error

	bs.logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
	bs.logger.Info("logger construction succeeded")
}

func (bs *Service) StartGin() {
	log.Printf("Server started")
	router := url.NewRouter()
	log.Fatal(router.Run(bs.cfg.Server.Http.Addr))
}

func (bs *Service) Run() {
	bs.InitLogger()
	defer func() {
		err := bs.logger.Sync()
		if err != nil {

		}
	}()
	bs.logger.Info("Using zap logger...")

	go bs.StartGin()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())

	bs.cfg.Service.WorkersCount = 10
	for i := 0; i < bs.cfg.Service.WorkersCount; i++ {
		go bs.worker(ctx)
	}

	for {
		s := <-c
		switch s {
		case syscall.SIGINT:
			bs.logger.Info("Get signal SIGUSR1")
			cancel()
		}
	}
}

func (bs *Service) worker(ctx context.Context) {
	bs.logger.Debug("worker", zap.String("message", "Start worker"))
	for {
		select {
		case <-ctx.Done():
			bs.logger.Debug("worker", zap.String("message", "Stop worker"))
			return
		}
	}
}
