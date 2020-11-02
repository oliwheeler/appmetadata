package store

import (
	"sync"

	"github.com/oliwheeler/appmetadata/models"
)

type Store interface {
	Get() []models.Metadata
	GetByCompany(company string) ([]models.Metadata, error)
	GetMetadata(title string) (models.Metadata, error)
	Add(models.Metadata) error
	Update(models.Metadata) error
}

func NewInMemStore() Store {
	return &InMemoryStore{
		mutex:        sync.RWMutex{},
		data:         map[string]models.Metadata{},
		companyIndex: map[string]map[string]bool{},
	}
}

type InMemoryStore struct {
	mutex        sync.RWMutex
	data         map[string]models.Metadata // key: title, value: application metadata
	companyIndex companyIndex
}

func (store *InMemoryStore) Get() []models.Metadata {
	results := []models.Metadata{}
	for _, app := range store.data {
		results = append(results, app)
	}
	return results
}

func (store *InMemoryStore) GetByCompany(company string) ([]models.Metadata, error) {
	if companies, exists := store.companyIndex[company]; !exists {
		return nil, &DoesNotExistError{"company", company}
	} else {
		results := []models.Metadata{}
		for title := range companies {
			results = append(results, store.data[title])
		}
		return results, nil
	}
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
		return models.Metadata{}, &DoesNotExistError{"title", title}
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
		addToCompanyIndex(store, payload.Company, payload.Title)
	}
	store.mutex.Unlock()
	return err
}

func (store *InMemoryStore) Update(payload models.Metadata) error {
	store.mutex.Lock()
	var err error
	if storedMetadata, exists := store.data[payload.Title]; !exists {
		err = &DoesNotExistError{"title", payload.Title}
	} else {
		if payload.Company != storedMetadata.Company {
			delete(store.companyIndex[storedMetadata.Company], payload.Title)
		}
		store.data[payload.Title] = payload
		addToCompanyIndex(store, payload.Company, payload.Title)
	}
	store.mutex.Unlock()
	return err
}

func addToCompanyIndex(store *InMemoryStore, company, title string) {
	if store.companyIndex[company] == nil {
		store.companyIndex[company] = map[string]bool{}
	}
	store.companyIndex[company][title] = true
}
