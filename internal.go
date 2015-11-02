package gw2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
)

//Timeout solution adapted from Volker on stackoverflow
func (gw2 *GW2Api) dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, gw2.timeout)
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
