package main

import (
	"AOC/internal/util"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func parsePoints(rows []string) []Point {
	points := make([]Point, 0)

	for _, row := range rows {
		parts := strings.Split(row, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		points = append(points, Point{X: x, Y: y})
	}

	return points
}

func main() {
	data := util.ReadLines("2025/09/input.txt")
	points := parsePoints(data)
	fmt.Println("Part 1:", part1(points))
	fmt.Println("Part 2:", part2(data))
}

func traversePrune(points []Point) (sw []Point, ne []Point) {

	p := make([]Point, len(points))
	copy(p, points)

	sort.Slice(p, func(i, j int) bool {
		if p[i].X == p[j].X {
			return p[i].Y < p[j].Y
		}
		return p[i].X < p[j].X
	})

	minY := math.MaxInt
	for _, pt := range p {
		if pt.Y < minY {
			sw = append(sw, pt)
			minY = pt.Y
		}
	}

	sort.Slice(p, func(i, j int) bool {
		if p[i].X == p[j].X {
			return p[i].Y > p[j].Y
		}
		return p[i].X > p[j].X
	})

	maxY := math.MinInt
	for _, pt := range p {
		if pt.Y > maxY {
			ne = append(ne, pt)
			maxY = pt.Y
		}
	}

	return sw, ne
}

func getMaxArea(points []Point) int {
	sw, ne := traversePrune(points)
	maxArea := 0

	for _, a := range sw {
		for _, b := range ne {
			if b.X <= a.X || b.Y <= a.Y {
				continue
			}
			dx := b.X - a.X + 1
			dy := b.Y - a.Y + 1
			area := dx * dy
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func part1(points []Point) int {

	maxA := getMaxArea(points)

	flip := make([]Point, len(points))
	for i, p := range points {
		flip[i] = Point{X: p.X, Y: -p.Y}
	}

	maxB := getMaxArea(flip)

	return max(maxA, maxB)
}

func part2(points []string) int {

	return 0
}
