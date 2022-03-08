package butty

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"testing"
	"url/internal/config"
	"url/internal/storage"
)

func TestGetService(t *testing.T) {
	tests := []struct {
		name string
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewButtyService(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewButtyService()
		})
	}
}

func TestService_Run(t *testing.T) {
	type fields struct {
		Cfg            *config.Cfg
		Logger         *zap.Logger
		Storage        storage.Storager
		UrlPostCounter prometheus.Counter
		UrlGetCounter  prometheus.Counter
		Router         *gin.Engine
		ServerButty    http.Server
		ServerProm     http.Server
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &Service{
				Cfg:            tt.fields.Cfg,
				Logger:         tt.fields.Logger,
				Storage:        tt.fields.Storage,
				UrlPostCounter: tt.fields.UrlPostCounter,
				UrlGetCounter:  tt.fields.UrlGetCounter,
				Router:         tt.fields.Router,
				ServerButty:    tt.fields.ServerButty,
				ServerProm:     tt.fields.ServerProm,
			}
			bs.Run()
		})
	}
}

func TestService_RunButty(t *testing.T) {
	type fields struct {
		Cfg            *config.Cfg
		Logger         *zap.Logger
		Storage        storage.Storager
		UrlPostCounter prometheus.Counter
		UrlGetCounter  prometheus.Counter
		Router         *gin.Engine
		ServerButty    http.Server
		ServerProm     http.Server
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &Service{
				Cfg:            tt.fields.Cfg,
				Logger:         tt.fields.Logger,
				Storage:        tt.fields.Storage,
				UrlPostCounter: tt.fields.UrlPostCounter,
				UrlGetCounter:  tt.fields.UrlGetCounter,
				Router:         tt.fields.Router,
				ServerButty:    tt.fields.ServerButty,
				ServerProm:     tt.fields.ServerProm,
			}
			bs.RunButty()
		})
	}
}

func TestService_RunPrometheus(t *testing.T) {
	type fields struct {
		Cfg            *config.Cfg
		Logger         *zap.Logger
		Storage        storage.Storager
		UrlPostCounter prometheus.Counter
		UrlGetCounter  prometheus.Counter
		Router         *gin.Engine
		ServerButty    http.Server
		ServerProm     http.Server
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &Service{
				Cfg:            tt.fields.Cfg,
				Logger:         tt.fields.Logger,
				Storage:        tt.fields.Storage,
				UrlPostCounter: tt.fields.UrlPostCounter,
				UrlGetCounter:  tt.fields.UrlGetCounter,
				Router:         tt.fields.Router,
				ServerButty:    tt.fields.ServerButty,
				ServerProm:     tt.fields.ServerProm,
			}
			bs.RunPrometheus()
		})
	}
}
