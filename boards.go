package gochan

// Board is a description of a 4chan board. Board's fields are equivalent
// to JSON returned by https://a.4chan.org/boards.json
type Board struct {
	board           string
	title           string
	wsBoard         bool
	perPage         int
	pages           int
	maxFilesize     int
	maxWebmFilesize int
	maxCommentChars int
	maxWebmDuration int
	bumpLimit       int
	imageLimit      int
	cooldowns       struct {
		threads int
		replies int
		images  int
	}
	metaDescription string
	isArchived      bool
}
