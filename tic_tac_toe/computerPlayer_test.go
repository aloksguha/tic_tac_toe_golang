package tic_tac_toe

import (
	"fmt"
	"testing"
)

func TestComputerPlayer(t *testing.T) {
	cp1 := NewComputerPlayer("Computer", "O")
	b := newBoard()

	b[0][0] = "O" ; b[0][1] = "X" ; b[0][2] = "X"
	b[1][0] = "X" ; b[1][1] = "X" ; //b[1][2] = "X"
	//b[2][1] = "X" ; // b[2][1] = "X" ; //b[2][2] = "X"
	i,j,_ := cp1.GetMove(b)
	fmt.Printf("Computed Move = %d , %d, \n", i, j)


	//b[1][1] = "O"
	//b[2][2] = "X"
	//i,j,_ = cp1.GetMove(b)
	//fmt.Printf("Computed Move = %d , %d, \n", i, j)
}