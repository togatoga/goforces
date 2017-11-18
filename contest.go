package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

//Contest represents a Codeforces Contest
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

//Div2 returns boolean whether contest'name contains "Div. 2"
func (c Contest) Div2() bool {
	return strings.Contains(c.Name, "Div. 2")
}

// Finished returns boolean whether contest was over
func (c Contest) Finished() bool {
	return c.Phase == "FINISHED"
}

// Before returns boolean whether contest hasn't started.
func (c Contest) Before() bool {
	return c.Phase == "BEFORE"
}

// Coding returns boolean whether contest is now being held.
func (c Contest) Coding() bool {
	return c.Phase == "CODING"
}

// ContestURL returns the contest's url.
func (c Contest) ContestURL() string {
	return fmt.Sprintf("http://codeforces.com/contest/%d", c.ID)
}

//GetContestHacks implements /contest.hacks
func (c *Client) GetContestHacks(ctx context.Context, contestID int) ([]Hack, error) {
	c.Logger.Println("GetContestHacks contestId: ", contestID)
	v := url.Values{}
	v.Add("contestId", strconv.Itoa(contestID))
	spath := "/contest.hacks" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
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

//ContestListOptions represents the option of /contest.list
type ContestListOptions struct {
	Gym bool `url:"gym"`
}

//GetContestList implements /contest.list
func (c *Client) GetContestList(ctx context.Context, options *ContestListOptions) ([]Contest, error) {
	c.Logger.Println("GetContestList : ", options)
	v := url.Values{}

	spath := "/contest.list" + "?" + v.Encode() + "&"
	req, err := c.newRequest(ctx, "GET", spath, options, nil)
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

//GetContestRatingChanges implements /contest.ratingChanges
func (c *Client) GetContestRatingChanges(ctx context.Context, contestID int) ([]RatingChange, error) {
	c.Logger.Println("GetContestRatingChange contestId: ", contestID)
	v := url.Values{}
	v.Add("contestId", strconv.Itoa(contestID))
	spath := "/contest.ratingChanges" + "?" + v.Encode()
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
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}

	var ratingChanges []RatingChange
	ratingChanges = resp.Result
	return ratingChanges, nil
}

//ContestStatndingsOptions represents the option of /contest.standings
type ContestStatndingsOptions struct {
	From           int
	Count          int
	Handles        []string
	Room           int
	ShowUnofficial bool
}

func (o *ContestStatndingsOptions) options() interface{} {
	if o == nil {
		return nil
	}
	type option struct {
		From           int    `url:"from"`
		Count          int    `url:"count"`
		Handles        string `url:"handles"`
		Room           int    `url:"room"`
		ShowUnofficial bool   `url:"showUnofficial"`
	}
	return &option{From: o.From, Count: o.Count, Handles: strings.Join(o.Handles, ";"), Room: o.Room, ShowUnofficial: o.ShowUnofficial}
}

//GetContestStandings implements /contest.standings
func (c *Client) GetContestStandings(ctx context.Context, contestID int, options *ContestStatndingsOptions) (*Standings, error) {
	c.Logger.Println("GetContestStandings: ", contestID, options)

	v := url.Values{}
	v.Add("contestId", strconv.Itoa(contestID))

	spath := "/contest.standings" + "?" + v.Encode() + "&"
	req, err := c.newRequest(ctx, "GET", spath, options.options(), nil)
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

//ContestStatusOptions represents the option of /contest.status
type ContestStatusOptions struct {
	From   int
	Count  int
	Handle string
}

func (o *ContestStatusOptions) options() interface{} {
	if o == nil {
		return nil
	}
	type option struct {
		From   int    `url:"from"`
		Count  int    `url:"count"`
		Handle string `url:"handle"`
	}
	return &option{From: o.From, Count: o.Count, Handle: o.Handle}
}

//GetContestStatus implements /contest.status
func (c *Client) GetContestStatus(ctx context.Context, contestID int, options *ContestStatusOptions) ([]Submission, error) {
	c.Logger.Println("GetContestStatus: ", contestID, options)

	v := url.Values{}
	v.Add("contestId", strconv.Itoa(contestID))

	//check options
	spath := "/contest.status" + "?" + v.Encode() + "&"
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
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}

	return resp.Result, nil
}
