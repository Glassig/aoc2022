package week3

import (
	"aoc2022/week1"
	"fmt"
	"strconv"
)

func Day20() {
	input, fileScanner := week1.GetInput("week3/input_20.txt")
	defer input.Close()
	arr := []crypted{}
	i := 0
	for fileScanner.Scan() {
		val, _ := strconv.Atoi(fileScanner.Text())
		arr = append(arr, crypted{value: val, valueBig: val * 811589153, moved: false, origPos: i})
		i++
	}
	// fmt.Println("Answer part 1: ", day20Part1(arr)) // 5904
	fmt.Println("Answer part 2: ", day20Part2(arr)) // 8332585833851
}

type crypted struct {
	value int
	moved bool
	valueBig int
	origPos int
}

func day20Part1(arr []crypted) int {
	maxIndex := len(arr) - 1
	for i := 0; i < len(arr); {
		popped := arr[i]
		if popped.moved {
			i++
			continue
		}
		arr = removeItemAtIndex(i, arr)
		newIndex := (i + popped.value) % maxIndex
		if newIndex <= 0 {
			newIndex += maxIndex
		}
		popped.moved = true
		arr = insertAtIndex(popped, newIndex, arr)
		if newIndex <= i {
			i++
		}
	}
	index := 0
	for i := 0; i <= maxIndex; i++ {
		if arr[i].value == 0 {
			fmt.Println("index for 0: ", i)
			index = i
			break
		}
	}
	fmt.Println(arr[(index + 1000) % (maxIndex + 1)].value, arr[(index + 2000) % (maxIndex + 1)].value, arr[(index + 3000) % (maxIndex + 1)].value)
	return arr[(index + 1000) % (maxIndex + 1)].value + arr[(index + 2000) % (maxIndex + 1)].value + arr[(index + 3000) % (maxIndex + 1)].value
}

func day20Part2(arr []crypted) int {
	// we need to keep track of the indexes, then we can iterate through it instead.
	maxIndex := len(arr) - 1
	for iter := 0; iter < 10; iter++ {
		for i := 0; i < len(arr); i++ {
			index := findIndex(arr, i)
			popped := arr[index]
			arr = removeItemAtIndex(index, arr)
			newIndex := (index + popped.valueBig) % maxIndex
			if newIndex < 0 {
				newIndex += maxIndex
			}
			arr = insertAtIndex(popped, newIndex, arr)
		}
	}
	index := 0
	for i := 0; i <= maxIndex; i++ {
		if arr[i].value == 0 {
			fmt.Println("index for 0: ", i)
			index = i
			break
		}
	}
	fmt.Println(arr[(index + 1000) % (maxIndex + 1)].valueBig, arr[(index + 2000) % (maxIndex + 1)].valueBig, arr[(index + 3000) % (maxIndex + 1)].valueBig)
	return arr[(index + 1000) % (maxIndex + 1)].valueBig + arr[(index + 2000) % (maxIndex + 1)].valueBig + arr[(index + 3000) % (maxIndex + 1)].valueBig
}

func printItems(arr []crypted) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].valueBig, ", ")
	}
	fmt.Println()
}

func removeItemAtIndex(index int, array []crypted)  []crypted {
	copy(array[index:], array[index+1:]) // Shift a[i+1:] left one index.
	// array[len(array)-1] = crypted{}    // Erase last element (write zero value).
	return array[:len(array)-1]     // Truncate slice.
}

func insertAtIndex(item crypted, index int, array []crypted) []crypted {
	array = append(array[:index + 1], array[index:]...)
	array[index] = item
	return array
}

func findIndex(arr []crypted, origPos int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i].origPos == origPos {
			return i
		}
	}
	return 0
}
