package main

import (
	"fmt"
	"strings"

	"github.com/harikaduyu/advent-of-code-2023/utils"
)

func parseRound(line string) int {
	// create a set of winning numbers
	winningNumbers := make(map[string]struct{})
	sumWinningCards := 0
	twoParts := strings.Split(line, ":")[1]
	splitted := strings.Split(twoParts, "|")
	// numbers before | are winning numbers
	for _, numStr := range strings.Fields(splitted[0]) {
		winningNumbers[numStr] = struct{}{}
	}
	// numbers after | are numbers you have
	for _, numYouHave := range strings.Fields(splitted[1]) {
		_, win := winningNumbers[numYouHave]
		if win {
			sumWinningCards++
		}
	}
	return sumWinningCards
}
func getNumWinningsPerRound(input string) []int {
	var numWinningsPerRound []int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sumWinningCards := parseRound(line)
		numWinningsPerRound = append(numWinningsPerRound, sumWinningCards)
	}
	return numWinningsPerRound
}

func part1(numWinningsPerRound []int) int {
	totalPoints := 0
	for _, matches := range numWinningsPerRound {
		if matches > 0 {
			totalPoints += 1 << matches >> 1
		}
	}
	return totalPoints
}

func sumPoints(copyCards []int) int {
	sum := 0
	for i := range copyCards {
		sum += copyCards[i]
	}
	return sum
}

func part2(numWinningsPerRound []int) int {
	cardCopies := make([]int, len(numWinningsPerRound))
	// initialize
	for i := range numWinningsPerRound {
		cardCopies[i] = 1
	}
	for cardNr, matches := range numWinningsPerRound {
		for i := cardNr + 1; i < cardNr+matches+1 && i < len(numWinningsPerRound); i++ {
			cardCopies[i] += cardCopies[cardNr]
		}
	}

	return sumPoints(cardCopies)
}

func main() {

	input := utils.ReadInput()
	numWinnings := getNumWinningsPerRound(input)

	res1 := part1(numWinnings)
	fmt.Println("Part1:", res1)

	res2 := part2(numWinnings)
	fmt.Println("Part2:", res2)
}
