package utils

import (
	"bufio"
	"os"
)

func ReadLinesToSlice(path string) []string {
	f, err := os.Open(path)

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

func GetDelta(a int, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}
