package goforces

type Contest struct {
	DurationSeconds     int64  `json:"durationSeconds"`
	Frozen              bool   `json:"frozen"`
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	Phase               string `json:"phase"`
	RelativeTimeSeconds int64  `json:"relativeTimeSeconds"`
	StartTimeSeconds    int64  `json:"startTimeSeconds"`
	Type                string `json:"type"`
}
