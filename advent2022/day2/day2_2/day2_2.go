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
  // Map according to strategy: X when p2 losses, Y when there is a draw, Z when p2 wins
  var choiceMap = map[string]map[string]string{"X":{"A": "C","B": "A","C": "B"},"Y":{"A": "A", "B": "B","C": "C"},"Z":{"A": "B","B": "C","C": "A"}}
	var scoreByStrategy = map[string]int{"X": 0, "Y": 3, "Z": 6}
	var scoreByChoice = map[string]int{"A": 1, "B": 2, "C": 3}
	for _, pair := range pairs {
		score += scoreByChoice[choiceMap[pair[1]][pair[0]]]
		score += scoreByStrategy[pair[1]]
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
	fmt.Println(score)
	t1 := time.Now()
	fmt.Printf("It took %v to run.\n", t1.Sub(t0))
}
