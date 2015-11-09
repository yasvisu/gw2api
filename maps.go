package gw2api

import (
	"fmt"
	"net/url"
)

type Map struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	MinLevel      int       `json:"min_level"`
	MaxLevel      int       `json:"max_level"`
	DefaultFloor  int       `json:"default_floor"`
	Floors        []int     `json:"floors"`
	RegionID      int       `json:"region_id"`
	RegionName    string    `json:"region_name"`
	ContinentID   int       `json:"continent_id"`
	ContinentName string    `json:"continent_name"`
	MapRect       [2][2]int `json:"map_rect"`
	ContinentRect [2][2]int `json:"continent_rect"`
}

// Returns a list of all current skin ids
func (gw2 *GW2Api) Maps() (res []int, err error) {
	ver := "v2"
	tag := "skins"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// Returns a list of skins as requested by the id parameter.
// Special id `all` is not permitted on this endpoint
func (gw2 *GW2Api) MapIds(lang string, ids ...int) (maps []Map, err error) {
	ver := "v2"
	tag := "maps"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &maps)
	return
}

func (gw2 *GW2Api) Continents() (res []string, err error) {
	ver := "v2"
	tag := "continents"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

type Continent struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ContinentDims []int  `json:"continent_dims"`
	MinZoom       int    `json:"min_zoom"`
	MaxZoom       int    `json:"max_zoom"`
	Floors        []int  `json:"floors"`
}

func (gw2 *GW2Api) ContinentIds(lang string) (conts []Continent, err error) {
	ver := "v2"
	tag := "continents"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	err = gw2.fetchEndpoint(ver, tag, params, &conts)
	return
}

func (gw2 *GW2Api) ContinentFloors(continent int) (floors []int, err error) {
	ver := "v2"
	tag := fmt.Sprintf("continents/%d/floors", continent)
	err = gw2.fetchEndpoint(ver, tag, nil, &floors)
	return
}

func (gw2 *GW2Api) ContinentFloorRegions(continent, floor int) (regions []int, err error) {
	ver := "v2"
	tag := fmt.Sprintf("continents/%d/floors/%d/regions", continent, floor)
	err = gw2.fetchEndpoint(ver, tag, nil, &regions)
	return
}

func (gw2 *GW2Api) ContinentFloorRegionMaps(continent, floor, region int) (maps []int, err error) {
	ver := "v2"
	tag := fmt.Sprintf("continents/%d/floors/%d/regions/%d/maps", continent, floor, maps)
	err = gw2.fetchEndpoint(ver, tag, nil, &maps)
	return
}

func (gw2 *GW2Api) ContinentFloorRegionMapSectors(continent, floor, region, mapID int) (sectors []int, err error) {
	ver := "v2"
	tag := fmt.Sprintf("continents/%d/floors/%d/regions/%d/maps/%d/sectors", continent, floor, region, mapID)
	err = gw2.fetchEndpoint(ver, tag, nil, &sectors)
	return
}

func (gw2 *GW2Api) ContinentFloorRegionMapPois(continent, floor, region, mapID int) (pois []int, err error) {
	ver := "v2"
	tag := fmt.Sprintf("continents/%d/floors/%d/regions/%d/maps/%d/pois", continent, floor, region, mapID)
	err = gw2.fetchEndpoint(ver, tag, nil, &pois)
	return
}

func (gw2 *GW2Api) ContinentFloorRegionMapTasks(continent, floor, region, mapID int) (tasks []int, err error) {
	ver := "v2"
	tag := fmt.Sprintf("continents/%d/floors/%d/regions/%d/maps/%d/tasks", continent, floor, region, mapID)
	err = gw2.fetchEndpoint(ver, tag, nil, &tasks)
	return
}
