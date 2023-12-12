package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
)

type MapDesc struct {
	DestRangeStart, SourceRangeStart, RangeLength int
}

func convertStringSliceToInt(s []string) []int {
	integers := make([]int, 0, len(s))
	for _, numStr := range s {
		n, _ := strconv.Atoi(numStr)
		integers = append(integers, n)
	}
	return integers
}

func parseMap(batch string) []MapDesc {
	lines := strings.Split(batch, "\n")
	mDescs := make([]MapDesc, 2)
	for _, line := range lines[1:] {
		ints := convertStringSliceToInt(strings.Fields(line))
		mDesc := MapDesc{ints[0], ints[1], ints[2]}
		mDescs = append(mDescs, mDesc)
	}
	return mDescs
}

func parse(input string) ([]int, [][]MapDesc) {
	maps := make([][]MapDesc, 7)
	batches := strings.Split(input, "\n\n")

	seedsStr := strings.Fields(strings.Split(batches[0], ": ")[1])
	seeds := convertStringSliceToInt(seedsStr)

	for i, batch := range batches[1:] {
		m := parseMap(batch)
		maps[i] = m
	}
	return seeds, maps
}

func findValFromMapDescs(seed int, mapDescs []MapDesc) int {
	item := seed
	for _, md := range mapDescs {
		if item >= md.SourceRangeStart && item < md.SourceRangeStart+md.RangeLength {
			return md.DestRangeStart + item - md.SourceRangeStart
		}
	}
	return item

}

// This is practically the same thing as above but for the reverse lookup
func findKeyFromMapDescs(seed int, mapDescs []MapDesc) int {
	item := seed
	for _, md := range mapDescs {
		if item >= md.DestRangeStart && item < md.DestRangeStart+md.RangeLength {
			return md.SourceRangeStart + item - md.DestRangeStart
		}
	}
	return item
}

func findLocation(seed int, maps [][]MapDesc) int {
	item := seed
	for _, m := range maps {
		item = findValFromMapDescs(item, m)
	}
	return item
}

func findSeed(location int, maps [][]MapDesc) int {

	item := location
	for i := len(maps) - 1; i >= 0; i-- {
		item = findKeyFromMapDescs(item, maps[i])
	}
	return item
}

func part1(seeds []int, maps [][]MapDesc) int {
	minLocation := 999999999999999999
	for _, seed := range seeds {
		location := findLocation(seed, maps)
		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}

func part2(seeds []int, maps [][]MapDesc) int {
	minLocation := 999999999999999999

	for i := 1; i < len(seeds); i += 2 {

		begin, end := seeds[i-1], seeds[i-1]+seeds[i]

		for seed := begin; seed < end; seed++ {
			// fmt.Println("seed", seed)
			location := findLocation(seed, maps)
			// fmt.Println("location", location)

			if location < minLocation {
				minLocation = location
			}
		}
	}
	return minLocation
}

// The optimization here is we do a reverse lookup starting
// from the location zero
func part2Reversed(seeds []int, maps [][]MapDesc) int {
	for location := 0; ; location++ {
		possibleSeed := findSeed(location, maps)
		for i := 1; i < len(seeds); i += 2 {
			begin, end := seeds[i-1], seeds[i-1]+seeds[i]
			if possibleSeed >= begin && possibleSeed < end {
				return location
			}
		}
	}
}

func main() {

	// input := utils.ReadExampleInput(1)
	input := utils.ReadInput()
	startTime := time.Now()

	seeds, maps := parse(input)
	res1 := part1(seeds, maps)
	endTime := time.Now()
	fmt.Println("Part1:", res1)
	fmt.Printf("Part1 took %v\n", endTime.Sub(startTime))
	// Part1 took 99.875µs

	startTime = time.Now()
	res2 := part2(seeds, maps)
	endTime = time.Now()
	fmt.Println("Part2:", res2)
	fmt.Printf("Part2 took %v\n", endTime.Sub(startTime))
	// Part2 took 2m18.43376325s

	startTime = time.Now()
	res2Rev := part2Reversed(seeds, maps)
	endTime = time.Now()
	fmt.Println("Part2 with reverse lookup:", res2Rev)
	fmt.Printf("Part2 with reverse lookup took %v\n", endTime.Sub(startTime))
	// Part2 with reverse lookup took 3.291763417s
}
