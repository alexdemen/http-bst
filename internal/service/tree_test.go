package service

import (
	"errors"
	"github.com/alexdemen/http-bst/internal/core"
	"testing"
)

func TestTreeService_Add(t *testing.T) {
	service := NewTreeService(nil)

	if err := service.Add(100); err != nil {
		t.Error("failed to add key")
	}
}

func TestTreeService_Find(t *testing.T) {
	service := NewTreeService(nil)
	if err := service.Add(100); err != nil {
		t.Error("failed to add key")
	}

	if _, err := service.Find(100); err != nil {
		t.Error("failed find key 100")
	}
}

func TestTreeService_Delete(t *testing.T) {
	service := NewTreeService(nil)
	if err := service.Add(100); err != nil {
		t.Error("failed to add key")
	}

	if _, err := service.Find(100); err != nil {
		t.Error("failed find key 100")
	}

	if err := service.Delete(100); err != nil {
		t.Error("failed delete key 100")
	}

	if _, err := service.Find(100); err != nil {
		if !errors.Is(err, core.NewKeyNotFoundError(100)) {
			t.Error("failed find key 100")
		}
	} else {
		t.Error("find deleted key 100")
	}
}
