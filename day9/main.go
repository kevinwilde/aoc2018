package main

import "fmt"

func main() {
	fmt.Println(winningScore(9, 25) == 32)
	fmt.Println(winningScore(10, 1618) == 8317)
	fmt.Println(winningScore(13, 7999) == 146373)
	fmt.Println(winningScore(17, 1104) == 2764)
	fmt.Println(winningScore(21, 6111) == 54718)
	fmt.Println(winningScore(30, 5807) == 37305)
	fmt.Println(winningScore(411, 71170))
	fmt.Println(winningScore(411, 71170*100))
}

func winningScore(numPlayers, numMarbles int) int {
	scores := playGame(numPlayers, numMarbles)
	return findMaxValue(scores)
}

type node struct {
	value int
	left  *node
	right *node
}

func playGame(numPlayers, numMarbles int) map[int]int {
	playerScores := make(map[int]int)
	n := node{0, nil, nil}
	n.left = &n
	n.right = &n
	circle := &n
	for marble := 1; marble <= numMarbles; marble++ {
		player := marble % numPlayers
		if marble%23 == 0 {
			// Score! Marble is added to score
			playerScores[player] += marble
			// Also  7 marbles counter-clockwise from the current marble is
			// removed from the circle and also added to the current player's
			// score. The marble located immediately clockwise of the marble
			// that was removed becomes the new current marble.
			for i := 0; i < 7; i++ {
				circle = circle.left
			}
			playerScores[player] += circle.value
			circle.left.right = circle.right
			circle.right.left = circle.left
			circle = circle.right
		} else {
			n := node{marble, circle.right, circle.right.right}
			circle.right.right.left = &n
			circle.right.right = &n
			circle = &n
		}
	}
	return playerScores
}

func findMaxValue(keyValMap map[int]int) int {
	maxVal := 0
	for _, val := range keyValMap {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
