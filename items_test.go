package gw2api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func ExampleItems() {
	//get list of item id's
	i, _ := Items("")
	fmt.Println(i[0:3])
	// Output:
	// [1 2 6]
}

func ExampleItemsIds() {
	//get specific items by their id's, in French
	i, _ := ItemsIds("fr", 12452)

	for _, val := range i {
		fmt.Printf("ID %d - %s - %s\n", val.ID, val.Type, val.Name)
	}
	// Output:
	// ID 12452 - Consumable - Barre aux baies d'Omnom
}

func ExampleItemsIds_icons() {
	//get a specific icon off the API
	i, _ := Items("")
	j, _ := ItemsIds("", i[0:3]...)
	resp, _ := http.Get(j[0].Icon)       //get the http response
	blob, _ := ioutil.ReadAll(resp.Body) //read all the bytes into a blob []byte
	//do interesting things!
	blob = blob
}

func ExampleItemsPages() {
	//get specific items by their pages
	i, _ := ItemsPages(3, 1, "") //get page 3 with page_size 1, in English

	for _, val := range i {
		fmt.Printf("%s - %s\n", val.Name, val.Icon)
	}
	// Output:
	// Undead Unarmed - https://render.guildwars2.com/file/C5B365D6105F76470106A61F4AB96F3E39D10E18/60991.png
}

func TestItems(t *testing.T) {
	//t.Skip()
	id := "Items"
	_, err := Items("es")
	if err != nil {
		t.Errorf("Failed to get %s in Spanish! Got:\n%s", id, err.Error())
	}

	i, err := Items("")
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
func TestItemsIds(t *testing.T) {
	//t.Skip()
	id := "ItemsIds"
	_, err := ItemsIds("es")
	if err == nil {
		t.Errorf("No error calling %s with only lang!", id)
	}

	_, err = ItemsIds("", 0, 0, 0, 0)
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := ItemsIds("es", 162, 1002, 3, 4, 5)
	if err != nil {
		t.Errorf("Error getting %s in Spanish with multiple parameters! Got:\n%s", id, err.Error())
	} else if i[0].ID == 0 {
		t.Errorf("Empty output for %s in Spanish with multiple parameters!")
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

//DEEP OK
func TestItemsPages(t *testing.T) {
	//t.Skip()
	id := "ItemsPages"
	_, err := ItemsPages(-1, -1, "es")
	if err == nil {
		t.Errorf("No error calling %s with only lang!", id)
	}

	_, err = ItemsPages(-1, 0, "")
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := ItemsPages(0, 0, "es")
	if err != nil {
		t.Errorf("Error getting %s in Spanish with multiple parameters! Got:\n%s", id, err.Error())
	} else if i == nil {
		t.Errorf("Empty output for %s in Spanish with multiple parameters!")
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

//DEEP OK
func TestItemsUnmarshalling(t *testing.T) {
	//t.Skip()
	id := "ItemsUnmarshalling"
	i, err := ItemsIds("", 162)
	if err != nil {
		t.Errorf("Error getting data for %s!", id)
	}
	if i[0].Name == "" {
		t.Errorf("%s: Empty Name!", id)
	}
	if i[0].Type == "" {
		t.Errorf("%s: Empty Type!", id)
	}
	if i[0].Level == 0 {
		t.Errorf("%s: Empty Level!", id)
	}
	if i[0].Rarity == "" {
		t.Errorf("%s: Empty Rarity!", id)
	}
	if i[0].VendorValue == 0 {
		t.Errorf("%s: Empty VendorValue!", id)
	}
	if i[0].DefaultSkin == 0 {
		t.Errorf("%s: Empty DefaultSkin!", id)
	}
	if i[0].GameTypes == nil {
		t.Errorf("%s: Empty GameTypes!", id)
	}
	if i[0].Flags == nil {
		t.Errorf("%s: Empty Flags!", id)
	}
	if i[0].ID == 0 {
		t.Errorf("%s: Empty ID!", id)
	}
	if i[0].Icon == "" {
		t.Errorf("%s: Empty Icon!", id)
	}
	if i[0].Details == nil {
		t.Errorf("%s: Empty Details!", id)
	}
	fmt.Printf("\t-%s\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}
