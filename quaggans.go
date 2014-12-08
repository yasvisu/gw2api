package gw2api

import (
"bytes"
"encoding/json"
"errors"
)

//Parent type for all Quaggan Resources.
type QuagganResource struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

//Returns list of quaggan ids.
func Quaggans() (res []string, err error) {
	ver := "v2"
	tag := "quaggans"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	return
}

//Returns list of quaggan resources.
func QuaggansIds(ids ...string) ([]QuagganResource, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Quaggans() instead?")
	}
	
	var appendix bytes.Buffer
	ver := "v2"
	tag := "quaggans"
	concatenator := "?"

	appendix.WriteString(concatenator)
	appendix.WriteString("ids=")
	concatenator = ""
	for _, i := range ids {
		appendix.WriteString(concatenator)
		appendix.WriteString(i)
		if concatenator == "" {
			concatenator = ","
		}
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []QuagganResource
	err = json.Unmarshal(data, &res)
	if err != nil {
		var gwerr GW2ApiError
		err = json.Unmarshal(data, &gwerr)
		if err != nil {
			return nil, err
		}
		return nil, gwerr
	}
	return res, err
}
