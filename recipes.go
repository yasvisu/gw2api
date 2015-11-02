package gw2api

import (
	"net/url"
	"strconv"
)

//Parent type for all Recipes endpoints.
type Recipe struct {
	ID              int                `json:"id"`
	Type            string             `json:"type"`
	OutputItemID    int                `json:"output_item_id"`
	OutputItemCount int                `json:"output_item_count"`
	TimeToCraftMS   int                `json:"time_to_craft_ms"`
	Disciplines     []string           `json:"disciplines"`
	MinRating       int                `json:"min_rating"`
	Flags           []string           `json:"flags"`
	Ingredients     []RecipeIngredient `json:"ingredients"`
}

//RecipeIngredient for Recipe.Ingredients.
type RecipeIngredient struct {
	ItemID int `json:"item_id"`
	Count  int `json:"count"`
}

//Returns list of recipe ids.
func (gw2 *GW2Api) Recipes(lang string) (res []int, err error) {
	_ = lang
	ver := "v2"
	tag := "recipes"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

//Returns list of recipes.
func (gw2 *GW2Api) RecipesIds(lang string, ids ...int) (recipes []Recipe, err error) {
	ver := "v2"
	tag := "recipes"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &recipes)
	return
}

//Internal recipes search (combined)
func (gw2 *GW2Api) recipesSearch(mode string, lang string, n int) (res []int, err error) {
	ver := "v2"
	tag := "recipes/search"

	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add(mode, strconv.Itoa(n))
	err = gw2.fetchEndpoint(ver, tag, params, &res)
	return
}

//A search interface for recipes - by input.
func (gw2 *GW2Api) RecipesSearchInput(lang string, input int) ([]int, error) {
	return gw2.recipesSearch("input", lang, input)
}

//A search interface for recipes - by output.
func (gw2 *GW2Api) RecipesSearchOutput(lang string, output int) ([]int, error) {
	return gw2.recipesSearch("output", lang, output)
}
