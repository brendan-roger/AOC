package util

import (
	"os"
	"strings"
	"sort"
)

func ReadFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}

func ReadLines(path string) []string {
	raw := ReadFile(path)
	return strings.Split(raw, "\n")
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func LowerBound(a []int64, x int64) int {
	return sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
}

func UpperBound(a []int64, x int64) int {
	return sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
}