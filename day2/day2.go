package day2

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getLine(line string) []int {
	splitted := strings.Split(line, " ")
	var nums []int

	for _, i := range splitted {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func increasing(a, b int) int {
	delta := utils.GetDelta(a, b)
	dPred := 1 <= delta && delta <= 3

	if dPred && a == b {
		return 0
	} else if dPred && a > b {
		return 1
	} else {
		return -1
	}
}

func decreasing(a, b int) int {
	delta := utils.GetDelta(a, b)
	dPred := 1 <= delta && delta <= 3

	if dPred && a == b {
		return 0
	} else if dPred && a < b {
		return 1
	} else {
		return -1
	}
}

func First(lines []string) {
	lines = utils.ReadLinesToSlice("./day2/input.txt")
	answer := 0

	for _, line := range lines {
		nums := getLine(line)
		increasing := slices.IsSortedFunc(nums, increasing)
		decreasing := slices.IsSortedFunc(nums, decreasing)

		if increasing || decreasing {
			answer += 1
		}
	}
	fmt.Println(answer)
}

func Second(lines []string) {
	lines = utils.ReadLinesToSlice("./day2/input.txt")

	answer := 0

	for _, line := range lines {
		nums := getLine(line)
		N := len(nums)
		dampened := false

		for i := 0; i < N; i++ {
			clone := slices.Clone(nums)
			clone = slices.Delete(clone, i, i+1)

			increasing := slices.IsSortedFunc(clone, increasing)
			decreasing := slices.IsSortedFunc(clone, decreasing)

			if !(increasing || decreasing) && !dampened {
				dampened = true
			}

			if increasing || decreasing || !dampened {
				answer += 1
				break
			}
		}

	}
	fmt.Println(answer)
}
