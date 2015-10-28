package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Points/Kills/Deaths per team
type TeamAssoc struct {
	Green int `json:"green"`
	Blue  int `json:"blue"`
	Red   int `json:"red"`
}

// Map bonuses and current owner indetified by Color/Neutral
type Bonus struct {
	Owner string `json:"owner"`
	Type  string `json:"type"`
}

// Map objectives such as towers, keeps, etc
type MatchObjective struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Owner       string `json:"owner"`
	LastFlipped string `json:"last_flipped"`
	ClaimedBy   string `json:"claimed_by"`
	ClaimedAt   string `json:"claimed_at"`
}

// One of the four maps and their status
type MapWvW struct {
	ID         int              `json:"id"`
	Type       string           `json:"type"`
	Scores     TeamAssoc        `json:"scores"`
	Bonuses    []Bonus          `json:"bonuses"`
	Deaths     TeamAssoc        `json:"deaths"`
	Kills      TeamAssoc        `json:"kills"`
	Objectives []MatchObjective `json:"objectives"`
}

// Match including overall stats and indivdual maps with stats
type Match struct {
	Id        string    `json:"id"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	Scores    TeamAssoc `json:"scores"`
	Worlds    TeamAssoc `json:"worlds"`
	Deaths    TeamAssoc `json:"deaths"`
	Kills     TeamAssoc `json:"kills"`
	Maps      []MapWvW  `json:"maps"`
}

// Returns a list of all current match ids in the form of %d-%d
func Matches() (res []string, err error) {
	ver := "v2"
	tag := "wvw/matches"

	err = fetchEndpoint(ver, tag, "", &res)
	return
}

// Returns matches as per requested by ids in the form
// provided by Matches(). Use special id `all` for every match in US/EU
func MatchIds(ids ...string) (match []Match, err error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil.")
	}
	var appendix bytes.Buffer
	ver := "v2"
	tag := "wvw/matches"

	appendix.WriteString("?ids=")
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
	err = json.Unmarshal(data, &match)
	return
}

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
	err = fetchEndpoint(ver, tag, "", &res)
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
