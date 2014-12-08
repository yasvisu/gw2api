package gw2api

import (
"fmt"
"testing"
)

func ExampleWorlds() {
	//get list of world id's
	i, _ := Worlds("")
	fmt.Println(i[0:2])
	
	// Output:
	// [1001 1002]
}

func ExampleWorldsIds() {
	//get specific worlds by their id's, in Spanish
	i, _ := WorldsIds("es", 1001, 1014, 2206, 13371337)
	
	for _, val := range i {
		fmt.Printf("%d - %s\n", val.ID, val.Name)
	}
	
	// Output:
	// 1001 - Roca del Yunque
	// 1014 - Desierto de Cristal
	// 2206 - Estrecho de Miller [DE]
}

//DEEP OK
func TestWorlds(t *testing.T) {
	//t.Skip()
	id := "Worlds"
	_, err := Worlds("es")
	if err != nil {
		t.Errorf("Failed to get %s in Spanish! Got:\n%s", id, err.Error())
	}

	i, err := Worlds("")
	if err != nil {
		t.Errorf("Failed to get %s with empty parameter! Got:\n%s", id, err.Error())
	} else if i == nil {
		t.Errorf("Empty output for %s with empty parameter!", id)
	}
	fmt.Printf("\t-%s\t\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

//DEEP OK
func TestWorldsIds(t *testing.T) {
	//t.Skip()
	id := "WorldsIds"
	_, err := WorldsIds("es")
	if err == nil {
		t.Errorf("No error calling %s with only lang!", id)
	}

	_, err = WorldsIds("", 0, 0, 0, 0)
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := WorldsIds("es", 1001, 1002, 3, 4, 5)
	if err != nil {
		t.Errorf("Error getting %s in Spanish with multiple parameters! Got:\n%s", id, err.Error())
	} else if i[0].Name == "" {
		t.Errorf("Empty output for %s in Spanish with multiple parameters!")
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}
