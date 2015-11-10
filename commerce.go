package gw2api

import (
	"fmt"
	"net/url"
	"strconv"
)

// CommerceListings returns a list of all current transactions ids
func (gw2 *GW2Api) CommerceListings() (res []int, err error) {
	ver := "v2"
	tag := "commerce/listings"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// ArticleListing includes the id of the listing and a list of buys and sells
type ArticleListing struct {
	ID    int       `json:"id"`
	Buys  []Listing `json:"buys"`
	Sells []Listing `json:"sells"`
}

// Listing includes number of listings at the unit price with respective
// quantity
type Listing struct {
	Listings  int `json:"listings"`
	UnitPrice int `json:"unit_price"`
	Quantity  int `json:"quantity"`
}

// CommerceListingIds returns the article listings for the provided ids
func (gw2 *GW2Api) CommerceListingIds(ids ...int) (articles []ArticleListing, err error) {
	ver := "v2"
	tag := "commerce/listings"
	params := url.Values{}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &articles)
	return
}

// CommerceListingPages for paginating through all existing listings
func (gw2 *GW2Api) CommerceListingPages(page int, pageSize int) (res []ArticleListing, err error) {
	if page < 0 {
		return nil, fmt.Errorf("Page parameter cannot be a negative number!")
	}

	ver := "v2"
	tag := "commerce/listings"
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	if pageSize >= 0 {
		params.Add("page_size", strconv.Itoa(pageSize))
	}
	err = gw2.fetchEndpoint(ver, tag, params, &res)
	return
}

// Exchange with CoinsPerGem and Quantity. Varies on request
type Exchange struct {
	CoinsPerGem int `json:"coins_per_gem"`
	Quantity    int `json:"quantity"`
}

// CommerceExchangeGems returns the amount of gold given for the quantity of
// coin
func (gw2 *GW2Api) CommerceExchangeGems(quantity int) (res Exchange, err error) {
	if quantity < 1 {
		return res, fmt.Errorf("Required parameter too low.")
	}

	ver := "v2"
	tag := "commerce/exchange/gems"
	params := url.Values{}
	params.Add("quantity", strconv.Itoa(quantity))
	err = gw2.fetchEndpoint(ver, tag, params, &res)
	return
}

// CommerceExchangeCoins returns the amount of gems given for the quantity of
// gems
func (gw2 *GW2Api) CommerceExchangeCoins(quantity int64) (res Exchange, err error) {
	if quantity < 1 {
		return res, fmt.Errorf("Required parameter too low.")
	}

	ver := "v2"
	tag := "commerce/exchange/coins"
	params := url.Values{}
	params.Add("quantity", strconv.FormatInt(quantity, 10))
	err = gw2.fetchEndpoint(ver, tag, params, &res)
	return
}

// ArticlePrice Listings for a specific item (highest buy, lowest sell)
type ArticlePrice struct {
	ID    int   `json:"id"`
	Buys  Price `json:"buys"`
	Sells Price `json:"sells"`
}

// Price includes quantity and the high/low price
type Price struct {
	Quantity  int `json:"quantity"`
	UnitPrice int `json:"unit_price"`
}

// CommercePrices returns a list of all ids
func (gw2 *GW2Api) CommercePrices() (res []int, err error) {
	ver := "v2"
	tag := "commerce/prices"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// CommercePriceIds returns price information about the requested ids
func (gw2 *GW2Api) CommercePriceIds(ids ...int) (artprices []ArticlePrice, err error) {
	ver := "v2"
	tag := "commerce/prices"
	params := url.Values{}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &artprices)
	return
}

// Transaction represents one of the accounts listed buy or sell orders
type Transaction struct {
	ID        int    `json:"id"`
	ItemID    int    `json:"item_id"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Created   string `json:"created"`
	Purchased string `json:"purchased"`
}

// CommerceTransactionsCurrentBuys returns all current buy orders of the account
func (gw2 *GW2Api) CommerceTransactionsCurrentBuys() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/current/buys"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}

// CommerceTransactionsCurrentSells returns all current sell orders of the account
func (gw2 *GW2Api) CommerceTransactionsCurrentSells() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/current/sells"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}

// CommerceTransactionsHistoryBuys returns all past buy orders of the account
// for the last 90 days
func (gw2 *GW2Api) CommerceTransactionsHistoryBuys() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/current/buys"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}

// CommerceTransactionsHistorySells returns all past sell orders of the account
// for the last 90 days
func (gw2 *GW2Api) CommerceTransactionsHistorySells() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/history/sells"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}
