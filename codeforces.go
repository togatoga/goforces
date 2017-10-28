//
//Package goforces provides tructs and functions for accessing the codeforces API.
//If the API query is successful, return Go struct.
//
//Queries
//
//Executing queries on Codeforces API is very simple.
//Almost all of the methods don't require authentication.
//
//	api, _ := goforces.NewClient(nil)
//	ctx := context.Background()
//	problems, _ := api.GetProblemSetProblems(ctx, map[string]interface{}{"tags": []string{"graph", "dp"}})
//	for _, problem := range problems.Problems {
//		fmt.Printf("%+v\n", problem)
//	}
//
//Endpoints
//
//Goforces implements almost all of the endpoints defined in the Codeforces API(http://codeforces.com/api/help)
//More detailed information about the behavior of endpoint and the parameters can be found at the official Codeforces API documentation.
package goforces

import (
	"context"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	baseURL = "http://codeforces.com/api"
)

// Client manages the access for the Codeforces API.
type Client struct {
	APIKey     string
	APISecret  string
	URL        *url.URL
	HTTPClient *http.Client
	Logger     *log.Logger
}

//NewClient takes a logger and return a Client struct.
//The Client struct can be use for accessing the endpoints
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

//SetAPIKey takes an user api key.
//If you use authorized methods, you must set it.
func (c *Client) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

//SetAPISecret takes an user key secret.
//If you use authorized methods, you must set it.
func (c *Client) SetAPISecret(apiSecret string) {
	c.APISecret = apiSecret
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

func generateAPISig(method, apiSecret string, urlValues url.Values) string {

	//set api sig
	rand.Seed(time.Now().UnixNano())

	randSixDigits := ""
	for i := 0; i < 6; i++ {
		randSixDigits += strconv.Itoa(rand.Intn(10))
	}

	type Param struct {
		Param string
		Value string
	}
	params := make([]Param, 0)
	for param, values := range urlValues {
		for _, value := range values {
			params = append(params, Param{Param: param, Value: value})
		}
	}

	//sorted
	sort.Slice(params, func(i, j int) bool {
		if params[i].Param != params[j].Param {
			return params[i].Param < params[j].Param
		}
		return params[i].Value < params[j].Value
	})

	//concat params
	s := ""
	for i, p := range params {
		s += p.Param + "=" + p.Value
		if i != len(params)-1 {
			s += "&"
		}
	}

	text := randSixDigits + "/" + method + "?" + s + "#" + apiSecret
	hash := sha512.Sum512([]byte(text))

	return randSixDigits + fmt.Sprintf("%x", hash)

}
