// This is simple tic-tac-toe game to be played by two human players.
// It's starter package of program
package main

import (
	"fmt"
	"github.com/aloksguha/tic_tac_toe_golang/tic_tac_toe"
)

func main() {
	//get user input for names
	player1 := "player-1"
	fmt.Print("Enter Player-1's name (player-1) : ")
	if n, err := fmt.Scanf("%s", &player1); err != nil || n != 1 {
		player1 = "player-1"
	}

	player2 := "player-2"
	fmt.Print("Enter Player-2's name (player-2) : ")
	if n, err := fmt.Scanf("%s", &player2); err != nil || n != 1 {
		player2 = "player-2"
	}

	//degreeOfMatrix := 3
	//fmt.Print("Enter Degree of Tic-Tac-Toe matrix (3) : ")
	//if n, err := fmt.Scanf("%d", &degreeOfMatrix); err != nil || n != 1 {
	//
	//}

	//initializing human players
	hp1 := tic_tac_toe.NewHumanPlayer(player1, "X")
	hp2 := tic_tac_toe.NewHumanPlayer(player2, "O")

	//starting the game

	g := tic_tac_toe.NewGame(hp1, hp2)
	g.Start()
}
