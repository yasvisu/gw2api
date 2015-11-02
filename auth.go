package gw2api

const (
	PermAccount = iota
	PermCharacter
	PermInventory
	PermTradingpost
	PermWallet
	PermUnlocks
	PermPvP
	PermBuilds
)

type TokenInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func (gw2 *GW2Api) TokenInfo() (token TokenInfo, err error) {
	ver := "v2"
	tag := "tokeninfo"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, nil, &token)
	return
}

func flagGet(n uint64, pos uint) bool {
	return n&1<<pos == 0
}

func flagSet(n uint64, pos uint) {
	n |= 1 << pos
}
