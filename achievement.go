package gw2api

import (
	"fmt"
	"net/url"
	"strconv"
)

// Achievements - List of all achievement ids
func (gw2 *GW2Api) Achievements() (res []int, err error) {
	ver := "v2"
	tag := "achievements"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// DailyAchievementLevel range for the achievement. Determining if it is visible
// to a character or not
type DailyAchievementLevel struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

// DailyAchievement for the gamemode
type DailyAchievement struct {
	ID          int                   `json:"id"`
	Level       DailyAchievementLevel `json:"level"`
	Requirement []string              `json:"required_access"`
}

// DailyAchievements structured by their gamemode
type DailyAchievements struct {
	PvE      []DailyAchievement `json:"pve"`
	PvP      []DailyAchievement `json:"pvp"`
	WvW      []DailyAchievement `json:"wvw"`
	Fractals []DailyAchievement `json:"fractals"`
	Special  []DailyAchievement `json:"special"`
}

// AchievementsDaily returns the daily achievements which can be completed today
func (gw2 *GW2Api) AchievementsDaily() (achievs DailyAchievements, err error) {
	ver := "v2"
	tag := "achievements/daily"
	err = gw2.fetchEndpoint(ver, tag, nil, &achievs)
	return
}

// AchievementsTomorrow returns the daily achievements which can be completed tomorrow
func (gw2 *GW2Api) AchievementsDailyTomorrow() (achievs DailyAchievements, err error) {
	ver := "v2"
	tag := "achievements/daily/tomorrow"
	err = gw2.fetchEndpoint(ver, tag, nil, &achievs)
	return
}

// Achievement detailed achievement information
type Achievement struct {
	ID          int      `json:"id"`
	Icon        string   `json:"icon"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	Type        string   `json:"type"`
	Flags       []string `json:"flags"`
}

// AchievementIds localized achievement details via ids
func (gw2 *GW2Api) AchievementIds(lang string, ids ...int) (achievs []Achievement, err error) {
	ver := "v2"
	tag := "achievements"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &achievs)
	return
}

// AchievementPages returns a paginated list of all achievements. Use it to grab
// all achievements from the API. Return values are localized to lang
func (gw2 *GW2Api) AchievementPages(lang string, page, pageSize int) (achievs []Achievement, err error) {
	if page < 0 {
		return nil, fmt.Errorf("Page parameter cannot be a negative number!")
	}
	ver := "v2"
	tag := "achievements"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("page", strconv.Itoa(page))
	if pageSize >= 0 {
		params.Add("page_size", strconv.Itoa(pageSize))
	}
	err = gw2.fetchEndpoint(ver, tag, params, &achievs)
	return
}

//AchievementGroups list of all achievement groups
func (gw2 *GW2Api) AchievementGroups() (res []string, err error) {
	ver := "v2"
	tag := "achievements/groups"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// AchievementGroup Group description
type AchievementGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	Categories  []int  `json:"categories"`
}

// AchievementGroupIds localized achievement groups by id
func (gw2 *GW2Api) AchievementGroupIds(lang string, ids ...string) (achievs []AchievementGroup, err error) {
	ver := "v2"
	tag := "achievements/groups"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &achievs)
	return
}

//AchievementCategories list of all achievement categories
func (gw2 *GW2Api) AchievementCategories() (res []int, err error) {
	ver := "v2"
	tag := "achievements/categories"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// AchievementCategory Category description
type AchievementCategory struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Order        int    `json:"order"`
	Icon         string `json:"icon"`
	Achievements []int  `json:"achievements"`
	// Either GuildWars2 or HeartOfThorns
	RequiredAccess []string `json:"required_access"`
}

// AchievementCategoryIds localized achievement categories by id
func (gw2 *GW2Api) AchievementCategoryIds(lang string, ids ...int) (achievs []AchievementCategory, err error) {
	ver := "v2"
	tag := "achievements/categories"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &achievs)
	return
}

// AchievementProgress represents detailed information about the accounts
// progress on an achievement.
type AchievementProgress struct {
	ID      int `json:"id"`
	Current int `json:"current"`
	// Max - The amount needed to complete the achievement. Most WvW achievements
	// have this set to -1.
	Max  int  `json:"max"`
	Done bool `json:"done"`
	// Bits - This attribute contains an array of numbers from 0 to 7, giving more
	// specific information on the progress for the achievement. The meaning of
	// each 0-7 value varies with each achievement. (
	Bits []int `json:"bits"`
}

// AccountAchievements returns a list of the accounts progress with all known
// achievements in the game
func (gw2 *GW2Api) AccountAchievements() (achievs []AchievementProgress, err error) {
	ver := "v2"
	tag := "account/achievements"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermProgression, nil, &achievs)
	return
}
