package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func getStacks(arr []string) (endOfSetupIndex int, stacks []*list.List) {
	stacks = []*list.List{list.New(), list.New(), list.New(), list.New(), list.New(), list.New(), list.New(), list.New(), list.New()}

	for index, value := range arr {
		if value[1] == 49 { // begin number
			endOfSetupIndex = index
			break
		}
		if value[1] != 32 {
			stacks[0].PushFront(value[1])
		}
		if value[5] != 32 {
			stacks[1].PushFront(value[5])
		}
		if value[9] != 32 {
			stacks[2].PushFront(value[9])
		}
		if value[13] != 32 {
			stacks[3].PushFront(value[13])
		}
		if value[17] != 32 {
			stacks[4].PushFront(value[17])
		}
		if value[21] != 32 {
			stacks[5].PushFront(value[21])
		}
		if value[25] != 32 {
			stacks[6].PushFront(value[25])
		}
		if value[29] != 32 {
			stacks[7].PushFront(value[29])
		}
		if value[33] != 32 {
			stacks[8].PushFront(value[33])
		}
	}
	return
}

func getInt(str string) int {
	ret, err := strconv.Atoi(str)
	check(err)
	return ret
}

func day5StackSimple(endOfSetupIndex int, stacks []*list.List, arr []string) {
	for index := endOfSetupIndex +2; index < len(arr); index++ {
		values := strings.Split(arr[index], " ")
		amount, from, to := getInt(values[1]), getInt(values[3])-1, getInt(values[5])-1
		for i := 0; i < amount; i++ {
			val := stacks[from].Back()
			stacks[from].Remove(val)
			stacks[to].PushBack(val.Value)
		}
	}
}

func day5StackStack(endOfSetupIndex int, stacks []*list.List , arr []string) {
	for index := endOfSetupIndex +2; index < len(arr); index++ {
		values := strings.Split(arr[index], " ")
		amount, from, to := getInt(values[1]), getInt(values[3])-1, getInt(values[5])-1
		tempStack := list.New()
		for i := 0; i < amount; i++ {
			val := stacks[from].Back()
			stacks[from].Remove(val)
			tempStack.PushFront(val.Value)
		}
		stacks[to].PushBackList(tempStack)
	}
}

func day5() {
	input, fileScanner := getInput("input_5.txt")
	defer input.Close()
	arr := convertFileScannerToArr(fileScanner)
	endOfSetupIndex, stacks := getStacks(arr)
	// day5StackSimple(endOfSetupIndex, stacks, arr)
	day5StackStack(endOfSetupIndex, stacks, arr)
	
	for i := 0; i < 9; i++ {
		fmt.Print(stacks[i].Back().Value, " ")
	}
}