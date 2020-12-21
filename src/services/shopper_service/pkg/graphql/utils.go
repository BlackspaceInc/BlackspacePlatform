package graphql

import (
	"errors"
)

func handleErrorIfPresent(id int) (*int, error) {
	if id == 0 {
		return nil, errors.New("failed to cast uint32 to int value")
	}
	return &id, nil
}
