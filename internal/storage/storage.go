package storage

import (
	"math"
	"strconv"
	"sync"
	"time"
)

// Storager interface for ...
type Storager interface {
	getFullUrlByButty(buttyUrl string) (string, error)
	newButtyUrl(url string) (string, error)
	clearOldLinks(days int) error
}

type Link struct {
	ShortUrl    string `json:"ShortUrl"`
	Url         string `json:"Url"`
	CreatedDate time.Time
	Id          int
}

// InMemoryStorage реализует хранилище ссылок в памяти
type InMemoryStorage struct {
	Storage     map[string]Link
	FreeLinksId map[int]bool
	mux         sync.RWMutex
}

func NewInMemoryStorage() (*InMemoryStorage, error) {
	s := &InMemoryStorage{}
	s.Storage = map[string]Link{}
	s.FreeLinksId = map[int]bool{}
	for j := 0; j < 10; j++ {
		s.FreeLinksId[j] = false
	}
	return s, nil
}

type NotFindError struct{}

func (n NotFindError) Error() string {
	return "url not find"
}

type OverFlowError struct{}

func (o OverFlowError) Error() string {
	return "Overflow memory"
}

// getFullUrlByButty достает из хранилища ссылку по красивой ссылке
func (i *InMemoryStorage) getFullUrlByButty(buttyUrl string) (string, error) {
	i.mux.RLock()
	defer i.mux.RUnlock()

	link, found := i.Storage[buttyUrl]
	if !found {
		return "", NotFindError{}
	}
	return link.Url, nil
}

//newButtyUrl генерит и добавляет в хранилище новую ссылку,
// считается, что у ссылки есть время жизни и в хранилище могут быть более красивые ссылки раньше
func (i *InMemoryStorage) newButtyUrl(url string) (string, error) {
	i.mux.Lock()
	defer i.mux.Unlock()
	for j := 0; j < 10; j++ {
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

func idToUrl(id int) string {
	return strconv.FormatUint(uint64(id), 36)
}

func (i *InMemoryStorage) clearOldLinks(days int) error {
	i.mux.Lock()
	defer i.mux.Unlock()

	for key, link := range i.Storage {
		if math.Abs(link.CreatedDate.Sub(time.Now()).Hours()/24) >= float64(days) {
			i.FreeLinksId[link.Id] = false
			delete(i.Storage, key)
		}
	}
	return nil
}
