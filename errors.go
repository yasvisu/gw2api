package gw2api

//APIError for API errors to unmarshal into.
type APIError struct {
	Err  string `json:"error"`
	Text string `json:"text`
}

//Error implements the error interface.
func (e APIError) Error() string {
	if e.Err == "" {
		return e.Text
	}

	return e.Err
}
