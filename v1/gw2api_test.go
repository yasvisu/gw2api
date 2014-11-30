/*
The MIT License (MIT)

Copyright (c) 2014 Yasen Visulchev

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

/* 
Written by Yasen Visulchev - Sofia University FMI/Golang
	20141130
*/
package gw2api

import (
	"fmt"
	"testing"
)

// Tests

/* GW2API UNEXPORTED FUNCTIONS */

//OK
func TestFetchJSON(t *testing.T) {
	//t.Skip()
	id := "fetchJSON"
	_, err := fetchJSON("v1", "world_names", "")
	if err != nil {
		t.Errorf("Failed %s for world_names! Got:\n%s", id, err.Error())
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

/* DYNAMIC EVENTS */

//EXPERIMENTAL
func TestEvents(t *testing.T) {
	t.Skip()
	id := "Events"
	es, err := Events(0, 0, "")
	if err != nil {
		t.Errorf("Failed to get %s! Got:\n%s", id, err.Error())
	}
	if es == nil {
		t.Errorf("%s is nil! API broken?: %s", id, es)
	}
	fmt.Printf("\t-%s\t\t\tOK\n", id)
}

//OK
func TestEventNames(t *testing.T) {
	//t.Skip()
	id := "EventNames"
	_, err := EventNames("es")
	if err != nil {
		t.Errorf("Failed to get %s! Got:\n%s", id, err.Error())
	}

	_, err = EventNames("")
	if err != nil {
		t.Errorf("Failed to get %s with empty appendix! Got:\n%s", id, err.Error)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

//OK
func TestMapNames(t *testing.T) {
	//t.Skip()
	id := "MapNames"
	_, err := MapNames("es")
	if err != nil {
		t.Errorf("Failed to get %s! Got:\n%s", id, err.Error())
	}

	_, err = MapNames("")
	if err != nil {
		t.Errorf("Failed to get %s with empty appendix! Got:\n%s", id, err.Error)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

//OK
func TestWorldNames(t *testing.T) {
	//t.Skip()
	id := "WorldNames"
	_, err := WorldNames("es")
	if err != nil {
		t.Errorf("Failed to get %s! Got:\n%s", id, err.Error())
	}

	_, err = WorldNames("")
	if err != nil {
		t.Errorf("Failed to get %s with empty appendix! Got:\n%s", id, err.Error)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

//OK
func TestEventDetails(t *testing.T) {
	//t.Skip()
	id := "EventDetails"
	_, err := EventDetails("", "es")
	if err != nil {
		t.Errorf("Failed to get %s in Spanish! Got:\n%s", id, err.Error())
	}

	_, err = EventDetails("EED8A79F-B374-4AE6-BA6F-B7B98D9D7142", "")
	if err != nil {
		t.Errorf("Failed to get %s for EED8A79F-B374-4AE6-BA6F-B7B98D9D7142", id)
	}

	_, err = EventDetails("EED8A79F-B374-4AE6-BA6F-B7B98D9D7142", "es")
	if err != nil {
		t.Errorf("Failed to get %s for EED8A79F-B374-4AE6-BA6F-B7B98D9D7142 in Spanish!", id)
	}

	_, err = EventDetails("", "")
	if err != nil {
		t.Errorf("Failed to get %s with empty appendix! Got:\n%s", id, err.Error)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

/* GUILDS */

//OK
func TestGuildDetails(t *testing.T) {
	//t.Skip()
	id := "GuildDetails"
	_, err := GuildDetails("", "")
	if err == nil {
		t.Errorf("No error getting %s with empty required parameters\n", id)
	}

	_, err = GuildDetails("", "When%20Time%20Flies")
	if err != nil {
		t.Errorf("Error getting %s for When Time Flies:\n\t%s\n", id, err)
	}
	_, err = GuildDetails("B986E58F-17A6-4D68-8C50-21E3C47FE103", "")
	if err != nil {
		t.Errorf("Error getting %s for B986E58F-17A6-4D68-8C50-21E3C47FE103:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

/* ITEMS */

//OK
func TestItems(t *testing.T) {
	//t.Skip()
	id := "Items"
	_, err := Items("")
	if err != nil {
		t.Errorf("Error getting %s with empty lang parameter:\n\t%s\n", id, err)
	}

	_, err = Items("es")
	if err != nil {
		t.Errorf("Error getting %s in Spanish:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\t\tOK\n", id)
}

//OK
func TestItemDetails(t *testing.T) {
	//t.Skip()
	id := "ItemDetails"
	_, err := ItemDetails("", "arf arf arf")
	if err == nil {
		t.Errorf("No error getting %s with empty required parameters\n", id)
	}

	_, err = ItemDetails("54770", "es")
	if err != nil {
		t.Errorf("Error getting %s for 54770 in Spanish:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

//OK
func TestRecipes(t *testing.T) {
	//t.Skip()
	id := "Recipes"
	_, err := Recipes("")
	if err != nil {
		t.Errorf("Error getting %s with empty lang parameter:\n\t%s\n", id, err)
	}

	_, err = Recipes("es")
	if err != nil {
		t.Errorf("Error getting %s in Spanish:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

//OK
func TestRecipeDetails(t *testing.T) {
	//t.Skip()
	id := "RecipeDetails"
	_, err := RecipeDetails("", "arf arf arf")
	if err == nil {
		t.Errorf("No error getting %s with empty required parameters\n", id)
	}

	_, err = RecipeDetails("3147", "es")
	if err != nil {
		t.Errorf("Error getting %s for 3147 in Spanish:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

//OK
func TestSkins(t *testing.T) {
	//t.Skip()
	id := "Skins"
	_, err := Skins()
	if err != nil {
		t.Errorf("Error getting %s:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\t\tOK\n", id)
}

//OK
func TestSkinDetails(t *testing.T) {
	//t.Skip()
	id := "SkinDetails"
	_, err := SkinDetails("", "arf arf arf")
	if err == nil {
		t.Errorf("No error getting %s with empty required parameters\n", id)
	}

	_, err = SkinDetails("900", "es")
	if err != nil {
		t.Errorf("Error getting %s for 900 in Spanish!:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

/* MAP INFORMATION */

//OK
func TestContinents(t *testing.T) {
	//t.Skip()
	id := "Continents"
	_, err := Continents("")
	if err != nil {
		t.Errorf("Error getting %s with empty lang parameter:\n\t%s\n", id, err)
	}

	_, err = Continents("es")
	if err != nil {
		t.Errorf("Error getting %s in Spanish:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

//OK
func TestMaps(t *testing.T) {
	//t.Skip()
	id := "Maps"
	_, err := Maps("", "")
	if err != nil {
		t.Errorf("Error getting %s with empty parameters:\n\t%s\n", id, err)
	}

	_, err = Maps("15", "es")
	if err != nil {
		t.Errorf("Error getting %s 15 in Spanish:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\t\tOK\n", id)
}

//OK
func TestMapFloor(t *testing.T) {
	//t.Skip()
	id := "MapFloor"
	_, err := MapFloor("", "", "arf arf arf")
	if err == nil {
		t.Errorf("No error getting %s with empty parameters!", id)
	}

	_, err = MapFloor("1", "2", "es")
	if err != nil {
		t.Errorf("Error getting %s, continent 1, floor 2 in Spanish:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

/* WORLD VS. WORLD */

//OK
func TestWVWMatches(t *testing.T) {
	//t.Skip()
	id := "WVWMatches"
	_, err := WVWMatches()
	if err != nil {
		t.Errorf("Error getting %s:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\tOK\n", id)
}

//OK
func TestWVWMatchDetails(t *testing.T) {
	//t.Skip()
	id := "WVWMatchDetails"
	_, err := WVWMatchDetails("")
	if err == nil {
		t.Errorf("No error getting %s with empty required parameters\n", id)
	}

	_, err = WVWMatchDetails("1-4")
	if err != nil {
		t.Errorf("Error getting %s for 1-4:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\tOK\n", id)
}

//OK
func TestWVWObjectiveNames(t *testing.T) {
	//t.Skip()
	id := "WVWObjectiveNames"
	_, err := WVWObjectiveNames("es")
	if err != nil {
		t.Errorf("Failed to get %s! Got:\n%s", id, err.Error())
	}

	_, err = WVWObjectiveNames("")
	if err != nil {
		t.Errorf("Failed to get %s with empty appendix! Got:\n%s", id, err.Error)
	}
	fmt.Printf("\t-%s\tOK\n", id)
}

/* MISCELLANEOUS */

//OK
func TestBuild(t *testing.T) {
	//t.Skip()
	id := "Build"
	_, err := Build()
	if err != nil {
		t.Errorf("Error getting %s:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\t\tOK\n", id)
}

//OK
func TestColors(t *testing.T) {
	//t.Skip()
	id := "Colors"
	_, err := Colors("")
	if err != nil {
		t.Errorf("Error getting %s with empty lang parameter:\n\t%s\n", id, err)
	}

	_, err = Colors("es")
	if err != nil {
		t.Errorf("Error getting %s in Spanish:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\t\tOK\n", id)
}

//OK
func TestFiles(t *testing.T) {
	//t.Skip()
	id := "Files"
	_, err := Files()
	if err != nil {
		t.Errorf("Error getting %s:\n\t%s\n", id, err)
	}
	fmt.Printf("\t-%s\t\t\tOK\n", id)
}
