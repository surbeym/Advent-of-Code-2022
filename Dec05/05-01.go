package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("Dec05/input.txt")

	if err == nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	columns := make([][]string, 0)

	row := 0
	for fileScanner.Scan() {

		li := fileScanner.Text()

		if len(li) > 0 {
			// Map init boxes
			if strings.Trim(li, " ")[0] == '[' {
				colsForRow := (len(li)) / 3
				columns = addBoxes(colsForRow, li, columns, row)
				row++
			}

			if strings.HasPrefix(strings.Trim(li, " "), "move") {
				columns = moveCommand(li, columns)
			}
		}

		if len(li) == 0 {
			columns = transpose(columns)
		}
	}

	for i := 0; i < len(columns); i++ {
		a := columns[i]

		if len(a) > 0 {
			print(a[len(a)-1])
		}
	}

	//CMZ
}

func transpose(columns [][]string) [][]string {
	res := make([][]string, 0)
	for c := 0; c < len(columns); c++ {
		col := columns[c]
		reversed := make([]string, 0)

		for r := len(col) - 1; r >= 0; r-- {
			reversed = append(reversed, col[r])
		}

		res = append(res, reversed)
	}

	return res
}

func addBoxes(colsForRow int, li string, columns [][]string, row int) [][]string {
	for i := 0; i < colsForRow; i++ {
		start := 4 * i
		end := 0

		if i+1 < colsForRow {
			end = 4 + 3*i
		} else {
			end = 5 + 3*i
		}

		box := li[start:end]
		columns = addBox(columns, box, row, colsForRow, i)

	}
	return columns
}

func addBox(columns [][]string, box string, row int, totalCol int, col int) [][]string {
	toAdd := totalCol - len(columns)
	if len(columns) < totalCol {
		for i := 0; i < toAdd; i++ {
			columns = append(columns, make([]string, 0))
		}
	} else {

	}
	column := columns[col]
	if box[0] == '[' {
		column = append(column, box)

	} else {
		//column = append(column, "-")
	}

	columns[col] = column

	return columns
}

func moveCommand(command string, columns [][]string) [][]string {
	mc := strings.Split(command, " ")

	count, _ := strconv.Atoi(mc[1])
	from, _ := strconv.Atoi(mc[3])
	to, _ := strconv.Atoi(mc[5])

	removeFrom := columns[from-1]
	addTo := columns[to-1]

	removed := make([]string, 0)

	for i := len(removeFrom) - 1; i > count-1; i-- {
		back := removeFrom[i]
		removed = append(removed, back)
		removeFrom = append(removeFrom[:i])
	}

	for i := 0; i < count; i++ {
		addTo = append(addTo, removed...)
	}

	columns[from-1] = removeFrom
	columns[to-1] = addTo

	return columns
}
