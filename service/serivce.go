package service

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/oliwheeler/appmetadata/models"
	"github.com/oliwheeler/appmetadata/service/store"
	"gopkg.in/yaml.v2"
)

type Service struct {
	store store.Store
}

type Filter struct {
	Company string
}

func New() *Service {
	store := store.NewInMemStore()
	return &Service{
		store,
	}
}

func (svc *Service) GetApps(filter Filter) (io.Reader, error) {
	var apps []models.Metadata
	if len(filter.Company) > 0 {
		if companyApps, err := svc.store.GetByCompany(filter.Company); err != nil {
			return nil, &CannotNotGetMetadataError{"company", filter.Company, err}
		} else {
			apps = companyApps
		}
	} else {
		apps = svc.store.Get()
	}
	data, err := yaml.Marshal(apps)
	if err != nil {
		return nil, &ServiceError{err}
	}
	return bytes.NewReader(data), nil
}

func (svc *Service) GetAppMetadata(title string) (io.Reader, error) {
	metadata, err := svc.store.GetMetadata(title)
	if err != nil {
		return nil, &CannotNotGetMetadataError{"title", title, err}
	}

	data, err := yaml.Marshal(&metadata)
	if err != nil {
		return nil, &ServiceError{err}
	}
	return bytes.NewReader(data), nil
}

func (svc *Service) CreateAppMetadata(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return &ServiceError{err}
	}
	metadata := &models.Metadata{}
	if err = yaml.UnmarshalStrict(data, metadata); err != nil {
		return &InValidYamlError{err}
	}
	if err = metadata.Validate(); err != nil {
		return &InValidYamlError{err}
	}
	err = svc.store.Add(*metadata)
	return err
}

func (svc *Service) UpdateAppMetadata(title string, r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	metadata := &models.Metadata{}
	if err = yaml.UnmarshalStrict(data, metadata); err != nil {
		return &InValidYamlError{err}
	}
	if title != metadata.Title {
		return CannotUpdateTitle
	}
	if err = metadata.Validate(); err != nil {
		return &InValidYamlError{err}
	}
	if err = svc.store.Update(*metadata); err != nil {
		return &CannotUpdateNonExistantMetadataError{err}
	}
	return nil
}
