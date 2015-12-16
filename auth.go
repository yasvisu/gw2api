package gw2api

// Permission abstracts the bitmask containing permission information
type Permission uint

// Perm* represent the specific permission as required by the API
const (
	_                      = iota
	PermAccount Permission = iota
	PermCharacter
	PermInventory
	PermTradingpost
	PermWallet
	PermUnlocks
	PermPvP
	PermBuilds
	PermProgression
	PermGuilds
	PermSize
)

//
var (
	permissionsMapping = map[string]Permission{
		"account":     PermAccount,
		"characters":  PermCharacter,
		"inventories": PermInventory,
		"tradingpost": PermTradingpost,
		"wallet":      PermWallet,
		"unlocks":     PermUnlocks,
		"pvp":         PermPvP,
		"builds":      PermBuilds,
		"progression": PermProgression,
		"guilds":      PermGuilds,
	}
)

// TokenInfo contains information about the provided API Key of the user.
// Including the name of the key as set by the user and the permissions
// associated with it
type TokenInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// TokenInfo requests the token information from the authenticated API
// Requires authentication
func (gw2 *GW2Api) TokenInfo() (token TokenInfo, err error) {
	ver := "v2"
	tag := "tokeninfo"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, 0, nil, &token)
	return
}
