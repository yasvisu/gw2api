package gw2api

import (
	"os"
	"testing"
)

func TestPvPStats(t *testing.T) {
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
}

func TestPvPGames(t *testing.T) {
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
	var testGames []string
	if testGames, err = api.PvPGames(); err != nil {
		t.Error("Failed to parse the pvp game ids: ", err)
	} else if len(testGames) < 1 {
		t.Error("Fetched an unlikely number of pvp game ids")
	}
	var games []PvPGameStats
	if games, err = api.PvPGameIds(testGames[0], testGames[1]); err != nil || len(games) != 2 {
		t.Error("Failed to parse the pvp game data: ", err)
	}
}
