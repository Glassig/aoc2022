package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput(path string) (*os.File, *bufio.Scanner) {
	input, err := os.Open(path)
	check(err)
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	return input, fileScanner
}

func convertFileScannerToArr(fs *bufio.Scanner) []string {
	result := []string{}
	for fs.Scan() {
		result = append(result, fs.Text())
	}
	return result
}

func main() {
	day4()
}
