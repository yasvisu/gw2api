package gw2api

import (
	"fmt"
	"net/url"
	"strconv"
)

//Returns article ids.
func (gw2 *GW2Api) CommerceListings() (res []int, err error) {
	ver := "v2"
	tag := "commerce/listings"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
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
func (gw2 *GW2Api) CommerceListingsIds(ids ...int) (articles []ArticleListings, err error) {
	ver := "v2"
	tag := "commerce/listings"
	params := url.Values{}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &articles)
	return
}

//Returns page of articles.
func (gw2 *GW2Api) CommerceListingsPages(page int, pageSize int) (res []ArticleListings, err error) {
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

//COMMERCE/EXCHANGE
//Returns coins and gems exchange information.

//Exchange type for all coin/gem exchanges.
type Exchange struct {
	CoinsPerGem int `json:"coins_per_gem"`
	Quantity    int `json:"quantity"`
}

//Returns gem exchange prices.
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

//Returns coin exchange prices.
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
func (gw2 *GW2Api) CommercePrices() (res []int, err error) {
	ver := "v2"
	tag := "commerce/prices"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

//Returns list of articles.
func (gw2 *GW2Api) CommercePricesIds(ids ...int) (artprices []ArticlePrices, err error) {
	ver := "v2"
	tag := "commerce/prices"
	params := url.Values{}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &artprices)
	return
}

type Transaction struct {
	ID        int    `json:"id"`
	ItemID    int    `json:"item_id"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Created   string `json:"created"`
	Purchased string `json:"purchased"`
}

func (gw2 *GW2Api) CommerceTransactionsCurrentBuys() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/current/buys"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}

func (gw2 *GW2Api) CommerceTransactionsCurrentSells() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/current/sells"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}

func (gw2 *GW2Api) CommerceTransactionsHistoryBuys() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/current/buys"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}

func (gw2 *GW2Api) CommerceTransactionsHistorySells() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/history/sells"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}
