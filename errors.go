package gw2api

//Error for API errors to unmarshal into.
type Error struct {
	Text string `json:"text"`
}

//Error implements the error interface.
func (e Error) Error() string {
	return e.Text
}
