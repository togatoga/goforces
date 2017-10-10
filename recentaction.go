package goforces

type RecentAction struct {
	TimeSeconds int       `json:"timeSeconds"`
	BlogEntry   BlogEntry `json:"blogEntry"`
	Comment     Comment   `json:"comment,omitempty"`
}
