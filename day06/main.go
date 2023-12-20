package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/harikaduyu/advent-of-code-2023/utils"
)

func getNumWaysCouldWin(time int, distance int) int {
	numWaysCouldWin := 0

	for i := time / 2; i >= 0; i-- {

		travel := i * (time - i)

		if travel > distance {
			numWaysCouldWin++
		} else {
			numWaysCouldWin *= 2
			if time%2 == 0 {
				numWaysCouldWin--
			}
			break
		}

	}

	return numWaysCouldWin

}

func parse(input string) ([]string, []string) {
	splitted := strings.Split(input, "\n")
	times := strings.Fields(strings.TrimSpace(strings.Split(splitted[0], ":")[1]))
	distances := strings.Fields(strings.TrimSpace(strings.Split(splitted[1], ":")[1]))
	return times, distances
}
func part1(times []string, distances []string) int {
	marginOfError := 1
	for i, ts := range times {
		time, _ := strconv.Atoi(ts)
		distance, _ := strconv.Atoi(distances[i])
		numWaysCouldWin := getNumWaysCouldWin(time, distance)

		marginOfError *= numWaysCouldWin
	}
	return marginOfError
}

func part2(times []string, distances []string) int {
	timeStr := strings.Join(times, "")
	distanceStr := strings.Join(distances, "")

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)
	return getNumWaysCouldWin(time, distance)
}

func main() {

	input := utils.ReadInput()
	times, distances := parse(input)
	startTime := time.Now()

	res1 := part1(times, distances)
	endTime := time.Now()
	fmt.Println("Part1:", res1)
	fmt.Printf("Part1 took %v\n", endTime.Sub(startTime))
	// 625ns

	startTime = time.Now()
	res2 := part2(times, distances)
	endTime = time.Now()
	fmt.Println("Part2:", res2)
	fmt.Printf("Part2 took %v\n", endTime.Sub(startTime))
	// 19.034ms
}
