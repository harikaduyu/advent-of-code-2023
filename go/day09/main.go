package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
)

func parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	oasisReport := make([][]int, len(lines))
	for i, line := range lines {
		numStrs := strings.Fields(line)
		for _, ns := range numStrs {
			n, _ := strconv.Atoi(ns)
			oasisReport[i] = append(oasisReport[i], n)
		}
	}
	return oasisReport
}

func isSequenceFullZeroes(sequence []int) bool {
	for _, num := range sequence {
		if num != 0 {
			return false
		}
	}
	return true
}

func getSequences(line []int) [][]int {
	sequences := [][]int{line}

	for {
		nextSequence := make([]int, 0)
		for i := range sequences[len(sequences)-1][1:] {
			nextSequence = append(nextSequence, sequences[len(sequences)-1][i+1]-sequences[len(sequences)-1][i])
		}

		sequences = append(sequences, nextSequence)
		if isSequenceFullZeroes(nextSequence) {
			break
		}
	}
	return sequences
}

func getNextValOfHistory(line []int) int {
	sequences := getSequences(line)
	lastVal := 0

	for i := len(sequences) - 1; i >= 0; i-- {
		lastVal += sequences[i][len(sequences[i])-1]
	}
	return lastVal
}

func getPrevValOfHistory(line []int) int {
	sequences := getSequences(line)
	lastVal := 0
	for i := len(sequences) - 1; i >= 0; i-- {
		lastVal = sequences[i][0] - lastVal
	}
	return lastVal
}

func part1(oasisReport [][]int) int {
	sum := 0
	for _, line := range oasisReport {
		val := getNextValOfHistory(line)
		sum += val
	}
	return sum
}

func part2(oasisReport [][]int) int {
	sum := 0
	for _, line := range oasisReport {
		val := getPrevValOfHistory(line)
		sum += val
	}
	return sum
}

func main() {

	input := utils.ReadInput()
	startTime := time.Now()

	oasisReport := parse(input)
	res1 := part1(oasisReport)

	endTime := time.Now()
	fmt.Println("Part1:", res1)
	fmt.Printf("Part1 took %v\n", endTime.Sub(startTime))

	startTime = time.Now()
	res2 := part2(oasisReport)
	endTime = time.Now()
	fmt.Println("Part2:", res2)
	fmt.Printf("Part2 took %v\n", endTime.Sub(startTime))
}
