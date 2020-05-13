package tic_tac_toe

import (
	"fmt"
	"math"
)

func NewHumanPlayer(name string, mark string) *HumanPlayer {
	return &HumanPlayer{
		name: name,
		mark: mark,
	}
}

type Player interface {
	GetMove(*board) (int, int, error)
	ConvertMove(int) (int, int, error)
	Mark() string
	Name() string
}

type HumanPlayer struct {
	name string
	mark string
}

//getMove from player
func (p *HumanPlayer) GetMove(b *board) (int, int, error) {
	//return 0, 0, nil
	fmt.Print("Enter Position [1 - 9]: ")
	var input, i, j int
	if n, err := fmt.Scanf("%d", &input); err != nil || n != 1 {
		return 0, 0, err
	}
	//fmt.Print("User input : ", input)
	i, j, err := p.ConvertMove(input)
	if err != nil {
		return 0, 0, err
	}
	return i, j, nil
}

// Converts move from 1-9 to indexes of a 2d array
func (p *HumanPlayer) ConvertMove(input int) (int, int, error) {
	degreeOfMatrix := 3
	rowIndex := 0
	colIndex := 0

	//matrix always be n x n
	if input > int(math.Pow(float64(int(degreeOfMatrix)), 2)) {
		return rowIndex, colIndex, fmt.Errorf("error : Invalid input")
	}

	if input < degreeOfMatrix {
		colIndex = input%degreeOfMatrix - 1
	} else if input == degreeOfMatrix {
		colIndex = degreeOfMatrix - 1
	} else if input > degreeOfMatrix {
		mod := input % degreeOfMatrix
		if mod == 0 {
			rowIndex = (input / degreeOfMatrix) - 1
			colIndex = degreeOfMatrix - 1
		} else {
			rowIndex = (input - mod) / degreeOfMatrix
			colIndex = input%degreeOfMatrix - 1
		}
	}
	return rowIndex, colIndex, nil
}

//get player's mark
func (p *HumanPlayer) Mark() string {
	return p.mark
}

//get players name
func (p *HumanPlayer) Name() string {
	return p.name
}
