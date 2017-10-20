package goforces

//ProblemResult represents codeforces ProblemResult
type ProblemResult struct {
	Points                    float64 `json:"points"`
	Penalty                   int     `json:"penalty,omitempty"`
	RejectedAttemptCount      int64   `json:"rejectedAttemptCount"`
	Type                      string  `json:"type"`
	BestSubmissionTimeSeconds int64   `json:"bestSubmissionTimeSeconds"`
}
