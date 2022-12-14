package week2

import (
	week1 "aoc2022/week1"
	"fmt"
	"strings"
	"strconv"
)

func Day14() {
	cave, topLeft, bottomRight := getCaveAndCorners("week2/input_14.txt")
	// fmt.Printf("topleft: (x:%d, y:%d)\n", topLeft.x, topLeft.y)
	// fmt.Printf("bottomRight: (x:%d, y:%d)\n", bottomRight.x, bottomRight.y)
	// create moving sand calc. when one get to pos.x < topleft.x or pos.x > bottomright.x, it's outside
	sandUnits := 0
	for !addsandToCave(topLeft.x, bottomRight.x, bottomRight.y, &cave) {
		sandUnits++
	}
	fmt.Println("Part1 ", sandUnits)
	setPath(Coords{x: 0, y: bottomRight.y + 2}, Coords{x: 1000, y: bottomRight.y + 2}, &cave)
	for !addSandToCaveHitSource(&cave) {
		sandUnits++
	}
	fmt.Println("Part2 ", sandUnits)
}

func getCoord(str string) Coords {
	coord:= strings.Split(str, ",")
	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])
	return Coords{x: x, y: y}
}

func getCaveAndCorners(filePath string) (cave map[Coords]bool, topLeft, bottomRight Coords) {
	input, fileScanner := week1.GetInput(filePath)
	defer input.Close()
	topLeft, bottomRight = Coords{x: 2147483647, y: 0}, Coords{x: 0, y: 0}
	cave = make(map[Coords]bool) // if true, it has something in it.
	for fileScanner.Scan() {
		rock := fileScanner.Text()
		path := strings.Split(rock, " -> ")
		for index, point := range path {
			coord := getCoord(point)
			if coord.x < topLeft.x {
				topLeft.x = coord.x
			} else if coord.x > bottomRight.x {
				bottomRight.x = coord.x
			}
			if coord.y > bottomRight.y {
				bottomRight.y = coord.y
			}
			if index == 0 {
				cave[coord] = true
			} else {
				setPath(getCoord(path[index - 1]), coord, &cave)
			}
		}
	}
	return cave, topLeft, bottomRight
}

//not including first, but second
func setPath(first, second Coords, cave *map[Coords]bool) {
	if first.x == second.x {
		// change y
		if first.y > second.y {
			// go from second to first
			for y := first.y - 1; y >= second.y; y-- {
				(*cave)[Coords{x: first.x, y: y}] = true
			}
		} else {
			for y := first.y + 1; y <= second.y; y++ {
				(*cave)[Coords{x: first.x, y: y}] = true
			}
		}
	} else {
		// change x
		if first.x > second.x {
			// go from second to first
			for x := first.x - 1; x >= second.x; x-- {
				(*cave)[Coords{x: x, y: first.y}] = true
			}
		} else {
			for x := first.x + 1; x <= second.x; x++ {
				(*cave)[Coords{x: x, y: first.y}] = true
			}
		}
	}
}

func addsandToCave(leftx, rightx, bottomy int, cave *map[Coords]bool) (isOutside bool) {
	// create sand, let it fall. if x ever gets outside, return true, if stop falling, return false
	// sandSource "500,0"
	x, y := 500, 0
	for leftx <= x && x <= rightx && y < bottomy {
		y += 1
		if (*cave)[Coords{x: x, y: y}] {
			if !(*cave)[Coords{x: x - 1, y: y}] {
			// we move diagonal left
				x -= 1
			} else if !(*cave)[Coords{x: x + 1, y: y}] {
				// we move diagonal right
				x += 1
			} else {
				(*cave)[Coords{x: x, y: y-1}] = true
				return false
			}
		}
	}
	return true
}

func addSandToCaveHitSource(cave *map[Coords]bool) (hitSource bool) {
	x, y := 500, 0
	for !(*cave)[Coords{x: 500, y: 0}] {
		y += 1
		if (*cave)[Coords{x: x, y: y}] {
			if !(*cave)[Coords{x: x - 1, y: y}] {
			// we move diagonal left
				x -= 1
			} else if !(*cave)[Coords{x: x + 1, y: y}] {
				// we move diagonal right
				x += 1
			} else {
				(*cave)[Coords{x: x, y: y-1}] = true
				return false
			}
		}
	}
	return true
}
