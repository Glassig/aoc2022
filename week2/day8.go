package week2

import (
	week1 "aoc2022/week1"
	"fmt"
	"strconv"
	"strings"
)

func getHighestInColumnBelow(matrix [][]string, row, col int) int {
	maxnum := 0
	for index := row + 1; index < len(matrix); index++ {
		value, _ := strconv.Atoi(matrix[index][col])
		if value > maxnum {
			maxnum = value
		}
	}
	return maxnum
}

func getHighestInRowAhead(rowArr []string, start int) int {
	maxnum := 0
	for _, value := range rowArr[start:] {
		valueInt, _ := strconv.Atoi(value)
		if valueInt > maxnum {
			maxnum = valueInt
		}
	}
	return maxnum
}

func day8Part1(matrix [][]string) {
	amountVisible := len(matrix) * 2 + (len(matrix[0]) - 2) * 2
	topsColumn := [99]int{}
	for index, value := range matrix[0] {
		topsColumn[index], _ = strconv.Atoi(value)
	}
	for row := 1; row < len(matrix)-1; row++ {
		topsLeft, _ := strconv.Atoi(matrix[row][0])
		for col := 1; col < len(matrix[row])-1; col++ {
			value, _ := strconv.Atoi(matrix[row][col])
			if value > topsLeft || value > topsColumn[col] || value > getHighestInColumnBelow(matrix, row, col) || value > getHighestInRowAhead(matrix[row], col + 1) {
				amountVisible++
				if value > topsLeft {
					topsLeft = value
				}
				if value > topsColumn[col] {
					topsColumn[col] = value
				}
			}
		}
	}
	fmt.Println("Answer part 1 is: ", amountVisible) // 1854
}

func day8Part2(matrix [][]string) {
	maxVisibility := 0
	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[row])-1; col++ {
			value, _ := strconv.Atoi(matrix[row][col])
			left, top, right, bottom := 0, 0, 0, 0
			for index := col - 1; index >= 0; index-- {
				left++
				val, _ := strconv.Atoi(matrix[row][index])
				 if val >= value {
					break
				 }
			}
			for index := col + 1; index < len(matrix[row]); index++ {
				right++
				val, _ := strconv.Atoi(matrix[row][index])
				 if val >= value {
					break
				 }
			}
			for index := row - 1; index >= 0; index-- {
				top++
				val, _ := strconv.Atoi(matrix[index][col])
				 if val >= value {
					break
				 }
			}
			for index := row + 1; index < len(matrix[row]); index++ {
				bottom++
				val, _ := strconv.Atoi(matrix[index][col])
				 if val >= value {
					break
				 }
			}
			maxVisibility = max(maxVisibility, left * top * right * bottom)
		}
	}

	fmt.Println("Answer part 2 is: ", maxVisibility) //527340
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func day8() {
	input, fileScanner := week1.GetInput("week2/input_8.txt")
	defer input.Close()
	matrix := [][]string{}
	for fileScanner.Scan() {
		arr := strings.Split(fileScanner.Text(), "")
		matrix = append(matrix, arr)
	}
	day8Part2(matrix)
}