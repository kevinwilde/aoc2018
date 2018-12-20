package main

import (
	"fmt"
	"math"
)

func main() {
	var pLevels [][]int
	var x, y, s int

	pLevels = powerLevels(300, 18)
	x, y = part1(pLevels)
	fmt.Println(x == 33 && y == 45)
	x, y, s = part2(pLevels)
	fmt.Println(x == 90 && y == 269 && s == 16)

	pLevels = powerLevels(300, 42)
	x, y = part1(pLevels)
	fmt.Println(x == 21 && y == 61)
	x, y, s = part2(pLevels)
	fmt.Println(x == 232 && y == 251 && s == 12)

	pLevels = powerLevels(300, 7315)
	x, y = part1(pLevels)
	fmt.Println(x, y)
	x, y, s = part2(pLevels)
	fmt.Println(x, y, s)
}

func part1(pLevels [][]int) (int, int) {
	_, x, y := findMaxArea(pLevels, 3)
	return x, y
}

func part2(pLevels [][]int) (int, int, int) {
	maxPower, maxX, maxY, maxS := math.MinInt64, 0, 0, 0
	for s := 1; s < len(pLevels); s++ {
		power, x, y := findMaxArea(pLevels, s)
		if power > maxPower {
			maxPower, maxX, maxY, maxS = power, x, y, s
		}
	}
	return maxX, maxY, maxS
}

func findMaxArea(pLevels [][]int, squareSize int) (int, int, int) {
	maxPower, maxX, maxY := math.MinInt64, 0, 0
	prevPower := 0
	for x := 1; x < len(pLevels)-squareSize; x++ {
		for y := 1; y < len(pLevels[x])-squareSize; y++ {
			var power int
			// Optimization: only recalculate power of area from scratch when y == 1
			// When y > 1, start based off prevPower. Subtract the row that is no
			// longer in area and add the row that is now in the area.
			if y == 1 {
				for i := 0; i < squareSize; i++ {
					for j := 0; j < squareSize; j++ {
						power += pLevels[x+i][y+j]
					}
				}
			} else {
				power = prevPower
				for i := 0; i < squareSize; i++ {
					power -= pLevels[x+i][y-1]
					power += pLevels[x+i][y+squareSize-1]
				}
			}

			if power > maxPower {
				maxPower, maxX, maxY = power, x, y
			}

			prevPower = power
		}
	}
	return maxPower, maxX, maxY
}

func powerLevels(squareSize, serialNo int) [][]int {
	pLevels := make([][]int, squareSize+1)
	for x := 1; x <= squareSize; x++ {
		if pLevels[x] == nil {
			pLevels[x] = make([]int, squareSize+1)
		}
		for y := 1; y <= squareSize; y++ {
			pLevels[x][y] = powerLevel(x, y, serialNo)
		}
	}
	return pLevels
}

func powerLevel(x, y, serialNo int) int {
	rackID := x + 10
	p := rackID * y
	p += serialNo
	p *= rackID
	p %= 1000
	p /= 100
	p -= 5
	return p
}
