package common

// Search is the request body for searching of
// users for later use in a Lookup request.
type Search struct {
	Term string `json:"term"` // The search term for the query.
}
