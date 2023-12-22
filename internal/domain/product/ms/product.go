// Package ms provides the repositories implemented with the Meilisearch
package ms

import (
	"github.com/meilisearch/meilisearch-go"
	"go.uber.org/zap"
)

// ProductEngine is a search engine.
type ProductEngine struct {
	logger *zap.Logger
	client *meilisearch.Client
}

// NewProductEngine creates a new search engine.
func NewProductEngine(logger *zap.Logger, client *meilisearch.Client) *ProductEngine {
	return &ProductEngine{logger: logger, client: client}
}
