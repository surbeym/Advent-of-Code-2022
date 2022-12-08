package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("Dec03/input.txt")

	if err == nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

	}
}
