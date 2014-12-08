package gw2api

import (
"fmt"
"testing"
)

//OK
func TestFetchJSON(t *testing.T) {
	//t.Skip()
	id := "fetchJSON"
	_, err := fetchJSON("v2", "worlds", "")
	if err != nil {
		t.Errorf("Failed %s for worlds! Got:\n%s", id, err.Error())
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}