package main

import (
	"aoc2018/utils"
	"fmt"
	"log"
	"math"
	"regexp"
)

func main() {
	coords := parseCoordinates(utils.GetInput("input.txt"))
	fmt.Println(part1(coords))
	fmt.Println(part2(coords, 10000))
}

type coord struct {
	x int
	y int
}

type mapEntry struct {
	coordIndex int
	dist       int
}

func parseCoordinates(input []string) []*coord {
	re := regexp.MustCompile(`^(\d+), (\d+)$`)
	var coords []*coord

	for _, s := range input {
		match := re.FindStringSubmatch(s)
		if len(match) < 1 {
			log.Fatal("Input did not match regex")
		}

		x := utils.ParseInt(match[1])
		y := utils.ParseInt(match[2])

		coords = append(coords, &coord{x, y})
	}

	return coords
}

func part1(coords []*coord) int {
	m := make(map[int]map[int]*mapEntry)

	corner1, corner2 := getCorners(coords)

	for x := corner1.x; x <= corner2.x; x++ {
		if m[x] == nil {
			m[x] = make(map[int]*mapEntry)
		}
		for y := corner1.y; y <= corner2.y; y++ {
			for i, c := range coords {
				d := manhattanDist(coord{x, y}, *c)
				if m[x][y] == nil {
					m[x][y] = &mapEntry{i, d}
				} else if d < m[x][y].dist {
					m[x][y] = &mapEntry{i, d}
				} else if d == m[x][y].dist {
					m[x][y] = &mapEntry{-1, d} // tie
				}
			}
		}
	}

	var area []int

	for i := range coords {
		count := 0
		infinite := false
		for x := range m {
			for y := range m[x] {
				if m[x][y].coordIndex == i {
					if isEdgePoint(coord{x, y}, corner1, corner2) {
						infinite = true
					}
					count++
				}
			}
		}

		if infinite {
			area = append(area, -1)
		} else {
			area = append(area, count)
		}
	}

	maxArea := 0
	for _, val := range area {
		if val > maxArea {
			maxArea = val
		}
	}

	return maxArea
}

func part2(coords []*coord, maxSum int) int {
	corner1, corner2 := getCorners(coords)

	count := 0

	for x := corner1.x; x <= corner2.x; x++ {
		for y := corner1.y; y <= corner2.y; y++ {
			sumToAll := 0
			for _, c := range coords {
				sumToAll += manhattanDist(coord{x, y}, *c)
			}
			if sumToAll < maxSum {
				count++
			}
		}
	}
	return count
}

func getCorners(coords []*coord) (coord, coord) {
	var minX, minY int = math.MaxInt64, math.MaxInt64
	maxX, maxY := 0, 0

	for _, c := range coords {
		minX = min(minX, c.x)
		maxX = max(maxX, c.x)
		minY = min(minY, c.y)
		maxY = max(maxY, c.y)
	}

	return coord{minX, minY}, coord{maxX, maxY}
}

func manhattanDist(c1 coord, c2 coord) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}

func isEdgePoint(c, corner1, corner2 coord) bool {
	return c.x == corner1.x || c.x == corner2.x || c.y == corner1.y || c.y == corner2.y
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
