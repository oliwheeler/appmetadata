package store

import "fmt"

type DoesNotExistError struct {
	field string
	value string
}

type AlreadyExistsError struct {
	title string
}

func (e *DoesNotExistError) Error() string {
	return fmt.Sprintf("No application with '%s: %s' exists", e.field, e.value)
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("Application with title: '%s' already exists", e.title)
}
