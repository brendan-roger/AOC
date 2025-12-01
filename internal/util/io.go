package util

import (
	"os"
	"strings"
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