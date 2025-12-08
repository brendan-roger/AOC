package main

import (
	"AOC/internal/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := util.ReadLines("2025/06/input.txt")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func part1(lines []string) int {

	res := 0
	length := len(lines)
	lastLine := lines[length-1]
	operators := strings.Fields(lastLine)

	var multiMap []int

	for _, operator := range operators {
		if operator == "*" {
			multiMap = append(multiMap, 1)
		}
	}

	for idx, line := range lines {

		if idx == length-1 {
			continue
		}

		numbers := strings.Fields(line)
		multiIdx := 0

		for numIdx, number := range numbers {

			if operators[numIdx] == "+" {
				inc, _ := strconv.Atoi(number)
				res += inc

			} else {
				inc, _ := strconv.Atoi(number)
				multiMap[multiIdx] = multiMap[multiIdx] * inc
				multiIdx += 1
			}
		}
	}

	for _, number := range multiMap {
		res += number
	}

	return res
}

type OpIndex struct {
	index    int
	operator rune
}

func part2(lines []string) int {

	res := 0
	var opMap []OpIndex
	var numberMap []int

	for lIdx, line := range lines {

		for cIdx, r := range line {

			// operators
			if lIdx == len(lines)-1 {
				if r != ' ' {
					opMap = append(opMap, OpIndex{cIdx, r})
				}
				continue
			}

			if cIdx == len(numberMap) {
				numberMap = append(numberMap, -1)
			}

			if r == ' ' {
				continue
			}

			number := int(r - '0')

			if numberMap[cIdx] == -1 {
				numberMap[cIdx] = number
			} else {
				numberMap[cIdx] = numberMap[cIdx]*10 + number
			}

		}
	}

	for _, op := range opMap {

		index := op.index
		operation := op.operator
		acc := 0

		if operation == '*' {
			acc += 1
		}

		for index < len(numberMap) && numberMap[index] != -1 {

			if operation == '*' {
				acc *= numberMap[index]
			} else {
				acc += numberMap[index]
			}

			index += 1
		}

		res += acc
	}

	return res
}
