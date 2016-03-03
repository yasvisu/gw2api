package gw2api

import (
	"os"
	"testing"
)

func TestAchievements(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testAchievements []int
	if testAchievements, err = api.Achievements(); err != nil {
		t.Error("Failed to fetch achievements")
	}

	var testDailyAchievments DailyAchievements
	if testDailyAchievments, err = api.AchievementsDaily(); err != nil && len(testDailyAchievments.PvE) < 1 {
		t.Error("Failed to parse daily achievements: ", err)
	}

	var achievements []Achievement
	if achievements, err = api.AchievementIds("en", testAchievements[0:2]...); err != nil {
		t.Error("Failed to parse the achievement data: ", err)
	} else if len(achievements) != 2 {
		t.Error("Failed to fetch existing achievements")
	}

	if _, err = api.AchievementPages("en", -1, 0); err == nil {
		t.Error("Failed to fetch error with wrong arguments")
	}

	if achievements, err = api.AchievementPages("en", 0, 20); err != nil {
		t.Error("Failed to parse the achievement data: ", err)
	} else if len(achievements) != 20 {
		t.Error("Failed to fetch correct amount of achievements")
	}

	var testGroups []string
	if testGroups, err = api.AchievementGroups(); err != nil {
		t.Error("Failed to fetch achievement groups")
	}
	var achievementGroups []AchievementGroup
	if achievementGroups, err = api.AchievementGroupIds("en", testGroups[0:2]...); err != nil {
		t.Error("Failed to parse the achievement group data: ", err)
	} else if len(achievementGroups) != 2 {
		t.Error("Failed to fetch existing achievement groups")
	}

	if testAchievements, err = api.AchievementCategories(); err != nil {
		t.Error("Failed to fetch achievement categories")
	}
	var achievementCats []AchievementCategory
	if achievementCats, err = api.AchievementCategoryIds("en", testAchievements[0:2]...); err != nil {
		t.Error("Failed to parse the achievement category data: ", err)
	} else if len(achievementCats) != 2 {
		t.Error("Failed to fetch existing achievement categories")
	}
}

func TestAccountAchievements(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermProgression) {
		t.Skip("API-Key does not have required permission for the test")
	}
	var achievements []AchievementProgress
	if achievements, err = api.AccountAchievements(); err != nil {
		t.Error("Failed to parse the achievement unlock data: ", err)
	} else if len(achievements) < 1 {
		t.Error("Fetched an unlikely number of achievement unlocks")
	}
}
