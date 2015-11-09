package gw2api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

//Timeout solution adapted from Volker on stackoverflow
func (gw2 *GW2Api) dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, gw2.timeout)
}

func (gw2 *GW2Api) fetchEndpoint(ver, tag string, params url.Values, result interface{}) (err error) {
	var endpoint *url.URL
	endpoint, _ = url.Parse("https://api.guildwars2.com")
	endpoint.Path += "/" + ver + "/" + tag
	if params != nil {
		endpoint.RawQuery = params.Encode()
	}

	var resp *http.Response
	if resp, err = gw2.client.Get(endpoint.String()); err != nil {
		return err
	}
	var data []byte
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.Unmarshal(data, &result); err != nil {
		var gwerr Error
		if err = json.Unmarshal(data, &gwerr); err != nil {
			return err
		}
		return fmt.Errorf("Endpoint returned error: %s", gwerr.Error())
	}
	return
}

func (gw2 *GW2Api) fetchAuthenticatedEndpoint(ver, tag string, perm Permission, params url.Values, result interface{}) (err error) {
	if len(gw2.auth) < 1 {
		return fmt.Errorf("API requires authentication")
	}

	if perm >= PermAccount && !flagGet(gw2.authFlags, uint(perm)) {
		return fmt.Errorf("Missing permissions for authenticated Endpoint")
	}

	if params == nil {
		params = url.Values{}
	}
	params.Add("access_token", gw2.auth)
	return gw2.fetchEndpoint(ver, tag, params, result)
}
