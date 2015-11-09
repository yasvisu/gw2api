package gw2api

import "testing"

func TestSkins(t *testing.T) {
	var skins []int
	var err error
	api := NewGW2Api()
	if skins, err = api.Skins(); err != nil {
		t.Error("Failed to fetch skins: ", err)
	} else if len(skins) < 1 {
		t.Error("Fetched an unlikely number of skins")
	}
}

func TestSkinIds(t *testing.T) {
	var err error
	var skins []Skin
	api := NewGW2Api()
	if skins, err = api.SkinIds("en", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10); err != nil {
		t.Error("Failed to parse the skin data: ", err)
	} else if len(skins) != 10 {
		t.Error("Fetched a non-matching number of skins")
	}
}
