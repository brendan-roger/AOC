package main

import (
	"AOC/internal/util"
	f "fmt"
)

func main() {
	data := util.ReadLines("2025/03/input.txt")
	f.Println("Part 1:", part(data, 2))
	f.Println("Part 2:", part(data, 12))
}

func part(data []string, minLength int) int {
	var res int

	for _, part := range data {

		stack := []int{}

		for idx, char := range part {
			num := int(char - '0')
			rem := len(part) - idx - 1

			for len(stack) > 0 && rem+len(stack) >= minLength {
				if stack[len(stack)-1] < num {
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}

			if len(stack) < minLength {
				stack = append(stack, num)
			}

		}

		var num int

		for _, d := range stack {
			num = num*10 + d
		}
		res += num
	}

	return res
}
