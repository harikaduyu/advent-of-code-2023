package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
)

func hasSymbolNearby(lines []string, i, j int) bool {
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[x]) && !(x == i && y == j) {
				if char := rune(lines[x][y]); !unicode.IsDigit(char) && char != '.' {
					return true
				}
			}
		}
	}
	return false
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	var currentNum int
	var totalNums int
	var symbolNearby bool
	for i, line := range lines {
		for j, char := range line {
			if unicode.IsDigit(char) {
				currentNum = (currentNum * 10) + int(char-'0')
				symbolNearby = symbolNearby || hasSymbolNearby(lines, i, j)

			}
			// We are at the end of the number if the current char is not digit
			// or it's the end of row
			if !unicode.IsDigit(char) || j == len(line)-1 {
				if symbolNearby {
					totalNums += currentNum
					symbolNearby = false
				}
				currentNum = 0
			}
		}
	}
	return totalNums
}

// func part2(input string) int {

// }

func main() {

	input := utils.ReadInput()
	res1 := part1(input)
	fmt.Println("Part1:", res1)
}
