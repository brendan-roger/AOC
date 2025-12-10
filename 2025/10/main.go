package main

import (
	"AOC/internal/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"math"
)

// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7} => 2 [(0,2) && (0,1)]
// [...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
// [.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}

// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7} => 2 [(0,2) && (0,1)]

func main() {
	data := util.ReadLines("2025/10/input.txt")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}



type Part struct {
	N       int
	target  string   // [.##.] => [0110]
	options [][]int // [[3], [1,3], [2], [2,3], [0,2], [0,1]]
	req     []int   // [3547]
}

var (
	reOps = regexp.MustCompile(`\(([^)]*)\)`)
	reReq = regexp.MustCompile(`\{([^}]*)\}`)
)

func parseLine(line string) Part {
	var p Part

	start := strings.Index(line, "[")
	end := strings.Index(line, "]")
	rawMask := line[start+1 : end]

	p.N = len(rawMask)

	buf := make([]byte, len(rawMask))

	for i, ch := range rawMask {
		if ch == '#' {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}

	p.target = string(buf)


	opsPart := line[end+1:]
	for _, match := range reOps.FindAllStringSubmatch(opsPart, -1) {
		raw := match[1]


		if raw == "" {
			p.options = append(p.options, []int{})
			continue
		}

		parts := strings.Split(raw, ",")
		ints := make([]int, 0, len(parts))

		for _, s := range parts {
			n, _ := strconv.Atoi(strings.TrimSpace(s))
			ints = append(ints, n)
		}
		p.options = append(p.options, ints)
	}

	reqMatch := reReq.FindStringSubmatch(line)
	if len(reqMatch) > 1 {
		parts := strings.Split(reqMatch[1], ",")
		for _, s := range parts {
			n, _ := strconv.Atoi(strings.TrimSpace(s))
			p.req = append(p.req, n)
		}
	}

	return p
}


func nextStateP1(state string, option []int) string {
    newState := []byte(state)
    for _, idx := range option {
        newState[idx] = ((newState[idx]-'0') ^ 1) + '0'
    }
    return string(newState)
}


func runP1(part Part) int {
    memo := make(map[string]int)
    var recur func(int, string) int

	target := strings.Repeat("0", len(part.target))

    recur = func(idx int, state string) int {
        if state == target {
            return 0
        }

        key := fmt.Sprintf("%d|%s", idx, state)
        if v, ok := memo[key]; ok {
            return v
        }

        res := math.MaxInt32

        for i := idx; i < len(part.options); i++ {
            op := part.options[i]
            newState := nextStateP1(state, op)

            count := 1 + recur(i+1, newState)
            if count < res {
                res = count
            }
        }

        memo[key] = res
        return res
    }

    return recur(0, part.target)
}

func part1(lines []string) int {
    res := 0

    for _, line := range lines {
        parsed := parseLine(line)
        count := runP1(parsed)
        res += count
    }

    return res
}


func part2(lines []string) int {
    res := 0


    return res
}