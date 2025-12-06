package main

import (
	"AOC/internal/util"
	f "fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	data := util.ReadLines("2025/05/input.txt")
	f.Println("Part 1:", part1(data))
	f.Println("Part 2:", part2(data))
}

// 3-5
// 10-14
// 16-20
// 12-18
// 50-60
// 20-30

// {3: 5}
// {10: 30}
// {50: 60}

type Range struct {
	Min int64
	Max int64
}

type node struct {
	key      int64
	data     Range
	priority int
	left     *node
	right    *node
}


func rotateRight(node *node) *node {
	prev := node.left
	next := prev.right

	prev.right = node
	node.left = next
	return prev
}

func rotateLeft(node *node) *node {
	next := node.right
	prev := next.left

	next.left = node
	node.right = prev
	return next
}

func insert(root *node, key int64, data Range) *node {
	if root == nil {
		return &node{
			key:      key,
			data:     data,
			priority: rand.Int(),
		}
	}

	if key < root.key {
		root.left = insert(root.left, key, data)

		if root.left.priority > root.priority {
			root = rotateRight(root)
		}
	} else if key > root.key {
		root.right = insert(root.right, key, data)

		if root.right.priority > root.priority {
			root = rotateLeft(root)

		}
	} else {
		root.data = data
	}
	return root
}

func delete(root *node, key int64) *node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = delete(root.left, key)

	} else if key > root.key {
		root.right = delete(root.right, key)

	} else {

		if root.left == nil {
			return root.right

		} else if root.right == nil {
			return root.left

		} else if root.left.priority > root.right.priority {
			root = rotateRight(root)
			root.right = delete(root.right, key)

		} else {
			root = rotateLeft(root)
			root.left = delete(root.left, key)

		}
	}

	return root
}

func getPrev(root *node, x int64) *node {
	var res *node
	for root != nil {
		if root.key <= x {
			res = root
			root = root.right

		} else {
			root = root.left
		}
	}
	return res
}

func getNext(root *node, x int64) *node {
	var res *node
	for root != nil {
		if root.key > x {
			res = root
			root = root.left

		} else {
			root = root.right
		}
	}
	return res
}

// root or updated range
func insertRange(root *node, r Range) *node {

	if prev := getPrev(root, r.Min); prev != nil {
		if prev.data.Max >= r.Min-1 {

			root = delete(root, prev.key)

			r.Min = min(prev.data.Min, r.Min)
			r.Max = max(prev.data.Max, r.Max)
		}
	}

	// merge with all next that overlap
	for {
		next := getNext(root, r.Min)
		if next == nil {
			break
		}
		if next.data.Min > r.Max+1 {
			break
		}

		// merge and delete
		root = delete(root, next.key)

		r.Min = min(r.Min, next.data.Min)
		r.Max = max(r.Max, next.data.Max)
	}


	return insert(root, r.Min, r)
}

func inorder(n *node, out *[]Range) {
	if n == nil {
		return
	}
	inorder(n.left, out)
	*out = append(*out, n.data)
	inorder(n.right, out)
}

func contains(root *node, x int64) bool {
	n := getPrev(root, x)
	if n == nil {
		return false
	}
	return x <= n.data.Max
}

func part1(lines []string) int {
	var root *node

	breakpoint := 0

	for idx, line := range lines {
		if line == "" {
			breakpoint = idx + 1
			break
		}

		parts := strings.Split(line, "-")
		min, _ := strconv.ParseInt(parts[0], 10, 64)
		max, _ := strconv.ParseInt(parts[1], 10, 64)

		root = insertRange(root, Range{Min: min, Max: max})
	}


	var merged []Range
	inorder(root, &merged)

	res := 0

	for _, line := range lines[breakpoint:] {

		digit, _ := strconv.ParseInt(line, 10, 64)

		if contains(root, digit) {
			res += 1
		}

	}

	return res
}

func part2(lines []string) int64 {
	var root *node

	for _, line := range lines {
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		min, _ := strconv.ParseInt(parts[0], 10, 64)
		max, _ := strconv.ParseInt(parts[1], 10, 64)

		root = insertRange(root, Range{Min: min, Max: max})
	}

	var merged []Range
	inorder(root, &merged)

	var res int64 = 0

	for _, item := range merged {

		res += item.Max - item.Min + 1

	}

	return res
}
