package main

import (
	"fmt"
	"testing"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
)

var day = 1

func TestResult(t *testing.T) {

	testCases := []struct {
		part     int
		expected int
	}{
		{1, 142},
		{2, 281},
	}

	for _, tc := range testCases {
		input := utils.ReadExampleInput(tc.part)
		t.Run(fmt.Sprintf("Input: %s", input), func(t *testing.T) {
			result := result(input, tc.part)

			if result != tc.expected {
				t.Errorf("For part %d and input %s, got %d, expected %d", tc.part, input, result, tc.expected)
			}
		})
	}
}
