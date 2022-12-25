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

	team := make(map[string]int)
	member := 1

	teams := make(map[int]string)
	for fileScanner.Scan() {
		backpack1 := fileScanner.Text()

		bpAlready := make(map[string]int)

		for i := 0; i < len(backpack1); i++ {

			charc := string(backpack1[i])
			_, currentBp := bpAlready[charc]

			if !currentBp {
				_, present := team[charc]

				if present {
					team[charc] = team[charc] + 1
				} else if member%3 == 1 {
					team[charc] = 1
				}

				bpAlready[charc] = 1
			}
		}
		if member%3 == 0 {
			teamNumber := member / 3

			for r, i := range team {
				if i == 3 {
					teams[teamNumber] = r
				}
			}

			team = make(map[string]int)
		}

		member++
	}

	sum := 0

	for _, r := range teams {

		n := int(rune(r[0]))

		if n > 96 {
			n = n - 96
		} else {
			n = n - 64 + 26
		}

		sum = sum + n
	}

	println(sum)
}
