package search

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/search"
)

type Service interface {
	Search(ctx context.Context, query search.Entity[search.Value]) ([]SearchedEntity[search.Value], error)
}

func NewService(logger log.Logger, entities []search.Entity[search.Value]) Service {

	fmt.Printf("v2search NewService(%d entity types)\n", len(entities)) //nolint:forbidigo

	return &service{
		logger:   logger,
		entities: entities,
	}
}

type service struct {
	logger   log.Logger
	entities []search.Entity[search.Value]
}

func (s *service) Search(ctx context.Context, query search.Entity[search.Value]) ([]SearchedEntity[search.Value], error) {
	for _, entity := range s.entities {
		if len(entity.Addresses) > 0 {
			bs, _ := json.Marshal(entity)
			fmt.Printf("\n\n %s \n", string(bs)) //nolint:forbidigo
			return nil, nil
		}
	}

	// TODO(adam): use SearchedEntity
	// type SearchedEntity[T any] struct {
	// 	search.Entity[T]
	// 	Match float64 `json:"match"`

	var out []SearchedEntity[search.Value]

	return out, nil
}
