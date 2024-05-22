package main

import (
	"fmt"
)

func main() {
	var boardState BoardState = initializeBoard()

	playerOne, playerTwo, err := getUserData()

	if err != nil {
		fmt.Println("Incorrect input for player data")
		return
	}

	printBoard(boardState, playerOne, playerTwo)
}

type Player struct {
	id     uint8
	name   string
	symbol string
	score  uint8
}

type BoardState [9]uint8

func getUserData() (Player, Player, error) {
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

func initializeBoard() BoardState {
	return BoardState{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func printBoard(boardState BoardState, playerOne Player, playerTwo Player) {
	boardStateCharacters := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}

	for i := 0; i < len(boardState); i++ {
		switch boardState[i] {
		case playerOne.id:
			boardStateCharacters[i] = playerOne.symbol
		case playerTwo.id:
			boardStateCharacters[i] = playerTwo.symbol
		}
	}

	fmt.Printf(`
	 %v | %v | %v
	-----------
	 %v | %v | %v
	-----------
	 %v | %v | %v
`, boardStateCharacters[0], boardStateCharacters[1], boardStateCharacters[2],
		boardStateCharacters[3], boardStateCharacters[4], boardStateCharacters[5],
		boardStateCharacters[6], boardStateCharacters[7], boardStateCharacters[8])
}
