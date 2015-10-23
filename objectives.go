package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Map objectives such as towers, keeps, etc
type Objective struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	SectorId int       `json:"sector_id"`
	Type     string    `json:"type"`
	MapType  string    `json:"map_type"`
	MapId    int       `json:"map_id"`
	Coord    []float32 `json:"coord"`
	Marker   string    `json:"marker"`
}

// Returns a list of all objectives on wvw maps
func Objectives() (res []string, err error) {
	ver := "v2"
	tag := "wvw/objectives"

	var data []byte
	if data, err = fetchJSON(ver, tag, ""); err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	return
}

// Returns a list of objectives as request by ids. Use special id `all` to request
// all objectives
func ObjectiveIds(lang string, ids ...string) (objs []Objective, err error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil.")
	}
	var appendix bytes.Buffer
	ver := "v2"
	tag := "wvw/objectives"

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
		appendix.WriteString(id)
	}
	var data []byte
	if data, err = fetchJSON(ver, tag, appendix.String()); err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &objs)
	return
}
