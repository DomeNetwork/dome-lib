package api

// Error represents a error response.
type Error struct {
	Code  int    `example:"000" json:"code"`                  // HTTP status code for error.
	Error string `example:"There was an error." json:"error"` // The error message string.
}

// Result response for service requests.
type Result struct {
	Code   int         `example:"200" json:"code"` // HTTP status code for response.
	Result interface{} `json:"result"`             // The response body if any.
}
