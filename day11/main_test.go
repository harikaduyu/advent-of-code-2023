package main

import (
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/utils"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	input := utils.ReadExampleInput(1)
	galaxies := parse(input)

	res := result(galaxies, 2)
	assert.Equal(t, 374, res)

}

func Test2(t *testing.T) {

	input := utils.ReadExampleInput(1)
	galaxies := parse(input)

	res := result(galaxies, 10)
	assert.Equal(t, 1030, res)

}

func Test3(t *testing.T) {

	input := utils.ReadExampleInput(1)
	galaxies := parse(input)

	res := result(galaxies, 100)
	assert.Equal(t, 8410, res)

}
