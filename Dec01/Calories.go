package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	readFile, err := os.Open("Dec01/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	highestElf := 0

	currentElf := 0
	for fileScanner.Scan() {
		currentCount := fileScanner.Text()

		c, _ := strconv.Atoi(currentCount)

		if len(currentCount) != 0 {
			currentElf += c
		} else {
			if currentElf > highestElf {
				highestElf = currentElf
			}
			currentElf = 0
		}

	}

	readFile.Close()

	fmt.Print(highestElf)
}
