package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/domenetwork/dome-lib/pkg/common"
	"github.com/gin-gonic/gin"
)

// Run a test by providing the method (DELETE,GET,POST,PUT), route (ex. "/health"), service
// handler function, and the body interface to send in the request.  A Result is returned
// from the run that will provide the results of the response.
func Run(method, route string, fn func(*gin.Context), body interface{}) (result Result) {
	// Instantiate the result object for testing.
	result = Result{}

	// Setup the body as a reader.
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		result.Error = err
		return
	}

	// Setup a test server to be used locally.
	r := gin.Default()

	// Setup a default user ID of TEST.
	r.Use(func(c *gin.Context) {
		c.Set("userID", common.ID("TEST"))
		c.Next()
	})

	// Setup our route with the correct method.
	switch strings.ToUpper(method) {
	case "DELETE":
		r.DELETE(route, fn)
	case "GET":
		r.GET(route, fn)
	case "POST":
		r.POST(route, fn)
	case "PUT":
		r.PUT(route, fn)
	}

	// Create a new request object.
	req, err := http.NewRequest(strings.ToUpper(method), route, buf)
	if err != nil {
		result.Error = err
		return
	}

	// Setup a recorder to capture the response.
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	// Attach our response code to the test result.
	result.Code = res.Code

	// Attempt to convert the response with JSON.
	data := make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		result.Error = err
		return
	}

	result.Data = data["result"]
	return
}
