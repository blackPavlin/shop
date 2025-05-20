package search

import (
	"fmt"

	"github.com/meilisearch/meilisearch-go"

	"github.com/blackPavlin/shop/internal/config"
)

func NewSearchEngine(config *config.SearchEngineConfig) (meilisearch.ServiceManager, error) {
	serviceManager := meilisearch.New(config.Host, meilisearch.WithAPIKey(config.APIKey))

	if _, err := serviceManager.Health(); err != nil {
		return nil, fmt.Errorf("health check meilisearch: %w", err)
	}

	return serviceManager, nil
}
