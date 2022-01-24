package storage

import (
	"math"
	"sync"
	"time"
)

// InMemoryStorage реализует хранилище ссылок в памяти
type InMemoryStorage struct {
	Storage     map[string]Link
	FreeLinksId map[int64]bool
	mux         sync.RWMutex
}

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

type NotFindError struct{}

func (n NotFindError) Error() string {
	return "url not find"
}

type OverFlowError struct{}

func (o OverFlowError) Error() string {
	return "Overflow memory"
}

// getFullUrlByButty достает из хранилища ссылку по красивой ссылке
func (i *InMemoryStorage) GetFullUrlByButty(buttyUrl string) (string, error) {
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

func (i *InMemoryStorage) Close() {
}
