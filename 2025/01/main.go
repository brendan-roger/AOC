package main

import (
	"fmt"
	"strconv"
	"AOC/internal/util"
)

func main() {
	data := util.ReadLines("2025/01/input.txt")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func part1(data []string) int {
	res := 0
	currIndex := 50

	for _, line := range data {

		dir := line[:1]
		num, _ := strconv.Atoi(line[1:])

		if dir == "L" {
			currIndex -= num
		} else {
			currIndex += num
		}

		currIndex %= 100

		if(currIndex == 0){
			res += 1
		}

	}
	return res
}

func part2(data []string) int {
	res := 0
	currIndex := 50

    for _, line := range data {
        dir := line[:1]
        num, _ := strconv.Atoi(line[1:])


        if dir == "R" {
			val := currIndex + num
            res += val / 100
            currIndex = val % 100

        } else {

			dist := (100 - currIndex) % 100
			res += (dist + num) / 100

			currIndex = ((currIndex - num) % 100 + 100) % 100
        }
    }

	return res
}
