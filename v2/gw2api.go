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

GW2 API Wrapper for Go

For GW2 API Version 2
*/
package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

/* GW2API UNEXPORTED FUNCTIONS */

//Fetcher function to return only the body of a HTTP request.
//Parameters: version, tag, appendix strings.
func fetchJSON(ver string, tag string, appendix string) ([]byte, error) {
	resp, err := http.Get("https://api.guildwars2.com/" + ver + "/" + tag + ".json" + appendix)
	if err != nil {
		return nil, err
	}
	var jsonBlob []byte
	jsonBlob, err = ioutil.ReadAll(resp.Body)
	return jsonBlob, err
}

/* ERRORS */

type GW2ApiError struct {
	Text string `json:"text"`
}

func (e GW2ApiError) Error() string {
	return e.Text
}

/* ITEMS */

//General Item type all items inherit from.
type Item struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	Icon         string                 `json:"icon"`
	Description  string                 `json:"description,omitempty"`
	Type         string                 `json:"type"`
	Rarity       string                 `json:"rarity"`
	Level        int                    `json:"level"`
	VendorValue  int                    `json:"vendor_value"`
	DefaultSkin  int                    `json:"default_skin"`
	GameTypes    []string               `json:"game_types"`
	Flags        []string               `json:"flags"`
	Restrictions []string               `json:"restrictions"`
	Details      map[string]interface{} `json:"details"`
}

//Returns list of item ids.
func Items(lang string) (res []int, err error) {
	var appendix bytes.Buffer
	ver := "v2"
	tag := "items"

	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	return
}

//Returns list of items.
func ItemsDetails(page int, pageSize int, lang string, ids ...int) ([]Item, error) {
	if ids != nil {
		return ItemsIds(lang, ids...)
	} else if page >= 0 {
		return ItemsPages(page, pageSize, lang)
	} else {
		return nil, errors.New("Invalid combination of parameters. Consider using Items() instead?")
	}
}

