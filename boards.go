package gochan

// Board is a description of a 4chan board. Board's fields are equivalent
// to JSON returned by https://a.4cdn.org/boards.json
type Board struct {
	Board           string
	Title           string
	WSBoard         bool
	PerPage         int
	Pages           int
	MaxFilesize     int
	MaxWebMFilesize int
	MaxCommentChars int
	MaxWebMDuration int
	BumpLimit       int
	ImageLimit      int
	Cooldowns       struct {
		Threads int
		Replies int
		Images  int
	}
	MetaDescription string
	IsArchived      bool
}
