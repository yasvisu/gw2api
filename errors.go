package gw2api


//GW2ApiError for API errors to unmarshal into.
type GW2ApiError struct {
	Text string `json:"text"`
}

//GW2ApiError implements the error interface.
func (e GW2ApiError) Error() string {
	return e.Text
}