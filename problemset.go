package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type Problems struct {
	Problems          []Problem           `json:"problems"`
	ProblemStatistics []ProblemStatistics `json:"problemStatistics"`
}

func (c *Client) GetProblems(ctx context.Context, tags []string) (*Problems, error) {

	c.Logger.Println("GetProblems tags: ", tags)
	type ProblemsResponse struct {
		Status string   `json:"status"`
		Result Problems `json:"result"`
	}

	var resp ProblemsResponse

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
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	//check status
	if res.Status != "OK" {
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}
	var problems Problems
	return &problemResponse, nil
}
