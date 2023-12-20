package main

import (
	"fmt"
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	input := utils.ReadExampleInput(1)
	rounds := getNumWinningsPerRound(input)
	res1 := part1(rounds)
	assert.Equal(t, 13, res1)

}

func Test2(t *testing.T) {

	input := utils.ReadExampleInput(1)
	rounds := getNumWinningsPerRound(input)
	fmt.Println(rounds)
	res2 := part2(rounds)
	assert.Equal(t, 30, res2)

}
