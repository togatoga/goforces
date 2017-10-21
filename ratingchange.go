package goforces

//RatingChange represents Codeforces RatingChange
type RatingChange struct {
	ContestID               int64  `json:"contestId"`
	ContestName             string `json:"contestName"`
	Handle                  string `json:"handle"`
	NewRating               int    `json:"newRating"`
	OldRating               int    `json:"oldRating"`
	Rank                    int64  `json:"rank"`
	RatingUpdateTimeSeconds int64  `json:"ratingUpdateTimeSeconds"`
}

//RatingDiff returns NewRating - OldRating
func (r RatingChange) RatingDiff() int {
	return r.NewRating - r.OldRating
}
