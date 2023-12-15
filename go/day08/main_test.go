package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	input := utils.ReadExampleInput(1)

	instructions, networkMap := parse(input)
	res := part1(instructions, networkMap)

	assert.Equal(t, 2, res)

}

func Test2(t *testing.T) {

	input := utils.ReadExampleInput(2)

	instructions, networkMap := parse(input)
	res := part1(instructions, networkMap)

	assert.Equal(t, 6, res)

}

func Test3(t *testing.T) {

	input := utils.ReadExampleInput(3)

	instructions, networkMap := parse(input)
	res := part2(instructions, networkMap)

	assert.Equal(t, 6, res)

}
