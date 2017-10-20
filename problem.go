package goforces

//Problem represents Codeforces problem
type Problem struct {
	ContestID int64    `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Points    float32  `json:"points"`
	Tags      []string `json:"tags"`
}
