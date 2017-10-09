package goforces

import (
	"context"
	"net/url"
	"strings"
)

type ProblemsResponse struct {
	Status string `json:"status"`
	Result struct {
		Problems          []Problem           `json:"problems"`
		ProblemStatistics []ProblemStatistics `json:"problemStatistics"`
	} `json:"result"`
}

func (c *Client) GetProblems(ctx context.Context, tags []string) (*ProblemsResponse, error) {
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
	var problemResponse ProblemsResponse
	if err := decodeBody(res, &problemResponse); err != nil {
		return nil, err
	}
	return &problemResponse, nil
}
