package main

import (
	"AOC/internal/util"
	f "fmt"
)

func main() {
	data := util.ReadLines("2025/04/input.txt")
	f.Println("Part 1:", part1(data))
	f.Println("Part 2:", part2(data))
}

func getPrefixMap(matrix []string) [][]int {

	rows := len(matrix)
	cols := len(matrix[0])

	prefixMap := make([][]int, rows)

	for i := range prefixMap {
		prefixMap[i] = make([]int, cols)
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

	return prefixMap
}

func part1(matrix []string) int {

	// can only access 4 < in 8 adjacent positions
	rows := len(matrix)
	cols := len(matrix[0])
	prefixMap := getPrefixMap(matrix)

	get := func(row, col int) int {
		if row < 0 || col < 0 {
			return 0
		}
		return prefixMap[row][col]
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

			// prefix sum for 3x3 - 1 for the center cell
			count :=
				get(row2, col2) -
					get(row0-1, col2) -
					get(row2, col0-1) +
					get(row0-1, col0-1) -
					1

			if count < 4 {
				res += 1
			}

		}
	}

	return res
}

type Cell struct {
	row int
	col int
}

func part2(matrix []string) int {

	rows := len(matrix)
	cols := len(matrix[0])
	prefixMap := getPrefixMap(matrix)
	var queue = []Cell{}

	get := func(row, col int) int {
		if row < 0 || col < 0 {
			return 0
		}
		return prefixMap[row][col]
	}

	neighborMap := make([][]int, rows)

	for i := range prefixMap {
		neighborMap[i] = make([]int, cols)
	}

	// Make a matrix with the number of "@" neighbours a cell has
	// If less that 4 add to queue
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			if matrix[i][j] != '@' {
				continue
			}

			row0 := max(i-1, 0)
			row2 := min(i+1, rows-1)
			col0 := max(j-1, 0)
			col2 := min(j+1, cols-1)

			count :=
				get(row2, col2) -
					get(row0-1, col2) -
					get(row2, col0-1) +
					get(row0-1, col0-1) -
					1

			if count < 4 {
				queue = append(queue, Cell{row: i, col: j})
				neighborMap[i][j] = -1
			} else {
				neighborMap[i][j] = count
			}

		}
	}

	res := 0

	// Iterate over queue and decrement neighbours.
	// If after decrementing neighbour is less than 4 enqueue cell
	for len(queue) != 0 {

		cell := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		operations := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, 0}, {1, -1}, {1, 1}}

		for _, op := range operations {

			row := cell.row + op[0]
			col := cell.col + op[1]

			if row < 0 || col < 0 || row >= rows || col >= cols || matrix[row][col] != '@' {
				continue
			}

			next := neighborMap[row][col]

			if next == 4 {
				queue = append(queue, Cell{row, col})
				neighborMap[row][col] = -1
			}

			if next > 4 {
				neighborMap[row][col] -= 1
			}

		}

		res += 1
	}

	return res
}
