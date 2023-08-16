package server

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var baseUrl = url.URL{
	Scheme: "http",
	Host:   "api.tvmaze.com",
}

type Client struct {
}

func NewClient() Client {
	return Client{}
}
func (c Client) Get(url url.URL, ret interface{}) (status int, err error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return 0, err
	}
	defer func(get io.ReadCloser) {
		if err != nil {
			log.Fatal(err)
		}

	}(resp.Body)
	return resp.StatusCode, json.NewDecoder(resp.Body).Decode(&ret)
}
func BaseUrlWithPath(path string) url.URL {
	ret := baseUrl
	ret.Path = path
	return ret
}

func BaseUrlWithPathQuery(path, key, val string) url.URL {
	ret := baseUrl
	ret.Path = path
	ret.RawQuery = fmt.Sprintf("%s=%s", key, url.QueryEscape(val))
	return ret
}
func BaseUrlWithPathQueries(path string, vals map[string]string) url.URL {
	ret := baseUrl
	ret.Path = path
	var queryStrings []string
	for key, val := range vals {
		queryStrings = append(queryStrings, fmt.Sprintf("%s=%s", key, url.QueryEscape(val)))
	}
	ret.RawQuery = strings.Join(queryStrings, "&")
	return ret
}
