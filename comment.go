package goforces

//Comment represents a Codeforces comment
type Comment struct {
	ID                  int    `json:"id"`
	CreationTimeSeconds int    `json:"creationTimeSeconds"`
	CommentatorHandle   string `json:"commentatorHandle"`
	Locale              string `json:"locale"`
	Text                string `json:"text"`
	Rating              int    `json:"rating"`
	ParentCommentID     int    `json:"parentCommentId,omitempty"`
}
