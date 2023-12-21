package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/harikaduyu/advent-of-code-2023/utils"
)

type coord struct {
	i, j, xi, xj int
}

func parse(input string) []coord {
	lines := strings.Split(input, "\n")
	space := make([][]string, 0)
	for _, line := range lines {
		space = append(space, strings.Split(line, ""))
	}
	rowIdx, colIdx := expandingSpace(space)
	galaxies := make([]coord, 0)
	xRows := 0
	for i := range space {
		if xRows < len(rowIdx) && i >= rowIdx[xRows] {
			xRows++
		}
		xCols := 0
		for j := range space[0] {
			if xCols < len(colIdx) && j >= colIdx[xCols] {
				xCols++
			}
			if space[i][j] == "#" {
				c := coord{i: i, j: j, xi: xRows, xj: xCols}
				galaxies = append(galaxies, c)
			}
		}
	}
	return galaxies
}

func expandingSpace(space [][]string) ([]int, []int) {
	colIdx := make([]int, 0)
	for j := range space[0] {
		colHasGalaxy := false
		for i := range space {
			if space[i][j] == "#" {
				colHasGalaxy = true
			}
		}
		if !colHasGalaxy {
			colIdx = append(colIdx, j)
		}
	}

	rowIdx := make([]int, 0)
	for i := range space {
		rowHasGalaxy := false
		for j := range space[0] {
			if space[i][j] == "#" {
				rowHasGalaxy = true
			}
		}
		if !rowHasGalaxy {
			rowIdx = append(rowIdx, i)
		}
	}
	return rowIdx, colIdx
}

func calcDistance(x coord, y coord) int {
	distance := 0
	if x.i > y.i {
		distance += x.i - y.i
	} else {
		distance += y.i - x.i
	}

	if x.j > y.j {
		distance += x.j - y.j
	} else {
		distance += y.j - x.j
	}
	return distance
}

func calcDistances(galaxies []coord) int {
	totalDistances := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			totalDistances += calcDistance(galaxies[i], galaxies[j])
		}
	}
	return totalDistances
}

func expandSpace(galaxies []coord, expansionRate int) []coord {
	g := make([]coord, 0)
	for _, c := range galaxies {
		new_c := coord{
			i: c.i + c.xi*(expansionRate-1),
			j: c.j + c.xj*(expansionRate-1),
		}
		g = append(g, new_c)
	}
	return g
}

func result(galaxies []coord, expansionRate int) int {
	galaxies_1 := expandSpace(galaxies, expansionRate)
	return calcDistances(galaxies_1)
}

func main() {

	input := utils.ReadInput()
	startTime := time.Now()

	galaxies := parse(input)

	res1 := result(galaxies, 2)
	endTime := time.Now()
	fmt.Println("Part1:", res1)
	fmt.Printf("Part1 took %v\n", endTime.Sub(startTime))

	startTime = time.Now()
	res2 := result(galaxies, 1000000)
	endTime = time.Now()
	fmt.Println("Part2:", res2)
	fmt.Printf("Part2 took %v\n", endTime.Sub(startTime))
}
