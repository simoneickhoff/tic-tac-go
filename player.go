package main

import "fmt"

type Player struct {
	id     uint8
	name   string
	symbol string
	score  uint8
}

func GetPlayerData() (Player, Player, error) {
	playerOne := Player{
		id:     1,
		name:   "",
		symbol: "x",
		score:  0,
	}

	playerTwo := Player{
		id:     2,
		name:   "",
		symbol: "o",
		score:  0,
	}

	fmt.Println("Player 1 - What is your name?")
	_, errPlayerOne := fmt.Scanln(&playerOne.name)

	if errPlayerOne != nil {
		return playerOne, playerTwo, errPlayerOne
	}

	fmt.Println("Player 2 - What is your name?")
	_, errPlayerTwo := fmt.Scanln(&playerTwo.name)

	if errPlayerTwo != nil {
		return playerOne, playerTwo, errPlayerTwo
	}

	return playerOne, playerTwo, nil
}
