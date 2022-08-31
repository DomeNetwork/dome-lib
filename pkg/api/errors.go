package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func fail(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, &Error{
		Code:  code,
		Error: err.Error(),
	})
}

// ErrBadGateway - 502 Bad Gateway
func ErrBadGateway(c *gin.Context, err error) {
	fail(c, http.StatusBadGateway, err)
}

// ErrBadRequest - 400 Bad Request
func ErrBadRequest(c *gin.Context, err error) {
	fail(c, http.StatusBadRequest, err)
}

// ErrForbidden - 403 Forbidden
func ErrForbidden(c *gin.Context, err error) {
	fail(c, http.StatusForbidden, err)
}

// ErrInternalServerError - 500 Internal Server Error
func ErrInternalServerError(c *gin.Context, err error) {
	fail(c, http.StatusInternalServerError, err)
}

// ErrNotFound - 404 Not Found
func ErrNotFound(c *gin.Context, err error) {
	fail(c, http.StatusNotFound, err)
}

// ErrUnauthorized - 401 Unauthorized
func ErrUnauthorized(c *gin.Context, err error) {
	fail(c, http.StatusUnauthorized, err)
}
