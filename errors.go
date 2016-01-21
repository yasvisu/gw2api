package gw2api

//APIError for API errors to unmarshal into.
type APIError struct {
	err string `json:"error"`
}

//Error implements the error interface.
func (e APIError) Error() string {
	return e.err
}
