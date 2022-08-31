package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// A convenience method for extracting the bearer token form the
// Authorization header.
func extractBearerFromHeader(c *gin.Context) (bearer string) {
	auth := c.Request.Header.Get("Authorization")
	bearer = strings.Replace(auth, "Bearer ", "", 1)
	return
}

// OK will respond to the request with a 200 status code.
func OK(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, &Result{Code: http.StatusOK, Result: v})
}
