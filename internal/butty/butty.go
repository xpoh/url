package butty

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"url/internal/config"
	"url/internal/storage"
	"url/pkg/logger"
)

var buttyService *Service

type Service struct {
	Cfg     *config.Cfg
	Logger  *zap.Logger
	Storage storage.Storager
}

func NewButtyService() {
	bs := Service{}
	buttyService = &bs

	bs.Cfg = config.NewConfig()
	bs.Logger = logger.New(bs.Cfg.Service.LogLevel)

	switch bs.Cfg.Data.Database.Driver {
	case "inmemory":
		bs.Logger.Info("NewButtyService", zap.String("database", "inmemory"))
		bs.Storage, _ = storage.NewInMemoryStorage()
		bs.Logger.Info("Create inmemory storage", zap.Any("storage", bs.Storage))
	case "mysql":
		bs.Logger.Info("NewButtyService", zap.String("database", "mysql"))
		bs.Storage = storage.NewMysqlStorage()
		bs.Logger.Info("Create mysql storage", zap.Any("storage", bs.Storage))
	default:
		bs.Logger.Panic("NewButtyService", zap.String("database", "Unknown database driver"))
		panic("Unknown database driver")
	}
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

	//ctx, cancel := context.WithCancel(context.Background())
	//bs.Cfg.Service.WorkersCount = 10
	//for i := 0; i < bs.Cfg.Service.WorkersCount; i++ {
	//	go bs.worker(ctx)
	//}

	for {
		s := <-c
		switch s {
		case syscall.SIGINT:
			bs.Logger.Info("Get signal SIGUSR1")
			bs.Storage.Close()
			//cancel()
		}
	}
}

//func (bs *Service) worker(ctx context.Context) {
//	bs.Logger.Debug("worker", zap.String("message", "Start worker"))
//	for {
//		select {
//		case <-ctx.Done():
//			bs.Logger.Debug("worker", zap.String("message", "Stop worker"))
//			return
//		}
//	}
//}
