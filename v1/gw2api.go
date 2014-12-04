/*
The MIT License (MIT)

Copyright (c) 2014 Yasen Visulchev

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

/* 
Written by Yasen Visulchev - Sofia University FMI/Golang
	20141130
	
GW2 API Wrapper for Go

For GW2 API Version 1
*/
package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

/* GW2API UNEXPORTED FUNCTIONS */

//Fetcher function to return only the body of a HTTP request.
//Parameters: version, tag, appendix strings.
func fetchJSON(ver string, tag string, appendix string) ([]byte, error) {
	resp, err := http.Get("https://api.guildwars2.com/" + ver + "/" + tag + ".json" + appendix)
	if err != nil {
		return nil, err
	}
	var jsonBlob []byte
	jsonBlob, err = ioutil.ReadAll(resp.Body)
	return jsonBlob, err
}

/* DYNAMIC EVENTS */

//Obsolete type since Events(...) is obsolete.
type Event struct {
	WorldID int    `json:"world_id"`
	MapID   int    `json:"map_id"`
	EventID string `json:"event_id"`
	State   string `json:"state"`
}

//Obsolete endpoint since the introduction of the Megaserver technology.
func Events(WorldID int, MapID int, EventID string) ([]Event, error) {
	return nil, errors.New("Obsolete endpoint. Contact the author if this is wrong.")
}

//Parent type for all -Names endpoints.
type Name struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Parent function for all -Names functions.
func names(ver string, tag string, lang string) ([]Name, error) {
	var appendix bytes.Buffer
	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res []Name
	err = json.Unmarshal(data, &res)
	return res, err
}

//Returns a list of localized event names.
func EventNames(lang string) ([]Name, error) {
	ver := "v1"
	tag := "event_names"
	res, err := names(ver, tag, lang)
	return res, err
}

//Returns a list of localized map names.
func MapNames(lang string) ([]Name, error) {
	ver := "v1"
	tag := "map_names"
	res, err := names(ver, tag, lang)
	return res, err
}

//Returns a list of localized world names.
func WorldNames(lang string) ([]Name, error) {
	ver := "v1"
	tag := "world_names"
	res, err := names(ver, tag, lang)
	return res, err
}

//EventDetail type for events - assumption is that number values are stored as integers (not true for all other endpoints).
type EventDetail struct {
	Name     string        `json:"name"`
	Level    int           `json:"level"`
	MapID    int           `json:"map_id"`
	Flags    []string      `json:"flags"`
	Location EventLocation `json:"location"`
}

//EventLocation type. Contains optional fields, which (for brevity) have been aggregated into one type. JSON tag set to omit empty fields.
type EventLocation struct {
	Type     string      `json:"type"`
	Center   []float64   `json:"center"`
	Height   float64     `json:"height,omitempty"`
	Radius   float64     `json:"radius,omitempty"`
	Rotation float64     `json:"rotation,omitempty"`
	ZRange   []float64   `json:"z_range,omitempty"`
	Points   [][]float64 `json:"points,omitempty"`
}

