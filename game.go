package main

import "fmt"

type BoardState [9]uint8

func StartGame(playerOne Player, playerTwo Player) {
	var boardState BoardState = initializeBoard()

	clearTerminal()
	printScoreBoard(playerOne, playerTwo)
	printBoard(boardState, playerOne, playerTwo)
}

func initializeBoard() BoardState {
	return BoardState{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func printScoreBoard(playerOne Player, playerTwo Player) {
	fmt.Printf(`
	------SCOREBOARD------
	%v: %v
	%v: %v
	----------------------
`, playerOne.name, playerOne.score, playerTwo.name, playerTwo.score)
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
