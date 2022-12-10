package week1

import (
	"fmt"
	"strconv"
	"strings"
)

func getNumberRanges(ranges string) (int, int) {
	numbers := strings.Split(ranges, "-")
	lower, err := strconv.Atoi(numbers[0])
	higher, err2 := strconv.Atoi(numbers[1])
	Check(err)
	Check(err2)
	return lower, higher
}

func isRangesCompletelyOverlap(ranges string) bool {
	values := strings.Split(ranges, ",")
	firstL, firstH := getNumberRanges(values[0])
	secondL, secondH := getNumberRanges(values[1])
	if firstL == secondL || firstH == secondH {
		return true
	} else if firstL < secondL {
		return firstH >= secondH
	} else if firstL > secondL {
		return firstH <= secondH
	}
	return false
}

func isRangesSomeOverlap(ranges string) bool {
	values := strings.Split(ranges, ",")
	firstL, firstH := getNumberRanges(values[0])
	secondL, secondH := getNumberRanges(values[1])
	return isOneRangeBetweenTheOther(firstL, firstH, secondL, secondH) || isOneRangeBetweenTheOther(secondL, secondH, firstL, firstH)
}

func isOneRangeBetweenTheOther(first, second, lower, upper int) bool {
	return first >= lower && first <= upper || second >= lower && second <= upper
}

func day4() {
	input, fileScanner := GetInput("week1/input_4.txt")
	defer input.Close()
	arr := ConvertFileScannerToArr(fileScanner)
	count := 0
	for _, value := range arr {
		// if isRangesCompletelyOverlap(value) {
		if isRangesSomeOverlap(value) {
			count++
		}
	}
	fmt.Println("Answer: ", count)
}