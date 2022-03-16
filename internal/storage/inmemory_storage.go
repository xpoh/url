// Copyright 2022. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package storage implement interface and functions
// for storage
package storage

import (
	"math"
	"sync"
	"time"
)

// InMemoryStorage implement in memory storage
type InMemoryStorage struct {
	Storage     map[string]Link
	FreeLinksId map[int64]bool
	mux         sync.RWMutex
}

// NewInMemoryStorage is a constructor for in memory storage
func NewInMemoryStorage() (*InMemoryStorage, error) {
	var j int64
	s := &InMemoryStorage{}
	s.Storage = map[string]Link{}
	s.FreeLinksId = map[int64]bool{}
	for j = 0; j < 10; j++ {
		s.FreeLinksId[j] = false
	}
	return s, nil
}

// NotFindError implement error for not found url in butty storage
type NotFindError struct{}

func (n NotFindError) Error() string {
	return "url not find"
}

type OverFlowError struct{}

func (o OverFlowError) Error() string {
	return "Overflow memory"
}

// GetFullUrlByButty return butty url from storage
func (i *InMemoryStorage) GetFullUrlByButty(buttyUrl string) (string, error) {
	i.mux.RLock()
	defer i.mux.RUnlock()

	link, found := i.Storage[buttyUrl]
	if !found {
		return "", NotFindError{}
	}
	return link.Url, nil
}

//NewButtyUrl add new url in storage
// in: full url for add
// out: butty url and error if needed
// function check if in url is in storage
func (i *InMemoryStorage) NewButtyUrl(url string) (string, error) {
	var j int64
	i.mux.Lock()
	defer i.mux.Unlock()
	for j = 0; j < 10; j++ {
		if !i.FreeLinksId[j] {
			i.FreeLinksId[j] = true
			i.Storage[idToUrl(j)] = Link{
				ShortUrl:    idToUrl(j),
				Url:         url,
				CreatedDate: time.Now(),
				Id:          j,
			}
			return idToUrl(j), nil
		}
	}
	return "", OverFlowError{}
}

// ClearOldLinks clean links with overflow time lives from config
func (i *InMemoryStorage) ClearOldLinks(days int) error {
	i.mux.Lock()
	defer i.mux.Unlock()

	for key, link := range i.Storage {
		if math.Abs(time.Until(link.CreatedDate).Hours()/24) >= float64(days) {
			i.FreeLinksId[link.Id] = false
			delete(i.Storage, key)
		}
	}
	return nil
}

// Close perfom actions for graceful shutdown inmemory storage
func (i *InMemoryStorage) Close() {
}
