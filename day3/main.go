package main

import (
	"aoc2018/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	claims := convertInputToClaims(utils.GetInput("input.txt"))
	coordToClaims := createCoordinateToClaimsMap(claims)
	fmt.Println(part1(coordToClaims))
	fmt.Println(part2(claims, coordToClaims))
}

func part1(coordToClaims map[int]map[int][]*claim) int {
	claimedByTwo := 0
	for x := range coordToClaims {
		for y := range coordToClaims[x] {
			if len(coordToClaims[x][y]) > 1 {
				claimedByTwo++
			}
		}
	}
	return claimedByTwo
}

func part2(allClaims map[int]*claim, coordToClaims map[int]map[int][]*claim) int {
	for x := range coordToClaims {
		for y := range coordToClaims[x] {
			if len(coordToClaims[x][y]) > 1 {
				for _, aClaim := range coordToClaims[x][y] {
					aClaim.overlapped = true
				}
			}
		}
	}

	for _, claim := range allClaims {
		if !claim.overlapped {
			return claim.id
		}
	}
	return -1
}

type claim struct {
	id         int
	left       int
	top        int
	width      int
	height     int
	overlapped bool
}

func convertInputToClaims(input []string) map[int]*claim {
	var claimMap = make(map[int]*claim)

	reg := regexp.MustCompile("[^0-9]+")

	for _, s := range input {
		numericStr := reg.ReplaceAllString(s, " ")
		splitClaim := strings.Split(numericStr, " ")
		var splitNumericClaim []int
		for _, s := range splitClaim {
			if len(s) > 0 {
				n, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				splitNumericClaim = append(splitNumericClaim, n)
			}
		}

		claimID := splitNumericClaim[0]
		claimMap[claimID] = &claim{
			claimID,
			splitNumericClaim[1],
			splitNumericClaim[2],
			splitNumericClaim[3],
			splitNumericClaim[4],
			false,
		}
	}
	return claimMap
}

func createCoordinateToClaimsMap(input map[int]*claim) map[int]map[int][]*claim {
	claimed := make(map[int](map[int][]*claim))

	for _, thisClaim := range input {
		for x := thisClaim.left; x < thisClaim.left+thisClaim.width; x++ {
			if claimed[x] == nil {
				claimed[x] = make(map[int][]*claim)
			}
			for y := thisClaim.top; y < thisClaim.top+thisClaim.height; y++ {
				claimed[x][y] = append(claimed[x][y], thisClaim)
			}
		}
	}

	return claimed
}
