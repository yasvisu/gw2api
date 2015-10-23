package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
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

func Version() (v int, err error) {
	ver := "v2"
	tag := "build"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return 0, err
	}

	type JsonVersion struct {
		ID int `json:"id"`
	}

	var res JsonVersion
	err = json.Unmarshal(data, &res)
	return res.ID, nil
}
