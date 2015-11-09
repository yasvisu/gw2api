package gw2api

import (
	"fmt"
	"testing"
)

func TestMatches(t *testing.T) {
	var matches []string
	var err error
	api := NewGW2Api()
	if matches, err = api.Matches(); err != nil {
		t.Error("Failed to fetch matches: ", err)
	} else if len(matches) < 1 {
		t.Error("Fetched an unlikely number of matches")
	}
}

func TestMatchIds(t *testing.T) {
	var matches []Match
	var err error
	api := NewGW2Api()
	if matches, err = api.MatchIds("all"); err != nil {
		t.Error("Failed to parse the match data: ", err)
	} else if len(matches) < 1 {
		t.Error("Fetched an unlikely number of matches")
	}
}

func TestObjectives(t *testing.T) {
	var objectives []string
	var err error
	api := NewGW2Api()
	if objectives, err = api.Objectives(); err != nil {
		t.Error("Failed to fetch objectives: ", err)
	} else if len(objectives) < 1 {
		t.Error("Fetched an unlikely number of objectives")
	}
}

func TestObjectiveIds(t *testing.T) {
	var objectives []Objective
	var err error
	api := NewGW2Api()
	if objectives, err = api.ObjectiveIds("all"); err != nil {
		t.Error("Failed to parse the objective data: ", err)
	} else if len(objectives) < 1 {
		t.Error("Fetched an unlikely number of objectives")
	}
}