//Returns detailed information about events.
//Both parameters are optional.
func EventDetails(EventID string, lang string) (map[string]Event, error) {
	type events struct {
		Events map[string]Event `json:"events"`
	}

	var appendix bytes.Buffer
	ver := "v1"
	tag := "event_details"
	concatenator := ""
	if EventID != "" || lang != "" {
		appendix.WriteString("?")
	}
	if EventID != "" {
		appendix.WriteString("event_id=")
		appendix.WriteString(EventID)
		concatenator = "&"
	}
	if lang != "" {
		appendix.WriteString(concatenator)
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res events
	err = json.Unmarshal(data, &res)
	return res.Events, err
}

/* GUILDS */

//Guild type for Guilds. The ID is alpha-numeric.
type Guild struct {
	GuildID   string      `json:"guild_id"`
	GuildName string      `json:"guild_name"`
	Tag       string      `json:"tag"`
	Emblem    GuildEmblem `json:"emblem"`
}

//GuildEmblem type for Guild.Emblem - assumption is that numbers are stored as integers.
type GuildEmblem struct {
	BackgroundID               int      `json:"background_id"`
	ForegroundID               int      `json:"foreground_id"`
	Flags                      []string `json:"flags"`
	BackgroundColorID          int      `json:"background_color_id"`
	ForegroundPrimaryColorID   int      `json:"foreground_primary_color_id"`
	ForegroundSecondaryColorID int      `json:"foreground_secondary_color_id"`
}

//Returns detailed information about a guild.
func GuildDetails(GuildID string, GuildName string) (Guild, error) {
	if GuildID == "" && GuildName == "" {
		var t Guild
		return t, errors.New("Required parameter empty.")
	}

	var appendix bytes.Buffer
	ver := "v1"
	tag := "guild_details"
	concatenator := ""

	appendix.WriteString("?")

	if GuildID != "" {
		appendix.WriteString("guild_id=")
		appendix.WriteString(GuildID)
		concatenator = "&"
	}
	if GuildName != "" {
		appendix.WriteString(concatenator)
		appendix.WriteString("guild_name=")
		appendix.WriteString(GuildName)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		var t Guild

		return t, err
	}

	var res Guild
	err = json.Unmarshal(data, &res)
	return res, err
}

/* ITEMS */

//Returns a list of discovered items.
func Items(lang string) ([]int, error) {
	type itemArray struct {
		Items []int `json:"items"`
	}
	var appendix bytes.Buffer
	ver := "v1"
	tag := "items"
	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res itemArray
	err = json.Unmarshal(data, &res)
	return res.Items, err
}

//Common interface for all Item types.
type ItemTyper interface {
	ItemType() reflect.Type
}

//General Item type all items inherit from.
type Item struct {
	ItemID       string   `json:"item_id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	Level        string   `json:"level"`
	Rarity       string   `json:"rarity"`
	VendorValue  string   `json:"vendor_value"`
	IconFileID   string   `json:"icon_file_id"`
	DefaultSkin  string   `json:"default_skin"`
	GameTypes    []string `json:"game_types"`
	Flags        []string `json:"flags"`
	Restrictions []string `json:"restrictions"`
}

func (u Item) ItemType() reflect.Type {
	return reflect.TypeOf(u)
}

//ArmorItem to get armor-specific fields.
type ArmorItem struct {
	Item
	Armor ArmorItemDetails `json:"armor"`
}

type flags struct {
	Flags []string `json:"flags"`
}

//Armor item details field.
type ArmorItemDetails struct {
	Type                  string           `json:"type"`
	WeightClass           string           `json:"weight_class"`
	Defense               string           `json:"defense"`
	InfusionSlots         []flags          `json:"infusion_slots"`
	InfixUpgrade          ItemInfixUpgrade `json:"infix_upgrade"`
	SuffixItemID          string           `json:"suffix_item_id"`
	SecondarySuffixItemID string           `json:"secondary_suffix_item_id"`
}

//Common structure for all items with InfixUpgrade.
type ItemInfixUpgrade struct {
	Buff       ItemInfixUpgradeBuff       `json:"buff"`
	Attributes ItemInfixUpgradeAttributes `json:"attributes"`
}

//Common structure for all InfixUpgrades with InfixUpgradeBuff.
type ItemInfixUpgradeBuff struct {
	SkillID     string `json:"skill_id"`
	Description string `json:"description"`
}

//Common structure for all InfixUpgrades with InfixUpgradeAttributes.
type ItemInfixUpgradeAttributes struct {
	Attribute string `json:"attribute"`
	Modifier  string `json:"modifier"`
}

//BackItem to get back-specific fields.
type BackItem struct {
	Item
	Back BackItemDetails `json:"back"`
}

//Back item details field.
type BackItemDetails struct {
	InfusionSlots         []ItemInfusionSlots `json:"infusion_slots"`
	InfixUpgrade          ItemInfixUpgrade    `json:"infix_upgrade"`
	SuffixItemID          string              `json:"suffix_item_id"`
	SecondarySuffixItemID string              `json:"secondary_suffix_item_id"`
}

//Different format from Armor infusion slots.
type ItemInfusionSlots struct {
	ItemID string   `json:"item_id"`
	Flags  []string `json:"flags"`
}

//BagItem to get bag-specific fields.
type BagItem struct {
	Item
	Bag BagItemDetails `json:"bag"`
}

//Bag item details field.
type BagItemDetails struct {
	NoSellOrSort string `json:"no_sell_or_sort"`
	Size         string `json:"size"`
}

//ConsumableItem to get consumable-specific fields.
type ConsumableItem struct {
	Item
	Consumable ConsumableItemDetails `json:"consumable"`
}

//Consumable item details field.
type ConsumableItemDetails struct {
	Type        string `json:"type"`
	DurationMS  string `json:"duration_ms,omitempty"`
	Description string `json:"description,omitempty"`
	UnlockType  string `json:"unlock_type,omitempty"`
	RecipeID    string `json:"recipe_id,omitempty"`
	ColorID     string `json:"color_id,omitempty"`
}

//ContainerItem to get container-specific fields.
type ContainerItem struct {
	Item
	Container ContainerItemDetails `json:"container"`
}

//Container item details field.
type ContainerItemDetails struct {
	Type string `json:"type"`
}

//ContainerItem to get CraftingMaterial-specific fields. None so far.
type CraftingMaterialItem struct {
	Item
}

//GatheringItem to get gathering-specific fields.
type GatheringItem struct {
	Item
	Gathering GatheringItemDetails `json:"gathering"`
}

//Gathering item details field.
type GatheringItemDetails struct {
	Type string `json:"type"`
}

//GizmoItem to get gizmo-specific fields.
type GizmoItem struct {
	Item
	Gizmo GizmoItemDetails `json:"gizmo"`
}

//Gizmo item details field.
type GizmoItemDetails struct {
	Type string `json:"type"`
}

//MiniPetItem to get minipet-specific fields. None so far.
type MiniPetItem struct {
	Item
}

//ToolItem to get tool-specific fields.
type ToolItem struct {
	Item
	Tool ToolItemDetails `json:"tool"`
}

//Tool item details field.
type ToolItemDetails struct {
	Type    string `json:"type"`
	Charges string `json:"charges"`
}

//TrinketItem to get trinket-specific fields.
type TrinketItem struct {
	Item
	Trinket TrinketItemDetails `json:"trinket"`
}

//Trinket item details field.
type TrinketItemDetails struct {
	Type          string              `json:"type"`
	InfusionSlots []ItemInfusionSlots `json:"infusion_slots"`
	InfixUpgrade  ItemInfixUpgrade    `json:"infix_upgrade"`
}

//TrophyItem to get trophy-specific fields. None so far.
type TrophyItem struct {
	Item
}

//UpgradeComponentItem to get UpgradeComponent-specific fields.
type UpgradeComponentItem struct {
	Item
	UpgradeComponent UpgradeComponentItemDetails `json:"upgrade_component"`
}

//Upgrade component item details field.
type UpgradeComponentItemDetails struct {
	Type                 string           `json:"type"`
	Flags                []string         `json:"flags"`
	InfusionUpgradeFlags []string         `json:"infusion_upgrade_flags"`
	Bonuses              []string         `json:"bonuses"`
	InfixUpgrade         ItemInfixUpgrade `json:"infix_upgrade"`
	Suffix               string           `json:"suffix"`
}

//WeaponItem to get weapon-specific fields.
type WeaponItem struct {
	Item
	Weapon WeaponItemDetails `json:"weapon"`
}

//Weapon item details field.
type WeaponItemDetails struct {
	Type                  string            `json:"type"`
	DamageType            string            `json:"damage_type"`
	MinPower              string            `json:"min_power"`
	MaxPower              string            `json:"max_power"`
	Defense               string            `json:"defense"`
	InfusionSlots         ItemInfusionSlots `json:"infusion_slots"`
	InfixUpgrade          ItemInfixUpgrade  `json:"infix_upgrade"`
	SuffixItemID          string            `json:"suffix_item_id"`
	SecondarySuffixItemID string            `json:"secondary_suffix_item_id"`
}

//Returns detailed information about an item.
func ItemDetails(ItemID string, lang string) (ItemTyper, error) {
	if ItemID == "" {
		var t Item
		return t, errors.New("Required parameter empty.")
	}

	var appendix bytes.Buffer
	ver := "v1"
	tag := "item_details"
	appendix.WriteString("?item_id=")
	appendix.WriteString(ItemID)

	if lang != "" {
		appendix.WriteString("&lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		var t Item
		return t, err
	}
	pattern := regexp.MustCompile(`"(armor|ba(g|ck)|c(on(sumable|tainer)|rafting_material)|g(athering|izmo)|minipet|t(ool|r(inket|ophy))|upgrade_component|weapon)"`)
	resultString := pattern.FindString(string(data))
	switch resultString {
	case "\"armor\"":
		var res ArmorItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"back\"":
		var res BackItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"bag\"":
		var res BagItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"consumable\"":
		var res ConsumableItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"container\"":
		var res ContainerItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"crafting_material\"":
		var res CraftingMaterialItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"gathering\"":
		var res GatheringItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"gizmo\"":
		var res GizmoItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"minipet\"":
		var res MiniPetItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"tool\"":
		var res ToolItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"trinket\"":
		var res TrinketItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"trophy\"":
		var res TrophyItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"upgrade_component\"":
		var res UpgradeComponentItem
		err = json.Unmarshal(data, &res)
		return res, err
	case "\"weapon\"":
		var res WeaponItem
		err = json.Unmarshal(data, &res)
		return res, err
	default:
		var res Item
		err = json.Unmarshal(data, &res)
		return res, err
	}
	var res Item
	err = json.Unmarshal(data, &res)
	return res, err
}

//Returns a list of discovered recipes.
func Recipes(lang string) ([]int, error) {
	type recipeArray struct {
		recipes []int `json:"recipes"`
	}
	var appendix bytes.Buffer
	ver := "v1"
	tag := "recipes"
	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res recipeArray
	err = json.Unmarshal(data, &res)
	return res.recipes, err
}

//Recipe type for RecipeDetails.
type Recipe struct {
	RecipeID        string             `json:"recipe_id"`
	Type            string             `json:"type"`
	OutputItem      string             `json:"output_item"`
	OutputItemCount string             `json:"output_item_count"`
	MinRating       string             `json:"min_rating"`
	TimeToCraftMS   string             `json:"time_to_craft_ms"`
	Disciplines     []string           `json:"disciplines"`
	Flags           []string           `json:"flags"`
	Ingredients     []RecipeIngredient `json:"ingredients"`
}

//RecipeIngredient for Recipe.Ingredients.
type RecipeIngredient struct {
	ItemID string `json:"item_id"`
	Count  string `json:"count"`
}

//Returns detailed information about a recipe.
func RecipeDetails(RecipeID string, lang string) (Recipe, error) {
	if RecipeID == "" {
		var t Recipe
		return t, errors.New("Required parameter empty.")
	}

	var appendix bytes.Buffer
	ver := "v1"
	tag := "recipe_details"
	appendix.WriteString("?recipe_id=")
	appendix.WriteString(RecipeID)

	if lang != "" {
		appendix.WriteString("&lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		var t Recipe
		return t, err
	}

	var res Recipe
	err = json.Unmarshal(data, &res)
	return res, err
}

//Returns a list of skins.
func Skins() ([]int, error) {
	type skinArray struct {
		Skins []int `json:"skins"`
	}
	ver := "v1"
	tag := "skins"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return nil, err
	}

	var res skinArray
	err = json.Unmarshal(data, &res)
	return res.Skins, err
}

//Common interface for all Skin types.
type SkinTyper interface {
	SkinType() reflect.Type
}

//General Skin type all skins inherit from.
type Skin struct {
	SkinID            string   `json:"skin_id"`
	Name              string   `json:"name"`
	Type              string   `json:"type"`
	Flags             []string `json:"flags"`
	Restrictions      []string `json:"restrictions"`
	IconFileID        string   `json:"icon_file_id"`
	IconFileSignature string   `json:"icon_file_signature"`
	Description       string   `json:"description"`
}

func (u Skin) SkinType() reflect.Type {
	return reflect.TypeOf(u)
}

//ArmorSkin to get armor-specific fields.
type ArmorSkin struct {
	Skin
	Armor ArmorSkinDetails `json:"armor"`
}

//Armor skin details field.
type ArmorSkinDetails struct {
	Type        string `json:"type"`
	WeightClass string `json:"weight_class"`
}

//WeaponSkin to get weapon-specific fields.
type WeaponSkin struct {
	Skin
	Weapon WeaponSkinDetails `json:"weapon"`
}

//Weapon skin details field.
type WeaponSkinDetails struct {
	Type       string `json:"type"`
	DamageType string `json:"damage_type"`
}

//Returns detailed information about a skin.
func SkinDetails(SkinID string, lang string) (SkinTyper, error) {
	if SkinID == "" {
		var t Skin
		return t, errors.New("Required parameter empty.")
	}

	var appendix bytes.Buffer
	ver := "v1"
	tag := "skin_details"
	appendix.WriteString("?skin_id=")
	appendix.WriteString(SkinID)

	if lang != "" {
		appendix.WriteString("&lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		var t Skin
		return t, err
	}

	if strings.Contains(string(data), "\"weapon\"") {
		var res WeaponSkin
		err = json.Unmarshal(data, &res)
		return res, err
	}
	var res ArmorSkin
	err = json.Unmarshal(data, &res)
	return res, err
}

/* MAP INFORMATION */

//Continent type - API fault in overlapping fields later on.
type Continent struct {
	Name          string    `json:"name"`
	ContinentDims []float64 `json:"continent_dims"`
	MinZoom       int       `json:"min_zoom"`
	MaxZoom       int       `json:"max_zoom"`
	Floors        []int     `json:"floors"`
}

//Returns a list of continents and information about each continent.
func Continents(lang string) (map[string]Continent, error) {
	type continents struct {
		Continents map[string]Continent `json:"continents"`
	}
	var appendix bytes.Buffer
	ver := "v1"
	tag := "continents"
	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res continents
	err = json.Unmarshal(data, &res)
	return res.Continents, err
}

//Map type - overlapping fields set to be omitted. Inherited below for brevity.
type Map struct {
	MapName       string  `json:"map_name,omitempty"`
	MinLevel      int     `json:"min_level"`
	MaxLevel      int     `json:"max_level"`
	DefaultFloor  int     `json:"default_floor"`
	Floors        []int   `json:"floors,omitempty"`
	RegionID      int     `json:"region_id,omitempty"`
	RegionName    string  `json:"region_name,omitempty"`
	ContinentID   int     `json:"continent_id,omitempty"`
	ContinentName string  `json:"continent_name"`
	MapRect       [][]int `json:"map_rect"`
	ContinentRect [][]int `json:"continent_rect"`
}

//Returns a list of maps in the game.
func Maps(MapID string, lang string) (map[string]Map, error) {
	type maps struct {
		Maps map[string]Map `json:"maps"`
	}
	var appendix bytes.Buffer
	ver := "v1"
	tag := "maps"
	concatenator := ""
	if MapID != "" || lang != "" {
		appendix.WriteString("?")
	}
	if MapID != "" {
		appendix.WriteString("map_id=")
		appendix.WriteString(MapID)
		concatenator = "&"
	}
	if lang != "" {
		appendix.WriteString(concatenator)
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res maps
	err = json.Unmarshal(data, &res)
	return res.Maps, err
}

//MapFloor type, containing sub-types - with redundant structures, so inherited parts from Map for brevity.
type MapFloorDetails struct {
	TextureDims []int             `json:"texture_dims"`
	ClampedView []int             `json:"clamped_view,omitempty"`
	Regions     map[string]Region `json:"regions"`
}

//Region type for MapFloorDetails.Regions.
type Region struct {
	Name       string               `json:"name"`
	LabelCoord []float64            `json:"label_coord"`
	Maps       map[string]RegionMap `json:"maps"`
}

//RegionMap type for Region.Maps.
type RegionMap struct {
	Map
	Name             string           `json:"name"`
	PointsOfInterest []POI            `json:"points_of_interest"`
	Tasks            []Task           `json:"task"`
	SkillChallenges  []SkillChallenge `json"skill_challenges"`
	Sectors          []Sector         `json:"sectors"`
}

//POI type for RegionMap.PointsOfInterest.
type POI struct {
	POIID int       `json:"poi_id"`
	Name  string    `json:"name"`
	Type  string    `json:"type"`
	Floor int       `json:"floor"`
	Coord []float64 `json:"coord"`
}

//Task type for RegionMap.Tasks.
type Task struct {
	TaskID    int       `json:"task_id"`
	Objective string    `json:"objective"`
	Level     int       `json:"level"`
	Coord     []float64 `json:"coord"`
}

//SkillChallenge type for RegionMap.SkillChallenges.
type SkillChallenge struct {
	Coord []float64 `json:"coord"`
}

//Sector type for RegionMap.Sectors.
type Sector struct {
	SectorID int       `json:"sector_id"`
	Name     string    `json:"name"`
	Level    int       `json:"level"`
	Coord    []float64 `json:"coord"`
}

//Returns detailed information about a map floor
func MapFloor(ContinentID string, Floor string, lang string) (MapFloorDetails, error) {
	if ContinentID == "" || Floor == "" {
		var t MapFloorDetails
		return t, errors.New("Required parameter empty.")
	}

	var appendix bytes.Buffer
	ver := "v1"
	tag := "map_floor"
	concatenator := "&"
	appendix.WriteString("?continent_id=")
	appendix.WriteString(ContinentID)
	appendix.WriteString(concatenator)
	appendix.WriteString("floor=")
	appendix.WriteString(Floor)

	if lang != "" {
		appendix.WriteString(concatenator)
		appendix.WriteString("lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		var t MapFloorDetails
		return t, err
	}

	var res MapFloorDetails
	err = json.Unmarshal(data, &res)
	return res, err
}

/* WORLD VS. WORLD */

//WVWMatchEntry type for WVWMatches.
type WVWMatchEntry struct {
	WVWMatchID   string `json:"wvw_match_id"`
	RedWorldID   int    `json:"red_world_id"`
	BlueWorldID  int    `json:"blue_world_id"`
	GreenWorldID int    `json:"green_world_id"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
}

//Returns the currently running WvW matches.
func WVWMatches() ([]WVWMatchEntry, error) {
	type wvwMatches struct {
		WVWMatches []WVWMatchEntry `json:"wvw_matches"`
	}
	ver := "v1"
	tag := "wvw/matches"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return nil, err
	}

	var res wvwMatches
	err = json.Unmarshal(data, &res)
	return res.WVWMatches, err
}

