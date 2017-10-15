package goforces

import (
	"context"
	"crypto/sha512"
	"fmt"
	rand2 "math/rand"
	"net/url"
	"strconv"
	"strings"
	time2 "time"
)

type User struct {
	LastName                string `json:"lastName"`
	Country                 string `json:"country"`
	LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds"`
	City                    string `json:"city"`
	Rating                  int    `json:"rating"`
	FriendOfCount           int    `json:"friendOfCount"`
	TitlePhoto              string `json:"titlePhoto"`
	Handle                  string `json:"handle"`
	Avatar                  string `json:"avatar"`
	FirstName               string `json:"firstName"`
	Contribution            int    `json:"contribution"`
	Organization            string `json:"organization"`
	Rank                    string `json:"rank"`
	MaxRating               int    `json:"maxRating"`
	RegistrationTimeSeconds int    `json:"registrationTimeSeconds"`
	MaxRank                 string `json:"maxRank"`
}

func (c *Client) GetUserBlogEntries(ctx context.Context, handle string) ([]BlogEntry, error) {
	c.Logger.Println("GetUserBlogEntries :", handle)

	v := url.Values{}
	v.Add("handle", handle)
	spath := "/user.blogEntries" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string      `json:"status"`
		Result []BlogEntry `json:"result"`
	}
	var resp Response
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}

	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error : %s", resp.Status)
	}

	return resp.Result, nil
}

func (c *Client) GetUserInfo(ctx context.Context, handles []string) ([]User, error) {
	c.Logger.Println("GetUserInfo :", handles)

	v := url.Values{}
	v.Add("handles", strings.Join(handles, ";"))
	spath := "/user.info" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string `json:"status"`
		Result []User `json:"result"`
	}
	var resp Response
	if err := decodeBody(
		res, &resp); err != nil {
		return nil, err
	}

	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error : %s", resp.Status)
	}

	return resp.Result, nil
}

func (c *Client) GetUserRatedList(ctx context.Context, options map[string]interface{}) ([]User, error) {
	c.Logger.Println("GetUserRatedList :", options)

	v := url.Values{}
	//check activeOnly
	activeOnly, ok := options["activeOnly"]
	if ok {
		activeOnlyVal := activeOnly.(bool)
		if activeOnlyVal {
			v.Add("activeOnly", "true")
		}
	}

	spath := "/user.ratedList" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string `json:"status"`
		Result []User `json:"result"`
	}
	var resp Response
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}

	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error : %s", resp.Status)
	}

	return resp.Result, nil
}

func (c *Client) GetUserRating(ctx context.Context, handle string) ([]RatingChange, error) {
	c.Logger.Println("GetUserRating :", handle)

	v := url.Values{}
	v.Add("handle", handle)

	spath := "/user.rating" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string         `json:"status"`
		Result []RatingChange `json:"result"`
	}
	var resp Response
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}

	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error : %s", resp.Status)
	}

	return resp.Result, nil
}

func (c *Client) GetUserStatus(ctx context.Context, handle string, options map[string]interface{}) ([]Submission, error) {
	c.Logger.Println("GetUserRating :", handle)

	v := url.Values{}
	v.Add("handle", handle)

	//check options
	from, ok := options["from"]
	if ok {
		fromVal := from.(int)
		if fromVal <= 0 {
			return nil, fmt.Errorf("from must be at least 1")
		}
		v.Add("from", strconv.Itoa(fromVal))
	}
	count, ok := options["count"]
	if ok {
		countVal := count.(int)
		if countVal <= 0 {
			return nil, fmt.Errorf("count must be at least 1")
		}
		v.Add("count", strconv.Itoa(countVal))
	}

	spath := "/user.status" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string       `json:"status"`
		Result []Submission `json:"result"`
	}
	var resp Response
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}

	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error : %s", resp.Status)
	}

	return resp.Result, nil
}

func (c *Client) GetUserFriends(ctx context.Context, options map[string]interface{}) ([]Member, error) {
	c.Logger.Println("GetUserFriends :", options)

	v := url.Values{}

	//check api key and scret
	if c.ApiKey == "" || c.ApiSecret == "" {
		return nil, fmt.Errorf("GetUserFriends requires your api key and api secret")
	}

	v.Add("apiKey", c.ApiKey)
	v.Add("time", strconv.FormatInt(time2.Now().Unix(), 10))
	//check options
	onlyOnline, ok := options["onlyOnline"]
	if ok {
		onlyOnlineVal := onlyOnline.(bool)
		if onlyOnlineVal {
			v.Add("onlyOnline", "true")
		}
	}
	//set api sig
	rand2.Seed(time2.Now().UnixNano())
	rand := ""
	for i := 0; i < 6; i++ {
		rand += strconv.Itoa(rand2.Intn(10))
	}
	hash := rand + "/user.friends?" + v.Encode() + "#" + c.ApiSecret
	apiSig := rand + string(sha512.Sum512([]byte(hash)))
	v.Add("apiSig", string(apiSig))

	spath := "/user.friends" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string       `json:"status"`
		Result []Submission `json:"result"`
	}
	var resp Response
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}

	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error : %s", resp.Status)
	}

	return resp.Result, nil
}
