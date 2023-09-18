// Package searchx contains utils to work with search engines.
package searchx

import (
	"context"
	"fmt"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

// SearchConfig is search engines configuration.
type SearchConfig struct {
	Host    string        `envconfig:"HOST" required:"true"`
	APIKey  string        `envconfig:"API_KEY" required:"true"`
	Timeout time.Duration `envconfig:"TIMEOUT" required:"true"`
}

// NewClient create instance of search engines.
func NewClient(ctx context.Context, config *SearchConfig) (*meilisearch.Client, error) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:    config.Host,
		APIKey:  config.APIKey,
		Timeout: config.Timeout,
	})
	if _, err := client.Health(); err != nil {
		return nil, fmt.Errorf("check health milisearch error: %w", err)
	}
	return client, nil
}
