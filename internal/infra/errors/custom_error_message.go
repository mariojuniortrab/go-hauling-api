package infra_errors

import (
	"fmt"
	"strings"
)

func IsRequired(field string) error {
	return fmt.Errorf("%s is required", strings.ToLower(field))
}

func AlreadyExists(entity string) error {
	return fmt.Errorf("%s already exists", strings.ToLower(entity))
}

func MustBeUUID(field string) error {
	return fmt.Errorf("%s must be uuid", strings.ToLower(field))
}
