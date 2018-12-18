package main

import (
	"aoc2018/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	points := parsePoints(utils.GetInput("input.txt"))
	part1And2(points)
}

type xy struct {
	x int
	y int
}

type point struct {
	startPosition xy
	velocity      xy
}

func part1And2(points []*point) {
	minTime := findTimeToMinimizeBoundingBox(points)
	fmt.Println("Min Time", minTime)
	printPoints(getPositionsAtTime(points, minTime))
}

func findTimeToMinimizeBoundingBox(points []*point) int {
	var minTime int
	minBoundingBox, prevBoundingBox, expansionsInARow := math.MaxInt64, 0, 0

	for t := 0; ; t++ {
		_, _, width, height := getBoundingBox(getPositionsAtTime(points, t))

		// Minimizing width + height of bounding box with the idea that the points
		// will form a word when the perimeter of the bounding box is smallest
		boundingBox := width + height
		if boundingBox < minBoundingBox {
			minBoundingBox = boundingBox
			minTime = t
		}

		if boundingBox > prevBoundingBox {
			expansionsInARow++
			// Heuristic: If points have been expanding for 10 ticks in a row,
			// they will continue to expand
			if expansionsInARow > 10 {
				break
			}
		} else {
			expansionsInARow = 0
		}

		prevBoundingBox = boundingBox
	}

	return minTime
}

func getBoundingBox(locations []*xy) (int, int, int, int) {
	minX, maxX, minY, maxY := math.MaxInt64, -math.MaxInt64, math.MaxInt64, -math.MaxInt64

	for _, p := range locations {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	return minX, minY, maxX - minX + 1, maxY - minY + 1
}

func getPositionsAtTime(points []*point, time int) []*xy {
	var locations []*xy
	for _, p := range points {
		locations = append(locations, getPositionAtTime(p, time))
	}
	return locations
}

func getPositionAtTime(p *point, time int) *xy {
	return &xy{
		p.startPosition.x + time*p.velocity.x,
		p.startPosition.y + time*p.velocity.y,
	}
}

func printPoints(locations []*xy) {
	locationMap := make(map[int]map[int]bool)
	for _, loc := range locations {
		if locationMap[loc.x] == nil {
			locationMap[loc.x] = make(map[int]bool)
		}
		locationMap[loc.x][loc.y] = true
	}

	x, y, width, height := getBoundingBox(locations)
	for j := y; j < y+height; j++ {
		for i := x; i < x+width; i++ {
			if locationMap[i] != nil && locationMap[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func parsePoints(input []string) []*point {
	var res []*point
	reg := regexp.MustCompile(`^position=\<(.+),(.+)\> velocity=\<(.+),(.+)\>$`)
	for _, line := range input {
		match := reg.FindStringSubmatch(line)
		if len(match) < 1 {
			panic("Input did not match regex")
		}

		var ints []int
		for i := 1; i < 5; i++ {
			ints = append(ints, utils.ParseInt(strings.Trim(match[i], " ")))
		}

		p := point{
			xy{ints[0], ints[1]},
			xy{ints[2], ints[3]},
		}
		res = append(res, &p)
	}

	return res
}
