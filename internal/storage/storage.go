// Copyright 2022. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package storage implement interface and functions
// for storage
package storage

import (
	"strconv"
	"time"
)

// Storager interface for storgage butty service
type Storager interface {
	// GetFullUrlByButty return full url by incoming butty url
	GetFullUrlByButty(buttyUrl string) (string, error)

	//NewButtyUrl add new url in storage
	// in: full url for add
	// out: butty url and error if needed
	// function check if in url is in storage
	NewButtyUrl(url string) (string, error)

	// ClearOldLinks clean links with overflow time lives from config
	ClearOldLinks(days int) error

	// Close perfom actions for graceful shutdown inmemory storage
	Close()
}

// Link is a data for store in storage
type Link struct {
	ShortUrl    string `json:"ShortUrl"`
	Url         string `json:"Url"`
	CreatedDate time.Time
	Id          int64
}

// idToUrl convert unique id number in short symbolic name
// using num 0..9 and letters a-z
func idToUrl(id int64) string {
	return strconv.FormatUint(uint64(id), 36)
}
