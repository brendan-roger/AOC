package main

import (
	"AOC/internal/util"
	f "fmt"
)

func main() {
	data := util.ReadLines("2025/04/input.txt")
	f.Println("Part 1:", part1(data))
	// f.Println("Part 2:", part(data, 12))
}

// 00@@.@@@@.
// @@@.@.@.@@
// @@@@@.@.@@
// @.@@@@..@.
// @@.@@@@.@@
// .@@@@@@@.@
// .@.@.@.@@@
// @.@@@.@@@@
// .@@@@@@@@.
// @.@.@@@.@.

// [
//  [0 0 1 2 2 3 4 5 6 6]
//  [1 2 4 5 6 7 9 10 12 13]
//  [2 4 7 9 11 12 15 16 19 21]
//  [3 5 9 12 15 17 20 21 25 27]
//  [4 7 11 15 19 22 26 27 32 35]
//  [4 8 13 18 23 27 32 34 39 43]
//  [4 9 14 20 25 30 35 38 44 49]
//  [5 10 16 23 29 34 40 44 51 57]
//  [5 11 18 26 33 39 46 51 59 65]
//  [6 12 20 28 36 43 51 56 65 71]
// ]


func part1(matrix []string) int {

	// can only access 4 < in 8 adjacent positions
	rows := len(matrix)
	cols := len(matrix[0])

	prefixMap := make([][]int, rows)

	for i := range prefixMap {
		prefixMap[i] = make([]int, cols)
	}

	get := func(row, col int) int {
		if row < 0 || col < 0 {
			return 0
		}
		return prefixMap[row][col]
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			top := 0
			left := 0
			topLeft := 0

			if i > 0 {
				top = prefixMap[i-1][j]
			}
			if j > 0 {
				left = prefixMap[i][j-1]
			}
			if i > 0 && j > 0 {
				topLeft = prefixMap[i-1][j-1]
			}

			var inc int

			if matrix[i][j] == '@' {
				inc = 1
			} else {
				inc = 0
			}

			prefixMap[i][j] = inc + top + left - topLeft
		}
	}

	res := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			if matrix[i][j] != '@' {
				continue
			}
			row0 := max(i-1, 0)
			row2 := min(i+1, rows-1)
			col0 := max(j-1, 0)
			col2 := min(j+1, cols-1)

			total :=
				get(row2, col2) -
				get(row0-1, col2) -
				get(row2, col0-1) +
				get(row0-1, col0-1)


			if total-1 < 4 {
				res += 1
			}

		}
	}

	return res
}
