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
	"url/internal/api"
	"url/internal/butty"
)

func main() {
	butty.NewButtyService()
	bs := butty.GetService()
	bs.Router = api.NewRouter()

	bs.Run()
}
