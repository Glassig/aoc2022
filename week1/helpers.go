package week1

import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInput(path string) (*os.File, *bufio.Scanner) {
	input, err := os.Open(path)
	Check(err)
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	return input, fileScanner
}

func ConvertFileScannerToArr(fs *bufio.Scanner) []string {
	result := []string{}
	for fs.Scan() {
		result = append(result, fs.Text())
	}
	return result
}
