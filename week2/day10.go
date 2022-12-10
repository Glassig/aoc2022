package week2

import (
	week1 "aoc2022/week1"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day10() {
	input, fileScanner := week1.GetInput("week2/input_10.txt")
	defer input.Close()
	arr := week1.ConvertFileScannerToArr(fileScanner)
	clock, x, sum := 0, 1, 0
	fmt.Println("Part 2: ")
	for _, value := range arr {
		command := strings.Split(value, " ")
		checkClockAndXPart2(clock, x)
		clock++
		checkClockAndXPart1(clock, x, &sum)
		if command[0] == "noop" {
			continue
		}
		number, _ := strconv.Atoi(command[1])
		checkClockAndXPart2(clock, x)
		clock++
		checkClockAndXPart1(clock, x, &sum)
		x += number
	}
	fmt.Println()
	fmt.Println("Part 1: ", sum)
}

func checkClockAndXPart1(clock, x int, sum *int) {
	if clock == 20 || clock == 60 || clock == 100 || clock == 140 || clock == 180 || clock == 220 {
		(*sum) += clock * x
	}
}

func checkClockAndXPart2(clock, x int) {
	if math.Abs(float64((clock % 40) - x)) <= 1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if (clock+1) % 40 == 0 {
		fmt.Println()
	}
}