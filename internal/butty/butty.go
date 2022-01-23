package butty

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"url/internal/config"
	"url/internal/storage"
)

var buttyService *Service

type Service struct {
	Cfg     *config.Cfg
	Logger  *zap.Logger
	Storage storage.Storager
}

func NewButtyService() {
	bs := Service{}
	cfg := config.NewConfig()

	bs.InitLogger()
	defer func() {
		err := bs.Logger.Sync()
		if err != nil {
		}
	}()
	bs.Logger.Debug("Init zap Logger.")

	bs.Cfg = cfg
	bs.Storage, _ = storage.NewInMemoryStorage()

	buttyService = &bs
}

func GetService() *Service {
	if buttyService == nil {
		panic("butty service is nil")
	}
	return buttyService
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

	bs.Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
	bs.Logger.Info("Logger construction succeeded")
}

func (bs *Service) Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())

	bs.Cfg.Service.WorkersCount = 10
	for i := 0; i < bs.Cfg.Service.WorkersCount; i++ {
		go bs.worker(ctx)
	}

	for {
		s := <-c
		switch s {
		case syscall.SIGINT:
			bs.Logger.Info("Get signal SIGUSR1")
			cancel()
		}
	}
}

func (bs *Service) worker(ctx context.Context) {
	bs.Logger.Debug("worker", zap.String("message", "Start worker"))
	for {
		select {
		case <-ctx.Done():
			bs.Logger.Debug("worker", zap.String("message", "Stop worker"))
			return
		}
	}
}
