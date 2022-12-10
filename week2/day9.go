package week2

import (
	week1 "aoc2022/week1"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coords struct {
	x, y int
}

func Day9() {
	input, fileScanner := week1.GetInput("week2/input_9.txt")
	defer input.Close()
	arr := week1.ConvertFileScannerToArr(fileScanner)
	visited := make(map[Coords]bool)

	head, tail := Coords{x: 0, y: 0}, Coords{x: 0, y: 0}
	rope := [8]Coords{}
	visited[tail] = true
	for _, motion := range arr {
		// day9Part1(motion, &head, &tail, &visited)
		day9Part2(motion, &head, &tail, &visited, &rope)
	}

	fmt.Println("Answer: ", len(visited))
}

func moveHead(direction string, amount int, head *Coords) {
	switch direction {
		case "U":
			(*head).y += amount
		case "D": 
			(*head).y -= amount
		case "L": 
			(*head).x -= amount
		case "R": 
			(*head).x += amount
		default:
			fmt.Println("oooops")
	}
}

func day9Part1(motion string, head, tail *Coords, visited *map[Coords]bool) {
	instruction := strings.Split(motion, " ")
	direction := instruction[0]
	amount, _ := strconv.Atoi(instruction[1])
	moveHead(direction, amount, head)
	moveTailPart1(tail, visited, (*head))
}

func day9Part2(motion string, head, tail *Coords, visited *map[Coords]bool, rope *[8]Coords) {
	instruction := strings.Split(motion, " ")
	direction := instruction[0]
	amount, _ := strconv.Atoi(instruction[1])
	for iter := 0; iter < amount; iter++{
		moveHead(direction, 1, head)
		moveTailPart2(tail, visited, (*head), rope)
	}
}

func moveTailPart1(tail *Coords, visited *map[Coords]bool, head Coords) {
	if isTwoCoordsTouching(head, (*tail)) {
		return
	}
	moveOneStep(head, tail)
	(*visited)[(*tail)] = true
	moveTailPart1(tail, visited, head)
}

func moveOneStep(head Coords, tail *Coords) {
	if head.x == (*tail).x {
		if head.y < (*tail).y {
			(*tail).y--
		} else {
			(*tail).y++
		}
	} else if head.y == (*tail).y {
		if head.x < (*tail).x {
			(*tail).x--
		} else {
			(*tail).x++
		}
	} else {
		if head.x > (*tail).x {
			(*tail).x++
		} else {
			(*tail).x--
		}
		if head.y > (*tail).y {
			(*tail).y++
		} else {
			(*tail).y--
		}
	}
}

func moveOneLinkToAnother(head Coords, tail *Coords) {
	if isTwoCoordsTouching(head, (*tail)) {
		return
	}
	moveOneStep(head, tail)
	moveOneLinkToAnother(head, tail)
}

func moveTailPart2(tail *Coords, visited *map[Coords]bool, head Coords, rope *[8]Coords) {
	moveOneLinkToAnother(head, &rope[0])
	for index := 1; index < 8; index++ {
		moveOneLinkToAnother((*rope)[index-1], &rope[index])
	}
	for !isTwoCoordsTouching((*rope)[7], (*tail)) {
		moveOneStep((*rope)[7], tail)
		(*visited)[(*tail)] = true
	}
}


func isTwoCoordsTouching(head, tail Coords) bool {
	xdiff, ydiff := math.Abs(float64(head.x - tail.x)), math.Abs(float64(head.y - tail.y))
	return xdiff <= 1 && ydiff <= 1
}