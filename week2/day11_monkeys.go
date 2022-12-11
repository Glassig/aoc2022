package week2

func GetMonkeys() []monkey {
	return []monkey{
		{
			id:        0,
			items:     []int{84, 72, 58, 51},
			operation: func(old int) int { return old * 3 },
			divisible: 13,
			ifTrue:    1,
			ifFalse:   7,
		},
		{
			id:        1,
			items:     []int{88, 58, 58},
			operation: func(old int) int { return old + 8 },
			divisible: 2,
			ifTrue:    7,
			ifFalse:   5,
		},
		{
			id:        2,
			items:     []int{93, 82, 71, 77, 83, 53, 71, 89},
			operation: func(old int) int { return old * old },
			divisible: 7,
			ifTrue:    3,
			ifFalse:   4,
		},
		{
			id:        3,
			items:     []int{81, 68, 65, 81, 73, 77, 96},
			operation: func(old int) int { return old + 2 },
			divisible: 17,
			ifTrue:    4,
			ifFalse:   6,
		},
		{
			id:        4,
			items:     []int{75, 80, 50, 73, 88},
			operation: func(old int) int { return old + 3 },
			divisible: 5,
			ifTrue:    6,
			ifFalse:   0,
		},
		{
			id:        5,
			items:     []int{59, 72, 99, 87, 91, 81},
			operation: func(old int) int { return old * 17 },
			divisible: 11,
			ifTrue:    2,
			ifFalse:   3,
		},
		{
			id:        6,
			items:     []int{86, 69},
			operation: func(old int) int { return old + 6 },
			divisible: 3,
			ifTrue:    1,
			ifFalse:   0,
		},
		{
			id:        7,
			items:     []int{91},
			operation: func(old int) int { return old + 1 },
			divisible: 19,
			ifTrue:    2,
			ifFalse:   5,
		},
	}
}

func GetTestMonkeys() []monkey {
	return []monkey{
		{
			id:        0,
			items:     []int{79, 98},
			operation: func(old int) int { return old * 19 },
			divisible: 23,
			ifTrue:    2,
			ifFalse:   3,
		},
		{
			id:        1,
			items:     []int{54, 65, 75, 74},
			operation: func(old int) int { return old + 6 },
			divisible: 19,
			ifTrue:    2,
			ifFalse:   0,
		},
		{
			id:        2,
			items:     []int{79, 60, 97},
			operation: func(old int) int { return old * old },
			divisible: 13,
			ifTrue:    1,
			ifFalse:   3,
		},
		{
			id:        3,
			items:     []int{74},
			operation: func(old int) int { return old + 3 },
			divisible: 17,
			ifTrue:    0,
			ifFalse:   1,
		},
	}
}

type monkey struct {
	id        int
	items     []int
	operation operation
	divisible int
	ifTrue    int
	ifFalse   int
}

type operation func(int) int

func (m monkey) test(item int) bool {
	return item%m.divisible == 0
}