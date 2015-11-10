package gw2api

import (
	"fmt"
	"net/url"
	"strconv"
)

// Account includes all general information
type Account struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	World   int      `json:"world"`
	Guilds  []string `json:"guilds"`
	Created string   `json:"created"`
}

// Account fetches the general account information
// Requires authentication
func (gw2 *GW2Api) Account() (acc Account, err error) {
	ver := "v2"
	tag := "account"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermAccount, nil, &acc)
	return
}

// BankItem describes an item stored in the players normal bank
type BankItem struct {
	ID        int   `json:"id"`
	Count     int   `json:"count"`
	Skin      int   `json:"skin"`
	Upgrades  []int `json:"upgrades"`
	Infusions []int `json:"infusions"`
}

// AccountBank returns an array of objects, each representing an item slot in
// the vault. If a slot is empty, it will return null. The amount of slots/bank
// tabs is implied by the length of the array.
// Requires authentication
func (gw2 *GW2Api) AccountBank() (items []BankItem, err error) {
	ver := "v2"
	tag := "account/bank"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermInventory, nil, &items)
	return
}

// AccountDyes returns an array of dyes unlocked by the player. The IDs can be
// resolved via ColorIds()
// Requires authentication
func (gw2 *GW2Api) AccountDyes() (dyes []int, err error) {
	ver := "v2"
	tag := "account/dyes"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermUnlocks, nil, &dyes)
	return
}

// MaterialItem represents a material stored in the material storage of the
// player
// Requires authentication
type MaterialItem struct {
	ID       int `json:"id"`
	Count    int `json:"count"`
	Category int `json:"category"`
}

// AccountMaterials returns an array of objects, each representing a material that
// can be stored in the vault. Every material will be returned, even if they
// have a count of 0
// Requires authentication
func (gw2 *GW2Api) AccountMaterials() (items []MaterialItem, err error) {
	ver := "v2"
	tag := "account/materials"
	params := url.Values{}
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermInventory, params, &items)
	return
}

// AccountSkins returns an array of skins unlocked by the player. The IDs can be
// resolved via SkinIds()
// Requires authentication
func (gw2 *GW2Api) AccountSkins() (skins []int, err error) {
	ver := "v2"
	tag := "account/dyes"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermUnlocks, nil, &skins)
	return
}

// WalletCurrency represents a currency and the amount owned by the player
type WalletCurrency struct {
	ID    int `json:"id"`
	Value int `json:"value"`
}

// AccountWallet returns an array of currencies owned by the player. The IDs can be
// resolved via CurrencyIds()
// Requires authentication
func (gw2 *GW2Api) AccountWallet() (currency []WalletCurrency, err error) {
	ver := "v2"
	tag := "account/materials"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermWallet, nil, &currency)
	return
}

// Characters returns an array of character names associated with the account.
// The names can be used to request detailed information via CharacterIds()
// Requires authentication
func (gw2 *GW2Api) Characters() (res []string, err error) {
	ver := "v2"
	tag := "characters"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermCharacter, nil, &res)
	return
}

// CraftingDiscipline contains information about learned crafting discplines on
// the character and shows their active state
type CraftingDiscipline struct {
	Discipline string `json:"discipline"`
	Rating     int    `json:"rating"`
	Active     bool   `json:"active"`
}

// Specialization represents the currently selected one and includes all active
// traits on the character
type Specialization struct {
	ID     int   `json:"id"`
	Traits []int `json:"trait"`
}

// Specalizations is a meta object which holds the specalizations associated
// with the different game types
type Specalizations struct {
	PvE Specialization `json:"pve"`
	PvP Specialization `json:"pvp"`
	WvW Specialization `json:"wvw"`
}

// InventoryItem Item in a characters inventory and their count
type InventoryItem struct {
	ID    int `json:"id"`
	Count int `json:"count"`
}

// Bag on the character with size and inventory
type Bag struct {
	ID        int             `json:"id"`
	Size      int             `json:"size"`
	Inventory []InventoryItem `json:"inventory"`
}

// Equipment worn by the character
type Equipment struct {
	ID        int    `json:"id"`
	Slot      string `json:"slot"`
	Upgrades  []int  `json:"upgrades"`
	Infusions []int  `json:"infusions"`
	Skin      int    `json:"skin"`
}

// Character combines all information about the character
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

// CharacterIds requests detailed information about the characters requested via
// the IDs. The IDs can also be character names as e.g. provided by the
// Characters() API Call
// Requires authentication
func (gw2 *GW2Api) CharacterIds(ids ...string) (chars []Character, err error) {
	ver := "v2"
	tag := "characters"
	params := url.Values{}
	params.Add("ids", commaList(ids))
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermCharacter, params, &chars)
	return
}

// CharactersPage allows pagination for the character objects and enables the
// request of all character objects
// Requires authentication
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
