package infra_errors

import (
	"errors"
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

func IsInvalid(field1 string) error {
	return fmt.Errorf("%s is invalid", strings.ToLower(field1))
}

func UserNotFound() error {
	return WrongPassword()
}

func WrongPassword() error {
	return errors.New("user not found / wrong password")
}
