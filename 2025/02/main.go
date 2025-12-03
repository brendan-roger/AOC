package main

import (
	"AOC/internal/util"
	f "fmt"
	"math"
	"sort"
	str "strings"
)

func main() {
	data := util.ReadFile("2025/02/input.txt")
	f.Println("Part 1:", part1(data))
	f.Println("Part 2:", part2(data))
}

func parseRange(data string) ([][]int64, int64) {
	var ranges [][]int64
	var max int64

	for _, part := range str.Split(data, ",") {

		var lower, upper int64

		f.Sscanf(part, "%d-%d", &lower, &upper)

		ranges = append(ranges, []int64{lower, upper})

		if upper > max {
			max = upper
		}
	}

	return ranges, max

}

func getPossibleValues1(max int64) []int64 {

	maxDigits := int(math.Log10(float64(max))) + 1

	var res []int64

	// Only even digit lengths are possible
	// Get even ranges and cycle through
	for i := 1; 2*i <= maxDigits; i++ {
		inc := int64(math.Pow10(i))

		start := int64(math.Pow10(i - 1))
		end := inc - 1

		for j := start; j <= end; j++ {
			val := j*inc + j

			if val > max {
				break
			}
			res = append(res, val)
		}
	}

	return res
}

func part1(data string) int64 {

	ranges, max := parseRange(data)

	values := getPossibleValues1(max)

	var res int64

	prefixMap := make([]int64, len(values)+1)

	for i, v := range values {
		prefixMap[i+1] = prefixMap[i] + v
	}

	for _, part := range ranges {
		i := util.LowerBound(values, part[0])
		j := util.UpperBound(values, part[1])
		res += prefixMap[j] - prefixMap[i]
	}

	return res
}

func getPossibleValues2(max int64) []int64 {

	maxDigits := int(math.Log10(float64(max))) + 1
	set := make(map[int64]struct{})

	var res []int64

	for i := 1; i <= maxDigits; i++ {
		inc := int64(math.Pow10(i))

		start := int64(math.Pow10(i - 1))
		end := inc - 1

		for t := 2; i*t <= maxDigits; t++ {

			factor := int64(0)
			for rep := 0; rep < t; rep++ {
				factor = factor*inc + 1
			}

			for p := start; p <= end; p++ {
				val := p * factor

				if val > max {
					break
				}

				_, found := set[val]

				if found {
					continue
				}

				set[val] = struct{}{}
				res = append(res, val)
			}

		}

	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })

	return res
}

func part2(data string) int64 {

	ranges, max := parseRange(data)

	values := getPossibleValues2(max)

	var res int64

	prefixMap := make([]int64, len(values)+1)

	for i, v := range values {
		prefixMap[i+1] = prefixMap[i] + v
	}

	for _, part := range ranges {
		i := util.LowerBound(values, part[0])
		j := util.UpperBound(values, part[1])
		res += prefixMap[j] - prefixMap[i]
	}

	return res
}
