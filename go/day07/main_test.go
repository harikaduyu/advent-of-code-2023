package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	input := utils.ReadExampleInput(1)

	hands := parse(input, 1)
	res1 := result(hands)
	assert.Equal(t, 6440, res1)

}

func Test2(t *testing.T) {

	input := utils.ReadExampleInput(2)

	hands := parse(input, 1)
	res1 := result(hands)
	assert.Equal(t, 6592, res1)

}

func Test3(t *testing.T) {

	input := utils.ReadExampleInput(1)

	hands := parse(input, 2)
	res2 := result(hands)
	assert.Equal(t, 5905, res2)

}

func Test4(t *testing.T) {

	input := utils.ReadExampleInput(2)

	hands := parse(input, 2)
	res2 := result(hands)
	assert.Equal(t, 6839, res2)

}
