package test

// Result of a test run that contains the response code, data or error.
type Result struct {
	Code  int         // HTTP status code returned for the response.
	Data  interface{} // Data if provided by the response.
	Error error       // Error if provided.
}
