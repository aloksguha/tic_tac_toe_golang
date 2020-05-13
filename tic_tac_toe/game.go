package tic_tac_toe

import (
	"fmt"
	"github.com/aloksguha/tic_tac_toe_golang/utils"
	"strconv"
)

func NewGame(p1 Player, p2 Player) *Game {
	return &Game{
		p1:       p1,
		p2:       p2,
		current:  p1,
		board:    newBoard(),
		round:    1,
		gameLogs: make([]string, 9),
	}
}

type Game struct {
	p1       Player
	p2       Player
	current  Player
	board    *board
	round    int
	gameLogs []string
}

// IsOver checks if a game is over.
func (g *Game) isOver() bool {
	return g.board.winner() != "" || g.board.emptyCount() == 0
}

//Prints game state
func (g *Game) printInfo() {
	fmt.Println(g.gameLogs)
	fmt.Println(utils.Info("Turn # " + strconv.Itoa(g.round)))
	fmt.Println(g.board)
	fmt.Println(utils.InfoBlue("Current player: ", g.current.Name()))
}

//logs activity of Players
func (g *Game) logActivity(log string) {
	g.gameLogs = append(g.gameLogs, log)
}

// starter function of Game, it continues till game is won or comes to an end
func (g *Game) Start() {
	fmt.Println(utils.Info("Welcome to Tic-Tac-Toe"))

	for !g.isOver() {
		g.printInfo()
		i, j, err := g.current.GetMove(g.board)
		if err != nil {
			fmt.Print(utils.Warn("Invalid inputs \n"))
			continue
		}
		if i < 0 || i > 2 || j < 0 || j > 2 {
			fmt.Print(utils.Warn("Invalid inputs \n"))
			continue
		}

		if g.board[i][j] != "_" {
			fmt.Print("Position occupied !!\n")
			continue
		}

		log := "\n Player " + g.current.Name() + " placed " + g.current.Mark() + " at [" + strconv.Itoa(i) + "," + strconv.Itoa(j) + "]"
		g.logActivity(log)

		//set on board
		g.board[i][j] = g.current.Mark()
		g.switchPlayer()
		g.round++
	}
	fmt.Println(g.board)
	if g.board.winner() == "" {
		fmt.Println(utils.Fatal("Game over, It's a draw"))
	} else {
		fmt.Println(utils.Fatal("Game over, Winner is ", g.board.winner()))
	}

}

// This function switches the chances of players
func (g *Game) switchPlayer() {
	if g.current == g.p1 {
		g.current = g.p2
	} else {
		g.current = g.p1
	}
}
