package main

import (
	"aoc2018/utils"
	"fmt"
	"strings"
)

func main() {
	rules := parseRules(utils.GetInput("input1.txt"))
	fmt.Println(part1And2("#..#.#..##......###...###", rules, 20))

	initialState := "..#..####.##.####...#....#######..#.#..#..#.#.#####.######..#.#.#.#..##.###.#....####.#.#....#.#####"
	rules = parseRules(utils.GetInput("input.txt"))
	fmt.Println(part1And2(initialState, rules, 20))

	rules = parseRules(utils.GetInput("input.txt"))
	fmt.Println(part1And2(initialState, rules, 50000000000))
}

func part1And2(initialState string, rules map[string]string, numIter int) int {
	state, firstNum := padState(strings.Split(initialState, ""))
	firstNum *= -1
	prevStateString, prevFirstNum := "", 0

	for i := 0; i < numIter; i++ {
		stateString := strings.Join(state, "")

		// Optimization: state hasn't changed from previous iteration, so it will
		// never change again. Since firstNum may still have changed though,
		// caculate what final value of firstNum will be. Then we have answer
		if stateString == prevStateString {
			finalFirstNum := (numIter-i)*(firstNum-prevFirstNum) + firstNum
			return sum(state, finalFirstNum)
		}
		prevStateString, prevFirstNum = stateString, firstNum

		nextState, amountPaddedOnLeft := transform(state, rules)
		state = nextState
		firstNum -= amountPaddedOnLeft
	}

	return sum(state, firstNum)
}

func transform(state []string, rules map[string]string) ([]string, int) {
	nextState := state
	stateString := strings.Join(state, "")

	for j := 0; j < len(state)-5; j++ {
		if rules[stateString[j:j+5]] == "#" {
			nextState[j+2] = "#"
		} else {
			nextState[j+2] = "."
		}
	}

	return padState(nextState)
}

func sum(state []string, firstNum int) int {
	sum := 0
	for i, c := range state {
		if c == "#" {
			sum += (i + firstNum)
		}
	}
	return sum
}

func parseRules(input []string) map[string]string {
	rules := make(map[string]string)
	for _, s := range input {
		split := strings.Split(s, " => ")
		rules[split[0]] = split[1]
	}
	return rules
}

func padState(state []string) ([]string, int) {
	i := 0
	for ; state[i] == "."; i++ {
	}
	padOnLeft := 5 - i
	j := 0
	for ; state[len(state)-1-j] == "."; j++ {
	}
	padOnRight := 5 - j

	paddedState := make([]string, padOnLeft+len(state)+padOnRight)

	for k := 0; k < padOnLeft; k++ {
		paddedState[k] = "."
	}
	for k := padOnLeft; k < len(state)+padOnLeft; k++ {
		if k >= 0 {
			paddedState[k] = state[k-padOnLeft]
		}
	}
	for j = 0; j < padOnRight; j++ {
		paddedState[padOnLeft+len(state)+j] = "."
	}

	return paddedState, padOnLeft
}
