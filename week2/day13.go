package week2

import (
	week1 "aoc2022/week1"
	"encoding/json"
	"fmt"
	"sort"
)

func Day13() {
	input, fileScanner := week1.GetInput("week2/input_13.txt")
	defer input.Close()
	var lines []string
	pair := 1
	count := 0
	for fileScanner.Scan() {
		leftLine :=  fileScanner.Text()
		fileScanner.Scan()
		rightLine :=  fileScanner.Text()
		fileScanner.Scan()
		lines = append(lines, leftLine)
		lines = append(lines, rightLine)
		count += pair * compareLines(leftLine, rightLine)
		pair++
	}
	//p1: 5843
	fmt.Println("P1: ", count)
	lines = append(lines, "[[2]]", "[[6]]")
	sortLines(lines)
	p2 := 1
	for index, val := range lines {
		if val == "[[2]]" || val == "[[6]]" {
			p2 *= (index + 1)
		}
	}
	fmt.Println("p2: ", p2)
}

func compareLines(left, right string) int {
	var leftInter interface{}
	var rightInter interface{}
	json.Unmarshal([]byte(left), &leftInter)
	json.Unmarshal([]byte(right), &rightInter)
	leftList, _ := leftInter.([]any)
	rightList, _ := rightInter.([]any)
	returnValue := compareTwoLists(leftList, rightList)
	if returnValue != 1 {
		return 0
	}
	return 1
}

func compareTwoLists(left, right []any) int {
	for index := 0; index < len(left) && index < len(right); index++ {
		leftVal := left[index]
		rightVal := right[index]
		leftAsFloat, isLeftNumber := leftVal.(float64)
		rightAsFloat, isRightNumber := rightVal.(float64)
		if isLeftNumber && isRightNumber {
			if leftAsFloat < rightAsFloat {
				return 1
			}
			if leftAsFloat > rightAsFloat {
				return -1
			}
		} else {
			var leftList []any
			var rightList []any
			if !isLeftNumber {
				leftList = leftVal.([]any)
			} else {
				leftList = []any{leftVal}
			}
			if !isRightNumber {
				rightList = rightVal.([]any)
			} else {
				rightList = []any{rightVal}
			}
			ret := compareTwoLists(leftList, rightList)
			if ret != 0 {
				return ret
			}
		}
	}
	if len(left) < len(right) {
		return 1
	}
	if len(left) > len(right) {
		return -1
	}
	return 0
}

func sortLines(lines []string) []string {
	sort.Slice(lines, func(a, b int) bool {
		if compareLines(lines[a], lines[b]) == 1 {
			return true
		}
		return false
	})
	return lines
}

