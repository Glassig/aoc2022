package week3

import (
	"aoc2022/week1"
	"fmt"
	"math"
	"regexp"
	"strconv"

	"golang.org/x/exp/maps"
)

func Day15() {
	// day15Part1()
	day15Part2()
}

func day15Part1() {
	// for each beacon, with the range of sensor, can it affect y=2000000?
	input, fileScanner := week1.GetInput("week3/input_15.txt")
	defer input.Close()
	re := regexp.MustCompile(`[-]?\d+`)
	rowToCheck := make(map[int]bool)
	row := 2000000
	beaconsInRow := make(map[int]bool)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := (re.FindAllString(line, -1))
		x1, _ := strconv.Atoi(numbers[0])
		y1, _ := strconv.Atoi(numbers[1])
		x2, _ := strconv.Atoi(numbers[2])
		y2, _ := strconv.Atoi(numbers[3])
		dist := manHattanDist(x1, y1, x2, y2)
		if (isSensorRelevant(y1, dist)) {
			for x:= x1 - dist + int(math.Abs(float64(row - y1))); x <= x1 + dist - int(math.Abs(float64(row - y1))); x++ {
				rowToCheck[x]=true
			}
		}
		if y2 == 2000000 {
			beaconsInRow[x2] = true
		}
	}
	// 6275922
	fmt.Println("Answer part 1:", len(rowToCheck) - len(beaconsInRow))
}

func manHattanDist(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1 - x2)) + math.Abs(float64(y1 - y2)))
}

func isSensorRelevant(y, beacondDist int) bool { 
	return (y - beacondDist) <= 2000000 && 2000000 <= (y + beacondDist)
}

// var maxCoord = 20
var maxCoord = 4000000
func day15Part2() {
	input, fileScanner := week1.GetInput("week3/input_15.txt")
	defer input.Close()
	// max := maxCoord
	re := regexp.MustCompile(`[-]?\d+`)
	possiblyCoords := make(map[Coords]bool)
	sensors := []DistCoords{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := re.FindAllString(line, -1)
		// add all coordinates for outside the diamond
		// if already added, remove. check at end
		x1, _ := strconv.Atoi(numbers[0])
		y1, _ := strconv.Atoi(numbers[1])
		x2, _ := strconv.Atoi(numbers[2])
		y2, _ := strconv.Atoi(numbers[3])
		dist := manHattanDist(x1, y1, x2, y2)
		sensors = append(sensors, DistCoords{x:x1, y:y1, dist: dist})
		ydiff := 1
		if x1 - 1 - dist >= 0 {
			possiblyCoords[Coords{x: x1 - 1 - dist, y: y1}] = true
		}
		for x := getMinOrZero(x1, dist); x < getMaxOrLimit(x1, dist); x++ {
			possiblyCoords[Coords{x: x, y: y1 - ydiff}] = true
			possiblyCoords[Coords{x: x, y: y1 + ydiff}] = true
			if x <= x1 {
				ydiff++
			} else {
				ydiff--
			}
		}
		if x1 + 1 + dist <= maxCoord {
			possiblyCoords[Coords{x: x1 + 1 + dist, y: y1}] = true
		}
	}

	for _, point := range maps.Keys(possiblyCoords) {
		if point.y < 0 || point.y > maxCoord {
			possiblyCoords[point] = false
			continue
		}
		for _, sensor := range sensors {
			if manHattanDist(sensor.x, sensor.y, point.x, point.y) <= sensor.dist {
				possiblyCoords[point] = false
				break
			}
		}
		if possiblyCoords[point] {
			fmt.Println("Answer part 2:", point.x * maxCoord + point.y)
			break
		}
	}
	
}

func getMinOrZero(pos, dist int) int {
	if pos - dist < 0 {
		return 0
	}
	return pos - dist
}

func getMaxOrLimit(pos, dist int) int {
	if pos + dist > maxCoord {
		return maxCoord
	}
	return pos + dist
}


type Coords struct {
	x, y int
}

type DistCoords struct {
	x, y, dist int
}

