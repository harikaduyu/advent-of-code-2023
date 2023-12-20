package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/utils"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	input := utils.ReadExampleInput(1)
	part11, part2 := result(input)
	assert.Equal(t, 4361, part11)
	assert.Equal(t, 467835, part2)

	input_2 := utils.ReadExampleInput(2)
	part12, _ := result(input_2)
	assert.Equal(t, 4361, part12)

}
