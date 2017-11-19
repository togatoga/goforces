package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

//RecentAction represents Codeforces RecentAction
type RecentAction struct {
	TimeSeconds int       `json:"timeSeconds"`
	BlogEntry   BlogEntry `json:"blogEntry"`
	Comment     Comment   `json:"comment,omitempty"`
}

//GetRecentActions implements /recentActions
func (c *Client) GetRecentActions(ctx context.Context, maxCount int) ([]RecentAction, error) {
	c.Logger.Println("GetRecentActions :", maxCount)

	v := url.Values{}
	//check options
	v.Add("maxCount", strconv.Itoa(maxCount))

	spath := "/recentActions" + "?" + v.Encode()
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
		Result []RecentAction `json:"result"`
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
