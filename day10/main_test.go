package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/utils"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	input := utils.ReadExampleInput(1)
	m, start := parse(input)
	res := part1(m, start)
	assert.Equal(t, 8, res)

}

func Test2(t *testing.T) {

	input := utils.ReadExampleInput(2)
	m, start := parse(input)
	res := part1(m, start)
	assert.Equal(t, 4, res)

}

// func Test3(t *testing.T) {

// 	input := utils.ReadExampleInput(3)

// 	instructions, networkMap := parse(input)
// 	res := part2(instructions, networkMap)

// 	assert.Equal(t, 6, res)

// }
