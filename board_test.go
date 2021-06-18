package menace

import "testing"

func Test_BoardMatch(t *testing.T) {
	t.Run("Test Matching Boards", func(t *testing.T) {
		var Board1 = NewBoard(X, X, X, O, EMPTY, O, X, X, X)
		var Board2 = NewBoard(X, X, X, O, EMPTY, O, X, X, X)
		if !BoardMatch(Board1, Board2) {
			t.Errorf("got %v want %v", false, true)
		}

	})
	t.Run("Test Rotated Boards", func(t *testing.T) {
		var Board1 = NewBoard(X, X, X, O, EMPTY, O, X, X, X)
		var Board2 = NewBoard(X, O, X, X, EMPTY, X, X, O, X)
		if !BoardMatch(Board1, Board2) {
			t.Errorf("got %v want %v", false, true)
		}
	})
	t.Run("Test Rotated Boards", func(t *testing.T) {
		var Board1 = NewBoard(X, X, X, O, EMPTY, O, X, X, X)
		var Board2 = NewBoard(O, X, X, X, EMPTY, X, O, X, X)
		if BoardMatch(Board1, Board2) {
			t.Errorf("got %v want %v", true, false)
		}
	})
	t.Run("Test Mirrored Boards", func(t *testing.T) {
		var Board1 = NewBoard(
			X, X, O,
			X, EMPTY, O,
			X, O, EMPTY,
		)
		var Board2 = NewBoard(
			EMPTY, O, O,
			O, EMPTY, X,
			X, X, X)
		if !BoardMatch(Board1, Board2) {
			t.Errorf("got %v want %v", false, true)
		}
	})
}

func Test_BoardGetID(t *testing.T) {

}
