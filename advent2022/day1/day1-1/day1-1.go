package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func parseInput(filename string) []string {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	var fileLines []string

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}

func maxCalories(calories []string) int {
	var total int = 0
	var maxVal int = -int(math.Pow(10, 6))
	for _, line := range calories {
		if line == "" {
			if total > maxVal {
				maxVal = total
			}
			total = 0
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("There was an error reading data")
			break
		} else {
			total += value
		}
	}
	return maxVal
}

func main() {
	t0 := time.Now()
	test := false
	var filename string
	if test {
		filename = "../day1-test1-input.txt"
	} else {
		filename = "../day1-input.txt"
	}
	calories := parseInput(filename)
	maxCal := maxCalories(calories)
	fmt.Println(maxCal)
	t1 := time.Now()
	fmt.Printf("It took %v to run.\n", t1.Sub(t0))
}
