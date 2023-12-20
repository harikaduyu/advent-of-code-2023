package main

import (
	"fmt"
	"strings"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
)

type round struct {
	Red, Green, Blue int
}

var bagContents = round{12, 13, 14}

func extractRound(batch string) round {
	var round round
	for _, c := range strings.Split(batch, ",") {
		var cubes int
		var color string

		_, _ = fmt.Sscanf(c, "%d %s", &cubes, &color)

		switch color {
		case "red":
			round.Red = cubes
		case "blue":
			round.Blue = cubes
		case "green":
			round.Green = cubes
		}
	}
	return round
}

func extractRounds(line string) []round {
	splitGame := strings.Split(line, ":")[1]
	stringBatches := strings.Split(splitGame, ";")
	var rounds []round
	for _, batch := range stringBatches {
		round := extractRound(batch)
		rounds = append(rounds, round)
	}
	return rounds
}

func extractGameNumber(line string) (int, error) {
	var gameNumber int
	_, err := fmt.Sscanf(line, "Game %d:", &gameNumber)
	if err != nil {
		return 0, fmt.Errorf("Error searching for game number: %v", err)
	}

	return gameNumber, nil
}

func isRoundPossible(round round) bool {
	return round.Red <= bagContents.Red && round.Green <= bagContents.Green && round.Blue <= bagContents.Blue

}

func extractGamesRounds(input string) map[int][]round {
	lines := strings.Split(input, "\n")
	gamesRounds := make(map[int][]round, len(lines))
	for _, line := range lines {
		gameNumber, err := extractGameNumber(line)
		if err != nil {
			panic(err)
		}
		gamesRounds[gameNumber] = extractRounds(line)
	}
	return gamesRounds
}

func isGamePossible(rounds []round) bool {
	gamePossible := true
	for _, round := range rounds {
		if !isRoundPossible(round) {
			gamePossible = false
		}
	}
	return gamePossible
}

func part1(gamesRounds map[int][]round) int {
	total := 0

	for gameNumber := range gamesRounds {
		if isGamePossible(gamesRounds[gameNumber]) {
			total += gameNumber
		}
	}
	return total

}

func getPossibleBagContents(gamesRounds []round) round {
	var possibleBagContents round
	for _, round := range gamesRounds {
		if round.Blue > possibleBagContents.Blue {
			possibleBagContents.Blue = round.Blue
		}
		if round.Green > possibleBagContents.Green {
			possibleBagContents.Green = round.Green
		}
		if round.Red > possibleBagContents.Red {
			possibleBagContents.Red = round.Red
		}
	}
	return possibleBagContents
}

func part2(gamesRounds map[int][]round) int {
	sumPower := 0
	for gameNumber := range gamesRounds {
		minPossible := getPossibleBagContents(gamesRounds[gameNumber])
		sumPower += minPossible.Blue * minPossible.Green * minPossible.Red
	}
	return sumPower
}

func main() {

	input := utils.ReadInput()
	gamesRounds := extractGamesRounds(input)
	result_1 := part1(gamesRounds)
	fmt.Println("Part1:", result_1)
	result_2 := part2(gamesRounds)
	fmt.Println("Part2:", result_2)
}
