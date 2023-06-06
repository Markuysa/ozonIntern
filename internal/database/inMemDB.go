package database

import (
	"context"
	"github.com/pkg/errors"
	internalErrors "ozonIntern/internal/errors"
	"sync"
)

type InMemDatabase struct {
	storage map[string]string
	mutex   *sync.Mutex
}

func NewInMemDatabase() *InMemDatabase {
	return &InMemDatabase{
		storage: make(map[string]string),
		mutex:   &sync.Mutex{},
	}
}

func (db *InMemDatabase) SaveLink(ctx context.Context, url, link string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, ok := db.storage[link]; ok {
		return internalErrors.ErrAlreadyExists
	}
	db.storage[link] = url
	return nil
}
func (db *InMemDatabase) GetURL(ctx context.Context, link string) (string, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	url, ok := db.storage[link]
	if ok {
		return url, nil
	}
	return "", errors.New("no url found")
}
