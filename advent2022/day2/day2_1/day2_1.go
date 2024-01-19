package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func parseInput(filename string) [][2]string {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	var fileLines [][2]string

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var pair [2]string
		idx := 0
		for _, val := range fileScanner.Text() {
			if val != ' ' {
				pair[idx] = string(val)
				idx++
			}
		}
		fileLines = append(fileLines, pair)
	}
	return fileLines
}

func getScore(pairs [][2]string) int {
	var score int = 0
	var scoreMap = map[[2]string]int{{"A", "X"}: 3, {"A", "Y"}: 6, {"A", "Z"}: 0,
		{"B", "X"}: 0, {"B", "Y"}: 3, {"B", "Z"}: 6,
		{"C", "X"}: 6, {"C", "Y"}: 0, {"C", "Z"}: 3}
	var scoreByChoice = map[string]int{"X": 1, "Y": 2, "Z": 3}
	for _, pair := range pairs {
		score += scoreMap[pair]
		score += scoreByChoice[pair[1]]
	}
	return score
}

func main() {
	t0 := time.Now()
	test := false
	var filename string
	if test {
		filename = "../day2-test1-input.txt"
	} else {
		filename = "../day2-input.txt"
	}
	pairs := parseInput(filename)
	score := getScore(pairs)
	fmt.Println(score, len(pairs))
	t1 := time.Now()
	fmt.Printf("It took %v to run.\n", t1.Sub(t0))
}
