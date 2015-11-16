package gw2api

import (
	"net/url"
	"time"
)

// WinLoss includes the detailed information over all exit stati
type WinLoss struct {
	Wins       int `json:"wins"`
	Losses     int `json:"losses"`
	Desertions int `json:"desertions"`
	Byes       int `json:"byes"`
	Forfeits   int `json:"forfeits"`
}

// ProfessionStats WinLoss per profession
type ProfessionStats struct {
	Elementalist WinLoss `json:"elementalist"`
	Engineer     WinLoss `json:"engineer"`
	Guardian     WinLoss `json:"guardian"`
	Mesmer       WinLoss `json:"mesmer"`
	Necromancer  WinLoss `json:"necromancer"`
	Ranger       WinLoss `json:"ranger"`
	Revenant     WinLoss `json:"revenant"`
	Thief        WinLoss `json:"thief"`
	Warrior      WinLoss `json:"warrior"`
}

// LadderStats WinLoss per queue
type LadderStats struct {
	Ranked   WinLoss `json:"ranked"`
	UnRanked WinLoss `json:"unranked"`
}

// PvPStats Meta object for the stats
type PvPStats struct {
	PvPRank     int             `json:"pvp_rank"`
	Aggregate   WinLoss         `json:"aggregate"`
	Professions ProfessionStats `json:"professions"`
	Ladders     LadderStats     `json:"ladders"`
}

// PvPStats all current players stats.
// Requires authentication
func (gw2 *GW2Api) PvPStats() (stats PvPStats, err error) {
	ver := "v2"
	tag := "pvp/stats"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermPvP, nil, &stats)
	return
}

// PvPGameScore single games final score
type PvPGameScore struct {
	Red  int `json:"red"`
	Blue int `json:"blue"`
}

// PvPGameStats all information about a match
type PvPGameStats struct {
	ID         string       `json:"id"`
	MapID      int          `json:"map_id"`
	Started    time.Time    `json:"started"`
	Ended      time.Time    `json:"ended"`
	Result     string       `json:"result"`
	Team       string       `json:"team"`
	Profession string       `json:"profession"`
	Scores     PvPGameScore `json:"scores"`
}

// PvPGames lists the last 10 match ids
func (gw2 *GW2Api) PvPGames() (res []string, err error) {
	ver := "v2"
	tag := "pvp/games"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermPvP, nil, &res)
	return
}

// PvPGameIds fetches the details for the requested ids
func (gw2 *GW2Api) PvPGameIds(ids ...string) (games []PvPGameStats, err error) {
	ver := "v2"
	tag := "pvp/games"
	params := url.Values{}
	params.Add("ids", commaList(ids))
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermPvP, params, &games)
	return
}
