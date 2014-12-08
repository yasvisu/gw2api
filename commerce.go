package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

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
