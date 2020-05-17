package tic_tac_toe

import (
	"fmt"
)

func NewComputerPlayer(name string, mark string) *ComputerPlayer {
	return &ComputerPlayer{
		name: name,
		mark: mark,
	}
}



type ComputerPlayer struct {
	name string
	mark string
}
//
//func selectMove(b *board) {
//	winningMoves := []string
//	drawMoves := [] string
//	losingMoves := [] string
//
//	b.getLegalMoves()
//
//	for i := range b.getLegalMoves() {
//		// b.applmove => get next state of board
//	}
//
//
//}

//getMove from player
func (p *ComputerPlayer) GetMove(b *board) (int, int, error) {
	fmt.Print("computing possible move... ")
	move := p.minimax(b, p.Mark(), 1)
	return move.i, move.j, nil
}

// Converts move from 1-9 to indexes of a 2d array
func (p *ComputerPlayer) ConvertMove(input int) (int, int, error) {

	return 0, 0, nil
}

func(p *ComputerPlayer) BestResult(game Game) {
	if !game.isOver(){

	}
}


//No need of this method in computer's game
func (p *ComputerPlayer) Mark() string {
	return p.mark
}

//get players name
func (p *ComputerPlayer) Name() string {
	return p.name
}

type move struct {
	value int
	i     int
	j     int
}

func (cp *ComputerPlayer) minimax(b *board, mark string, depth int) move {
	//fmt.Println("\n----- Called with ", b , mark, depth )
	//fmt.Println("\ngetAvailablePos = ",b.getAvailablePos())
	//fmt.Print("\n")
	if b.winner() != "" || b.emptyCount() == 0 {
		m := move{}
		if b.winner() == cp.Mark() {
			m.value = 10 - depth
		} else {
			m.value = depth - 10
		}
		//fmt.Print("Game ends !! score is %d \n",)
		return m
	}

	moves := []move{}

	for _, pos := range b.getAvailablePos() {
		newBoard := b.copy()
		newBoard.String()
		i, j := pos[0], pos[1]
		newBoard[i][j] = mark

		var opponent string
		//switch opponents
		if mark == "X" {
			opponent = "O"
		} else {
			opponent = "X"
		}

		m := cp.minimax(newBoard, opponent, depth+1)
		m.i = i
		m.j = j

		moves = append(moves, m)
	}

	//fmt.Println(moves)
	//fmt.Println(len(moves))

	// maximize move value
	if mark == cp.Mark() {
		max := moves[0]
		for _, m := range moves {
			if max.value < m.value {
				max = m
			}
		}
		//fmt.Println("max : ",max)
		return max
	}

	// minimize move value
	min := moves[0]

	for _, m := range moves {
		if min.value > m.value {
			min = m
		}
	}
	//fmt.Println("min : ", min)
	return min
}

