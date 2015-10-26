package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

type Map struct {
	ID            int
	Name          string
	MinLevel      int
	MaxLevel      int
	DefaultFloor  int
	Floors        []int
	RegionId      int
	RegionName    string
	ContinentId   int
	ContinentName string
	MapRect       [2][2]int
	ContinentRect [2][2]int
}

// Returns a list of all current skin ids
func Maps() (res []int, err error) {
	ver := "v2"
	tag := "skins"

	var data []byte
	if data, err = fetchJSON(ver, tag, ""); err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	return
}

// Returns a list of skins as requested by the id parameter.
// Special id `all` is not permitted on this endpoint
func MapIds(lang string, ids ...int) (maps []Map, err error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil.")
	}
	var appendix bytes.Buffer
	ver := "v2"
	tag := "maps"

	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
		appendix.WriteString("&ids=")
	} else {
		appendix.WriteString("?ids=")
	}

	for i, id := range ids {
		if i > 0 {
			appendix.WriteString(",")
		}
		appendix.WriteString(strconv.Itoa(id))
	}

	var data []byte
	if data, err = fetchJSON(ver, tag, appendix.String()); err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &maps)
	return

}
