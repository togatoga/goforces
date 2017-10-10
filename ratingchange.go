package goforces

type RatingChange struct {
	ContestID               int64  `json:"contestId"`
	ContestName             string `json:"contestName"`
	Handle                  string `json:"handle"`
	NewRating               int64  `json:"newRating"`
	OldRating               int64  `json:"oldRating"`
	Rank                    int64  `json:"rank"`
	RatingUpdateTimeSeconds int64  `json:"ratingUpdateTimeSeconds"`
}
