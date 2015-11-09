package gw2api

import "testing"

func TestRecipesIds(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testRecipes []int
	if testRecipes, err = api.Recipes(); err != nil {
		t.Error("Failed to fetch recipe ids")
	}

	var recipes []Recipe
	if recipes, err = api.RecipeIds(testRecipes[0], testRecipes[1]); err != nil {
		t.Error("Failed to parse the recipe data: ", err)
	} else if len(recipes) != 2 {
		t.Error("Failed to fetch existing recipes")
	}
}

func TestRecipesSearchInput(t *testing.T) {
	var err error
	api := NewGW2Api()

	var recipes []int
	if recipes, err = api.RecipeSearchInput(46731); err != nil {
		t.Error("Failed to parse the recipe data: ", err)
	} else if len(recipes) < 9 {
		t.Error("Failed to fetch existing recipes: 9 > ", len(recipes))
	}
}

func TestRecipesSearchOutput(t *testing.T) {
	var err error
	api := NewGW2Api()

	var recipes []int
	if recipes, err = api.RecipeSearchOutput(50065); err != nil {
		t.Error("Failed to parse the recipe data: ", err)
	} else if len(recipes) < 3 {
		t.Error("Failed to fetch existing recipes: 3 > ", len(recipes))
	}
}
