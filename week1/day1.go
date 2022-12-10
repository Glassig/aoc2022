package week1

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

func getArr(fileScanner *bufio.Scanner) []int {
	list, current := []int{}, 0
	for fileScanner.Scan() {
		tmp := fileScanner.Text()
		if tmp == "" {
			list = append(list, current)
			current = 0
		} else {
			tmpInt, err := strconv.Atoi(tmp)
			Check(err)
			current += tmpInt
		}
	}
	return list
}

func day1() {
	input, fileScanner := GetInput("week1/input_1.txt")
	defer input.Close()
	list := getArr(fileScanner)
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println("first: ", list[0])
	fmt.Println("second: ", list[0] + list[1] + list[2])
}
