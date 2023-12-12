package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	input := utils.ReadExampleInput(1)
	seeds, maps := parse(input)
	res1 := part1(seeds, maps)
	assert.Equal(t, 35, res1)

}

func Test2(t *testing.T) {

	input := utils.ReadExampleInput(1)
	seeds, maps := parse(input)
	res2 := part2(seeds, maps)
	assert.Equal(t, 46, res2)

}

func Test2Reversed(t *testing.T) {

	input := utils.ReadExampleInput(1)
	seeds, maps := parse(input)
	res2 := part2Reversed(seeds, maps)
	assert.Equal(t, 46, res2)

}
