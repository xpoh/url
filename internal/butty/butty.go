package butty

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	"url/cmd/url"
	"url/internal/Config"
)

type ButtyService struct {
	cfg    *Config.Cfg
	logger *zap.Logger
}

type Link struct {
	shortUrl string `json:"shortUrl"`
	url      string `json:"url"`
}

func NewButtyService() *ButtyService {
	bs := ButtyService{}
	cfg := Config.NewConfig()
	bs.cfg = cfg
	return &bs
}

func (bs *ButtyService) InitLogger() {
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

func (bs *ButtyService) Run() {
	bs.InitLogger()
	defer bs.logger.Sync()
	bs.logger.Info("Using zap logger...")

	go func() {
		log.Printf("Server started")
		router := url.NewRouter()
		gin.SetMode(gin.ReleaseMode)
		log.Fatal(router.Run(":80"))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	ctx, cancel := context.WithCancel(context.Background())

	bs.cfg.Service.WorkersCount = 10
	for i := 0; i < bs.cfg.Service.WorkersCount; i++ {
		go bs.worker(ctx)
	}

	for {
		s := <-c
		switch s {
		case syscall.SIGUSR1:
			bs.logger.Info("Get signal SIGUSR1")
			cancel()
		}
	}
}

func (bs *ButtyService) worker(ctx context.Context) {
	bs.logger.Debug("worker", zap.String("message", "Start worker"))
	for {
		select {
		case <-ctx.Done():
			bs.logger.Debug("worker", zap.String("message", "Stop worker"))
			return
		}
	}
}
