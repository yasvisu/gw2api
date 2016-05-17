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
	PvPRank          int             `json:"pvp_rank"`
	PvPRankPoints    int             `json:"pvp_rank_points"`
	PvPRankRollovers int             `json:"pvp_rank_rollovers"`
	Aggregate        WinLoss         `json:"aggregate"`
	Professions      ProfessionStats `json:"professions"`
	Ladders          LadderStats     `json:"ladders"`
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
	RatingType string       `json:"rating_type"`
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

// PvPStanding TotalPoints is the definitive amount
type PvPStanding struct {
	TotalPoints int `json:"total_points"`
	Division    int `json:"division"`
	Tier        int `json:"tier"`
	Points      int `json:"points"`
	Repeats     int `json:"repeats"`
}

// PvPSeasonStanding meta object for the call
type PvPSeasonStanding struct {
	Current  PvPStanding `json:"current"`
	Best     PvPStanding `json:"best"`
	SeasonID string      `json:"season_id"`
}

// PvPStandings fetches the current and previsous standings during the pvp seasons
func (gw2 *GW2Api) PvPStandings() (res []PvPSeasonStanding, err error) {
	ver := "v2"
	tag := "pvp/standings"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermPvP, nil, &res)
	return
}

// PvPSeasons returns a list of season ids
func (gw2 *GW2Api) PvPSeasons() (res []string, err error) {
	ver := "v2"
	tag := "pvp/seasons"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// PvPTier unnecessary really
type PvPTier struct {
	Points int `json:"points"`
}

// PvPDivision mostly for graphical interfaces
type PvPDivision struct {
	Name      string    `json:"name"`
	Flags     []string  `json:"flags"`
	LargeIcon string    `json:"large_icon"`
	SmallIcon string    `json:"small_icon"`
	PipIcon   string    `json:"pip_icon"`
	Tiers     []PvPTier `json:"tiers"`
}

// PvPSeason Meta object for pvp seasons
type PvPSeason struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Start     time.Time     `json:"start"`
	End       time.Time     `json:"end"`
	Active    bool          `json:"active"`
	Divisions []PvPDivision `json:"divisions"`
}

// PvPSeasonID fetches information on a season id
func (gw2 *GW2Api) PvPSeasonID(lang string, id string) (res PvPSeason, err error) {
	ver := "v2"
	tag := "pvp/seasons"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("id", id)
	err = gw2.fetchEndpoint(ver, tag, params, &res)
	return
}
