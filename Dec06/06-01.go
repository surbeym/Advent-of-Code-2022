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
			chars := line
			if len(line) > i+4 {
				chars = line[i : i+4]
				d := checkDups(chars)
				if d > -1 {
					print(chars + ": ")
					println(d + 4)
					break
				}
			}
		}
	}
}

func checkDups(chars string) int {
	hodl := make(map[string]int)

	res := -1

	for i := 0; i < len(chars); i++ {
		ch := chars[i : i+1]
		_, v := hodl[ch]

		if v {
			hodl[ch] = hodl[ch] + 1
			res = -1
			break
		} else {
			hodl[ch] = 1
			res = i
		}
	}

	return res
}
