package entity

import "errors"

type Brand struct {
	ID   string
	name string
}

func (b *Brand) Validate() error {
	if b.name == "" {
		return errors.New("name is required")
	}

	return nil
}
