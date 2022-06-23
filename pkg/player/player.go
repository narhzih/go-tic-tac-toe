package player

import "fmt"

type Player struct {
	PlayerName  string
	PlayerIndex uint
}
type Players struct {
	List []Player
}

var players Players

func CreatePlayer(messagePrompt string) Player {
	var player Player
	var err error
	if messagePrompt != "" {
		fmt.Printf("%+v", messagePrompt)
	} else {
		fmt.Println("Please enter your name: ")
	}
	_, err = fmt.Scanln(&player.PlayerName)
	if err != nil {
		panic("An error occurred while trying to scan player name")
	}
	player.PlayerIndex = uint(len(players.List))
	players.List = append(players.List, player)
	return player
}

func GetPlayers() []Player {
	return players.List
}

func GetPlayer(playerIndex uint) Player {
	// loop through the created players and
	// find player with the given index
	return Player{}
}
