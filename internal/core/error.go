package core

import "fmt"

type KeyNotFoundError struct {
	key int
}

func NewKeyNotFoundError(key int) KeyNotFoundError {
	return KeyNotFoundError{key: key}
}

func (k KeyNotFoundError) Error() string {
	return fmt.Sprintf("item with key %d not found", k.key)
}
