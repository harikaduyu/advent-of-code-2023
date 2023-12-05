package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
	"github.com/stretchr/testify/assert"
)

var day = 2

func TestPart1(t *testing.T) {

	input := utils.ReadExampleInput(day, 1)
	gamesRounds := extractGamesRounds(input)

	assert.Equal(t, 8, part1(gamesRounds))
}

func TestPart2(t *testing.T) {

	input := utils.ReadExampleInput(day, 1)
	gamesRounds := extractGamesRounds(input)

	assert.Equal(t, 2286, part2(gamesRounds))
}
