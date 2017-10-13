package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type Contest struct {
	DurationSeconds     int64  `json:"durationSeconds"`
	Frozen              bool   `json:"frozen"`
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	Phase               string `json:"phase"`
	RelativeTimeSeconds int64  `json:"relativeTimeSeconds"`
	StartTimeSeconds    int64  `json:"startTimeSeconds"`
	Type                string `json:"type"`
}

func (c *Client) GetContestHacks(ctx context.Context, contestId int) ([]Hack, error) {
	c.Logger.Println("GetContestHacks contestId: ", contestId)
	v := url.Values{}
	v.Add("contestId", strconv.Itoa(contestId))
	spath := "/contest.hacks" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type HacksResponse struct {
		Status string `json:"status"`
		Result []Hack `json:"result"`
	}
	var resp HacksResponse
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}
	var hacks []Hack
	hacks = resp.Result
	return hacks, nil
}
