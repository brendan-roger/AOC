package main

import (
	"AOC/internal/util"
	"fmt"
	"strings"
)

func main() {
	data := util.ReadLines("2025/07/input.txt")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func part1(lines []string) int {

	queue := make(map[int]int)
	res := 0

	for idx, line := range lines {

		if idx == 0 {
			queue[strings.Index(line, "S")] = 1
			continue
		}

		next := make(map[int]int)

		for key, count := range queue {
			if line[key] == '^' {
				res += 1
				if key > 0 {
					next[key-1] += count
				}
				if key < len(line)-1 {
					next[key+1] += count
				}
			} else {
				next[key] += count
			}
			queue = next
		}
	}

	return res
}

func part2(lines []string) int {

	queue := make(map[int]int)
	res := 0

	for idx, line := range lines {

		if idx == 0 {
			queue[strings.Index(line, "S")] = 1
			continue
		}

		next := make(map[int]int)

		for key, count := range queue {

			if idx < len(line) {

				if line[key] == '^' {

					if key > 0 {
						next[key-1] += count
					}
					if key < len(line)-1 {
						next[key+1] += count
					}
				} else {
					next[key] += count
				}

				queue = next
			} else {
				res += count
			}
		}
	}

	return res
}
