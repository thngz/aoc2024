package day1

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getColumns(lines []string) ([]int, []int) {
	var firstCol []int
	var secondCol []int

	for _, line := range lines {
		parts := strings.Split(line, "  ")
		first, second := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

		firstInt, err := strconv.Atoi(first)

		if err != nil {
			panic(err)
		}

		secondInt, err := strconv.Atoi(second)

		if err != nil {
			panic(err)
		}

		firstCol = append(firstCol, firstInt)
		secondCol = append(secondCol, secondInt)
	}

	return firstCol, secondCol
}

func First(lines []string) {
	lines = utils.ReadLinesToSlice("day1/input.txt")
	var answer int = 0

	firstCol, secondCol := getColumns(lines)

	slices.Sort(firstCol)
	slices.Sort(secondCol)

	for i := range len(firstCol) {
		var ans int
		first, second := firstCol[i], secondCol[i]

		if first < second {
			ans = second - first
		} else {
			ans = first - second
		}

		answer += ans
	}
	fmt.Println(answer)
}

func Second(lines []string) {
    lines = utils.ReadLinesToSlice("day1/input.txt")
	firstCol, secondCol := getColumns(lines)
	answer := 0
	var curr int

	slices.Sort(firstCol)
	slices.Sort(secondCol)

	for _, i := range firstCol {
		for _, j := range secondCol {
			if i == j {
				curr += 1
			}
		}
		answer += curr * i
		curr = 0
	}

	fmt.Println(answer)
}
