//Package gw2api provides bindings for the Guild Wars 2 v2 API
//
//Using this package is as simple as calling the functions on the API struct
//Further examples can be found with the function definitions
//
//   func main() {
//     api := NewG2Api()
//     b, _ := api.Build()
//     fmt.Println(b)
//   }
package gw2api

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
)

type GW2Api struct {
	timeout   time.Duration
	client    http.Client
	auth      string
	authFlags uint
}

func NewGW2Api() *GW2Api {
	api := &GW2Api{
		client: http.Client{},
	}
	api.SetTimeout(time.Duration(15 * time.Second))
	return api
}

func NewAuthenticatedGW2Api(auth string) (api *GW2Api, err error) {
	api = NewGW2Api()
	err = api.SetAuthentication(auth)
	return
}

func (gw2 *GW2Api) SetTimeout(t time.Duration) {
	gw2.timeout = t
	gw2.client.Transport = &http.Transport{
		Dial: gw2.dialTimeout,
	}
}

func (gw2 *GW2Api) SetAuthentication(auth string) (err error) {
	if m, err := regexp.Match(`^(?:[A-F\d]{4,20}-?){8,}$`, []byte(auth)); !m {
		return fmt.Errorf("Provided API Key doesn't match expectations: %s", err)
	}

	gw2.auth = auth
	gw2.authFlags = 0
	var token TokenInfo
	if token, err = gw2.TokenInfo(); err != nil {
		gw2.auth = ""
		return fmt.Errorf("Failed to fetch Token Info")
	}
	for _, perm := range token.Permissions {
		if p, e := PermissionsMapping[perm]; e {
			flagSet(&gw2.authFlags, uint(p))
		} else {
			log.Print("Found new permission: ", perm)
		}
	}
	return
}

func (gw2 *GW2Api) HasPermission(p Permission) bool {
	return len(gw2.auth) > 1 && flagGet(gw2.authFlags, uint(p))
}
