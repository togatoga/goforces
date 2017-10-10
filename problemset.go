package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Problems struct {
	Problems          []Problem           `json:"problems"`
	ProblemStatistics []ProblemStatistics `json:"problemStatistics"`
}

func (c *Client) GetProblems(ctx context.Context, tags []string) (*Problems, error) {
	c.Logger.Println("GetProblems tags: ", tags)
	v := url.Values{}
	v.Add("tags", strings.Join(tags, ";"))
	spath := "/problemset.problems" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type ProblemsResponse struct {
		Status string   `json:"status"`
		Result Problems `json:"result"`
	}
	var resp ProblemsResponse
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}
	var problems Problems
	problems = resp.Result
	return &problems, nil
}

func (c *Client) GetRecentStatus(ctx context.Context, count int) ([]Submission, error) {
	c.Logger.Println("GetRecentStatus count: ", count)
	if count <= 0 || count > 1000 {
		return nil, fmt.Errorf("count value must be between 1 and 1000: %d", count)
	}

	v := url.Values{}
	v.Add("count", strconv.Itoa(count))
	spath := "/problemset.recentStatus" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type RecentStatusResponse struct {
		Status string       `json:"status"`
		Result []Submission `json:"result"`
	}
	var resp RecentStatusResponse
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}
	var submissions []Submission
	submissions = resp.Result
	return submissions, nil
}
