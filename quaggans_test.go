package gw2api

import (
	"testing"
	"fmt"
)

func ExampleQuaggans() {
	//get list of quaggan id's
	i, _ := Quaggans()
	fmt.Println(i[0:2])
	
	// Output:
	// [404 aloha]
}

func ExampleQuaggansIds() {
	//get specific quaggans by their id's
	i, _ := QuaggansIds("404", "aloha")
	
	for _, val := range i {
		fmt.Printf("%s - %s\n", val.ID, val.URL)
	}
	
	// Output:
	// 404 - https://static.staticwars.com/quaggans/404.jpg
	// aloha - https://static.staticwars.com/quaggans/aloha.jpg
}

//DEEP OK
func TestQuaggans(t *testing.T) {
	//t.Skip()
	id := "Quaggans"
	i, err := Quaggans()
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

//DEEP OK
func TestQuaggansIds(t *testing.T) {
	//t.Skip()
	id := "QuaggansIds"
	_, err := QuaggansIds()
	if err == nil {
		t.Errorf("No error calling %s with empty ids!", id)
	}

	_, err = QuaggansIds("", "", "", "")
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := QuaggansIds("box", "lollipop", "arf", "arf_arf", "arf_arf_arf")
	if err != nil {
		t.Errorf("Error getting %s with multiple parameters! Got:\n%s", id, err.Error())
	} else if i[0].URL == "" {
		t.Errorf("Empty output for %s with multiple parameters!")
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}