//Returns list of items.
func ItemsIds(lang string, ids ...int) ([]Item, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Items() instead?")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "items"
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

	var res []Item
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

//Returns page of items.
func ItemsPages(page int, pageSize int, lang string) ([]Item, error) {
	if page < 0 {
		return nil, errors.New("Page parameter cannot be a negative number!")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "items"
	appendix.WriteString("?")
	concatenator := ""
	if lang != "" {
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
		concatenator = "&"
	}
	appendix.WriteString(concatenator)
	appendix.WriteString("page=")
	appendix.WriteString(strconv.Itoa(page))
	if pageSize >= 0 {
		appendix.WriteString("&page_size=")
		appendix.WriteString(strconv.Itoa(pageSize))
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []Item
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

//RECIPES

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
	var appendix bytes.Buffer
	ver := "v2"
	tag := "recipes"

	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	return
}

//Returns recipes.
func RecipesDetails(page int, pageSize int, lang string, ids ...int) ([]Recipe, error) {
	if ids != nil {
		return RecipesIds(lang, ids...)
	} else if page >= 0 {
		return RecipesPages(page, pageSize, lang)
	} else {
		return nil, errors.New("Invalid combination of parameters. Consider using Recipes() instead?")
	}
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

//Returns page of recipes.
func RecipesPages(page int, pageSize int, lang string) ([]Recipe, error) {
	if page < 0 {
		return nil, errors.New("Page parameter cannot be a negative number!")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "recipes"
	appendix.WriteString("?")
	concatenator := ""
	if lang != "" {
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
		concatenator = "&"
	}
	appendix.WriteString(concatenator)
	appendix.WriteString("page=")
	appendix.WriteString(strconv.Itoa(page))
	concatenator = "&"
	if pageSize >= 0 {
		appendix.WriteString("&page_size=")
		appendix.WriteString(strconv.Itoa(pageSize))
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
	tag := "recipes"
	concatenator := "?"

	if lang != "" {
		appendix.WriteString(concatenator)
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
		concatenator = "&"
	}
	appendix.WriteString(concatenator)
	appendix.WriteString(mode)
	appendix.WriteString("=")
	appendix.WriteString(strconv.Itoa(n))

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

/* TRADING POST */

//Returns article ids.
func CommerceListings() (res []int, err error) {
	ver := "v2"
	tag := "commerce/listings"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	return
}

//Article type to contain listings of an item.
type ArticleListings struct {
	ID    int       `json:"id"`
	Buys  []Listing `json:"buys"`
	Sells []Listing `json:"sells"`
}

//Listing type for each listing.
type Listing struct {
	Listings  int `json:"listings"`
	UnitPrice int `json:"unit_price"`
	Quantity  int `json:"quantity"`
}

//Returns list of articles.
func CommerceListingsIds(ids ...int) ([]ArticleListings, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using CommerceListings() instead?")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "commerce/listings"

	appendix.WriteString("?ids=")
	concatenator := ""
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

	var res []ArticleListings
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

//Returns page of articles.
func CommerceListingsPages(page int, pageSize int) ([]ArticleListings, error) {
	if page < 0 {
		return nil, errors.New("Page parameter cannot be a negative number!")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "commerce/listings"
	appendix.WriteString("?page=")
	appendix.WriteString(strconv.Itoa(page))
	if pageSize >= 0 {
		appendix.WriteString("&page_size=")
		appendix.WriteString(strconv.Itoa(pageSize))
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []ArticleListings
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

//COMMERCE/EXCHANGE
//Returns coins and gems exchange information.

//Exchange type for all coin/gem exchanges.
type Exchange struct {
	CoinsPerGem int `json:"coins_per_gem"`
	Quantity    int `json:"quantity"`
}

//Returns gem exchange prices.
func CommerceExchangeGems(quantity int) (Exchange, error) {
	var res Exchange
	if quantity < 1 {
		return res, errors.New("Required parameter too low.")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "commerce/exchange/gems"

	appendix.WriteString("?quantity=")
	appendix.WriteString(strconv.Itoa(quantity))

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		var gwerr GW2ApiError
		err = json.Unmarshal(data, &gwerr)
		if err != nil {
			return res, err
		}
		return res, gwerr
	}
	return res, err
}

//Returns coin exchange prices.
func CommerceExchangeCoins(quantity int64) (Exchange, error) {
	var res Exchange
	if quantity < 1 {
		return res, errors.New("Required parameter too low.")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "commerce/exchange/coins"

	appendix.WriteString("?quantity=")
	appendix.WriteString(strconv.FormatInt(quantity, 10))

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		var gwerr GW2ApiError
		err = json.Unmarshal(data, &gwerr)
		if err != nil {
			return res, err
		}
		return res, gwerr
	}
	return res, err
}

//COMMERCE/PRICES
//Returns buy and sell listing information.

//Article type to contain prices of an item.
type ArticlePrices struct {
	ID    int   `json:"id"`
	Buys  Price `json:"buys"`
	Sells Price `json:"sells"`
}

//Price type for each buy/sell price.
type Price struct {
	Quantity  int `json:"quantity"`
	UnitPrice int `json:"unit_price"`
}

//Returns list of article ids.
func CommercePrices() (res []int, err error) {
	ver := "v2"
	tag := "commerce/prices"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	return
}

//Returns list of articles.
func CommercePricesIds(ids ...int) ([]ArticlePrices, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using CommerceListings() instead?")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "commerce/prices"

	appendix.WriteString("?ids=")
	concatenator := ""
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

	var res []ArticlePrices
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

/* MISCELLANEOUS */

//QUAGGANS

//Parent type for all Quaggan Resources.
type QuagganResource struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

//Returns list of quaggan ids.
func Quaggans() (res []string, err error) {
	ver := "v2"
	tag := "quaggans"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	return
}

//Returns quaggan resources.
func QuaggansDetails(page int, pageSize int, ids ...string) ([]QuagganResource, error) {
	if ids != nil {
		return QuaggansIds(ids...)
	} else if page >= 0 {
		return QuaggansPages(page, pageSize)
	} else {
		return nil, errors.New("Invalid combination of parameters. Consider using Quaggans() instead?")
	}
}

//Returns list of quaggan resources.
func QuaggansIds(ids ...string) ([]QuagganResource, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Quaggans() instead?")
	}
	var appendix bytes.Buffer
	ver := "v2"
	tag := "quaggans"
	concatenator := "?"

	appendix.WriteString(concatenator)
	appendix.WriteString("ids=")
	concatenator = ""
	for _, i := range ids {
		appendix.WriteString(concatenator)
		appendix.WriteString(i)
		if concatenator == "" {
			concatenator = ","
		}
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []QuagganResource
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

//Returns page of quaggan resources.
func QuaggansPages(page int, pageSize int) ([]QuagganResource, error) {
	var appendix bytes.Buffer
	ver := "v2"
	tag := "quaggans"
	appendix.WriteString("?")
	concatenator := ""

	appendix.WriteString(concatenator)
	appendix.WriteString("page=")
	appendix.WriteString(strconv.Itoa(page))
	concatenator = "&"
	if pageSize >= 0 {
		appendix.WriteString("&page_size=")
		appendix.WriteString(strconv.Itoa(pageSize))
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []QuagganResource
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

//WORLDS

//Parent type for all -Names endpoints.
type Name struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Returns list of world ids.
func Worlds(lang string) (res []int, err error) {
	var appendix bytes.Buffer
	ver := "v2"
	tag := "worlds"

	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	return
}

//Returns world names.
func WorldsDetails(page int, pageSize int, lang string, ids ...int) ([]Name, error) {
	if ids != nil {
		return WorldsIds(lang, ids...)
	} else if page >= 0 {
		return WorldsPages(page, pageSize, lang)
	} else {
		return nil, errors.New("Invalid combination of parameters. Consider using Worlds() instead?")
	}
}

//Returns list of world names.
func WorldsIds(lang string, ids ...int) ([]Name, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Worlds() instead?")
	}
	var appendix bytes.Buffer
	ver := "v2"
	tag := "worlds"
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

	var res []Name
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

//Returns page of world names.
func WorldsPages(page int, pageSize int, lang string) ([]Name, error) {
	var appendix bytes.Buffer
	ver := "v2"
	tag := "worlds"
	appendix.WriteString("?")
	concatenator := ""
	if lang != "" {
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
		concatenator = "&"
	}
	appendix.WriteString(concatenator)
	appendix.WriteString("page=")
	appendix.WriteString(strconv.Itoa(page))
	concatenator = "&"
	if pageSize >= 0 {
		appendix.WriteString("&page_size=")
		appendix.WriteString(strconv.Itoa(pageSize))
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []Name
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
