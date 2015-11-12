package gw2api

import (
	"fmt"
	"net/url"
	"strconv"
)

// Quaggan id and url of the image
type Quaggan struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// Quaggans list of all quaggan ids
func (gw2 *GW2Api) Quaggans() (res []string, err error) {
	ver := "v2"
	tag := "quaggans"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// QuagganIds ids -> urls for the quaggan images
func (gw2 *GW2Api) QuagganIds(ids ...string) (quag []Quaggan, err error) {
	ver := "v2"
	tag := "quaggans"
	params := url.Values{}
	params.Add("ids", commaList(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &quag)
	return
}

// World id <-> name
type World struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Worlds list of all worlds
func (gw2 *GW2Api) Worlds() (res []int, err error) {
	ver := "v2"
	tag := "worlds"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// WorldIds localized names of the worlds requested
func (gw2 *GW2Api) WorldIds(lang string, ids ...int) (worlds []World, err error) {
	ver := "v2"
	tag := "worlds"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &worlds)
	return
}

type version struct {
	ID int `json:"id"`
}

// Build returns the current Guild Wars 2 build
func (gw2 *GW2Api) Build() (v int, err error) {
	ver := "v2"
	tag := "build"

	var res version
	if err = gw2.fetchEndpoint(ver, tag, nil, &res); err != nil {
		return 0, err
	}
	return res.ID, nil
}

// Achievements - List of all achievement ids
func (gw2 *GW2Api) Achievements() (res []int, err error) {
	ver := "v2"
	tag := "achievements"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// DailyAchievementLevel range for the achievement. Determining if it is visible
// to a character or not
type DailyAchievementLevel struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

// DailyAchievement for the gamemode
type DailyAchievement struct {
	ID    int                   `json:"id"`
	Level DailyAchievementLevel `json:"level"`
}

// DailyAchievements structured by their gamemode
type DailyAchievements struct {
	PvE []DailyAchievement `json:"pve"`
	PvP []DailyAchievement `json:"pvp"`
	WvW []DailyAchievement `json:"wvw"`
}

// AchievementsDaily returns the daily achievements which can be completed today
func (gw2 *GW2Api) AchievementsDaily() (achievs DailyAchievements, err error) {
	ver := "v2"
	tag := "achievements/daily"
	err = gw2.fetchEndpoint(ver, tag, nil, &achievs)
	return
}

// Achievement detailed achievement information
type Achievement struct {
	ID          int      `json:"id"`
	Icon        string   `json:"icon"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	Type        string   `json:"type"`
	Flags       []string `json:"flags"`
}

// AchievementIds localized achievement details via ids
func (gw2 *GW2Api) AchievementIds(lang string, ids ...int) (achievs []Achievement, err error) {
	ver := "v2"
	tag := "achievements"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &achievs)
	return
}

// AchievementPages returns a paginated list of all achivements. Use it to grab
// all achievements from the API. Return values are localized to lang
func (gw2 *GW2Api) AchievementPages(lang string, page, pageSize int) (achievs []Achievement, err error) {
	if page < 0 {
		return nil, fmt.Errorf("Page parameter cannot be a negative number!")
	}
	ver := "v2"
	tag := "achievements"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("page", strconv.Itoa(page))
	if pageSize >= 0 {
		params.Add("page_size", strconv.Itoa(pageSize))
	}
	err = gw2.fetchEndpoint(ver, tag, params, &achievs)
	return
}

// Colors list of color ids
func (gw2 *GW2Api) Colors() (res []int, err error) {
	ver := "v2"
	tag := "colors"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// ColorDetail per armor type
type ColorDetail struct {
	Brightness int     `json:"brightness"`
	Contrast   float32 `json:"contrast"`
	Hue        int     `json:"hue"`
	Saturation float32 `json:"saturation"`
	Lightness  float32 `json:"lightness"`
	RGB        [3]int  `json:"rgb"`
}

// Color with specific values for the three armor types
type Color struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	BaseRGB [3]int      `json:"base_rgb"`
	Cloth   ColorDetail `json:"cloth"`
	Leather ColorDetail `json:"leather"`
	Metal   ColorDetail `json:"metal"`
}

// ColorIds fetch localized details on the requested colors
func (gw2 *GW2Api) ColorIds(lang string, ids ...int) (colors []Color, err error) {
	ver := "v2"
	tag := "colors"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &colors)
	return
}

// Currencies list of them
func (gw2 *GW2Api) Currencies() (res []int, err error) {
	ver := "v2"
	tag := "currencies"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// Currency detail struct
type Currency struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Order       int    `json:"order"`
}

// CurrencyIds localized currency information
func (gw2 *GW2Api) CurrencyIds(lang string, ids ...int) (currencies []Currency, err error) {
	ver := "v2"
	tag := "colors"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &currencies)
	return
}

// Files - List of available icons on the API
func (gw2 *GW2Api) Files() (res []string, err error) {
	ver := "v2"
	tag := "files"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// File - Id <-> icon
type File struct {
	ID   string `json:"id"`
	Icon string `json:"icon"`
}

// FileIds grabs the icon names to the id
func (gw2 *GW2Api) FileIds(ids ...string) (files []File, err error) {
	ver := "v2"
	tag := "files"
	params := url.Values{}
	params.Add("ids", commaList(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &files)
	return
}

// Minis returns a list of all mini ids
func (gw2 *GW2Api) Minis() (res []int, err error) {
	ver := "v2"
	tag := "minis"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// Mini has the basic information and the associated item id for the mini
type Mini struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Order  int    `json:"order"`
	ItemID int    `json:"item_id"`
}

// MiniIds returns the detailed information about requested minis localized to
// lang
func (gw2 *GW2Api) MiniIds(lang string, ids ...int) (minis []Mini, err error) {
	ver := "v2"
	tag := "minis"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &minis)
	return
}
