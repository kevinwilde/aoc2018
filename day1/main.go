package main

import (
	"aoc2018/utils"
	"fmt"
	"log"
	"strconv"
)

func main() {
	nums := stringsToInts(utils.GetInput("input.txt"))
	fmt.Println(part1(nums))
	fmt.Println(part2(nums))
}

func stringsToInts(strings []string) []int {
	var arr []int

	for _, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, n)
	}

	return arr
}

func part1(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func part2(nums []int) int {
	sum := 0
	seen := make(map[int]bool)
	for i := 0; ; i = (i + 1) % len(nums) {
		sum += nums[i]
		ok := seen[sum]
		if ok {
			return sum
		}
		seen[sum] = true
	}
}
