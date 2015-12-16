package gw2api

import (
	"os"
	"testing"
)

func TestGuilds(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testPermissions []string
	if testPermissions, err = api.GuildPermissions(); err != nil {
		t.Error("Failed to fetch permissions")
	}

	var permissions []GuildPermission
	if permissions, err = api.GuildPermissionIds(testPermissions[0:2]...); err != nil {
		t.Error("Failed to parse the permission data: ", err)
	} else if len(permissions) != 2 {
		t.Error("Failed to fetch existing permissions")
	}

	var testUpgrades []int
	if testUpgrades, err = api.GuildUpgrades(); err != nil {
		t.Error("Failed to fetch upgrades")
	}

	var upgrades []GuildUpgrade
	if upgrades, err = api.GuildUpgradeIds("en", testUpgrades[0:2]...); err != nil {
		t.Error("Failed to parse the upgrade data: ", err)
	} else if len(upgrades) != 2 {
		t.Error("Failed to fetch existing upgrades")
	}
}

func TestAuthenticatedGuilds(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermGuilds) {
		t.Skip("API-Key does not have required permission for the test")
	}

	var acc Account
	if acc, err = api.Account(); err != nil {
		t.Error("Failed to fetch account data")
	}

	var ranks []GuildRank
	var members []GuildMember
	for _, g := range acc.Guilds {
		if ranks, err = api.GuildRanks(g); err != nil {
			continue
		} else if len(ranks) < 1 {
			t.Error("Fetched an unlikely number of ranks")
		}

		if members, err = api.GuildMembers(g); err != nil {
			continue
		} else if len(members) < 1 {
			t.Error("Fetched an unlikely number of members")
		}
		return
	}
	t.Error("Failed to parse private guild data or user is not guild leader in any guild: ", err)
}
