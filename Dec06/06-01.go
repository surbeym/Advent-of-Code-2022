package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("Dec06/input.txt")

	if err == nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		for i := 0; i < len(line)-4; i++ {
			char1 := line[i+0 : i+1]
			char2 := line[i+1 : i+2]
			char3 := line[i+2 : i+3]
			char4 := line[i+3 : i+4]
			//char5 := line[i+4 : i+5]
			//println(char1)
			pack4 := mapChars(char1, char2, char3, char4)
			//pack4 = mapChars(pack4, char5)
			println(pack4)
		}
	}
}

func mapChars(char1 string, char2 string, char3 string, char4 string) int {
	pack4 := make(map[string]int)

	pack4 = add(char1, pack4)

	if check(char1, pack4) {
		return 1
	}

	pack4 = add(char2, pack4)

	if check(char2, pack4) {
		return 2
	}
	pack4 = add(char3, pack4)

	if check(char3, pack4) {
		return 3
	}
	pack4 = add(char4, pack4)

	if check(char4, pack4) {
		return 4
	}

	return 0
}

func check(char string, pac map[string]int) bool {
	v, _ := pac[char]

	if v > 1 {
		return true
	} else {
		return false
	}
}

func add(char string, pack4 map[string]int) map[string]int {
	_, v := pack4[char]

	if v {
		pack4[char] = pack4[char] + 1
	} else {
		pack4[char] = 1
	}

	return pack4
}
