package store

import "fmt"

type DoesNotExistError struct {
	title string
}

type AlreadyExistsError struct {
	title string
}

func (e *DoesNotExistError) Error() string {
	return fmt.Sprintf("No application with title: '%s' exists", e.title)
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("Application with title: '%s' already exists", e.title)
}
