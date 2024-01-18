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

func topThreeCalories(calories []string) int {
	var total int = 0
  var tmp int = -int(math.Pow(10, 6))
	var maxes [3]int = [3]int{tmp,tmp, tmp}
	for _, line := range calories {
		if line == "" {
			for idx, val := range maxes {
				if val < total {
					for i := 2; i > idx; i-- {
						maxes[i] = maxes[i-1]
					}
					maxes[idx] = total
					break
				}
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
	var net int = 0
	for i := 0; i < len(maxes); i++ {
		net += maxes[i]
	}
	return net
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
	maxCal := topThreeCalories(calories)
	fmt.Println(maxCal)
  t1 := time.Now()
  fmt.Printf("It tooks %v to run\n", t1.Sub(t0))
}
