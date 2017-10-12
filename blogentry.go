package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type BlogEntry struct {
	OriginalLocale          string   `json:"originalLocale"`
	AllowViewHistory        bool     `json:"allowViewHistory"`
	CreationTimeSeconds     int      `json:"creationTimeSeconds"`
	Rating                  int      `json:"rating"`
	AuthorHandle            string   `json:"authorHandle"`
	ModificationTimeSeconds int      `json:"modificationTimeSeconds"`
	ID                      int      `json:"id"`
	Title                   string   `json:"title"`
	Locale                  string   `json:"locale"`
	Content                 string   `json:"content"`
	Tags                    []string `json:"tags"`
}

func (c *Client) GetBlogEntryComments(ctx context.Context, blogEntryId int) ([]Comment, error) {
	c.Logger.Println("GetBlogEntryComments blogEntryId: ", blogEntryId)
	v := url.Values{}
	v.Add("blogEntryId", strconv.Itoa(blogEntryId))
	spath := "/blogEntry.comments" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type EntryCommentsResponse struct {
		Status string    `json:"status"`
		Result []Comment `json:"result"`
	}
	var resp EntryCommentsResponse
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}
	var comments []Comment
	comments = resp.Result
	return comments, nil
}
