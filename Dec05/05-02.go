package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var boxNum2 = 0
var moveNum2 = 0

var columns2 = make([][]string, 0)

func main() {
	readFile, err := os.Open("Dec05/input.txt")

	if err == nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	row := 0
	for fileScanner.Scan() {

		li := fileScanner.Text()

		len := len(li)
		if len > 0 {
			// Map init boxes
			if strings.Trim(li, " ")[0] == '[' {
				x := float64(len) / 4
				y := math.Ceil(x)
				colsForRow := int(y)
				columns2 = addBoxes2(colsForRow, li, columns2, row)
				row++
			}

			if strings.HasPrefix(strings.Trim(li, " "), "move") {
				moveNum2++
				moveCommand2(li)
				println("Move: " + strconv.Itoa(moveNum2))
			}
		}

		if len == 0 {
			columns2 = transpose2(columns2)
		}
	}

	printStack2(columns2)

	//CMZ
}

func printStack2(columns2 [][]string) {
	for i := 0; i < len(columns2); i++ {
		a := columns2[i]

		if len(a) > 0 {
			b := a[len(a)-1]
			print(b[1:2])
		}
	}
}

func transpose2(columns2 [][]string) [][]string {
	res := make([][]string, 0)
	for c := 0; c < len(columns2); c++ {
		col := columns2[c]
		reversed := make([]string, 0)

		for r := len(col) - 1; r >= 0; r-- {
			reversed = append(reversed, col[r])
		}

		res = append(res, reversed)
	}

	return res
}

func addBoxes2(colsForRow int, li string, columns2 [][]string, row int) [][]string {
	for i := 0; i < colsForRow; i++ {
		start := 4 * i
		end := 0

		end = start + 4

		if len(li) < end {
			end = len(li)
		}

		box := li[start:end]
		columns2 = addBox2(columns2, box, row, colsForRow, i)

	}
	return columns2
}

func addBox2(columns2 [][]string, box string, row int, totalCol int, col int) [][]string {
	toAdd := totalCol - len(columns2)
	if len(columns2) < totalCol {
		for i := 0; i < toAdd; i++ {
			columns2 = append(columns2, make([]string, 0))
		}
	} else {

	}
	column := columns2[col]
	if len(box) > 0 && box[0] == '[' {
		column = append(column, box)

	} else {
		//column = append(column, "-")
	}

	columns2[col] = column

	return columns2
}

func moveCommand2(command string) [][]string {
	mc := strings.Split(command, " ")

	count, _ := strconv.Atoi(mc[1])
	from, _ := strconv.Atoi(mc[3])
	to, _ := strconv.Atoi(mc[5])

	removeFrom := columns2[from-1]
	addTo := columns2[to-1]

	removed := make([]string, 0)

	rmLen := len(removeFrom)
	for i := rmLen - 1; i >= rmLen-count; i-- {
		back := removeFrom[i]
		removed = append(removed, back)
		removeFrom = append(removeFrom[:i])
	}

	removed = ReverseSlice(removed)
	addTo = append(addTo, removed...)

	columns2[from-1] = removeFrom
	columns2[to-1] = addTo

	return columns2
}

func ReverseSlice[T comparable](s []T) []T {
	var r []T
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
