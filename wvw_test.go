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
		fmt.Errorf("Failed to fetch matches: %s", err)
	} else if len(matches) < 1 {
		fmt.Errorf("Fetched an unlikely number of matches")
	}
}

func TestMatchIds(t *testing.T) {
	var matches []Match
	var err error
	api := NewGW2Api()
	if matches, err = api.MatchIds("all"); err != nil {
		fmt.Errorf("Failed to parse the match data: %s", err)
	} else if len(matches) < 1 {
		fmt.Errorf("Fetched an unlikely number of matches")
	}
}

func TestObjectives(t *testing.T) {
	var objectives []string
	var err error
	api := NewGW2Api()
	if objectives, err = api.Objectives(); err != nil {
		fmt.Errorf("Failed to fetch objectives: %s", err)
	} else if len(objectives) < 1 {
		fmt.Errorf("Fetched an unlikely number of objectives")
	}
}

func TestObjectiveIds(t *testing.T) {
	var objectives []Objective
	var err error
	api := NewGW2Api()
	if objectives, err = api.ObjectiveIds("all"); err != nil {
		fmt.Errorf("Failed to parse the objective data: %s", err)
	} else if len(objectives) < 1 {
		fmt.Errorf("Fetched an unlikely number of objectives")
	}
}
