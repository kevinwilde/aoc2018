package main

import (
	"aoc2018/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	tracks := parseTracks(utils.GetInput("input.txt"))
	cars := findCars(tracks)
	tracks = replaceCars(cars, tracks)
	// fmt.Println(part1(cars, tracks))
	fmt.Println(part2(cars, tracks))
}

type car struct {
	x                int
	y                int
	direction        string
	numIntersections int
	alive            bool
}

func part1(cars []*car, tracks [][]string) (int, int) {
	for {
		orderTopLeft(cars)
		for i, c := range cars {
			move(c, tracks)
			for j, c2 := range cars {
				if i != j && c.x == c2.x && c.y == c2.y {
					// Collision
					return c.x, c.y
				}
			}
		}
	}
}

func part2(cars []*car, tracks [][]string) (int, int) {
	for {
		orderTopLeft(cars)
		for i, c := range cars {
			if c.alive {
				move(c, tracks)
				for j, c2 := range cars {
					if c2.alive && i != j && c.x == c2.x && c.y == c2.y {
						// Collision!
						c.alive = false
						c2.alive = false
					}
				}
			}
		}

		// Check if there is only one car left. If so, return location
		carsAlive := 0
		for _, c := range cars {
			if c.alive {
				carsAlive++
			}
		}
		if carsAlive == 1 {
			for _, c := range cars {
				if c.alive {
					return c.x, c.y
				}
			}
		}
	}
}

func orderTopLeft(cars []*car) {
	sort.Slice(cars, func(i, j int) bool {
		if cars[i].y < cars[j].y {
			return true
		}
		if cars[i].y > cars[j].y {
			return false
		}
		return cars[i].x < cars[j].x
	})
}

func parseTracks(input []string) [][]string {
	tracks := make([][]string, len(input))
	for i := range input {
		tracks[i] = strings.Split(input[i], "")
	}
	return tracks
}

func findCars(tracks [][]string) []*car {
	var cars []*car

	for y := range tracks {
		for x := range tracks[y] {
			switch tracks[y][x] {
			case "v":
				fallthrough
			case "^":
				fallthrough
			case "<":
				fallthrough
			case ">":
				c := car{x, y, tracks[y][x], 0, true}
				cars = append(cars, &c)
			}
		}
	}

	return cars
}

func replaceCars(cars []*car, tracks [][]string) [][]string {
	for _, c := range cars {
		var rep string
		if c.y > 0 && (tracks[c.y-1][c.x] == "|" || tracks[c.y-1][c.x] == "+") {
			rep = "|"
		} else if c.y < len(tracks)-1 && (tracks[c.y+1][c.x] == "|" || tracks[c.y+1][c.x] == "+") {
			rep = "|"
		} else if c.x > 0 && (tracks[c.y][c.x-1] == "-" || tracks[c.y][c.x-1] == "+") {
			rep = "-"
		} else if c.x < len(tracks[c.y])-1 && (tracks[c.y][c.x+1] == "-" || tracks[c.y][c.x+1] == "+") {
			rep = "-"
		} else {
			fmt.Println(c.x, c.y, tracks[c.y][c.x])
			panic("Unable to replace car")
		}
		tracks[c.y][c.x] = rep
	}
	return tracks
}

func move(c *car, tracks [][]string) {
	switch c.direction {
	case "v":
		switch tracks[c.y+1][c.x] {
		case "|":
			c.y++
		case "\\":
			c.y++
			c.direction = ">"
		case "/":
			c.y++
			c.direction = "<"
		case "+":
			c.y++
			switch c.numIntersections % 3 {
			case 0:
				c.direction = ">"
			case 1:
				c.direction = "v"
			case 2:
				c.direction = "<"
			}
			c.numIntersections++
		default:
			panic("Unexpected tracks")
		}

	case "^":
		switch tracks[c.y-1][c.x] {
		case "|":
			c.y--
		case "\\":
			c.y--
			c.direction = "<"
		case "/":
			c.y--
			c.direction = ">"
		case "+":
			c.y--
			switch c.numIntersections % 3 {
			case 0:
				c.direction = "<"
			case 1:
				c.direction = "^"
			case 2:
				c.direction = ">"
			}
			c.numIntersections++
		default:
			panic("Unexpected tracks")
		}

	case "<":
		switch tracks[c.y][c.x-1] {
		case "-":
			c.x--
		case "\\":
			c.x--
			c.direction = "^"
		case "/":
			c.x--
			c.direction = "v"
		case "+":
			c.x--
			switch c.numIntersections % 3 {
			case 0:
				c.direction = "v"
			case 1:
				c.direction = "<"
			case 2:
				c.direction = "^"
			}
			c.numIntersections++
		default:
			panic("Unexpected tracks")
		}

	case ">":
		switch tracks[c.y][c.x+1] {
		case "-":
			c.x++
		case "\\":
			c.x++
			c.direction = "v"
		case "/":
			c.x++
			c.direction = "^"
		case "+":
			c.x++
			switch c.numIntersections % 3 {
			case 0:
				c.direction = "^"
			case 1:
				c.direction = ">"
			case 2:
				c.direction = "v"
			}
			c.numIntersections++
		default:
			panic("Unexpected tracks")
		}
	}
}

func drawTracks(cars []*car, tracks [][]string) {
	for y := 0; y < len(tracks); y++ {
		for x := 0; x < len(tracks[y]); x++ {
			carThere := false
			for _, c := range cars {
				if c.x == x && c.y == y {
					carThere = true
					fmt.Print(c.direction)
				}
			}
			if !carThere {
				fmt.Print(tracks[y][x])
			}
		}
		fmt.Printf("\n")
	}
}
