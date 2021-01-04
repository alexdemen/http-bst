package service

import (
	"github.com/alexdemen/http-bst/internal/core/log"
	"github.com/alexdemen/http-bst/internal/model"
)

type TreeService struct {
	tree   *model.Tree
	logger *log.Logger
}

func NewTreeService(logger *log.Logger) *TreeService {
	return &TreeService{
		tree:   model.NewIntTree(),
		logger: logger,
	}
}

func (s *TreeService) Add(key int) error {
	s.tree.Add(key)

	s.logger.Info(map[string]interface{}{
		"key":       key,
		"operation": "add",
	})

	return nil
}

func (s *TreeService) Find(key int) (int, error) {
	node, err := s.tree.Find(key)

	logData := map[string]interface{}{
		"key":       key,
		"operation": "find",
	}

	if err != nil {
		s.logger.Error(logData)
		return 0, err
	}

	s.logger.Info(logData)
	return node.Key, nil
}

func (s *TreeService) Delete(key int) error {
	s.tree.Delete(key)

	s.logger.Info(map[string]interface{}{
		"key":       key,
		"operation": "delete",
	})

	return nil
}
