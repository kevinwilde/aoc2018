package main

import (
	"aoc2018/utils"
	"fmt"
)

// Dist between lowercase and uppercase version of same letter
const lowerUpperDist = byte('a' - 'A')

func main() {
	input := utils.GetInput("input.txt")
	if len(input) != 1 {
		panic("Expected one string as input")
	}
	polymer := input[0]
	fmt.Println(part1(polymer))
	fmt.Println(part2(polymer))
}

func part1(polymer string) int {
	return len(react(polymer))
}

func part2(polymer string) int {
	minPolymer := len(polymer)
	for b := byte('A'); b <= byte('Z'); b++ {
		newPolymer := removeType(polymer, b)
		reacted := react(newPolymer)
		if len(reacted) < minPolymer {
			minPolymer = len(reacted)
		}
	}
	return minPolymer
}

func react(polymer string) string {
	s := polymer
	for i := 0; i < len(s)-1; {
		if willReact(s[i], s[i+1]) {
			s = s[:i] + s[i+2:]
			i--
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}
	return s
}

func willReact(c1 byte, c2 byte) bool {
	d := c1 - c2
	return d == lowerUpperDist || -d == lowerUpperDist
}

func removeType(polymer string, typeToRemove byte) string {
	s := polymer
	for i := 0; i < len(s); {
		if s[i] == typeToRemove || s[i] == typeToRemove+lowerUpperDist {
			s = s[:i] + s[i+1:]
		} else {
			i++
		}
	}
	return s
}
