package tic_tac_toe

import (
	"testing"
)

func TestGame(t *testing.T) {
	hp1 := NewHumanPlayer("player1", "X")
	hp2 := NewHumanPlayer("player2", "O")

	//starting the game
	g := NewGame(hp1, hp2)

	if g.isOver() {
		t.Errorf("New Game should not be in over state")
	}

	currentPlayer := g.current
	if currentPlayer.Name() != hp1.Name() {
		t.Errorf("Name attribute is not initialized properly, should be %s found : %s", hp1.Name(), currentPlayer.Name())
	}
	if currentPlayer.Mark() != hp1.Mark() {
		t.Errorf("Mark attribute is not initialized properly, should be %s found : %s", hp1.Mark(), currentPlayer.Mark())
	}

	g.switchPlayer()
	currentPlayer = g.current

	if currentPlayer.Name() != hp2.Name() {
		t.Errorf("Name attribute is not initialized properly, should be %s found : %s", hp2.Name(), currentPlayer.Name())
	}
	if currentPlayer.Mark() != hp2.Mark() {
		t.Errorf("Mark attribute is not initialized properly, should be %s found : %s", hp2.Mark(), currentPlayer.Mark())
	}

}
