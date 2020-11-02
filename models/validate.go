package models

type Validation interface {
	validate() error // nil error value means valid
}
