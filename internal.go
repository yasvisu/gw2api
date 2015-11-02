package gw2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

//Timeout solution adapted from Volker on stackoverflow

type GW2Api struct {
	timeout   time.Duration
	client    http.Client
	auth      string
	authFlags uint
}

func (gw2 *GW2Api) dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, gw2.timeout)
}

//Initialize http client (for timeout)
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

//Fetcher function to return only the body of a HTTP request.
//Parameters: version, tag, appendix strings.
func (gw2 *GW2Api) fetchJSON(ver string, tag string, appendix string) ([]byte, error) {
	resp, err := gw2.client.Get("https://api.guildwars2.com/" + ver + "/" + tag + appendix)
	if err != nil {
		return nil, err
	}
	var jsonBlob []byte
	jsonBlob, err = ioutil.ReadAll(resp.Body)
	return jsonBlob, err
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

func (gw2 *GW2Api) fetchEndpoint(ver, tag string, params url.Values, result interface{}) (err error) {
	var data []byte
	args := ""
	if len(params) > 0 {
		args = "?" + params.Encode()
	}
	if data, err = gw2.fetchJSON(ver, tag, args); err != nil {
		return err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		var gwerr Error
		if err = json.Unmarshal(data, &gwerr); err != nil {
			return err
		}
		return errors.New("Endpoint returned error: " + gwerr.Error())
	}
	return
}

func (gw2 *GW2Api) fetchAuthenticatedEndpoint(ver, tag string, perm Permission, params url.Values, result interface{}) (err error) {
	if len(gw2.auth) < 1 {
		return fmt.Errorf("API requires authentication")
	}

	var i uint
	for i = 0; i < uint(PermSize); i++ {
		if !flagGet(gw2.authFlags, i) {
			return fmt.Errorf("Missing permissions for authenticated Endpoint")
		}
	}

	params.Add("access_token", gw2.auth)
	return gw2.fetchEndpoint(ver, tag, params, result)
}
