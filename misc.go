package gw2api

import "net/url"

//Parent type for all Quaggan Resources.
type QuagganResource struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

//Returns list of quaggan ids.
func Quaggans() (res []string, err error) {
	ver := "v2"
	tag := "quaggans"
	err = fetchEndpoint(ver, tag, nil, &res)
	return
}

//Returns list of quaggan resources.
func QuaggansIds(ids ...string) (quag []QuagganResource, err error) {
	ver := "v2"
	tag := "quaggans"
	params := url.Values{}
	params.Add("ids", commaList(ids))
	err = fetchEndpoint(ver, tag, params, &quag)
	return
}

//Parent type for all -Names endpoints.
type Name struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Returns list of world ids.
func Worlds(lang string) (res []int, err error) {
	_ = lang
	ver := "v2"
	tag := "worlds"
	err = fetchEndpoint(ver, tag, nil, &res)
	return
}

//Returns list of world names.
func WorldsIds(lang string, ids ...int) (worlds []Name, err error) {
	ver := "v2"
	tag := "worlds"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = fetchEndpoint(ver, tag, params, &worlds)
	return
}

type Version struct {
	ID int `json:"id"`
}

func Build() (v int, err error) {
	ver := "v2"
	tag := "build"

	var res Version
	if err = fetchEndpoint(ver, tag, nil, &res); err != nil {
		return 0, err
	}
	return res.ID, nil
}

func Achievements() (res []string, err error) {
	ver := "v2"
	tag := "achievements"
	err = fetchEndpoint(ver, tag, nil, &res)
	return
}

type Achievement struct {
	ID          int      `json:"id"`
	Icon        string   `json:"icon"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	Type        string   `json:"type"`
	Flags       []string `json:"flags"`
}

func AchievmentIds(lang string, ids ...int) (achievs []Achievement, err error) {
	ver := "v2"
	tag := "achievements"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = fetchEndpoint(ver, tag, params, &achievs)
	return
}

func Colors() (res []int, err error) {
	ver := "v2"
	tag := "colors"
	err = fetchEndpoint(ver, tag, nil, &res)
	return
}

type ColorDetail struct {
	Brightness int     `json:"brightness"`
	Contrast   float32 `json:"contrast"`
	Hue        int     `json:"hue"`
	Saturation float32 `json:"saturation"`
	Lightness  float32 `json:"lightness"`
	RGB        [3]int  `json:"rgb"`
}

type Color struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	BaseRGB [3]int      `json:"base_rgb"`
	Cloth   ColorDetail `json:"cloth"`
	Leather ColorDetail `json:"leather"`
	Metal   ColorDetail `json:"metal"`
}

func ColorIds(lang string, ids ...int) (colors []Color, err error) {
	ver := "v2"
	tag := "colors"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = fetchEndpoint(ver, tag, params, &colors)
	return
}

func Currencies() (res []string, err error) {
	ver := "v2"
	tag := "currencies"
	err = fetchEndpoint(ver, tag, nil, &res)
	return
}

type Currency struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Order       int    `json:"order"`
}

func CurrencyIds(lang string, ids ...int) (currencies []Currency, err error) {
	ver := "v2"
	tag := "colors"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = fetchEndpoint(ver, tag, params, &currencies)
	return
}

func Files() (res []string, err error) {
	ver := "v2"
	tag := "files"
	err = fetchEndpoint(ver, tag, nil, &res)
	return
}

type File struct {
	ID   string `json:"id"`
	Icon string `json:"icon"`
}

func FileIds(ids ...string) (files []File, err error) {
	ver := "v2"
	tag := "files"
	params := url.Values{}
	params.Add("ids", commaList(ids))
	err = fetchEndpoint(ver, tag, params, &files)
	return
}
