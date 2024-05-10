package infra_errors

import (
	"errors"
	"fmt"
)

func IsRequired(field string) error {
	return fmt.Errorf("%s is required", field)
}

func AlreadyExists(entity string) error {
	return fmt.Errorf("%s already exists", entity)
}

func MustBeUUID(field string) error {
	return fmt.Errorf("%s must be uuid", field)
}

func MustMatch(field1, field2 string) error {
	return fmt.Errorf("%s does not match %s", field1, field2)
}

func IsInvalid(field1 string) error {
	return fmt.Errorf("%s is invalid", field1)
}

func UserNotFound() error {
	return WrongPassword()
}

func WrongPassword() error {
	return errors.New("user not found / wrong password")
}

func IsPasswordConfirmationInvalid() error {
	return fmt.Errorf("password and passwordConfirmation do not match")
}

func MustBeString(field string) error {
	return fmt.Errorf("%s must be a string", field)
}

func LengthMustBe(field string, length int) error {
	return fmt.Errorf("length of field %s must be %d", field, length)
}

func LengthMustBeOrLess(field string, length int) error {
	return fmt.Errorf("length of field %s must be %d or less", field, length)
}

func MustBeDateFormat(field string) error {
	return fmt.Errorf("%s must be YYYY-MM-DD format", field)
}
