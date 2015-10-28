package gw2api

type Map struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	MinLevel      int       `json:"min_level"`
	MaxLevel      int       `json:"max_level"`
	DefaultFloor  int       `json:"default_floor"`
	Floors        []int     `json:"floors"`
	RegionId      int       `json:"region_id"`
	RegionName    string    `json:"region_name"`
	ContinentId   int       `json:"continent_id"`
	ContinentName string    `json:"continent_name"`
	MapRect       [2][2]int `json:"map_rect"`
	ContinentRect [2][2]int `json:"continent_rect"`
}

// Returns a list of all current skin ids
func Maps() (res []int, err error) {
	ver := "v2"
	tag := "skins"
	err = fetchEndpoint(ver, tag, "", &res)
	return
}

// Returns a list of skins as requested by the id parameter.
// Special id `all` is not permitted on this endpoint
func MapIds(lang string, ids ...int) (maps []Map, err error) {
	ver := "v2"
	tag := "maps"
	err = fetchDetailEndpoint(ver, tag, lang, stringSlice(ids), &maps)
	return
}
