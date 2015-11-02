package gw2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//Timeout solution adapted from Volker on stackoverflow

type GW2Api struct {
	timeout time.Duration
	client  http.Client
}

func (gw2 *GW2Api) dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, gw2.timeout)
}

//Initialize http client (for timeout)
func NewGW2Api() *GW2Api {
	var api = &GW2Api{}
	api.timeout = time.Duration(15 * time.Second)
	transport := http.Transport{
		Dial: api.dialTimeout,
	}

	api.client = http.Client{
		Transport: &transport,
	}
	return api
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

func stringSlice(ids []int) []string {
	newIds := make([]string, len(ids))
	for i, id := range ids {
		newIds[i] = strconv.Itoa(id)
	}
	return newIds
}

func commaList(ids []string) string {
	var appendix bytes.Buffer
	for i, id := range ids {
		if i > 0 {
			appendix.WriteString(",")
		}
		appendix.WriteString(id)
	}
	return appendix.String()
}
