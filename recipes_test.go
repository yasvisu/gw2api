package gw2api

import (
	"fmt"
	"testing"
)

func ExampleRecipes() {
	//get list of recipe id's
	i, _ := Recipes("")
	fmt.Println(i[0:2])
	// Output:
	// [1 2]
}

func ExampleRecipesIds() {
	//get specific recipes by their id's, in Spanish
	i, _ := RecipesIds("es", 7319)

	for _, val := range i {
		fmt.Printf("%s - %d\n", val.Type, val.OutputItemID)
	}
	// Output:
	// RefinementEctoplasm - 46742
}

func ExampleRecipeSearch() {
	//search specific recipes by their id's, in Spanish
	i, _ := RecipesSearchInput("es", 46731) //recipes taking id 46731 as input, in Spanish
	j, _ := RecipesSearchOutput("", 50065)  //recipes producing id 50065 as output, in English
	fmt.Printf("46731 - %d\n", i)
	fmt.Printf("50065 - %d\n", j)
	// Output:
	// 46731 - [7314 7839 7840 7841 7846 7847 7848 7849 7850 10296 10718 11234]
	// 50065 - [8455 8459 8460]
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
