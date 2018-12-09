package main

import (
	"aoc2018/utils"
	"fmt"
)

func main() {
	input := utils.GetInput("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	var twos, threes int
	for _, id := range input {
		if hasExactly(id, 2) {
			twos++
		}
		if hasExactly(id, 3) {
			threes++
		}
	}
	return twos * threes
}

func hasExactly(id string, n int) bool {
	letterCounts := make(map[rune]int)
	for _, lett := range id {
		letterCounts[lett]++
	}
	for _, val := range letterCounts {
		if val == n {
			return true
		}
	}
	return false
}

func part2(input []string) string {
	for i, id := range input {
		for _, id2 := range input[i+1:] {
			common := commonChars(id, id2)
			if len(common) == len(id)-1 {
				return common
			}
		}
	}
	return ""
}

func commonChars(s1 string, s2 string) string {
	if len(s1) != len(s2) {
		return ""
	}
	res := ""
	for i := range s1 {
		if s1[i] == s2[i] {
			res += string(s1[i])
		}
	}
	return res
}
