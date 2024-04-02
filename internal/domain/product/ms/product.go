// Package ms provides the repositories implemented with the Meilisearch
package ms

import (
	"context"
	"strconv"

	"github.com/meilisearch/meilisearch-go"
	"go.uber.org/zap"

	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/errorx"
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

// Search products in search engine.
func (p *ProductEngine) Search(
	_ context.Context,
	criteria *product.SearchQueryCriteria,
) (*product.SearchQueryResult, error) {
	result, err := p.client.Index("").Search(criteria.Search, &meilisearch.SearchRequest{})
	if err != nil {
		p.logger.Error("search products error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return &product.SearchQueryResult{
		Count: int(result.EstimatedTotalHits),
	}, nil
}

// Index products.
func (p *ProductEngine) Index(_ context.Context, products product.Products) error {
	documents := make([]map[string]interface{}, 0, len(products))
	for _, p := range products {
		documents = append(documents, map[string]interface{}{
			"id":          strconv.FormatInt(int64(p.ID), 10),
			"name":        p.Props.Name,
			"description": p.Props.Description,
		})
	}
	if _, err := p.client.Index("").AddDocuments(documents); err != nil {
		p.logger.Error("index products error", zap.Error(err))
		return errorx.ErrInternal
	}
	return nil
}
