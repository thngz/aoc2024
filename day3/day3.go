package day3

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
	"strconv"
	// "strings"
)

func First(lines []string) {

	lines = utils.ReadLinesToSlice("day3/input.txt")
	rString := `mul\(\d+,\d+\)`
	r, _ := regexp.Compile(rString)
	rnums, _ := regexp.Compile(`\d+`)
	answer := 0

	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			numMatches := rnums.FindAllString(match, -1)
			a, _ := strconv.Atoi(numMatches[0])
			b, _ := strconv.Atoi(numMatches[1])
			answer += a * b
		}

		// fmt.Println(matches)
	}

	fmt.Println(answer)
}

func Second(lines []string) {
	lines = utils.ReadLinesToSlice("day3/input.txt")
	rString := `mul\(\d+,\d+\)|do\(\)|don't\(\)`
	r, _ := regexp.Compile(rString)
	rnums, _ := regexp.Compile(`\d+`)
	enabled := true
	answer := 0

	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		// fmt.Println(matches)
		for _, match := range matches {
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else {
				numMatches := rnums.FindAllString(match, -1)
				a, _ := strconv.Atoi(numMatches[0])
				b, _ := strconv.Atoi(numMatches[1])
				if enabled {
					answer += a * b
				}
			}
		}
	}
    fmt.Println(answer)
}
