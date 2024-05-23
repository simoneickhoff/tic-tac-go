package main

import "fmt"

func main() {
	playerOne, playerTwo, err := GetPlayerData()

	if err != nil {
		fmt.Println("Incorrect input for player data")
		return
	}

	StartGame(playerOne, playerTwo)
}