//WVWMatch type for WVWMatchDetails.
type WVWMatch struct {
	MatchID string   `json:"match_id"`
	Scores  []int    `json:"scores"`
	Maps    []WVWMap `json:"maps"`
}

//WVWMap type for WVWMatch.Maps.
type WVWMap struct {
	Type       string         `json:"type"`
	Scores     []int          `json:"scores"`
	Objectives []WVWObjective `json:"objectives"`
	Bonuses    []WVWBonus     `json:"bonuses"`
}

//WVWObjective type for WVWMap.Objectives.
type WVWObjective struct {
	ID         int    `json:"id"`
	Owner      string `json:"owner"`
	OwnerGuild string `json:"owner_guild"`
}

//WVWBonus type for WVWMap.Bonuses.
type WVWBonus struct {
	Type  string `json:"type"`
	Owner string `json:"owner"`
}

//Returns details about a WvW match.
func WVWMatchDetails(MatchID string) (WVWMatch, error) {
	if MatchID == "" {
		var t WVWMatch
		return t, errors.New("Required parameter empty.")
	}

	var appendix bytes.Buffer
	ver := "v1"
	tag := "wvw/match_details"
	appendix.WriteString("?match_id=")
	appendix.WriteString(MatchID)

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		var t WVWMatch
		return t, err
	}

	var res WVWMatch
	err = json.Unmarshal(data, &res)
	return res, err
}

