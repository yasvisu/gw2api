package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
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
func Recipes(lang string) (res []int, err error) {
	ver := "v2"
	tag := "recipes"
	err = fetchEndpoint(ver, tag, lang, &res)
	return
}

//Returns list of recipes.
func RecipesIds(lang string, ids ...int) ([]Recipe, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Recipes() instead?")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "recipes"
	concatenator := "?"

	if lang != "" {
		appendix.WriteString(concatenator)
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
		concatenator = "&"
	}
	appendix.WriteString(concatenator)
	appendix.WriteString("ids=")
	concatenator = ""
	for _, i := range ids {
		appendix.WriteString(concatenator)
		appendix.WriteString(strconv.Itoa(i))
		if concatenator == "" {
			concatenator = ","
		}
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []Recipe
	err = json.Unmarshal(data, &res)
	if err != nil {
		var gwerr GW2ApiError
		err = json.Unmarshal(data, &gwerr)
		if err != nil {
			return nil, err
		}
		return nil, gwerr
	}
	return res, err
}

//Internal recipes search (combined)
func recipesSearch(mode string, lang string, n int) ([]int, error) {
	var appendix bytes.Buffer
	ver := "v2"
	tag := "recipes/search"
	concatenator := "?"

	appendix.WriteString(concatenator)
	appendix.WriteString(mode)
	appendix.WriteString("=")
	appendix.WriteString(strconv.Itoa(n))

	if lang != "" {
		appendix.WriteString(concatenator)
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
		concatenator = "&"
	}
	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []int
	err = json.Unmarshal(data, &res)
	if err != nil {
		var gwerr GW2ApiError
		err = json.Unmarshal(data, &gwerr)
		if err != nil {
			return nil, err
		}
		return nil, gwerr
	}
	return res, err
}

//A search interface for recipes - by input.
func RecipesSearchInput(lang string, input int) ([]int, error) {
	return recipesSearch("input", lang, input)
}

//A search interface for recipes - by output.
func RecipesSearchOutput(lang string, output int) ([]int, error) {
	return recipesSearch("output", lang, output)
}
