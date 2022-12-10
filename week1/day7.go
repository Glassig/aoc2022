package week1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Dir struct {
	Up *Dir
	Files map[string]int
	Folders map[string]*Dir
}

func DirNew() *Dir {
	return &Dir{
		Up: nil,
		Files: make(map[string]int),
		Folders: make(map[string]*Dir),
	}
}

// $ cd <path>
// $ ls
// dir <name>
// <int> <filename>
func parseComand(command string, root *Dir, current **Dir) {
	switch command[0:4] {
		case "$ cd":
			if string(command[5]) == "/" {
				*current = root
			} else if string(command[5]) == "." {
				*current = (*current).Up
			} else {
				parts := strings.Split(command, " ")
				*current = (*current).Folders[parts[2]]  
			}
			
		case "$ ls":
			// we dont need to do anything
		case "dir ":
			parts := strings.Split(command, " ")
			newFolder := DirNew()
			newFolder.Up = (*current)
			(*current).Folders[parts[1]] = newFolder
		default:
			parts := strings.Split(command, " ")
			var errors error
			(*current).Files[parts[1]], errors = strconv.Atoi(parts[0])
			Check(errors)
	}
}

func sumBranches(branch *Dir, sums *[]int) int {
	currentSize := 0
	for _, file := range maps.Values(branch.Files) {
		currentSize += file
	}
	for _, folder := range maps.Values(branch.Folders) {
		currentSize += sumBranches(folder, sums)
	}
	*sums = append(*sums, currentSize)
	return currentSize
}

func sumFiles(branch *Dir) int {
	currentSize := 0
	for _, file := range maps.Values(branch.Files) {
		currentSize += file
	}
	for _, folder := range maps.Values(branch.Folders) {
		currentSize += sumFiles(folder)
	}
	return currentSize
}


func day7Part1(sums []int) int {
	sum := 0
	for _, size := range sums {
		if size <= 100000 {
			sum += size
		}
	}
	return sum
}

func day7Part2(root *Dir, sums []int) int {
	sum := sumFiles(root)
	freeSpace := 70000000 - sum
	spaceNeeded := 30000000 - freeSpace
	sort.Ints(sums)
	for _, size := range sums {
		if size >= spaceNeeded {
			return size
		}
	}
	return 0
}

func day7() {
	input, fileScanner := GetInput("week1/input_7.txt")
	defer input.Close()
	arr := ConvertFileScannerToArr(fileScanner)
	root := DirNew()
	current := root
	for _, value := range arr {
		parseComand(value, root, &current)
	}
	sums := []int{}
	for _, folder := range maps.Values(root.Folders) {
		sumBranches(folder, &sums)
	}
	fmt.Println("Answer part1: ", day7Part1(sums))
	fmt.Println("Answer part2: ", day7Part2(root, sums))
}