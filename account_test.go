package gw2api

import (
	"os"
	"testing"
)

func TestAccountBank(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermInventory) {
		t.Skip("API-Key does not have required permission for the test")
	}
	var items []BankItem
	if items, err = api.AccountBank(); err != nil {
		t.Error("Failed to parse the bank item data: ", err)
	} else if len(items) < 1 {
		t.Error("Fetched an unlikely number of bank items")
	}
}

func TestAccountDyes(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermUnlocks) {
		t.Skip("API-Key does not have required permission for the test")
	}
	var dyes []int
	if dyes, err = api.AccountDyes(); err != nil {
		t.Error("Failed to parse the dye unlock data: ", err)
	} else if len(dyes) < 1 {
		t.Error("Fetched an unlikely number of dye unlocks")
	}
}

func TestAccountMaterials(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermInventory) {
		t.Skip("API-Key does not have required permission for the test")
	}
	var materials []MaterialItem
	if materials, err = api.AccountMaterials(); err != nil {
		t.Error("Failed to parse the material item data: ", err)
	} else if len(materials) < 1 {
		t.Error("Fetched an unlikely number of material items")
	}
}

func TestAccountSkins(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermUnlocks) {
		t.Skip("API-Key does not have required permission for the test")
	}
	var skins []int
	if skins, err = api.AccountSkins(); err != nil {
		t.Error("Failed to parse the skin unlock data: ", err)
	} else if len(skins) < 1 {
		t.Error("Fetched an unlikely number of skin unlocks")
	}
}

func TestAccountWallet(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermWallet) {
		t.Skip("API-Key does not have required permission for the test")
	}
	var currencies []WalletCurrency
	if currencies, err = api.AccountWallet(); err != nil {
		t.Error("Failed to parse the wallet data: ", err)
	} else if len(currencies) < 1 {
		t.Error("Fetched an unlikely number of wallet currencies")
	}
}

func TestCharacterIds(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermCharacter) {
		t.Skip("API-Key does not have required permission for the test")
	}
	var testChars []string
	if testChars, err = api.Characters(); err != nil || len(testChars) < 1 {
		t.Error("Failed to parse the characters data: ", err)
	}
	var characters []Character
	if characters, err = api.CharacterIds(testChars[0]); err != nil {
		t.Error("Failed to parse the characters data: ", err)
	} else if len(characters) < 1 {
		t.Error("Fetched an unlikely number characters")
	}
}

func TestCharactersPage(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermCharacter) {
		t.Skip("API-Key does not have required permission for the test")
	}

	var characters []Character
	if characters, err = api.CharactersPage(0, 20); err != nil {
		t.Error("Failed to parse the characters data: ", err)
	} else if len(characters) < 1 {
		t.Error("Fetched an unlikely number characters")
	}
}

func TestAccountMinis(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermUnlocks) {
		t.Skip("API-Key does not have required permission for the test")
	}
	var minis []int
	if minis, err = api.AccountMinis(); err != nil {
		t.Error("Failed to parse the mini unlock data: ", err)
	} else if len(minis) < 1 {
		t.Error("Fetched an unlikely number of mini unlocks")
	}
}
