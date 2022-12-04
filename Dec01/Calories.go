package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	readFile, err := os.Open("Dec01/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	highestElves := make(map[int]int)

	currentElf := 0
	for fileScanner.Scan() {
		currentCount := fileScanner.Text()

		c, _ := strconv.Atoi(currentCount)

		if len(currentCount) != 0 {
			i := highestElves[currentElf] + c
			highestElves[currentElf] = i
		} else {
			currentElf++
			highestElves[currentElf] = 0
		}
	}
	readFile.Close()

	keys := make([]int, 0, len(highestElves))

	for key := range highestElves {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return highestElves[keys[i]] > highestElves[keys[j]]
	})

	sum := 0
	for i := 0; i < 3; i++ {

		i2 := keys[i]
		sum += highestElves[i2]
	}

	total := sum

	fmt.Println(total)
}
