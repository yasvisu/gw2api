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
}

// GuildMembers returns a list of all members
func (gw2 *GW2Api) GuildMembers(id string) (member []GuildMember, err error) {
	ver := "v2"
	tag := "guild/" + id + "/members"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermGuilds, nil, &member)
	return
}
