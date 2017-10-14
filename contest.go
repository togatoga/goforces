package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
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

func (c *Client) GetContestList(ctx context.Context, gym bool) ([]Contest, error) {
	c.Logger.Println("GetContestList gym: ", gym)
	v := url.Values{}
	if gym {
		v.Add("gym", "true")
	}
	spath := "/contest.list" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string    `json:"status"`
		Result []Contest `json:"result"`
	}
	var resp Response
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}

	var contests []Contest
	contests = resp.Result
	return contests, nil
}

func (c *Client) GetContestRatingChanges(ctx context.Context, contestId int) ([]RatingChange, error) {
	c.Logger.Println("GetContestRatingChange contestId: ", contestId)
	v := url.Values{}
	v.Add("contestId", strconv.Itoa(contestId))
	spath := "/contest.ratingChanges" + "?" + v.Encode()
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
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}

	var ratingChanges []RatingChange
	ratingChanges = resp.Result
	return ratingChanges, nil
}

func (c *Client) GetContestStandings(ctx context.Context, contestId int, options map[string]interface{}) (*Standings, error) {
	c.Logger.Println("GetContestStandings: ", contestId, options)

	v := url.Values{}
	v.Add("contestId", strconv.Itoa(contestId))

	//check options
	form, ok := options["from"]
	if ok {
		formVal := form.(int)
		if formVal <= 0 {
			return nil, fmt.Errorf("from must starts with 1-based index")
		}
		v.Add("from", strconv.Itoa(formVal))
	}

	count, ok := options["count"]
	if ok {
		countVal := count.(int)
		if countVal <= 0 {
			return nil, fmt.Errorf("count must be at least 1")
		}
		v.Add("count", strconv.Itoa(countVal))
	}

	handles, ok := options["handles"]
	if ok {
		handlesVal := handles.([]string)
		v.Add("handles", strings.Join(handlesVal, ";"))
	}

	room, ok := options["room"]
	if ok {
		roomVal := room.(int)
		if roomVal <= 0 {
			return nil, fmt.Errorf("room must be at least 1")
		}
		v.Add("room", strconv.Itoa(roomVal))
	}

	showUnofficial, ok := options["showUnofficial"]
	if ok {
		showUnofficialVal := showUnofficial.(bool)
		if showUnofficialVal {
			v.Add("showUnofficial", "true")
		}
	}

	spath := "/contest.standings" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string    `json:"status"`
		Result Standings `json:"result"`
	}
	var resp Response
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}

	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}

	return &resp.Result, nil
}
