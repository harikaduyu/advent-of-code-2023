package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
)

type node struct {
	value, right, left string
}

func parseNode(line string) node {
	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	match := re.FindStringSubmatch(line)
	value, left, right := match[1], match[2], match[3]
	return node{value: value, right: right, left: left}
}

func parse(input string) (string, map[string]node) {
	networkMap := make(map[string]node)
	lines := strings.Split(input, "\n")

	instructions := lines[0]
	for _, line := range lines[2:] {
		node := parseNode(line)
		networkMap[node.value] = node
	}
	return instructions, networkMap
}

func findNextNode(networkMap map[string]node, instruction byte, curNode string) string {
	n := networkMap[curNode]
	if instruction == 'L' {
		return n.left
	} else if instruction == 'R' {
		return n.right
	} else {
		panic("Instruction is not L or R.")
	}
}

func part1(instructions string, networkMap map[string]node) int {
	numSteps := 0
	for nodeVal := "AAA"; nodeVal != "ZZZ"; {
		instruction := instructions[numSteps%len(instructions)]
		nodeVal = findNextNode(networkMap, instruction, nodeVal)
		numSteps++
	}
	return numSteps
}

func endWithZ(nodes []string) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node, "Z") {
			return false
		}
	}
	return true
}

func getNumSteps(nodeVal string, instructions string, networkMap map[string]node) int {
	var numSteps int
	for numSteps = 0; !strings.HasSuffix(nodeVal, "Z"); numSteps++ {
		instruction := instructions[numSteps%len(instructions)]

		nodeVal = findNextNode(networkMap, instruction, nodeVal)

	}
	return numSteps

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func part2(instructions string, networkMap map[string]node) int {

	// find all nodes that end with A
	result := 1
	for nodeVal := range networkMap {
		if strings.HasSuffix(nodeVal, "A") {
			numSteps := getNumSteps(nodeVal, instructions, networkMap)
			result = lcm(result, numSteps)
		}
	}
	return result
}

func main() {

	input := utils.ReadInput()
	startTime := time.Now()

	instructions, networkMap := parse(input)
	res1 := part1(instructions, networkMap)
	endTime := time.Now()
	fmt.Println("Part1:", res1)
	fmt.Printf("Part1 took %v\n", endTime.Sub(startTime))

	startTime = time.Now()
	res2 := part2(instructions, networkMap)
	endTime = time.Now()
	fmt.Println("Part2:", res2)
	fmt.Printf("Part2 took %v\n", endTime.Sub(startTime))
}
