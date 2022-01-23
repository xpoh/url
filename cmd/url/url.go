/*
 * url butty maker
 *
 * url butty maker
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"go.uber.org/zap"
	"url/internal/api"
	"url/internal/butty"
)

func main() {
	butty.NewButtyService()
	bs := butty.GetService()
	bs.Logger.Debug("service", zap.String("message", "Create new butty service"))

	go bs.Run()
	bs.Logger.Debug("service", zap.String("message", "Butty service run"))

	router := api.NewRouter()
	bs.Logger.Debug("gin", zap.String("message", "gin created"))

	bs.Logger.Fatal("Error start gin http api", zap.Error(router.Run(bs.Cfg.Server.Http.Addr)))
}