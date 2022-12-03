package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
)

func findElem(f, s []byte) byte {
	for _, value1 := range f {
		for _, value2 := range s {
			if (value1 == value2) {
				return value1
			}
		}
	}
	return 0
}

func findElements(f, s []byte) []byte {
	arr := make(map[byte]bool)
	for _, value1 := range f {
		if !arr[value1] {
			for _, value2 := range s {
				if (value1 == value2) {
					arr[value1] = true
					break
				}
			}
		}
	}
	return maps.Keys(arr)
}

func getBytePrio(b byte) int {
	if b < 91 {
		return int(b - 38)
	}
	return int(b - 96)
}

func day3Part1(fs *bufio.Scanner) {
	result := 0
	for fs.Scan() {
		line := fs.Text()
		first, second := []byte(line[0:len(line)/2]), []byte(line[len(line)/2:])
		elem := findElem(first, second)
		result += getBytePrio(elem)
	}
	fmt.Println("Answer: ", result)
}

func day3Part2(fs *bufio.Scanner) {
	result := 0
	for fs.Scan() {
		line1 := []byte(fs.Text())
		fs.Scan()
		line2 := []byte(fs.Text())
		fs.Scan()
		line3 := []byte(fs.Text())
		tmp := findElements(line1, line2)
		tmp = findElements(tmp, line3)
		result += getBytePrio(tmp[0])
	}
	fmt.Println("Answer: ", result)
}

// a, 97 = 1
// A, 65 = 27
// Z, 90 = 52

func main() {
	input, fileScanner := getInput("input_3.txt")
	defer input.Close()
	day3Part2(fileScanner)
}