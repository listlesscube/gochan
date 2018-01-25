package gochan_test

import (
	"testing"

	"github.com/listlesscube/gochan"
)

func TestGetBoards_MainBoardsExist(t *testing.T) {
	boardNames := []string{
		"a", "b", "c", "d", "e", "f", "g", "gif", "h", "hr", "k", "m", "o", "p",
		"r", "s", "t", "u", "v", "vg", "vr", "w", "wg",
	}

	boards, err := gochan.GetBoards()
	if err != nil {
		t.Fatalf("GetBoards() returned the error: %s", err)
	}

	boardsAvailable := make(map[string]bool)
	for _, board := range boards {
		boardsAvailable[board.Board] = true
	}

	for _, boardName := range boardNames {
		if !boardsAvailable[boardName] {
			t.Errorf("GetBoards() did not return board /%s/", boardName)
		}
	}
}

func TestGetBoards_Populated(t *testing.T) {
	boards, err := gochan.GetBoards()
	if err != nil {
		t.Fatalf("GetBoards() returned the error: %s", err)
	}

	if len(boards) == 0 {
		t.Fatal("GetBoards() returned an empty slice")
	}

	for _, board := range boards {
		if board.Board == "" {
			t.Error("GetBoards() returned a board with an empty board name")
		}
		if board.Title == "" {
			t.Error("GetBoards() returned a board with an empty title")
		}
		if board.PerPage < 1 {
			t.Error("GetBoards() returned a board with non-positive threads per page")
		}
		if board.Pages < 1 {
			t.Error("GetBoards() returned a board with a non-positive number pages")
		}
		if board.MaxFilesize < 1 {
			t.Error("GetBoards() returned a board with a non-positive max filesize")
		}
		if board.MaxWebMFilesize < 1 {
			t.Error("GetBoards() returned a board with a non-positive max WebM filesize")
		}
		if board.MaxCommentChars < 1 {
			t.Error("GetBoards() returned a board with a non-positive comment character limit")
		}
		if board.MaxWebMDuration < 1 {
			t.Error("GetBoards() returned a board with a non-positive max WebM duration")
		}
		if board.BumpLimit < 1 {
			t.Error("GetBoards() returned a board with a non-positive bump limit")
		}
		if board.ImageLimit < 1 {
			t.Error("GetBoards() returned a board with a non-positive image limit")
		}
		if board.Cooldowns.Threads < 1 {
			t.Error("GetBoards() returned a board with a non-positive thread cooldown")
		}
		if board.Cooldowns.Replies < 1 {
			t.Error("GetBoards() returned a board with a non-positive reply cooldown")
		}
		if board.Cooldowns.Images < 1 {
			t.Error("GetBoards() returned a board with a non-positive image cooldown")
		}
		if board.MetaDescription == "" {
			t.Error("GetBoards() returned a board with an empty description")
		}
	}
}
