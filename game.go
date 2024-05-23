package main

import (
	"fmt"
	"math/rand/v2"
)

type BoardState [9]uint8

func StartGame(playerOne Player, playerTwo Player) {
	var boardState BoardState = initializeBoard()

	getCurrentPlayerId := getCurrentPlayerIdFunc(playerOne, playerTwo)

	currentPlayerId := getCurrentPlayerId()
	clearTerminal()
	printScoreBoard(playerOne, playerTwo)
	printCurrentPlayer(playerOne, playerTwo, currentPlayerId)
	printBoard(boardState, playerOne, playerTwo)
}

func initializeBoard() BoardState {
	return BoardState{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func getCurrentPlayerIdFunc(playerOne Player, playerTwo Player) func() uint8 {
	currentPlayerId := playerOne.id

	if rand.Float64() < 0.5 {
		currentPlayerId = playerTwo.id
	}

	getCurrentPlayer := func() uint8 {
		if currentPlayerId == playerOne.id {
			currentPlayerId = playerTwo.id
		} else {
			currentPlayerId = playerOne.id
		}

		return currentPlayerId
	}

	return getCurrentPlayer
}

func printScoreBoard(playerOne Player, playerTwo Player) {
	fmt.Printf(`
	------SCOREBOARD------
	%v: %v
	%v: %v
	----------------------
`, playerOne.name, playerOne.score, playerTwo.name, playerTwo.score)
}

func printCurrentPlayer(playerOne Player, playerTwo Player, currentPlayerId uint8) {
	currentPlayer := playerOne

	if playerTwo.id == currentPlayerId {
		currentPlayer = playerTwo
	}

	fmt.Printf(`
	Hey %v, it's your turn! Pick a field!
	`, currentPlayer.name)
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

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
