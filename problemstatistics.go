package goforces

//ProblemStatistics represents Codeforces ProblemStatistics
type ProblemStatistics struct {
	ContestID   int    `json:"contestId"`
	Index       string `json:"index"`
	SolvedCount int    `json:"solvedCount"`
}
