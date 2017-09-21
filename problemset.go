package goforces

type ProblemsResponse struct {
	Status string `json:"status"`
	Result struct {
		Problems          []Problem `json:"problems"`
		ProblemStatistics ProblemStatistics `json:"problemStatistics"`
	} `json:"result"`
}
