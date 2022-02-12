package storage

import (
	"strconv"
	"time"
)

// Storager interface for ...
type Storager interface {
	GetFullUrlByButty(buttyUrl string) (string, error)
	NewButtyUrl(url string) (string, error)
	ClearOldLinks(days int) error
	Close()
}

type Link struct {
	ShortUrl    string `json:"ShortUrl"`
	Url         string `json:"Url"`
	CreatedDate time.Time
	Id          int64
}

func idToUrl(id int64) string {
	return strconv.FormatUint(uint64(id), 36)
}
