package week3

import (
	"aoc2022/week1"
	"fmt"
	"strconv"
	"strings"
)

func Day21() {
	input, fileScanner := week1.GetInput("week3/input_21.txt")
	defer input.Close()
	commands := make(map[string]string)
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), ": ")
		commands[line[0]] = line[1]
	}
	fmt.Println("Answer part 1: ", day21Part1(commands, "root"))
	fmt.Println("Answer part 2: ", day21Part2(commands))
}

func day21Part1(commands map[string]string, key string) int {
	// if key == "humn" {
	// 	fmt.Println("humn is find")
	// }
	command := commands[key]
	number, err := strconv.Atoi(command)
	if err == nil {
		return number
	}
	parts := strings.Split(command, " ")
	switch parts[1] {
		case "+":
			return day21Part1(commands, parts[0]) + day21Part1(commands, parts[2])
		case "*":
			return day21Part1(commands, parts[0]) * day21Part1(commands, parts[2])
		case "/":
			return day21Part1(commands, parts[0]) / day21Part1(commands, parts[2])
		case "-":
			return day21Part1(commands, parts[0]) - day21Part1(commands, parts[2])
		default:
			fmt.Println("unknown: ", parts)
	}
	return 0
}

func day21Part2(commands map[string]string) int {
	root := strings.Split(commands["root"], " ")
	right := day21Part1(commands, root[2])
	// left := day21Part1(commands, root[0]) // humn is here
	low := 0
	high := 20
	val := right
	left := 0
	for {
		commands["humn"] = strconv.Itoa(val)
		left = day21Part1(commands, root[0])
		if left == right {
			break
		} else if left < right {
			low = val + 1
		} else {
			high = val - 1
		}
		val = (low + high) / 2
	}
	return val
}
