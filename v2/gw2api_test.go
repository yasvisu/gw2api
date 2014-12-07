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
	20141207
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

/* ITEMS */

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

//DEEP OK
func TestRecipes(t *testing.T) {
	//t.Skip()
	id := "Recipes"
	_, err := Recipes("es")
	if err != nil {
		t.Errorf("Failed to get %s in Spanish! Got:\n%s", id, err.Error())
	}

	i, err := Recipes("")
	if err != nil {
		t.Errorf("Failed to get %s with empty parameter! Got:\n%s", id, err.Error())
	} else if i == nil {
		t.Errorf("Empty output for %s with empty parameter!", id)
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

//DEEP OK
func TestRecipesIds(t *testing.T) {
	//t.Skip()
	id := "RecipesIds"
	_, err := RecipesIds("es")
	if err == nil {
		t.Errorf("No error calling %s with only lang!", id)
	}

	_, err = RecipesIds("", 0, 0, 0, 0)
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := RecipesIds("es", 7319, 1002, 3, 4, 5)
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
func TestRecipesPages(t *testing.T) {
	//t.Skip()
	id := "RecipesPages"
	_, err := RecipesPages(-1, -1, "es")
	if err == nil {
		t.Errorf("No error calling %s with only lang!", id)
	}

	_, err = RecipesPages(-1, 0, "")
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := RecipesPages(0, 0, "es")
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
func TestRecipesUnmarshalling(t *testing.T) {
	//t.Skip()
	id := "RecipesUnmarshalling"
	i, err := RecipesIds("", 7319)
	if err != nil {
		t.Errorf("Error getting data for %s!", id)
	}
	switch {
	case i[0].Type == "":
		t.Errorf("%s: Empty Type!", id)
	case i[0].OutputItemID == 0:
		t.Errorf("%s: Empty OutputItemID!", id)
	case i[0].OutputItemCount == 0:
		t.Errorf("%s: Empty OutputItemCount!", id)
	case i[0].MinRating == 0:
		t.Errorf("%s: Empty MinRating!", id)
	case i[0].TimeToCraftMS == 0:
		t.Errorf("%s: Empty TimeToCraftMS!", id)
	case i[0].Disciplines == nil:
		t.Errorf("%s: Empty Disciplines!", id)
	case i[0].Flags == nil:
		t.Errorf("%s: Empty Flags!", id)
	case i[0].Ingredients == nil:
		t.Errorf("%s: Empty Ingredients!", id)
	case i[0].ID == 0:
		t.Errorf("%s: Empty ID!", id)
	}
	fmt.Printf("\t-%s\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

//DEEP OK
func TestRecipesSearch(t *testing.T) {
	//t.Skip()
	id := "RecipesSearch"
	i, err := RecipesSearchInput("es", 46731)
	if err != nil {
		t.Errorf("Error getting %sInput for 46731 in Spanish!", id)
	} else if i == nil {
		t.Errorf("Empty output for %sInput for 46731 in Spanish!", id)
	}

	i, err = RecipesSearchOutput("", 50065)
	if err != nil {
		t.Errorf("Error getting %sOutput for 50065!", id)
	} else if i == nil {
		t.Errorf("Empty output for %sOutput for 50065!", id)
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

/* TRADING POST */

//DEEP OK
func TestCommerceExchangeGems(t *testing.T) {
	//t.Skip()
	id := "CommerceExchangeGems"

	_, err := CommerceExchangeGems(-1)
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := CommerceExchangeGems(2000)
	if err != nil {
		t.Errorf("Error getting %s with multiple parameters! Got:\n%s", id, err.Error())
	} else if i.Quantity == 0 {
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
func TestCommerceExchangeCoins(t *testing.T) {
	//t.Skip()
	id := "CommerceExchangeCoins"

	_, err := CommerceExchangeCoins(-1)
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := CommerceExchangeCoins(2000000)
	if err != nil {
		t.Errorf("Error getting %s with multiple parameters! Got:\n%s", id, err.Error())
	} else if i.Quantity == 0 {
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
func TestCommerceListings(t *testing.T) {
	//t.Skip()
	id := "CommerceListings"

	i, err := CommerceListings()
	if err != nil {
		t.Errorf("Failed to get %s! Got:\n%s", id, err.Error())
	} else if i == nil {
		t.Errorf("Empty output for %s with empty parameter!", id)
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

//DEEP OK
func TestCommerceListingsIds(t *testing.T) {
	//t.Skip()
	id := "CommerceListingsIds"
	_, err := CommerceListingsIds()
	if err == nil {
		t.Errorf("No error calling %s with empty parameters!", id)
	}

	_, err = CommerceListingsIds(0, 0, 0, 0)
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := CommerceListingsIds(19684, 1002, 3, 4, 5)
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
func TestCommerceListingsPages(t *testing.T) {
	//t.Skip()
	id := "CommerceListingsPages"
	_, err := CommerceListingsPages(-1, -1)
	if err == nil {
		t.Errorf("No error calling %s with invalid parameters!", id)
	}

	_, err = CommerceListingsPages(-1, 0)
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := CommerceListingsPages(0, 0)
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
func TestCommercePrices(t *testing.T) {
	//t.Skip()
	id := "CommercePrices"

	i, err := CommercePrices()
	if err != nil {
		t.Errorf("Failed to get %s! Got:\n%s", id, err.Error())
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
func TestCommercePricesIds(t *testing.T) {
	//t.Skip()
	id := "CommercePricesIds"
	_, err := CommercePricesIds()
	if err == nil {
		t.Errorf("No error calling %s with empty parameters!", id)
	}

	_, err = CommercePricesIds(0, 0, 0, 0)
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := CommercePricesIds(19684, 1002, 3, 4, 5)
	if err != nil {
		t.Errorf("Error getting %s with multiple parameters! Got:\n%s", id, err.Error())
	} else if i[0].ID == 0 {
		t.Errorf("Empty output for %s with multiple parameters!")
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
}

/* MISCELLANEOUS */

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

//DEEP OK
func TestQuaggansPages(t *testing.T) {
	//t.Skip()
	id := "QuaggansPages"
	_, err := QuaggansPages(-1, 0)
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := QuaggansPages(0, 0)
	if err != nil {
		t.Errorf("Error getting %s with multiple parameters! Got:\n%s", id, err.Error())
	} else if i == nil {
		t.Errorf("Empty output for %s with multiple parameters!")
	}
	fmt.Printf("\t-%s\t\t", id)
	if !t.Failed() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("FAILED\n")
	}
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

//DEEP OK
func TestWorldsPages(t *testing.T) {
	//t.Skip()
	id := "WorldsPages"
	_, err := WorldsPages(-1, -1, "es")
	if err == nil {
		t.Errorf("No error calling %s with only lang!", id)
	}

	_, err = WorldsPages(-1, 0, "")
	if err == nil {
		t.Errorf("No error unmarshalled when calling %s with invalid parameters!", id)
	}

	i, err := WorldsPages(0, 0, "es")
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
