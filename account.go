package gw2api

import (
	"fmt"
	"net/url"
	"strconv"
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

type CraftingDiscipline struct {
	Discipline string `json:"discipline"`
	Rating     int    `json:"rating"`
	Active     bool   `json:"active"`
}
type Specalization struct {
	ID     int   `json:"id"`
	Traits []int `json:"trait"`
}
type Specalizations struct {
	PvE Specalization `json:"pve"`
	PvP Specalization `json:"pvp"`
	WvW Specalization `json:"wvw"`
}
type InventoryItem struct {
	ID    int `json:"id"`
	Count int `json:"count"`
}
type Bag struct {
	ID        int             `json:"id"`
	Size      int             `json:"size"`
	Inventory []InventoryItem `json:"inventory"`
}

type Equipment struct {
	ID        int    `json:"id"`
	Slot      string `json:"slot"`
	Upgrades  []int  `json:"upgrades"`
	Infusions []int  `json:"infusions"`
	Skin      int    `json:"skin"`
}

type Character struct {
	Name          string               `json:"name"`
	Race          string               `json:"race"`
	Gender        string               `json:"gender"`
	Profession    string               `json:"profession"`
	Level         int                  `json:"level"`
	Guild         string               `json:"guild"`
	Created       string               `json:"created"`
	Age           int                  `json:"age"`
	Deaths        int                  `json:"deaths"`
	Crafting      []CraftingDiscipline `json:"crafting"`
	Specalization Specalizations       `json:"specalization"`
	Bags          []Bag                `json:"bags"`
	Equipment     []Equipment          `json:"equipment"`
}

func (gw2 *GW2Api) CharacterIds(ids ...string) (chars []Character, err error) {
	ver := "v2"
	tag := "characters"
	params := url.Values{}
	params.Add("ids", commaList(ids))
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermCharacter, params, &chars)
	return
}

func (gw2 *GW2Api) CharactersPage(page, pageSize int) (chars []Character, err error) {
	if page < 0 {
		return nil, fmt.Errorf("Page parameter cannot be a negative number!")
	}

	ver := "v2"
	tag := "characters"
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	if pageSize >= 0 {
		params.Add("page_size", strconv.Itoa(pageSize))
	}
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermCharacter, params, &chars)
	return
}
