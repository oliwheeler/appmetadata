package models

import (
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/mod/semver"
)

type Metadata struct {
	Title       string `yaml:"title"` // Primary key
	Version     string `yaml:"version"`
	Maintainers []struct {
		Name  string `yaml:"name"`
		Email string `yaml:"email"`
	} `yaml:"maintainers"`
	Company     string `yaml:"company"`
	Website     string `yaml:"website"`
	Source      string `yaml:"source"`
	License     string `yaml:"license"`
	Description string `yaml:"description"`
}

type ValidationError struct {
	Errors []string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid metadata:\n %s", strings.Join(e.Errors, "\n "))
}

func (metadata Metadata) Validate() error {
	invalids := []string{}
	if len(metadata.Title) == 0 {
		invalids = append(invalids, "Title is missing")
	}
	if len(metadata.Version) == 0 {
		invalids = append(invalids, "Version is missing")
	} else {
		semverTest := metadata.Version
		if metadata.Version[0] != 'v' {
			semverTest = "v" + semverTest
		}
		if !semver.IsValid(semverTest) {
			invalids = append(invalids, fmt.Sprintf("'%s' is not a valid semantic version", metadata.Version))
		}
	}
	if len(metadata.Maintainers) == 0 {
		invalids = append(invalids, "A Maintainer list is missing")
	}
	for _, maintainer := range metadata.Maintainers {
		if err := checkmail.ValidateFormat(maintainer.Email); err != nil {
			invalids = append(invalids, fmt.Sprintf("Maintainer email error: %w", err))
		}
		if len(maintainer.Name) == 0 {
			invalids = append(invalids, "Maintainer name is missing")
		}
	}
	if len(metadata.Company) == 0 {
		invalids = append(invalids, "Company is missing")
	}
	if len(metadata.Website) == 0 {
		invalids = append(invalids, "Website is missing")
	}
	if len(metadata.Source) == 0 {
		invalids = append(invalids, "Source is missing")
	}
	if len(metadata.License) == 0 {
		invalids = append(invalids, "License is missing")
	}
	if len(metadata.Description) == 0 {
		invalids = append(invalids, "Description is missing")
	}

	if len(invalids) > 0 {
		return &ValidationError{Errors: invalids}
	}
	return nil
}
