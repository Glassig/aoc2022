package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coords struct {
	x, y int
}

func day9() {
	input, fileScanner := getInput("input_9.txt")
	defer input.Close()
	arr := convertFileScannerToArr(fileScanner)
	visited := make(map[Coords]bool)

	head, tail := Coords{x: 0, y: 0}, Coords{x: 0, y: 0}
	rope := [8]Coords{}
	visited[tail] = true
	for _, motion := range arr {
		
		moveHead(motion, &head)
		// moveTailPart1(&tail, &visited, head)
		moveTailPart2(&tail, &visited, head, &rope)
		// for i := 17; i >= -5; i-- {
		// 	for j := -11; j < 15; j++ {
		// 		if head.x == j && head.y == i {
		// 			fmt.Print("H")
		// 		} else if tail.x == j && tail.y == i {
		// 			fmt.Print("T")
		// 		} else if i == 0 && j == 0{
		// 			fmt.Print("s")
		// 		}else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
			// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", motion)
	}

	fmt.Println("Answer: ", len(visited))
}

func moveHead(motion string, head *Coords) {
	instruction := strings.Split(motion, " ")
	direction := instruction[0]
	amount, _ := strconv.Atoi(instruction[1])
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