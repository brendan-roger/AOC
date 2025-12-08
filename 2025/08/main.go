package main

import (
	"AOC/internal/util"
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := util.ReadLines("2025/08/input.txt")
	fmt.Println("Part 1:", part1(data, 1000))
	fmt.Println("Part 2:", part2(data))
}

func getDist(a []int, b []int) int {
	dist := 0

	for i := range a {
		d := a[i] - b[i]
		dist += d * d
	}

	return dist
}

func parsePoint(s string) []int {
	parts := strings.Split(s, ",")

	arr := make([]int, len(parts))

	for i, p := range parts {
		arr[i], _ = strconv.Atoi(p)
	}

	return arr
}

func dfs(start int, set []bool, adjMap [][]int) int {
	stack := []int{start}
	size := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if set[node] {
			continue
		}
		set[node] = true
		size++

		for _, nb := range adjMap[node] {
			if !set[nb] {
				stack = append(stack, nb)
			}
		}
	}

	return size
}

func part1(points []string, max int) int {

	n := len(points)
	h := &util.IntMaxHeap{}

	heap.Init(h)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {

			a := parsePoint(points[i])
			b := parsePoint(points[j])

			dist := getDist(a, b)

			if h.Len() < max {
				heap.Push(h, util.HeapItem{Priority: dist, I: i, J: j})
			} else if (*h)[0].Priority > dist {
				heap.Pop(h)
				heap.Push(h, util.HeapItem{Priority: dist, I: i, J: j})
			}

		}
	}

	// dfs through graph to find sizes of connected components
	set := make([]bool, n)
	adjMap := make([][]int, n)

	for h.Len() > 0 {
		item := heap.Pop(h).(util.HeapItem)
		u, v := item.I, item.J

		adjMap[u] = append(adjMap[u], v)
		adjMap[v] = append(adjMap[v], u)
	}

	neighbours := make([]int, 0)

	for i := 0; i < n; i++ {
		if !set[i] {
			size := dfs(i, set, adjMap)
			neighbours = append(neighbours, size)
		}
	}

	sort.Slice(neighbours, func(a, b int) bool {
		return neighbours[a] > neighbours[b]
	})

	res := 1

	for _, num := range neighbours[:3] {
		res *= num
	}

	return res
}

type Edge struct {
	dist int
	U    int
	V    int
}

func part2(points []string) int {

	n := len(points)

	edges := make([]Edge, 0)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {

			a := parsePoint(points[i])
			b := parsePoint(points[j])

			dist := getDist(a, b)

			edges = append(edges, Edge{dist: dist, U: i, V: j})

		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a].dist < edges[b].dist
	})

	uf := util.InitUnionFind(n)
	components := n

	var lastU, lastV int

	// find the last edge that connects all components
	for _, e := range edges {
		u, v := e.U, e.V

		if uf.Union(u, v) {
			components--

			if components == 1 {
				lastU, lastV = u, v
				break
			}
		}
	}

	x1 := parsePoint(points[lastU])
	x2 := parsePoint(points[lastV])

	return x1[0] * x2[0]
}
