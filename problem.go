package goforces

import "fmt"

//Problem represents Codeforces problem
type Problem struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Points    float32  `json:"points"`
	Tags      []string `json:"tags"`
}

//ProblemURL returns problem's url
func (p Problem) ProblemURL() string {
	return fmt.Sprintf("http://codeforces.com/contest/%d/problem/%s", p.ContestID, p.Index)
}
