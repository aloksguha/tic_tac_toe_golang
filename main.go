// This is simple tic-tac-toe game to be played by two human players.
// It's starter package of program
package main

import (
	"fmt"
	"github.com/aloksguha/tic_tac_toe_golang/tic_tac_toe"
)

func main() {

	gameMode := 1
	fmt.Print("1) Enter 1 for Human Vs. Human \n2) Enter 2 for Human Vs.Computer \n Enter Game Mode (1) : ")
	if n, err := fmt.Scanf("%d", &gameMode); err != nil || n != 1 {
		gameMode = 1
	}
	fmt.Print(gameMode)


	//get user input for names
	player1 := "player-1"
	fmt.Print("Enter Player-1's name (player-1) : ")
	if n, err := fmt.Scanf("%s", &player1); err != nil || n != 1 {
		player1 = "player-1"
	}

	player2 := "player-2"
	if gameMode == 1 {
		fmt.Print("Enter Player-2's name (player-2) : ")
		if n, err := fmt.Scanf("%s", &player2); err != nil || n != 1 {
			player2 = "player-2"
		}
	}


	//degreeOfMatrix := 3
	//fmt.Print("Enter Degree of Tic-Tac-Toe matrix (3) : ")
	//if n, err := fmt.Scanf("%d", &degreeOfMatrix); err != nil || n != 1 {
	//
	//}

	//initializing human players
	p1 := tic_tac_toe.NewHumanPlayer(player1, "X")

	if gameMode == 1 {
		p2 := tic_tac_toe.NewHumanPlayer(player2, "O")
		g := tic_tac_toe.NewGame(p1, p2)
		g.Start()
	} else {
		p2 := tic_tac_toe.NewComputerPlayer("Computer", "O")
		g := tic_tac_toe.NewGame(p1, p2)
		g.Start()
	}
}
