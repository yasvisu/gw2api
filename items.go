package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

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
	ver := "v2"
	tag := "items"
	err = fetchEndpoint(ver, tag, lang, &res)
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
func ItemsIds(lang string, ids ...int) (items []Item, err error) {
	ver := "v2"
	tag := "items"
	err = fetchDetailEndpoint(ver, tag, lang, stringSlice(ids), &items)
	return
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
