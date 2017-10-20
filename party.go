package goforces

//Party represents Codeforces Party
type Party struct {
	ContestID        int      `json:"contestId"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	Ghost            bool     `json:"ghost"`
	StartTimeSeconds int      `json:"startTimeSeconds"`
}
