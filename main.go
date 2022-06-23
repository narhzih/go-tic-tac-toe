package main

import (
	"fmt"
	"github.com/narhzih/go-tic-tac-toe/pkg/player"
	"github.com/narhzih/go-tic-tac-toe/pkg/ui"
)

func main() {
	//cmd := os.Args[1]
	//fmt.Printf("The value for cmd is -> %+v", cmd)
	var gameIsOn bool
	gameIsOn = true
	for gameIsOn {
		// 1. Create participating players
		playerOne := player.CreatePlayer("Player 1, enter your name: ")
		playerTwo := player.CreatePlayer("Player 2, enter your name: ")

		// 2. Display participating players information
		fmt.Printf("The match is between %+v and %+v \n", playerOne.PlayerName, playerTwo.PlayerName)
		fmt.Println("Let the game begin!!!")
		ui.PrintBoard()
		gameIsOn = false
		// 3. Display the Board
	}
}
