package gw2api

import (
	"fmt"
	"testing"
)

func ExampleCommerceExchangeGems() {
	//get current coin-to-gem exchange ratio
	i, _ := CommerceExchangeCoins(2000000) //make sure to include a big enough quantity of coins
	i = i                                  //do something with i
	fmt.Printf("Coins per gem - Quantity of coins\n")
	// Output:
	// Coins per gem - Quantity of coins
}

func ExampleCommerceExchangeCoins() {
	//get current gem-to-coin exchange ratio
	i, _ := CommerceExchangeCoins(1000)
	i = i //do something with i
	fmt.Printf("Coins per gem - Quantity of gems\n")
	// Output:
	// Coins per gem - Quantity of gems
}

func ExampleCommerceListings() {
	//get list of listings id's
	i, _ := CommerceListings()
	i = i //do something with i
	fmt.Println("Listings")
	// Output:
	// Listings
}

func ExampleCommerceListingsIds() {
	//get specific listings by their id's
	i, _ := CommerceListingsIds(24, 68, 69)
	i = i //do something with i
	fmt.Println("Listings")
	// Output:
	// Listings
}

func ExampleCommerceListingsPages() {
	//get specific listings by their pages
	i, _ := CommerceListingsPages(3, 2) //get page 3 with page_size 2
	i = i                               //do something with i
	fmt.Println("Listings")
	// Output:
	// Listings
}

func ExampleCommercePrices() {
	//get list of prices id's
	i, _ := CommercePrices()
	i = i //do something with i
	fmt.Println("Listings")
	// Output:
	// Listings
}

func ExampleCommercePricesIds() {
	//get specific prices by their id's
	i, _ := CommercePricesIds(24, 68, 69)
	i = i //do something with i
	fmt.Println("Prices")
	// Output:
	// Prices
}

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
