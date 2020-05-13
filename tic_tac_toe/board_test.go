package tic_tac_toe

import (
	"testing"
)

func TestBoard(t *testing.T) {
	b := newBoard()
	colArray := b.columns()
	ColArrayLength := len(colArray)
	if ColArrayLength != 3 {
		t.Errorf("Array lngth should be 3, found : %d", ColArrayLength)
	}

	rowArray := b.rows()
	rowArrayLength := len(rowArray)
	if rowArrayLength != 3 {
		t.Errorf("Array lngth should be 3, found : %d", rowArrayLength)
	}

	diagonalArray := b.diagonals()
	diagonalArrayLength := len(diagonalArray)
	if diagonalArrayLength != 2 {
		t.Errorf("Array lngth should be 2, found : %d", diagonalArrayLength)
	}

	//Game is not started yet, so empty cells should be nxn
	if b.emptyCount() != ColArrayLength*rowArrayLength {
		t.Errorf("emptyCount should be multipication of ColArrayLength * rowArrayLength, found : %d", rowArrayLength)
	}

	//At start of game, board should not be won
	if b.isWon() {
		t.Errorf("At start of game, board should not be won, found : true")
	}

	//There should be no winner at start
	if b.winner() != "" {
		t.Errorf("There should be no winner at start, found : %s", b.winner())
	}

	b[0][0] = "X"
	//Game is not started yet, so empty cells should be nxn
	if b.emptyCount() != (ColArrayLength*rowArrayLength)-1 {
		t.Errorf("emptyCount should be multipication of ColArrayLength * rowArrayLength, found : %d", rowArrayLength)
	}

	if b.isWon() {
		t.Errorf("Game is not won yet")
	}

}
