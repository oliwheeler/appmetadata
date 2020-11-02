package store

import (
	"sync"

	"github.com/oliwheeler/appmetadata/models"
)

type Store interface {
	GetMetadata(title string) (models.Metadata, error)
	Add(models.Metadata) error
	Update(models.Metadata) error
}

func NewInMemStore() Store {
	return &InMemoryStore{
		mutex: sync.RWMutex{},
		data:  map[string]models.Metadata{},
	}
}

type InMemoryStore struct {
	mutex sync.RWMutex
	data  map[string]models.Metadata // key: title, value: application metadata
}

func (store *InMemoryStore) GetMetadata(title string) (models.Metadata, error) {
	store.mutex.RLock()
	var result *models.Metadata
	if payload, exists := store.data[title]; exists {
		result = &payload
	}
	store.mutex.RUnlock()
	var err error
	if result == nil {
		return models.Metadata{}, &DoesNotExistError{title}
	}
	return *result, err
}

func (store *InMemoryStore) Add(payload models.Metadata) error {
	store.mutex.Lock()
	var err error
	if _, exists := store.data[payload.Title]; exists {
		err = &AlreadyExistsError{payload.Title}
	} else {
		store.data[payload.Title] = payload
	}
	store.mutex.Unlock()
	return err
}

func (store *InMemoryStore) Update(payload models.Metadata) error {
	store.mutex.Lock()
	var err error
	if _, exists := store.data[payload.Title]; !exists {
		err = &DoesNotExistError{payload.Title}
	} else {
		store.data[payload.Title] = payload
	}
	store.mutex.Unlock()
	return err
}
