package gw2api

import (
	"fmt"
	"testing"
)

func TestSkins(t *testing.T) {
	//t.Skip()
	id := "Skins"

	i, err := Skins()
	if err != nil {
		t.Errorf("Failed to get %s! Got:\n%s", id, err.Error())
	} else if i == nil {
		t.Errorf("Empty output for %s!", id)
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

func TestSkinIds(t *testing.T) {
	//t.Skip()
	id := "SkinIds"

	i, err := SkinIds("de", 10, 12, 0)
	if err != nil {
		t.Errorf("Error getting %s with multiple parameters! Got:\n%s", id, err.Error())
	} else if i[0].Name == "" {
		t.Errorf("Empty output for %s with multiple parameters!")
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}
