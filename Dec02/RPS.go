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

	//Rock
	aWhat := make(map[string]string)
	aWhat["X"] = "Z"
	aWhat["Y"] = "X"
	aWhat["Z"] = "Y"

	//Paper
	bWhat := make(map[string]string)
	bWhat["X"] = "X"
	bWhat["Y"] = "Y"
	bWhat["Z"] = "Z"

	//Scissor
	cWhat := make(map[string]string)
	cWhat["X"] = "Y"
	cWhat["Y"] = "Z"
	cWhat["Z"] = "X"

	//X means you need to lose,
	//Y means you need to end the round in a draw, and
	//Z means you need to win. Good luck!"

	playMatrix := make(map[string]map[string]string)
	playMatrix["A"] = aWhat
	playMatrix["B"] = bWhat
	playMatrix["C"] = cWhat

	playerWin := make(map[string]int)
	playerWin["A X"] = 3 //Rock tied  Player Rock
	playerWin["A Y"] = 6 //Rock covered by Player Paper
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
		//round := "A Y"
		playerSymbol := round[2:]
		otherSymbol := round[0:1]

		toPlayMap := playMatrix[otherSymbol]
		platIt := toPlayMap[playerSymbol]

		playerWinScore := playerWin[otherSymbol+" "+platIt]
		playerSymbolScore := shapeWin[platIt]

		roundScore := playerWinScore + playerSymbolScore

		total += roundScore
		//total = roundScore

	}

	readFile.Close()

	fmt.Println(total)
}
