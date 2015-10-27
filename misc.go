package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

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

func Build() (v int, err error) {
	ver := "v2"
	tag := "build"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return 0, err
	}

	type JsonVersion struct {
		ID int `json:"id"`
	}

	var res JsonVersion
	err = json.Unmarshal(data, &res)
	return res.ID, nil
}

func Achievements(lang string) (res []string, err error) {
	var appendix bytes.Buffer
	ver := "v2"
	tag := "achievements"

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

type Achievement struct {
	ID          int      `json:"id"`
	Icon        string   `json:"icon"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	Type        string   `json:"type"`
	Flags       []string `json:"flags"`
}

func AchievmentIds(lang string, ids ...int) ([]Achievement, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Achievements() instead?")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "achievements"

	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
		appendix.WriteString("&ids=")
	} else {
		appendix.WriteString("?ids=")
	}

	for i, id := range ids {
		if i > 0 {
			appendix.WriteString(",")
		}
		appendix.WriteString(strconv.Itoa(id))
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []Achievement
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

func Colors(lang string) (res []string, err error) {
	var appendix bytes.Buffer
	ver := "v2"
	tag := "colors"

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

func ColorIds(lang string, ids ...int) ([]Color, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Colors() instead?")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "colors"

	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
		appendix.WriteString("&ids=")
	} else {
		appendix.WriteString("?ids=")
	}

	for i, id := range ids {
		if i > 0 {
			appendix.WriteString(",")
		}
		appendix.WriteString(strconv.Itoa(id))
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []Color
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

func Currencies(lang string) (res []string, err error) {
	var appendix bytes.Buffer
	ver := "v2"
	tag := "currencies"

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

type Currency struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Order       int    `json:"order"`
}

func CurrencyIds(lang string, ids ...int) ([]Currency, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Currencies() instead?")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "colors"

	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
		appendix.WriteString("&ids=")
	} else {
		appendix.WriteString("?ids=")
	}

	for i, id := range ids {
		if i > 0 {
			appendix.WriteString(",")
		}
		appendix.WriteString(strconv.Itoa(id))
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []Currency
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

func Files() (res []string, err error) {
	ver := "v2"
	tag := "files"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	return
}

type File struct {
	ID   string `json:"id"`
	Icon string `json:"icon"`
}

func FileIds(ids ...string) ([]File, error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil. Consider using Currencies() instead?")
	}

	var appendix bytes.Buffer
	ver := "v2"
	tag := "files"

	appendix.WriteString("?ids=")

	for i, id := range ids {
		if i > 0 {
			appendix.WriteString(",")
		}
		appendix.WriteString(id)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []File
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
