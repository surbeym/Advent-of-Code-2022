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

	sum := 0
	for fileScanner.Scan() {
		backpack := fileScanner.Text()
		packSize := len(backpack)

		c1Size := packSize / 2

		c1 := backpack[0:c1Size]
		c2 := backpack[c1Size:]

		packContainsTwo := make(map[rune]bool)

		for k := range c1 {
			u := rune(c1[k])
			packContainsTwo[u] = false
		}

		for k := range c2 {
			u := rune(c2[k])
			_, present := packContainsTwo[u]

			if present {
				packContainsTwo[u] = true
			}
		}

		for k := range packContainsTwo {
			if ok := !packContainsTwo[k]; ok {
				delete(packContainsTwo, k)
			}
		}

		//a == 97
		//A == 65

		for k := range packContainsTwo {
			n := int(k)

			if n > 96 {
				n = n - 96
			} else {
				n = n - 64 + 26
			}

			print(string(k) + " - ")
			println(n)

			sum = sum + n
		}

		print(sum)

	}
}
