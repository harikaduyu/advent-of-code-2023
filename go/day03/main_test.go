package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	input := utils.ReadExampleInput(1)
	assert.Equal(t, 4361, part1(input))

	input_2 := utils.ReadExampleInput(2)
	assert.Equal(t, 4361, part1(input_2))

}

func TestPart2(t *testing.T) {

	input := utils.ReadExampleInput(1)
	assert.Equal(t, 4361, part1(input))

	input_2 := utils.ReadExampleInput(2)
	assert.Equal(t, 4361, part1(input_2))

}
