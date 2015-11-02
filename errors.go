package gw2api

//GW2ApiError for API errors to unmarshal into.
type Error struct {
	Text string `json:"text"`
}

//GW2ApiError implements the error interface.
func (e Error) Error() string {
	return e.Text
}
