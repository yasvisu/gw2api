//Package gw2api provides bindings for the Guild Wars 2 v2 API
//
//Using this package is as simple as calling the functions on the API struct
//Further examples can be found with the function definitions
//
//   func main() {
//     api := NewGW2Api()
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

// GW2Api is the state holder of the API. It includes authentication information
type GW2Api struct {
	timeout   time.Duration
	client    http.Client
	auth      string
	authFlags uint
}

// NewGW2Api returns a simple GW2Api with a HTTP timeout of 15 seconds
func NewGW2Api() *GW2Api {
	api := &GW2Api{
		client: http.Client{},
	}
	api.SetTimeout(time.Duration(15 * time.Second))
	return api
}

// NewAuthenticatedGW2Api returns an authenticated GW2Api or an error if
// authentication failed
func NewAuthenticatedGW2Api(auth string) (api *GW2Api, err error) {
	api = NewGW2Api()
	err = api.SetAuthentication(auth)
	return
}

// SetTimeout set HTTP timeout for all new HTTP connections started from this
// instance
func (gw2 *GW2Api) SetTimeout(t time.Duration) {
	gw2.timeout = t
	gw2.client.Transport = &http.Transport{
		Dial: gw2.dialTimeout,
	}
}

// SetAuthentication adds authentication to a previously un-authenticated
// instance of GW2Api
func (gw2 *GW2Api) SetAuthentication(auth string) (err error) {
	if m, err := regexp.Match(`^(?:[A-F\d]{4,20}-?){8,}$`, []byte(auth)); !m {
		return fmt.Errorf("Provided API Key doesn't match expectations: %s", err)
	}

	gw2.auth = auth
	gw2.authFlags = 0
	var token TokenInfo
	if token, err = gw2.TokenInfo(); err != nil {
		gw2.auth = ""
		return fmt.Errorf("Failed to fetch Token Info: %s", err)
	}
	for _, perm := range token.Permissions {
		if p, e := permissionsMapping[perm]; e {
			flagSet(&gw2.authFlags, uint(p))
		} else {
			log.Print("Found new permission: ", perm)
		}
	}
	return
}

// HasPermission checks for permissions of the API Key provided by the user
func (gw2 *GW2Api) HasPermission(p Permission) bool {
	return len(gw2.auth) > 1 && flagGet(gw2.authFlags, uint(p))
}
