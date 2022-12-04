package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Rock is X for player
	//Paper is Y for player
	//Scissor is Z for player

	//Rock (A)
	//Paper (B)
	//Scissor (C)

	shapeWin := make(map[string]int)
	shapeWin["X"] = 1
	shapeWin["Y"] = 2
	shapeWin["Z"] = 3

	playerWin := make(map[string]int)
	playerWin["A Y"] = 6 //Rock covered by Player Paper
	playerWin["A X"] = 3 //Rock tied  Player Rock
	playerWin["A Z"] = 0 //Rock smashes Player Scissor

	playerWin["B X"] = 0 //Paper covers Player rock
	playerWin["B Y"] = 3 //Paper tid Player Paper
	playerWin["B Z"] = 6 //Paper cut by Player Scissor

	playerWin["C X"] = 6 //Scissor smashed by Player Rock
	playerWin["C Y"] = 0 //Scissor cut Player Paper
	playerWin["C Z"] = 3 //Scissor ties Player Scissor

	readFile, err := os.Open("Dec02/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {
		round := fileScanner.Text()
		playerSymbol := round[2:]
		//otherSymbol := round[1:2]
		playerWinScore := playerWin[round]
		playerSymbolScore := shapeWin[playerSymbol]

		roundScore := playerWinScore + playerSymbolScore

		total += roundScore

	}

	readFile.Close()

	fmt.Println(total)
}
