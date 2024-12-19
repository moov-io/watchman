package search

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/search"
)

type Service interface {
	Search(ctx context.Context)
}

func NewService[T any](logger log.Logger, entities []search.Entity[T]) Service {
	return &service[T]{
		logger:   logger,
		entities: entities,
	}
}

type service[T any] struct {
	logger   log.Logger
	entities []search.Entity[T]
}

func (s *service[T]) Search(ctx context.Context) {
	for _, entity := range s.entities {
		if len(entity.Addresses) > 0 {
			bs, _ := json.Marshal(entity)
			fmt.Printf("\n\n %s \n", string(bs)) //nolint:forbidigo
			return
		}
	}
}
