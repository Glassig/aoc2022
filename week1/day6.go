package week1

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/maps"
)

func findFirstBlockOfUnique(blockSize int, fileScanner *bufio.Scanner) int {
	fileScanner.Split(bufio.ScanRunes)
	bytes := []string{}
	for fileScanner.Scan() {
		bytes = append(bytes, fileScanner.Text())
	}

	for index, value := range bytes {
		uniqueKeys := make(map[string]bool)
		uniqueKeys[value] = true
		for i := index + 1; i < index + blockSize; i++ {
			if uniqueKeys[bytes[i]] {
				break
			}
			uniqueKeys[bytes[i]] = true
		}
		if len(maps.Keys(uniqueKeys)) == blockSize {
			return index + blockSize
		}
	}
	return 0
}

func day6() {
	input, err := os.Open("input_6.txt")
	Check(err)
	fileScanner := bufio.NewScanner(input)
	// fmt.Println("Answer part1: ", findFirstBlockOfUnique(4, fileScanner))
	fmt.Println("Answer part2: ", findFirstBlockOfUnique(14, fileScanner))
	
}