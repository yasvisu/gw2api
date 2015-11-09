package gw2api

import "testing"

func TestQuaggans(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testQuaggans []string
	if testQuaggans, err = api.Quaggans(); err != nil {
		t.Error("Failed to fetch quaggans")
	}

	var quaggans []Quaggan
	if quaggans, err = api.QuagganIds(testQuaggans[0], testQuaggans[1]); err != nil {
		t.Error("Failed to parse the quaggan data: ", err)
	} else if len(quaggans) != 2 {
		t.Error("Failed to fetch existing quaggans")
	}
}

func TestWorlds(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testWorlds []int
	if testWorlds, err = api.Worlds(); err != nil {
		t.Error("Failed to fetch worlds")
	}

	var worlds []World
	if worlds, err = api.WorldIds("en", testWorlds[0], testWorlds[1]); err != nil {
		t.Error("Failed to parse the world data: ", err)
	} else if len(worlds) != 2 {
		t.Error("Failed to fetch existing worlds")
	}
}

func TestAchievements(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testAchievements []int
	if testAchievements, err = api.Achievements(); err != nil {
		t.Error("Failed to fetch achievements")
	}

	var achievements []Achievement
	if achievements, err = api.AchievementIds("en", testAchievements[0], testAchievements[1]); err != nil {
		t.Error("Failed to parse the achievement data: ", err)
	} else if len(achievements) != 2 {
		t.Error("Failed to fetch existing achievements")
	}
}

func TestColors(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testColors []int
	if testColors, err = api.Colors(); err != nil {
		t.Error("Failed to fetch colors")
	}

	var colors []Color
	if colors, err = api.ColorIds("en", testColors[0], testColors[1]); err != nil {
		t.Error("Failed to parse the color data: ", err)
	} else if len(colors) != 2 {
		t.Error("Failed to fetch existing colors")
	}
}

func TestCurrencies(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testCurrencies []int
	if testCurrencies, err = api.Currencies(); err != nil {
		t.Error("Failed to fetch currencies")
	}

	var currencies []Currency
	if currencies, err = api.CurrencyIds("en", testCurrencies[0], testCurrencies[1]); err != nil {
		t.Error("Failed to parse the currency data: ", err)
	} else if len(currencies) != 2 {
		t.Error("Failed to fetch existing currencies")
	}
}

func TestFiles(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testFiles []string
	if testFiles, err = api.Files(); err != nil {
		t.Error("Failed to fetch files")
	}

	var files []File
	if files, err = api.FileIds("en", testFiles[0], testFiles[1]); err != nil {
		t.Error("Failed to parse the file data: ", err)
	} else if len(files) != 2 {
		t.Error("Failed to fetch existing files")
	}
}
