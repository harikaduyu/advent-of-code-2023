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
			if !(x == i && y == j) {
				if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[x]) {
					if char := rune(lines[x][y]); !unicode.IsDigit(char) && char != '.' {
						return true
					}
				}
			}
		}
	}
	return false
}

type position struct {
	x, y int
}

func checkStarNearby(lines []string, i, j int) (position, bool) {
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if !(x == i && y == j) {
				if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[x]) {
					if char := lines[x][y]; char == '*' {
						return position{x, y}, true
					}
				}
			}
		}
	}
	return position{-1, -1}, false
}

func result(input string) (int, int) {

	lines := strings.Split(input, "\n")
	currentNum := 0
	totalNums := 0
	symbolNearby := false
	starMap := make(map[position][]int)
	starNearby := position{}
	hasStarNearby := false
	totalGearRatio := 0
	for i, line := range lines {
		for j, char := range line {
			if unicode.IsDigit(char) {
				currentNum = (currentNum * 10) + int(char-'0')
				symbolNearby = symbolNearby || hasSymbolNearby(lines, i, j)
				Pos, exists := checkStarNearby(lines, i, j)
				if exists {
					starNearby = Pos
					hasStarNearby = exists
				}

			}
			// We are at the end of the number if the current char is not digit
			// or it's the end of row
			if !unicode.IsDigit(char) || j == len(line)-1 {
				if symbolNearby {
					totalNums += currentNum
					symbolNearby = false
				}
				if hasStarNearby && currentNum > 0 {
					starMap[starNearby] = append(starMap[starNearby], currentNum)
				}
				currentNum = 0
				hasStarNearby = false
			}
		}
	}
	for _, gears := range starMap {
		if len(gears) == 2 {
			gr := 1
			for _, gear := range gears {
				gr *= gear
			}
			totalGearRatio += gr
		}

	}

	return totalNums, totalGearRatio
}

func main() {

	input := utils.ReadInput()
	res1, res2 := result(input)
	fmt.Println("Part1:", res1)
	fmt.Println("Part2:", res2)
}
