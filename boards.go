package gochan

import (
	"encoding/json"
)

// Board is a description of a 4chan board. Board's fields are equivalent
// to JSON returned by https://a.4cdn.org/boards.json
type Board struct {
	Board           string
	Title           string
	WSBoard         bool `json:"ws_board"`
	PerPage         int  `json:"per_page"`
	Pages           int
	MaxFilesize     int `json:"max_filesize"`
	MaxWebMFilesize int `json:"max_webm_filesize"`
	MaxCommentChars int `json:"max_comment_chars"`
	MaxWebMDuration int `json:"max_webm_duration"`
	BumpLimit       int `json:"bump_limit"`
	ImageLimit      int `json:"image_limit"`
	Cooldowns       struct {
		Threads int
		Replies int
		Images  int
	}
	MetaDescription string `json:"meta_description"`
	IsArchived      bool   `json:"is_archived"`
}

type boardJSON struct {
	Boards []Board
}

// GetBoards returns a slice containing all boards. An error is returned if an
// error is encountered while attempting to acquire the list of boards.
func GetBoards() ([]Board, error) {
	resp, err := client.Get("https://a.4cdn.org/boards.json")
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(resp.Body)
	_, err = decoder.Token()
	if err != nil {
		return nil, err
	}

	var boards []Board
	err = decoder.Decode(&boards)
	if err != nil {
		return nil, err
	}

	return boards, nil
}
