package gw2api

import "testing"

func TestSkins(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testSkins []int
	if testSkins, err = api.Skins(); err != nil {
		t.Error("Failed to fetch skins: ", err)
	} else if len(testSkins) < 1 {
		t.Error("Fetched an unlikely number of skins")
	}
	var skins []Skin
	if skins, err = api.SkinIds("en", testSkins[0], testSkins[1], testSkins[2]); err != nil {
		t.Error("Failed to parse the skin data: ", err)
	} else if len(skins) != 3 {
		t.Error("Fetched a non-matching number of skins")
	}
}
