package gw2api

import (
	"fmt"
	"testing"
)

func TestMatches(t *testing.T) {
	//t.Skip()
	id := "Matches"

	i, err := Matches()
	if err != nil {
		t.Errorf("Failed to get %s! Got:\n%s", id, err.Error())
	} else if i == nil {
		t.Errorf("Empty output for %s!", id)
	}
	fmt.Printf("\t-%s\t\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

func TestMatchIds(t *testing.T) {
	//t.Skip()
	id := "MatchIds"

	i, err := MatchIds("2-1", "1-1")
	if err != nil {
		t.Errorf("Error getting %s with multiple parameters! Got:\n%s", id, err.Error())
	} else if i[0].StartTime == "" {
		t.Errorf("Empty output for %s with multiple parameters!")
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}