//Returns a list of WvW objective names. Uses common Name type.
func WVWObjectiveNames(lang string) ([]Name, error) {
	ver := "v1"
	tag := "wvw/objective_names"
	res, err := names(ver, tag, lang)
	return res, err
}

/* MISCELLANEOUS */

//Returns the current build id.
func Build() (int, error) {
	type build struct {
		buildID int `json:"build_id"`
	}
	data, err := fetchJSON("v1", "build", "")
	if err != nil {
		return -1, err
	}

	var res build
	err = json.Unmarshal(data, &res)
	return res.buildID, err
}

//Color type for Colors.
type Color struct {
	Name    string        `json:"name"`
	BaseRGB []byte        `json:"base_rgb"`
	Cloth   ColorMaterial `json:"cloth"`
	Leather ColorMaterial `json:"leather"`
	Metal   ColorMaterial `json:"metal"`
}

//Commonly structured ColorMaterial type for Color.Cloth, Color.Leather, Color.Metal.
type ColorMaterial struct {
	Brightness int     `json:"brightness"`
	Contrast   float64 `json:"contrast"`
	Hue        float64 `json:"hue"`
	Saturation float64 `json:"saturation"`
	Lightness  float64 `json:"lightness"`
	RGB        []byte  `json:"rgb"`
}

//Returns a list of dyes in the game.
func Colors(lang string) (map[string]Color, error) {
	type colors struct {
		Colors map[string]Color `json:"colors"`
	}
	var appendix bytes.Buffer
	ver := "v1"
	tag := "colors"
	if lang != "" {
		appendix.WriteString("?lang=")
		appendix.WriteString(lang)
	}

	data, err := fetchJSON(ver, tag, appendix.String())
	if err != nil {
		return nil, err
	}

	var res colors
	err = json.Unmarshal(data, &res)
	return res.Colors, err
}

//File type for Files.
type File struct {
	FileID    int    `json:"file_id"`
	Signature string `json:"signature"`
}

//Returns commonly requested assets.
func Files() (map[string]File, error) {
	ver := "v1"
	tag := "files"

	data, err := fetchJSON(ver, tag, "")
	if err != nil {
		return nil, err
	}

	var res map[string]File
	err = json.Unmarshal(data, &res)
	return res, err
}
