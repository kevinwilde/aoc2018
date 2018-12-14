package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// GetInput get each line from file as string array
func GetInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arr []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}

// ParseInt takes a string, returns int
func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
