package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("Dec04/input.txt")

	if err == nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	count := 0

	for fileScanner.Scan() {
		pair := fileScanner.Text()

		p := strings.Split(pair, ",")

		c1 := sectionNumbers(p[0])
		c2 := sectionNumbers(p[1])

		all := make([]int, 0)

		all = append(all, c1...)
		all = append(all, c2...)

		overlap := dupicates(all)

		if len(c1) > len(c2) {
			if len(overlap) == len(c2) {
				count++
			}
		} else {
			if len(overlap) == len(c1) {
				count++
			}
		}
	}

	println(count)
}

func dupicates(l []int) []int {
	res := make(map[int]int)
	for i := 0; i < len(l); i++ {
		n := l[i]
		_, present := res[n]

		if present {
			res[n] = res[n] + 1
		} else {
			res[n] = 1
		}
	}

	count := make([]int, 0)
	for k, v := range res {
		if v > 1 {
			count = append(count, k)
		}
	}

	return count
}

func sectionNumbers(s string) []int {
	sl := strings.Split(s, "-")
	start, _ := strconv.Atoi(sl[0])
	end, _ := strconv.Atoi(sl[1])

	c := end - start

	res := make([]int, end-start+1)
	for i := 0; i <= c; i++ {
		res[i] = i + start
	}

	return res
}
