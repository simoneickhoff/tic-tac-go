package main

import (
	"fmt"
)

func main() {
	var boardState BoardState = initializeBoard()

	// TODO: Get these from user input
	playerOne := Player{
		id:     1,
		name:   "Alice",
		symbol: "x",
		score:  0,
	}

	playerTwo := Player{
		id:     2,
		name:   "Bob",
		symbol: "o",
		score:  0,
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
