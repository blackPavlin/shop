package rest

import (
	"errors"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/errs"
	"github.com/gin-gonic/gin"
)

// HandleError
func HandleError(ctx *gin.Context, err error) {
	var e *errs.Error

	if errors.As(err, &e) {
		ctx.AbortWithStatusJSON(e.Code(), gin.H{"error": e.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
