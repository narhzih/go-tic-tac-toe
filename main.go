package main

import (
	"fmt"
	"github.com/narhzih/go-tic-tac-toe/pkg/player"
	"github.com/narhzih/go-tic-tac-toe/pkg/ui"
)

func main() {

	//var gameIsOn bool
	//gameIsOn = true
	playerOne := player.CreatePlayer("Player 1, enter your name: ")
	playerTwo := player.CreatePlayer("Player 2, enter your name: ")
	currentPlayer := playerOne
	nextPlayer := playerTwo

	for i := 0; i < 6; i++ {
		var playerInput int
		ui.PrintBoard()
		fmt.Printf("%+v, make your move. (enter a number between 1-9): \n", currentPlayer.PlayerName)
		_, err := fmt.Scanln(&playerInput)
		if err != nil {
			panic(err)
		}
		ui.UpdateBoard(playerInput, currentPlayer.GameChar)
		ui.PrintBoard()
		fmt.Printf("%+v has made a move. \n", currentPlayer.PlayerName)
		store := currentPlayer
		currentPlayer = nextPlayer
		nextPlayer = store
		//gameIsOn = false
		// 3. Display the Board
	}
}

func currentPlayerPleasePlay(currentPlayer player.Player) {

}
