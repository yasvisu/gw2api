package gw2api

import (
	"os"
	"testing"
	"time"
)

func TestPvPAuthenticated(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermPvP) {
		t.Skip("API-Key does not have required permission for the test")
	}

	var stats PvPStats
	if stats, err = api.PvPStats(); err != nil {
		t.Error("Failed to parse the pvp stats: ", err)
	} else if stats.PvPRank < 0 {
		t.Error("Failed with wrong pvp stat data: ", err)
	}

	var testGames []string
	if testGames, err = api.PvPGames(); err != nil {
		t.Error("Failed to parse the pvp game ids: ", err)
	} else if len(testGames) < 1 {
		t.Error("Fetched an unlikely number of pvp game ids")
	}
	var games []PvPGameStats
	if games, err = api.PvPGameIds(testGames[0:2]...); err != nil || len(games) != 2 {
		t.Error("Failed to parse the pvp game data: ", err)
	}

	if time.Since(games[0].Started) < 0 {
		t.Error("Impossible time difference since match start")
	}

	var standings []PvPSeasonStanding
	if standings, err = api.PvPStandings(); err != nil {
		t.Error("Failed to parse the pvp standings: ", err)
	} else if len(standings) < 1 {
		t.Error("Fetched an unlikely number of pvp standings")
	}
	if standings[0].Best.TotalPoints < 1 {
		t.Error("Fetched an unset pvp standings object")
	}
}

func TestPvP(t *testing.T) {
	var err error
	api := NewGW2Api()

	var seasons []string
	if seasons, err = api.PvPSeasons(); err != nil {
		t.Error("Failed to fetch season ids")
	}

	var season PvPSeason
	if season, err = api.PvPSeasonID("", seasons[0]); err != nil {
		t.Error("Failed to fetch season with id")
	}

	if len(season.Divisions) < 1 {
		t.Error("Fetched a broken pvp season object")
	}
}
