package goforces

//Submission represents a Codeforces Submission
type Submission struct {
	ID                  int     `json:"id"`
	ContestID           int     `json:"contestId"`
	CreationTimeSeconds int64   `json:"creationTimeSeconds"`
	RelativeTimeSeconds int64   `json:"relativeTimeSeconds"`
	Problem             Problem `json:"problem"`
	Author              Party   `json:"author"`
	ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict             string  `json:"verdict"`
	Testset             string  `json:"testset"`
	PassedTestCount     int     `json:"passedTestCount"`
	TimeConsumedMillis  int     `json:"timeConsumedMillis"`
	MemoryConsumedBytes int     `json:"memoryConsumedBytes"`
}

//AC returns boolean whether submission passed all test cases
func (s *Submission) AC() bool {
	if s.Verdict == "OK" {
		return true
	}
	return false
}
