package main

import (
	"aoc2018/utils"
	"fmt"
	"strings"
)

func main() {
	root := buildTree(stringsToInts(utils.GetInput("input.txt")))
	fmt.Println(part1(root))
	fmt.Println(part2(root))
}

type node struct {
	metadata []int
	children []*node
}

func part1(root *node) int {
	return dfsSum(root)
}

func dfsSum(n *node) int {
	sum := sumMetadata(n)
	for _, child := range n.children {
		sum += dfsSum(child)
	}
	return sum
}

func sumMetadata(n *node) int {
	sum := 0
	for _, val := range n.metadata {
		sum += val
	}
	return sum
}

func part2(root *node) int {
	return value(root)
}

func value(n *node) int {
	if len(n.children) == 0 {
		return sumMetadata(n)
	}

	val := 0
	for _, index := range n.metadata {
		// 1 based indexing...
		if index > 0 && index <= len(n.children) {
			val += value(n.children[index-1])
		}
	}
	return val
}

func stringsToInts(input []string) []int {
	var arr []int

	if len(input) != 1 {
		panic("Unexpected input")
	}

	for _, s := range strings.Split(input[0], " ") {
		n := utils.ParseInt(s)
		arr = append(arr, n)
	}

	return arr
}

func buildTree(nums []int) *node {
	tree, _ := buildTreeRecur(nums, 0, 1)
	return tree[0]
}

func buildTreeRecur(nums []int, i int, numChild int) ([]*node, int) {
	var res []*node
	for c := 0; c < numChild; c++ {
		children := nums[i]
		i++
		numMeta := nums[i]
		i++
		n := node{
			make([]int, numMeta),
			make([]*node, children),
		}
		if children > 0 {
			ch, j := buildTreeRecur(nums, i, children)
			n.children = ch
			i = j
		}
		for k := 0; k < numMeta; k++ {
			n.metadata[k] = nums[i]
			i++
		}
		res = append(res, &n)
	}
	return res, i
}
