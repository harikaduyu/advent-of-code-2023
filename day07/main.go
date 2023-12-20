package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/harikaduyu/advent-of-code-2023/utils"
)

type hand struct {
	Bid, Score int
	Cards      string
}

func getSortedKeys(counter map[rune]int) []rune {
	sortedKeys := make([]rune, 0)
	for k := range counter {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return counter[sortedKeys[i]] > counter[sortedKeys[j]]
	})
	return sortedKeys
}

func getScore(cards string, part int) int {

	score := 0

	counter := make(map[rune]int)
	for _, char := range cards {

		counter[char]++

	}
	sortedKeys := getSortedKeys(counter)

	mostCommonNum := counter[sortedKeys[0]]
	for _, k := range sortedKeys {
		if part == 2 && len(counter) > 1 && k == 'J' {
			// Find the most common card which is not J
			// and add joker values to it
			if sortedKeys[0] == 'J' {
				mostCommonNum = counter[sortedKeys[1]] + counter[k]
			} else {
				mostCommonNum += counter[k]
			}
			delete(counter, 'J')
		}
	}

	switch len(counter) {
	case 1: // Five of a kind
		score = 7
	case 2:
		//Two possible cases
		// Four of a kind 4,1
		if mostCommonNum == 4 {
			score = 6
		}
		// Full house 3,2
		if mostCommonNum == 3 {
			score = 5
		}

	case 3:
		//Also two possible cases

		// Three of a kind 3,1,1
		if mostCommonNum == 3 {
			score = 4
		}
		// Two pair 2,2,1
		if mostCommonNum == 2 {
			score = 3
		}

	case 4: // One pair 2, 1,1,1
		score = 2
	case 5: // High card 1,1,1,1,1
		score = 1
	}

	var singleCardScores map[rune]int
	singleCardScores = map[rune]int{'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2}
	if part == 2 {

		singleCardScores['J'] = 1
	}
	for _, char := range cards {
		score = score*100 + singleCardScores[char]
	}
	return score
}

func parseHand(line string, part int) hand {
	h := hand{}
	splitted := strings.Fields(line)
	h.Cards = splitted[0]
	h.Bid, _ = strconv.Atoi(splitted[1])
	h.Score = getScore(h.Cards, part)
	return h
}

func parse(input string, part int) []hand {
	hands := make([]hand, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		hands = append(hands, parseHand(line, part))
	}
	return hands
}

func result(hands []hand) int {
	totalWinnings := 0
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Score <= hands[j].Score
	})
	for i, hand := range hands {
		rank := i + 1
		totalWinnings += rank * hand.Bid
	}
	return totalWinnings
}

func main() {

	input := utils.ReadInput()
	startTime := time.Now()

	p1_hands := parse(input, 1)
	res1 := result(p1_hands)
	endTime := time.Now()
	fmt.Println("Part1:", res1)
	fmt.Printf("Part1 took %v\n", endTime.Sub(startTime))

	startTime = time.Now()
	p2_hands := parse(input, 2)
	res2 := result(p2_hands)
	endTime = time.Now()
	fmt.Println("Part2:", res2)
	fmt.Printf("Part2 took %v\n", endTime.Sub(startTime))
}
