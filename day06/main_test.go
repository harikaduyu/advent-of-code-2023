package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/utils"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	input := utils.ReadExampleInput(1)
	times, distances := parse(input)

	res1 := part1(times, distances)
	assert.Equal(t, 288, res1)

}

func Test2(t *testing.T) {

	input := utils.ReadExampleInput(1)
	times, distances := parse(input)
	res2 := part2(times, distances)
	assert.Equal(t, 71503, res2)

}
