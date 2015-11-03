package gw2api

import (
	"net/url"
)

type Account struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	World   int      `json:"world"`
	Guilds  []string `json:"guilds"`
	Created string   `json:"created"`
}

func (gw2 *GW2Api) Account() (acc Account, err error) {
	ver := "v2"
	tag := "account"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermAccount, nil, &acc)
	return
}

type BankItem struct {
	ID        int   `json:"id"`
	Count     int   `json:"count"`
	Skin      int   `json:"skin"`
	Upgrades  []int `json:"upgrades"`
	Infusions []int `json:"infusions"`
}

func (gw2 *GW2Api) AccountBank() (items []BankItem, err error) {
	ver := "v2"
	tag := "account/bank"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermInventory, nil, &items)
	return
}

func (gw2 *GW2Api) AccountDyes() (dyes []int, err error) {
	ver := "v2"
	tag := "account/dyes"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermUnlocks, nil, &dyes)
	return
}

type MaterialItem struct {
	ID       int `json:"id"`
	Count    int `json:"count"`
	Category int `json:"category"`
}

func (gw2 *GW2Api) AccountMaterials() (items []MaterialItem, err error) {
	ver := "v2"
	tag := "account/materials"
	params := url.Values{}
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermInventory, params, &items)
	return
}

func (gw2 *GW2Api) AccountSkins() (skins []int, err error) {
	ver := "v2"
	tag := "account/dyes"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermUnlocks, nil, &skins)
	return
}

type WalletCurrency struct {
	ID    int `json:"id"`
	Value int `json:"value"`
}

func (gw2 *GW2Api) AccountWallet() (currency []WalletCurrency, err error) {
	ver := "v2"
	tag := "account/materials"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermWallet, nil, &currency)
	return
}

func (gw2 *GW2Api) Characters() (res []string, err error) {
	ver := "v2"
	tag := "characters"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermCharacter, nil, &res)
	return
}
