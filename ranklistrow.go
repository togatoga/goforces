package goforces

//RanklistRow represents Codeforces RanklistRow
type RanklistRow struct {
	Party                     Party           `json:"party"`
	Rank                      int64           `json:"rank"`
	Points                    float64         `json:"points"`
	Penalty                   int64           `json:"penalty"`
	SuccessfulHackCount       int64           `json:"successfulHackCount"`
	UnsuccessfulHackCount     int64           `json:"unsuccessfulHackCount"`
	ProblemResults            []ProblemResult `json:"problemResults"`
	LastSubmissionTimeSeconds int64           `json:"lastSubmissionTimeSeconds"`
}
