package gw2api

import (
"bytes"
"encoding/json"
"errors"
"strconv"
)

//Parent type for all -Names endpoints.
type Name struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Returns list of world ids.
func Worlds(lang string) (res []int, err error) {
	var appendix bytes.Buffer
	ver := "v2"
	tag := "worlds"

	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	return
}

//Returns list of world names.
func WorldsIds(lang string, ids ...int) ([]Name, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Worlds() instead?")
	}
	
	var appendix bytes.Buffer
	ver := "v2"
	tag := "worlds"
	concatenator := "?"

	if lang != "" {
		appendix.WriteString(concatenator)
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
		concatenator = "&"
	}
	appendix.WriteString(concatenator)
	appendix.WriteString("ids=")
	concatenator = ""
	for _, i := range ids {
		appendix.WriteString(concatenator)
		appendix.WriteString(strconv.Itoa(i))
		if concatenator == "" {
			concatenator = ","
		}
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []Name
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
