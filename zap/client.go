// Zed Attack Proxy (ZAP) and its related class files.
//
// ZAP is an HTTP/HTTPS proxy for assessing web application security.
//
// Copyright 2017 the ZAP development team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zap

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	DefaultBase           = "http://zap/JSON/"
	DefaultBaseOther      = "http://zap/OTHER/"
	DefaultHTTPSBase      = "https://zap/JSON/"
	DefaultHTTPSBaseOther = "https://zap/OTHER/"
	DefaultProxy          = "tcp://127.0.0.1:8080"
	ZAP_API_KEY_PARAM     = "apikey"
	ZAP_API_KEY_HEADER    = "X-ZAP-API-Key"
)

// Config defines the config of ZAP client
type Config struct {
	Base      string
	BaseOther string
	Proxy     string
	APIKey    string
	TLSConfig tls.Config
}

// Client is a ZAP client that allows you to access ZAP API
type Client struct {
	*Config
	httpClient *http.Client
}

// NewClient returns a new ZAP client based on the passed in config
func NewClient(cfg *Config) (Interface, error) {
	if cfg.Base == "" {
		cfg.Base = DefaultBase
	}
	if cfg.BaseOther == "" {
		cfg.BaseOther = DefaultBaseOther
	}
	if cfg.Proxy == "" {
		cfg.Proxy = DefaultProxy
	}

	proxyURL, err := url.Parse(cfg.Proxy)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy:           http.ProxyURL(proxyURL),
			TLSClientConfig: &cfg.TLSConfig,
		},
	}
	return &Client{
		Config:     cfg,
		httpClient: httpClient,
	}, nil
}

// Request sends HTTP request to zap base("http://zap/JSON/") API group
func (c *Client) Request(path string, queryParams map[string]string) (map[string]interface{}, error) {
	body, err := c.request(c.Base+path, queryParams)
	if err != nil {
		return nil, err
	}
	// NOTE: since Golang can not unmarshal a json without knowing the exact struct
	// so we can only unmarshal json into a  map[string]interface{} here.
	var obj map[string]interface{}
	if err := json.Unmarshal(body, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// RequestOther sends HTTP request to zap other("http://zap/OTHER/") API group
func (c *Client) RequestOther(path string, queryParams map[string]string) ([]byte, error) {
	return c.request(c.BaseOther+path, queryParams)
}

func (c *Client) request(path string, queryParams map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	if len(queryParams) == 0 {
		queryParams = map[string]string{}
	}
	// Send the API key even if there are no parameters,
	// older ZAP versions might need API key as (query) parameter.
	queryParams[ZAP_API_KEY_PARAM] = c.APIKey

	// add url query parameter
	query := req.URL.Query()
	for k, v := range queryParams {
		if v == "" {
			continue
		}
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	// add HTTP Accept header
	req.Header.Add("Accept", "application/json")
	// add API Key header
	req.Header.Add(ZAP_API_KEY_HEADER, c.APIKey)

	// Close the connection
	req.Close = true

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Errored when sending request to the server: %v", err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
