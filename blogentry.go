package goforces

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

//BlogEntry represents a Codeforces BlogEntry
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

//GetBlogEntryComments implements /blogEntry.comments
func (c *Client) GetBlogEntryComments(ctx context.Context, blogEntryID int) ([]Comment, error) {
	c.Logger.Println("GetBlogEntryComments : ", blogEntryID)
	v := url.Values{}
	v.Add("blogEntryId", strconv.Itoa(blogEntryID))
	spath := "/blogEntry.comments" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Status string    `json:"status"`
		Result []Comment `json:"result"`
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

//GetBlogEntryView implements /blogEntry.view
func (c *Client) GetBlogEntryView(ctx context.Context, blogEntryID int) (*BlogEntry, error) {
	c.Logger.Println("GetBlogEntryView : ", blogEntryID)
	v := url.Values{}
	v.Add("blogEntryId", strconv.Itoa(blogEntryID))
	spath := "/blogEntry.view" + "?" + v.Encode()
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	type EntryViewResponse struct {
		Status string    `json:"status"`
		Result BlogEntry `json:"result"`
	}
	var resp EntryViewResponse
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	//check status
	if resp.Status != "OK" {
		return nil, fmt.Errorf("Status Error: %s", res.Status)
	}

	return &resp.Result, nil
}
