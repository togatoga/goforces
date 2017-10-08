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
	"path"
)

const (
	baseURL = "http://codeforces.com/api"
)

type Client struct {
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
	return &Client{URL: parsedURL, Logger: logger}, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
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
