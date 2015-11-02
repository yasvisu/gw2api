//Package gw2api provides bindings for the Guild Wars 2 API.
//
//Using this package is as simple as calling its functions.
//
//   func main() {
//     i, _ := gw2api.Items()
//     fmt.Println(i)
//   }
package gw2api

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type GW2Api struct {
	timeout   time.Duration
	client    http.Client
	auth      string
	authFlags uint
}

func NewGW2Api() (api *GW2Api) {
	api.timeout = time.Duration(15 * time.Second)
	transport := http.Transport{
		Dial: api.dialTimeout,
	}

	api.client = http.Client{
		Transport: &transport,
	}
	return api
}

func NewAuthenticatedGW2Api(auth string) (api *GW2Api, err error) {
	api = NewGW2Api()
	err = api.SetAuthentication(auth)
	return
}

//Set the timeout for each HTTP request.
func (gw2 *GW2Api) SetTimeout(t time.Duration) {
	gw2.timeout = t
}

func (gw2 *GW2Api) SetAuthentication(auth string) (err error) {
	gw2.auth = auth
	gw2.authFlags = 0
	var token TokenInfo
	if token, err = gw2.TokenInfo(); err != nil {
		gw2.auth = ""
		return fmt.Errorf("Failed to fetch Token Info")
	}
	for _, perm := range token.Permissions {
		if p, e := PermissionsMapping[perm]; e {
			flagSet(gw2.authFlags, uint(p))
		} else {
			log.Print("Found new permission: ", perm)
		}
	}
	return
}
