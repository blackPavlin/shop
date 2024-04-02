package product

import (
	"context"

	"github.com/blackPavlin/shop/pkg/paging"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "search_mock.go" -package "product"

// SearchEngine represents product search engine.
type SearchEngine interface {
	Search(ctx context.Context, criteria *SearchQueryCriteria) (*SearchQueryResult, error)
	Index(ctx context.Context, products Products) error
}

// SearchQueryCriteria represents criteria for products search.
type SearchQueryCriteria struct {
	Search     string
	Pagination paging.Pagination
}

// SearchQueryResult represents a result for products search.
type SearchQueryResult struct {
	Count int
}
