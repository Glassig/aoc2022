package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func part1(input *os.File) {
	most, current := 0, 0

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		tmp := fileScanner.Text()
		if tmp == "" {
			if current > most {
				most = current
			}
			current = 0
		} else {
			tmpInt, err := strconv.Atoi(tmp)
			check(err)
			current += tmpInt
		}
	}
	
	fmt.Println(most)
}

func part2(input *os.File) {
	most, current := []int{0, 0, 0}, 0

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		tmp := fileScanner.Text()
		if tmp == "" {
			sort.Ints(most)
			if current > most[0] {
				most[0] = current
			}
			current = 0
		} else {
			tmpInt, err := strconv.Atoi(tmp)
			check(err)
			current += tmpInt
		}
	}
	
	fmt.Println(most, most[0] + most[1] + most[2])
}

func main1() {
	input, err := os.Open("input_1.txt")
	defer input.Close()
	check(err)
	// part1(input)
	part2(input)
}
