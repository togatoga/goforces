package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//User represents Codeforces User
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

const (
	divsionRatingThreshold = 1900
)

// Div1 returns boolean whether user belog to div1
func (u User) Div1() bool {
	if u.Rating >= divsionRatingThreshold {
		return true
	}
	return false
}

// Div2 returns boolean whether user belog to div2
func (u User) Div2() bool {
	if u.Rating < divsionRatingThreshold {
		return true
	}
	return false
}

// Color returns user's color of rating.
// http://codeforces.com/blog/entry/20638
func (u User) Color() string {
	//div1
	if u.Div1() {
		if u.Rating >= 1900 && u.Rating < 2200 {
			return "Violet"
		}
		if u.Rating >= 2200 && u.Rating < 2400 {
			return "Orange"
		}
		// rating >= 2400
		return "Red"
	}
	//div2
	if u.Rating < 1200 {
		return "Gray"
	}
	if u.Rating >= 1200 && u.Rating < 1400 {
		return "Green"
	}
	if u.Rating >= 1400 && u.Rating < 1599 {
		return "Cyan"
	}
	//1600 <= rating < 1899
	return "Blue"
}

//GetUserBlogEntries implements /user.blogEntries
func (c *Client) GetUserBlogEntries(ctx context.Context, handle string) ([]BlogEntry, error) {
	c.Logger.Println("GetUserBlogEntries :", handle)

	v := url.Values{}
	v.Add("handle", handle)
	spath := "/user.blogEntries" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
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

//GetUserInfo implements /user.info
func (c *Client) GetUserInfo(ctx context.Context, handles []string) ([]User, error) {
	c.Logger.Println("GetUserInfo :", handles)

	v := url.Values{}
	v.Add("handles", strings.Join(handles, ";"))
	spath := "/user.info" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
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

//UserRatedListOptions represents the options of /user.ratedlist
type UserRatedListOptions struct {
	ActiveOnly bool
}

func (o *UserRatedListOptions) options() interface{} {
	if o == nil {
		return nil
	}
	type option struct {
		ActiveOnly bool `url:"activeOnly,omitempty"`
	}
	return &option{ActiveOnly: o.ActiveOnly}
}

//GetUserRatedList implements /user.ratedList
func (c *Client) GetUserRatedList(ctx context.Context, options *UserRatedListOptions) ([]User, error) {
	c.Logger.Println("GetUserRatedList :", options)
	v := url.Values{}

	spath := "/user.ratedList" + "?" + v.Encode() + "&"
	req, err := c.newRequest(ctx, "GET", spath, options.options(), nil)
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

//GetUserRating implements /user.rating
func (c *Client) GetUserRating(ctx context.Context, handle string) ([]RatingChange, error) {
	c.Logger.Println("GetUserRating :", handle)

	v := url.Values{}
	v.Add("handle", handle)

	spath := "/user.rating" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
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

//UserStatusOptions represents the opetions of /user.status
type UserStatusOptions struct {
	From  int
	Count int
}

func (o *UserStatusOptions) options() interface{} {
	if o == nil {
		return nil
	}
	type option struct {
		From  int `url:"from"`
		Count int `url:"count"`
	}
	return &option{From: o.From, Count: o.Count}
}

//GetUserStatus implements /user.status
func (c *Client) GetUserStatus(ctx context.Context, handle string, options *UserStatusOptions) ([]Submission, error) {
	c.Logger.Println("GetUserRating :", handle)

	v := url.Values{}
	v.Add("handle", handle)
	spath := "/user.status" + "?" + v.Encode() + "&"
	req, err := c.newRequest(ctx, "GET", spath, options.options(), nil)
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

//UserFriendsOptions represents the opetions of /user.friends
type UserFriendsOptions struct {
	OnlyOnline bool
}

func (o *UserFriendsOptions) options() interface{} {
	if o == nil {
		return nil
	}
	type option struct {
		OnlyOnline bool `url:"onlyOnline"`
	}
	return &option{OnlyOnline: o.OnlyOnline}
}

//GetUserFriends implements /user.friends
//You must your api key and secret key before call this method.
func (c *Client) GetUserFriends(ctx context.Context, options *UserFriendsOptions) ([]string, error) {
	c.Logger.Println("GetUserFriends :", options)

	v := url.Values{}
	//check api key and scret
	if c.APIKey == "" || c.APISecret == "" {
		return nil, fmt.Errorf("GetUserFriends requires your api key and api secret")
	}

	v.Add("apiKey", c.APIKey)
	v.Add("time", strconv.FormatInt(time.Now().Unix(), 10))
	apiSig := generateAPISig("user.friends", c.APISecret, v)
	v.Add("apiSig", apiSig)

	spath := "/user.friends" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, options.options(), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string   `json:"status"`
		Result []string `json:"result"`
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
