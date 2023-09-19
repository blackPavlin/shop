// Package restx contains utils for rest handlers.
package restx

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/blackPavlin/shop/pkg/paging"
)

const (
	defaultPaginationLimit  = 20
	defaultPaginationOffset = 0
)

// GetIDFromURLParams return id as an int64 from get params by its name.
func GetIDFromURLParams(r *http.Request, paramName string) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, paramName), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed parse id from url: %w", err)
	}
	return id, nil
}

// GetPaginationParams parses request pagination parameters.
func GetPaginationParams(r *http.Request) (*paging.Pagination, error) {
	var (
		pagination = &paging.Pagination{
			Limit:  defaultPaginationLimit,
			Offset: defaultPaginationOffset,
		}
		err error
	)
	if limit := r.URL.Query().Get("limit"); limit != "" {
		if pagination.Limit, err = strconv.ParseUint(limit, 10, 64); err != nil {
			return nil, fmt.Errorf("failed parse limit from query params: %w", err)
		}
	}
	if offset := r.URL.Query().Get("offset"); offset != "" {
		if pagination.Offset, err = strconv.ParseUint(offset, 10, 64); err != nil {
			return nil, fmt.Errorf("failed parse offset from query params: %w", err)
		}
	}
	return pagination, nil
}
