package gw2api

import (
	"net/url"
	"strconv"
	"time"
)

// TeamAssoc Points/Kills/Deaths per team
type TeamAssoc struct {
	Green int `json:"green"`
	Blue  int `json:"blue"`
	Red   int `json:"red"`
}

// TeamMulti Used in multi association for teams
type TeamMulti struct {
	Green []int `json:"green"`
	Blue  []int `json:"blue"`
	Red   []int `json:"red"`
}

// Bonus - Map bonuses and current owner indetified by Color/Neutral
type Bonus struct {
	Owner string `json:"owner"`
	Type  string `json:"type"`
}

// MatchObjective - Map objectives such as towers, keeps, etc
type MatchObjective struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Owner       string `json:"owner"`
	LastFlipped string `json:"last_flipped"`
	ClaimedBy   string `json:"claimed_by"`
	ClaimedAt   string `json:"claimed_at"`
}

// MapWvW - One of the four maps and their status
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
	ID        string    `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Scores    TeamAssoc `json:"scores"`
	Worlds    TeamAssoc `json:"worlds"`
	AllWorlds TeamMulti `json:"all_worlds"`
	Deaths    TeamAssoc `json:"deaths"`
	Kills     TeamAssoc `json:"kills"`
	Maps      []MapWvW  `json:"maps"`
}

// Matches returns a list of all current match ids in the form of %d-%d
func (gw2 *GW2Api) Matches() (res []string, err error) {
	ver := "v2"
	tag := "wvw/matches"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// MatchIds returns matches as requested by ids in the form
// provided by Matches(). Use special id `all` for every match in US/EU
func (gw2 *GW2Api) MatchIds(ids ...string) (match []Match, err error) {
	ver := "v2"
	tag := "wvw/matches"
	params := url.Values{}
	params.Add("ids", commaList(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &match)
	return
}

// MatchWorld finds the match the server id is participating in
func (gw2 *GW2Api) MatchWorld(worldID int) (match Match, err error) {
	ver := "v2"
	tag := "wvw/matches"
	params := url.Values{}
	params.Add("world", strconv.Itoa(worldID))
	err = gw2.fetchEndpoint(ver, tag, params, &match)
	return
}

// Objective - Map objectives such as towers, keeps, etc
type Objective struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	SectorID int       `json:"sector_id"`
	Type     string    `json:"type"`
	MapType  string    `json:"map_type"`
	MapID    int       `json:"map_id"`
	Coord    []float32 `json:"coord"`
	Marker   string    `json:"marker"`
}

// Objectives returns a list of all objectives on wvw maps
func (gw2 *GW2Api) Objectives() (res []string, err error) {
	ver := "v2"
	tag := "wvw/objectives"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// ObjectiveIds returns a list of objectives as request by ids. Use special id
// `all` to request all objectives
func (gw2 *GW2Api) ObjectiveIds(lang string, ids ...string) (objs []Objective, err error) {
	ver := "v2"
	tag := "wvw/objectives"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &objs)
	return
}
