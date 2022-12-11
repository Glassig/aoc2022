package week2

import (
	"fmt"
	"sort"
)

//comments are part 1
func Day11() {
	listOfMonkeys := GetMonkeys()
	amountOfInspections := []int{}
	commonModulo := 1
	for id := 0; id < len(listOfMonkeys); id++ {
			commonModulo *= listOfMonkeys[id].divisible
		}
	// for iter := 1; iter <= 20 ; iter++ {
	for iter := 1; iter <= 10000 ; iter++ { 
		for id := 0; id < len(listOfMonkeys); id++ {
			monkey := listOfMonkeys[id]
			if len(amountOfInspections) <= id {
				amountOfInspections = append(amountOfInspections, 0)
			}
			amountOfInspections[id] += len(monkey.items)
			for len(listOfMonkeys[id].items) > 0 {
				item := listOfMonkeys[id].items[0]
				listOfMonkeys[id].items = listOfMonkeys[id].items[1:]
				// itemNewWorry := monkey.operation(item) / 3
				itemNewWorry := monkey.operation(item) % commonModulo
				if monkey.test(itemNewWorry) {
					listOfMonkeys[monkey.ifTrue].items = append(listOfMonkeys[monkey.ifTrue].items, itemNewWorry)
				} else {
					listOfMonkeys[monkey.ifFalse].items = append(listOfMonkeys[monkey.ifFalse].items, itemNewWorry)
				}
			}
		}
		if iter % 1000 == 0 || iter == 1 || iter == 20 {
			fmt.Println("After round ", iter, ": ", amountOfInspections)
		}
	}
	sort.Ints(amountOfInspections)
	fmt.Println("Answer: ", amountOfInspections[len(amountOfInspections) - 1] * amountOfInspections[len(amountOfInspections) - 2])
}

