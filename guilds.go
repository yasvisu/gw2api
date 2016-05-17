package gw2api

import (
	"net/url"
	"time"
)

// GuildPermission Name/Desc basically
type GuildPermission struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GuildPermissions Grab them all
func (gw2 *GW2Api) GuildPermissions() (res []string, err error) {
	ver := "v2"
	tag := "guild/permissions"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// GuildPermissionIds Names/Description of permissions
func (gw2 *GW2Api) GuildPermissionIds(ids ...string) (guild []GuildPermission, err error) {
	ver := "v2"
	tag := "guild/permissions"
	params := url.Values{}
	params.Add("ids", commaList(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &guild)
	return
}

// GuildUpgradeCost specific costs of an upgrade. Type like Collection or Item
type GuildUpgradeCost struct {
	Type   string `json:"type"`
	Count  int    `json:"count"`
	Name   string `json:"name"`
	ItemID int    `json:"item_id"`
}

// GuildUpgrade All the Upgrade Information
type GuildUpgrade struct {
	ID            int                `json:"id"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	BuildTime     int                `json:"build_time"`
	Icon          string             `json:"icon"`
	Type          string             `json:"type"`
	RequiredLevel int                `json:"required_level"`
	Experience    int                `json:"experience"`
	Prerequisites []int              `json:"prerequisites"`
	Costs         []GuildUpgradeCost `json:"costs"`
}

// GuildUpgrades Get them all
func (gw2 *GW2Api) GuildUpgrades() (res []int, err error) {
	ver := "v2"
	tag := "guild/upgrades"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// GuildUpgradeIds Translated information
func (gw2 *GW2Api) GuildUpgradeIds(lang string, ids ...int) (guild []GuildUpgrade, err error) {
	ver := "v2"
	tag := "guild/upgrades"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &guild)
	return
}

// GuildRank Granted Permissions and Order to sorty by
type GuildRank struct {
	ID          string   `json:"id"`
	Order       int      `json:"order"`
	Permissions []string `json:"permissions"`
}

// GuildRanks Show all ranks in the guild
func (gw2 *GW2Api) GuildRanks(id string) (guild []GuildRank, err error) {
	ver := "v2"
	tag := "guild/" + id + "/ranks"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermGuilds, nil, &guild)
	return
}

// GuildMember Accountname, current rank and date they joined
type GuildMember struct {
	Name   string    `json:"name"`
	Rank   string    `json:"rank"`
	Joined time.Time `json:"joined"`

	// Only as part of guild team
	Role string `json:"role"`
}

// GuildMembers returns a list of all members
func (gw2 *GW2Api) GuildMembers(id string) (member []GuildMember, err error) {
	ver := "v2"
	tag := "guild/" + id + "/members"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermGuilds, nil, &member)
	return
}

// GuildStashItem represents a single slot in one of the bank tabs
type GuildStashItem struct {
	ID    int `json:"id"`
	Count int `json:"count"`
}

// GuildStash represents a bank tab in the guild bank
type GuildStash struct {
	UpgradeID int              `json:"upgrade_id"`
	Size      int              `json:"size"`
	Coins     int              `json:"coins"`
	Inventory []GuildStashItem `json:"inventory"`
}

// GuildStashes returns a list of all guild bank tabs and their content
func (gw2 *GW2Api) GuildStashes(id string) (stash []GuildStash, err error) {
	ver := "v2"
	tag := "guild/" + id + "/stash"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermGuilds, nil, &stash)
	return
}

// GuildTreasuryUpgrade represents an upgrade which requires the parent ressource
type GuildTreasuryUpgrade struct {
	UpgradeID int `json:"upgrade_id"`
	Count     int `json:"count"`
}

// GuildTreasury represents the need for an item. NeededBy show the upgrades needing it
type GuildTreasury struct {
	ID       int                    `json:"id"`
	Count    int                    `json:"count"`
	NeededBy []GuildTreasuryUpgrade `json:"needed_by"`
}

// GuildTreasuries returns a list of all currently needed items for any guild upgrade
func (gw2 *GW2Api) GuildTreasuries(id string) (treasury []GuildTreasury, err error) {
	ver := "v2"
	tag := "guild/" + id + "/treasury"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermGuilds, nil, &treasury)
	return
}

// GuildEmblems returns two lists for all emblem layers
func (gw2 *GW2Api) GuildEmblems() (foreground []int, background []int, err error) {
	ver := "v2"
	tag := "emblem/foregrounds"
	if err = gw2.fetchEndpoint(ver, tag, nil, &foreground); err != nil {
		return
	}
	tag = "emblem/backgrounds"
	err = gw2.fetchEndpoint(ver, tag, nil, &background)
	return
}

// EmblemLayers for the requested Fore/Background
type EmblemLayers struct {
	ID     int      `json:"id"`
	Layers []string `json:"layers"`
}

// GuildEmblemForegroundIds returns the layers for the requested foregrounds
func (gw2 *GW2Api) GuildEmblemForegroundIds(ids ...int) (layers []EmblemLayers, err error) {
	ver := "v2"
	tag := "emblem/foregrounds"
	params := url.Values{}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &layers)
	return
}

// GuildEmblemBackgroundIds returns the layers for the requested backgrounds
func (gw2 *GW2Api) GuildEmblemBackgroundIds(ids ...int) (layers []EmblemLayers, err error) {
	ver := "v2"
	tag := "emblem/backgrounds"
	params := url.Values{}
	params.Add("ids", commaList(stringSlice(ids)))
	err = gw2.fetchEndpoint(ver, tag, params, &layers)
	return
}

// GuildLogEntry returns information about certain events in a guild's log
type GuildLogEntry struct {
	ID   int       `json:"id"`
	Time time.Time `json:"time"`
	Type string    `json:"type"`
	//type = treasury
	ItemID int `json:"item_id"`
	Count  int `json:"count"`
	//type = motd
	Motd string `json:"motd"`
	//type = (motd|joined|invited)
	User string `json:"user"`
	//type = influence
	Activity          string   `json:"activity"`
	TotalParticipants int      `json:"total_participants"`
	Participants      []string `json:"participants"`
	//type = rank_change
	ChangedBy string `json:"changed_by"`
	OldRank   string `json:"old_rank"`
	NewRank   string `json:"new_rank"`
	//type = invited
	InvitedBy string `json:"invited_by"`
	//type = stash
	Operation string `json:"operation"`
	Coints    int    `json:"coins"`
	//type = upgrade
	UpgradeID int    `json:"upgrade_id"`
	Action    string `json:"action"`
}

// GuildLog returns up to a 100 of each type of log entry
func (gw2 *GW2Api) GuildLog(id string) (log []GuildLogEntry, err error) {
	ver := "v2"
	tag := "guild/" + id + "/log"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermGuilds, nil, &log)
	return
}

// GuildTeamStats per season stats of a guild team
type GuildTeamStats struct {
	SeasonID string `json:"id"`
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Rating   int    `json:"rating"`
}

// GuildTeam a designated team for PvP
type GuildTeam struct {
	ID        int            `json:"id"`
	Members   []GuildMember  `json:"members"`
	Name      string         `json:"name"`
	Aggregate WinLoss        `json:"aggregate"`
	Ladders   LadderStats    `json:"ladders"`
	Games     PvPGameStats   `json:"games"`
	Seasons   GuildTeamStats `json:"seasons"`
}

// GuildPvPTeams returns a list of teams for a guild
func (gw2 *GW2Api) GuildPvPTeams(id string) (teams []GuildTeam, err error) {
	ver := "v2"
	tag := "guild/" + id + "/teams"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermGuilds, nil, &teams)
	return
}
