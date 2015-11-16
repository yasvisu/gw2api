package gw2api

import (
	"testing"
	"time"
)

func TestMatches(t *testing.T) {
	var err error
	api := NewGW2Api()
	var testMatches []string
	if testMatches, err = api.Matches(); err != nil {
		t.Error("Failed to fetch matches: ", err)
	} else if len(testMatches) < 1 {
		t.Error("Fetched an unlikely number of matches")
	}
	var matches []Match
	if matches, err = api.MatchIds(testMatches[0:2]...); err != nil || len(matches) != 2 {
		t.Error("Failed to parse the match data: ", err)
	}
	if time.Since(matches[0].StartTime) < 0 {
		t.Error("Impossbile time difference since match start")
	}
	if _, err := api.MatchWorld(1001); err != nil {
		t.Error("Couldn't find match for world id: ", err)
	}
}

func TestObjectives(t *testing.T) {
	var err error
	api := NewGW2Api()
	var testObjectives []string
	if testObjectives, err = api.Objectives(); err != nil {
		t.Error("Failed to fetch objectives: ", err)
	} else if len(testObjectives) < 1 {
		t.Error("Fetched an unlikely number of objectives")
	}
	var objectives []Objective
	if objectives, err = api.ObjectiveIds("en", testObjectives[0:2]...); err != nil || len(objectives) != 2 {
		t.Error("Failed to parse the objective data: ", err)
	}
}
