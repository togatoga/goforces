package goforces

import (
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Logger     *log.Logger
}

func NewClient(urlStr string, logger *log.Logger) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}
	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}
	return &Client{URL: parsedURL, Logger: logger}, nil
}
