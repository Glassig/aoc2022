package week2

import (
	week1 "aoc2022/week1"
	"fmt"
	"strings"
)

func Day12() {
	input, fileScanner := week1.GetInput("week2/input_12.txt")
	defer input.Close()
	matrix := [][]string{}
	start, goal := Coords{x: 0, y: 0}, Coords{x: 0, y: 0}
	row := 0
	for fileScanner.Scan() {
		arr := strings.Split(fileScanner.Text(), "")
		matrix = append(matrix, arr)
		for index, value := range arr {
			if value == "S" {
				start.x, start.y = row, index
			} else if value == "E" {
				goal.x, goal.y = row, index
			}
		}
		row++
	}
	//lets go from end to start
	fmt.Print("Part 1")
	bfs(matrix, goal, "S")
	fmt.Print("Part 2")
	bfs(matrix, goal, "a")
}

type GraphPoint struct {
	visited bool
	distance int
}

func bfs(matrix [][]string, start Coords, goal string) {
	queue, visited := []Coords{start}, [41][173]GraphPoint{}
	visited[start.x][start.y].visited = true
	visited[start.x][start.y].distance = 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		x, y := current.x, current.y
		if matrix[x][y] == goal {
			fmt.Println("Found it! x:", current.x, ", y:", current.y, "Distance:", visited[x][y].distance)
			return
		}
		if x > 0 && !visited[x - 1][y].visited && canMoveToPoint(matrix[x][y], matrix[x-1][y]) {
			visited[x - 1][y] = GraphPoint{visited: true, distance: visited[x][y].distance + 1}
			queue = append(queue, Coords{x: x - 1, y: y})
		}
		if x < 40 && !visited[x + 1][y].visited && canMoveToPoint(matrix[x][y], matrix[x+1][y]) {
			visited[x + 1][y] = GraphPoint{visited: true, distance: visited[x][y].distance + 1}
			queue = append(queue, Coords{x: x + 1, y: y})
		}
		if y > 0 && !visited[x][y - 1].visited && canMoveToPoint(matrix[x][y], matrix[x][y-1]) {
			visited[x][y - 1] = GraphPoint{visited: true, distance: visited[x][y].distance + 1}
			queue = append(queue, Coords{x: x, y: y - 1})
		}
		if y < 172 && !visited[x][y + 1].visited && canMoveToPoint(matrix[x][y], matrix[x][y+1]) {
			visited[x][y + 1] = GraphPoint{visited: true, distance: visited[x][y].distance + 1}
			queue = append(queue, Coords{x: x, y: y + 1})
		}
	}
}

func canMoveToPoint(current, next string) bool {
	if (next == "S" && current == "a") || (next == "S" && current == "b") {
		return true
	}
	if current == "E" && next != "z" {
		return false
	}
	return next[0] >= current[0] || current[0] - next[0] == 1
}