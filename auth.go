package gw2api

type Permission uint

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
	PermSize
)

var (
	PermissionsMapping = map[string]Permission{
		"account":     PermAccount,
		"characters":  PermCharacter,
		"inventories": PermInventory,
		"tradingpost": PermTradingpost,
		"wallet":      PermWallet,
		"unlocks":     PermUnlocks,
		"pvp":         PermPvP,
		"builds":      PermBuilds,
	}
)

type TokenInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func (gw2 *GW2Api) TokenInfo() (token TokenInfo, err error) {
	ver := "v2"
	tag := "tokeninfo"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermAccount, nil, &token)
	return
}
