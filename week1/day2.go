// A = rock 65
// B = Paper 66
// C = Scissors 67

// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 if you lost, 3 if the round was a draw, and 6 if you won
package week1

import (
	"fmt"
)

func part2calc(s string) int {
	// X means loose 88
	// Y draw 89
	// Z win 90
	enemy, outcome := int(s[0]) - 64, (int(s[2]) - 88) * 3
	if outcome == 0 {
		if enemy == 1 {
			return outcome + 3
		}
		return outcome + enemy - 1
	} else if outcome == 3 {
		return outcome + enemy
	}
	return outcome + enemy % 3 + 1
}

func part1calc(s string) int {
	// X for Rock, 88
	// Y for Paper, 89
	// Z for Scissors 90
	enemy, mine := int(s[0]), int(s[2])
	myChoice := mine - 87
	if enemy == mine - 23 { // draw
		return 3 + myChoice
	} else if mine - enemy == 24 || mine - enemy == 21 { //victory
		return 6 + myChoice
	}
	return 0 + myChoice
}

func day2() {
	input, fileScanner := GetInput("week1/input_2.txt")
	defer input.Close()
	result1, result2, arr := 0, 0, ConvertFileScannerToArr(fileScanner)
	for _, value := range arr {
		result1 += part1calc(value)
		result2 += part2calc(value)
	}
	fmt.Println("1:", result1, ": 2:", result2)
}
