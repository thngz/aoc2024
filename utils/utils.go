package utils

import (
	"bufio"
	"os"
)

func ReadLinesToSlice(path string) []string {
	f, err := os.Open("day1/input.txt")

	var lines []string

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
