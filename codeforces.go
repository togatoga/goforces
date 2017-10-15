package goforces

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	baseURL = "http://codeforces.com/api"
)

type Client struct {
	ApiKey     string
	ApiSecret  string
	URL        *url.URL
	HTTPClient *http.Client
	Logger     *log.Logger
}

func NewClient(logger *log.Logger) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, err
	}
	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}
	return &Client{URL: parsedURL, HTTPClient: http.DefaultClient, Logger: logger}, nil
}

func (c *Client) SetApiKey(apiKey string) {
	c.ApiKey = apiKey
}

func (c *Client) SetApiSecret(apiSecret string) {
	c.ApiSecret = apiSecret
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.URL.String()+spath, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNoContent {
		//empty response
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request Error: %s", resp.Status)
	}
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
