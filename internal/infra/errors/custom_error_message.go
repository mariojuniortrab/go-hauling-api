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

func MustMatch(field1, field2 string) error {
	return fmt.Errorf("%s does not match %s", strings.ToLower(field1), strings.ToLower(field2))
}
