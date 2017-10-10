package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strings"
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

func (c *Client) GetUserInfo(ctx context.Context, handles []string) ([]User, error) {
	type UserInfoResponse struct {
		Status string `json:"status"`
		Result []User `json:"result"`
	}

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

	var resp UserInfoResponse
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}

	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error : %s", resp.Status)
	}
	var user []User
	user = resp.Result
	return user, nil
}

func (c *Client) GetUserRatedList(ctx context.Context, activeOnly bool) ([]User, error) {
	type UserRatedListResponse struct {
		Status string `json:"status"`
		Result []User `json:"result"`
	}

	v := url.Values{}

	//check activeOnly
	if activeOnly {
		v.Add("activeOnly", "true")
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

	var resp UserRatedListResponse
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}

	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error : %s", resp.Status)
	}
	var user []User
	user = resp.Result
	return user, nil
}

func (c *Client) GetUserRating(ctx context.Context, handle string) (interface{}, error) {
	//TODO
	return nil,nil
}
