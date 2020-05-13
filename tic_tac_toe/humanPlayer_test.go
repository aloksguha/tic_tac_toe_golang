package tic_tac_toe

import (
	"testing"
)

func TestHuman(t *testing.T) {
	name := "Alok"
	mark := "X"
	humanPlayer := NewHumanPlayer(name, mark)

	if humanPlayer.Name() != name {
		t.Errorf("Name attribute is not initialized properly, should be %s found : %s", name, humanPlayer.Name())
	}

	if humanPlayer.Mark() != mark {
		t.Errorf("Mark attribute is not initialized properly, should be %s found : %s", mark, humanPlayer.Mark())
	}

	rowIndex, colIndex, err := humanPlayer.ConvertMove(1)
	if rowIndex != 0 || colIndex != 0 || err != nil {
		t.Errorf("rowIndex should be %d, got %d", 0, rowIndex)
		t.Errorf("colIndex should be %d, got %d", 0, colIndex)
	}

	rowIndex, colIndex, err = humanPlayer.ConvertMove(2)
	if rowIndex != 0 || colIndex != 1 || err != nil {
		t.Errorf("rowIndex should be %d, got %d", 0, rowIndex)
		t.Errorf("colIndex should be %d, got %d", 1, colIndex)
	}

	rowIndex, colIndex, err = humanPlayer.ConvertMove(3)
	if rowIndex != 0 || colIndex != 2 || err != nil {
		t.Errorf("rowIndex should be %d, got %d", 0, rowIndex)
		t.Errorf("colIndex should be %d, got %d", 2, colIndex)
	}

	rowIndex, colIndex, err = humanPlayer.ConvertMove(4)
	if rowIndex != 1 || colIndex != 0 || err != nil {
		t.Errorf("rowIndex should be %d, got %d", 1, rowIndex)
		t.Errorf("colIndex should be %d, got %d", 0, colIndex)
	}

	rowIndex, colIndex, err = humanPlayer.ConvertMove(5)
	if rowIndex != 1 || colIndex != 1 || err != nil {
		t.Errorf("rowIndex should be %d, got %d", 1, rowIndex)
		t.Errorf("colIndex should be %d, got %d", 1, colIndex)
	}

	rowIndex, colIndex, err = humanPlayer.ConvertMove(6)
	if rowIndex != 1 || colIndex != 2 || err != nil {
		t.Errorf("rowIndex should be %d, got %d", 1, rowIndex)
		t.Errorf("colIndex should be %d, got %d", 2, colIndex)
	}

	rowIndex, colIndex, err = humanPlayer.ConvertMove(7)
	if rowIndex != 2 || colIndex != 0 || err != nil {
		t.Errorf("rowIndex should be %d, got %d", 2, rowIndex)
		t.Errorf("colIndex should be %d, got %d", 0, colIndex)
	}

	rowIndex, colIndex, err = humanPlayer.ConvertMove(8)
	if rowIndex != 2 || colIndex != 1 || err != nil {
		t.Errorf("rowIndex should be %d, got %d", 2, rowIndex)
		t.Errorf("colIndex should be %d, got %d", 1, colIndex)
	}

	rowIndex, colIndex, err = humanPlayer.ConvertMove(9)
	if rowIndex != 2 || colIndex != 2 || err != nil {
		t.Errorf("rowIndex should be %d, got %d", 2, rowIndex)
		t.Errorf("colIndex should be %d, got %d", 2, colIndex)
	}

}
