package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

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
func Skins() (res []int, err error) {
	ver := "v2"
	tag := "skins"

	var data []byte
	if data, err = fetchJSON(ver, tag, ""); err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	return
}

// Returns a list of skins as requested by the id parameter.
// Special id `all` is not permitted on this endpoint
func SkinIds(lang string, ids ...int) (skins []Skin, err error) {
	if ids == nil {
		return nil, errors.New("Required ids parameters nil.")
	}
	var appendix bytes.Buffer
	ver := "v2"
	tag := "skins"

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

	var data []byte
	if data, err = fetchJSON(ver, tag, appendix.String()); err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &skins)
	return

}
