package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// input go run . 25 01 - 2025 day 1

	year := os.Args[1]
	day := os.Args[2]
	cmd := exec.Command("go", "run", fmt.Sprintf("20%s/%s/main.go", year, day))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
