package service

import (
	"errors"
	"fmt"
)

var CannotUpdateTitle = errors.New("Cannot update the title, must create a new application registration")

type CannotNotGetMetadataError struct {
	Title string
	Err   error
}

func (e *CannotNotGetMetadataError) Error() string {
	return fmt.Sprintf("Could not get metadata with title: '%s'", e.Title)
}

func (e *CannotNotGetMetadataError) Unwrap() error {
	return e.Err
}

type CannotUpdateNonExistantMetadataError struct {
	Err error
}

func (e *CannotUpdateNonExistantMetadataError) Error() string {
	return fmt.Sprintf("Cannot update metadata: %s", e.Err)
}

func (e *CannotUpdateNonExistantMetadataError) Unwrap() error {
	return e.Err
}

type MetadataNameAlreadyExistsError struct {
	Name string
	Err  error
}

func (e *MetadataNameAlreadyExistsError) Error() string {
	return fmt.Sprintf("Metadata with name '%s' already exists", e.Name)
}

type InValidYamlError struct {
	Err error
}

func (e *InValidYamlError) Error() string {
	return fmt.Sprintf("Invalid yaml: %s", e.Err)
}

func (e *InValidYamlError) Unwrap() error {
	return e.Err
}

type ServiceError struct {
	Err error
}

func (e *ServiceError) Error() string {
	return "Server error"
}

func (e *ServiceError) Unwrap() error {
	return e.Err
}
