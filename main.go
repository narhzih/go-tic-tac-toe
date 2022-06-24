package main

import (
	"fmt"
	"github.com/narhzih/go-tic-tac-toe/pkg/player"
	"github.com/narhzih/go-tic-tac-toe/pkg/ui"
)

func main() {
	ui.PlayerHasWon()
	var gameIsOn bool
	gameIsOn = true
	playerOne := player.CreatePlayer("Player 1, enter your name: ")
	playerTwo := player.CreatePlayer("Player 2, enter your name: ")
	currentPlayer := playerOne
	nextPlayer := playerTwo

	for gameIsOn {

		// 1. Receive input from current player and display board accordingly
		var playerInput int
		ui.PrintBoard()
		fmt.Printf("%+v, make your move. (enter a number between 1-9): \n", currentPlayer.PlayerName)
		_, err := fmt.Scanln(&playerInput)
		if err != nil {
			panic(err)
		}
		ui.UpdateBoard(playerInput, currentPlayer.GameChar)
		ui.PrintBoard()

		// 2. Check for a win or tie before allowing the next player to make a move
		if ui.PlayerHasWon() {
			fmt.Printf("%+v has won the game \n", currentPlayer.PlayerName)
			gameIsOn = false
			return
		} else if ui.GameTie() {
			fmt.Println("The game ended in a Tie! Both players were fierce")
			gameIsOn = false
			return
		}

		// 3. At this point, there's neither a win nor a tie so, allow the next player to make a move
		fmt.Printf("%+v has made a move. \n", currentPlayer.PlayerName)
		store := currentPlayer
		currentPlayer = nextPlayer
		nextPlayer = store

	}
}

func currentPlayerPleasePlay(currentPlayer player.Player) {

}
