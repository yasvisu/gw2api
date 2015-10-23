package gw2api

import (
	"fmt"
	"testing"
)

func TestObjectives(t *testing.T) {
	//t.Skip()
	id := "Objectives"

	i, err := Objectives()
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

func TestObjectiveIds(t *testing.T) {
	//t.Skip()
	id := "ObjectiveIds"
	_, err := ObjectiveIds("de")
	if err == nil {
		t.Errorf("No error calling %s with only lang!", id)
	}

	i, err := ObjectiveIds("de", "95-48", "1-1")
	if err != nil {
		t.Errorf("Error getting %s in German with multiple parameters! Got:\n%s", id, err.Error())
	} else if i[0].Name == "" {
		t.Errorf("Empty output for %s in German with multiple parameters!")
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}
