package goforces

type BlogEntry struct {
	OriginalLocale          string   `json:"originalLocale"`
	AllowViewHistory        bool     `json:"allowViewHistory"`
	CreationTimeSeconds     int      `json:"creationTimeSeconds"`
	Rating                  int      `json:"rating"`
	AuthorHandle            string   `json:"authorHandle"`
	ModificationTimeSeconds int      `json:"modificationTimeSeconds"`
	ID                      int      `json:"id"`
	Title                   string   `json:"title"`
	Locale                  string   `json:"locale"`
	Content                 string   `json:"content"`
	Tags                    []string `json:"tags"`
}
