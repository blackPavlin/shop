// Package restx contains utils for rest handlers.
package restx

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// GetIDFromURLParams return id as an int64 from get params by its name.
func GetIDFromURLParams(r *http.Request, paramName string) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, paramName), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed parse id from url: %w", err)
	}
	return id, nil
}
