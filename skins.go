package gw2api

import "net/url"

// Either WeightClass(Armor) or DamageType(Weapon) is set
type SkinDetails struct {
	Type        string `json:"type"`
	WeightClass string `json:"weight_class"`
	DamageType  string `json:"damage_type"`
}

type Skin struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Flags        []string    `json:"flags"`
	Restrictions []string    `json:"restrictions"`
	Icon         string      `json:"icon"`
	Description  string      `json:"description"`
	Details      SkinDetails `json:"details"`
}

// Returns a list of all current skin ids
func (gw2 *GW2Api) Skins() (res []int, err error) {
	ver := "v2"
	tag := "skins"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// Returns a list of skins as requested by the id parameter.
// Special id `all` is not permitted on this endpoint
func (gw2 *GW2Api) SkinIds(lang string, ids ...int) (skins []Skin, err error) {
	ver := "v2"
	tag := "skins"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &skins)
	return
}
