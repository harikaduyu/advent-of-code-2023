package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/utils"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	input := utils.ReadExampleInput(1)
	m, start, _, _ := parse(input)
	res, _, _ := part1(m, start)
	assert.Equal(t, 8, res)

}

func Test2(t *testing.T) {

	input := utils.ReadExampleInput(2)
	m, start, _, _ := parse(input)
	res, _, _ := part1(m, start)
	assert.Equal(t, 4, res)

}

func Test3(t *testing.T) {

	input := utils.ReadExampleInput(3)
	m, start, c, r := parse(input)
	_, m, visited := part1(m, start)

	res := part2(m, visited, c, r)
	assert.Equal(t, 4, res)
}

func Test4(t *testing.T) {

	input := utils.ReadExampleInput(4)
	m, start, c, r := parse(input)
	_, m, visited := part1(m, start)

	res := part2(m, visited, c, r)
	assert.Equal(t, 8, res)

}

func Test5(t *testing.T) {

	input := utils.ReadExampleInput(5)
	m, start, c, r := parse(input)
	_, m, visited := part1(m, start)

	res := part2(m, visited, c, r)
	assert.Equal(t, 10, res)

}
