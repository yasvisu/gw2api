package gw2api

import "net/url"

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
	tag := "currencies"
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
