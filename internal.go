package gw2api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//Timeout solution adapted from Volker on stackoverflow
var timeout time.Duration
var client http.Client

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

//Initialize http client (for timeout)
func init() {
	timeout = time.Duration(15 * time.Second)
	transport := http.Transport{
		Dial: dialTimeout,
	}

	client := http.Client{
		Transport: &transport,
	}
	client = client
}

//Fetcher function to return only the body of a HTTP request.
//Parameters: version, tag, appendix strings.
func fetchJSON(ver string, tag string, appendix string) ([]byte, error) {
	resp, err := client.Get("https://api.guildwars2.com/" + ver + "/" + tag + appendix)
	if err != nil {
		return nil, err
	}
	var jsonBlob []byte
	jsonBlob, err = ioutil.ReadAll(resp.Body)
	return jsonBlob, err
}

//Set the timeout for each HTTP request.
func SetTimeout(t time.Duration) {
	timeout = t
}

//Fetch function for the endpoints which only serve ids as []string or []int
//e.g. Worlds returning [1001,2001] etc
func fetchEndpoint(ver, tag, lang string, objects interface{}) (err error) {
	var appendix bytes.Buffer
	var data []byte
	if data, err = fetchJSON(ver, tag, appendix.String()); err != nil {
		return
	}
	err = json.Unmarshal(data, &objects)
	if err != nil {
		var gwerr GW2ApiError
		if err = json.Unmarshal(data, &gwerr); err != nil {
			return err
		}
		return gwerr
	}
	return
}
