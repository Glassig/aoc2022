package week1

import (
	"fmt"
	"strconv"
	"strings"
)

func getStacks(arr []string) (int, [9][]byte) {
	endOfSetupIndex := 0
	stacks := [9][]byte{}
	for index, value := range arr {
		if value[1] == 49 { // begin number
			endOfSetupIndex = index
			break
		}
		for a, i := 0, 1; i < 34; i = i + 4 {
			if value[i] != 32 {
				stacks[a] = append([]byte{value[i]}, stacks[a]...)
			}
			a++
		}
	}
	return endOfSetupIndex, stacks
}

func getInt(str string) int {
	ret, err := strconv.Atoi(str)
	Check(err)
	return ret
}

func day5StackSimple(endOfSetupIndex int, stacks [9][]byte, arr []string) [9][]byte {
	for index := endOfSetupIndex +2; index < len(arr); index++ {
		values := strings.Split(arr[index], " ")
		amount, from, to := getInt(values[1]), getInt(values[3])-1, getInt(values[5])-1
		for i := 0; i < amount; i++ {
			popped := stacks[from][len(stacks[from])-1]
			stacks[from] = stacks[from][:len(stacks[from])-1]
			stacks[to] = append(stacks[to], popped)
		}
	}
	return stacks
}

func day5StackStack(endOfSetupIndex int, stacks [9][]byte , arr []string) [9][]byte {
	for index := endOfSetupIndex +2; index < len(arr); index++ {
		values := strings.Split(arr[index], " ")
		amount, from, to := getInt(values[1]), getInt(values[3])-1, getInt(values[5])-1
		popped := stacks[from][len(stacks[from])-amount:]
		stacks[from] = stacks[from][:len(stacks[from])-amount]
		stacks[to] = append(stacks[to], popped...)
	}
	return stacks
}

func day5() {
	input, fileScanner := GetInput("week1/input_5.txt")
	defer input.Close()
	arr := ConvertFileScannerToArr(fileScanner)
	endOfSetupIndex, stacks := getStacks(arr)
	one := day5StackSimple(endOfSetupIndex, stacks, arr)
	two := day5StackStack(endOfSetupIndex, stacks, arr)
	fmt.Print("Answer part 1 is: ")
	for i := 0; i < 9; i++ {
		fmt.Print(string(one[i][len(one[i])-1]))
	}
	fmt.Println()
	fmt.Print("Answer part 2 is: ")
	for i := 0; i < 9; i++ {
		fmt.Print(string(two[i][len(two[i])-1]))
	}
}