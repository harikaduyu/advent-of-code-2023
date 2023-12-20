package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/harikaduyu/advent-of-code-2023/utils"
)

type coord struct {
	i, j int
}

func (c coord) north() coord {
	return coord{i: c.i - 1, j: c.j}
}
func (c coord) south() coord {
	return coord{i: c.i + 1, j: c.j}
}
func (c coord) east() coord {
	return coord{i: c.i, j: c.j + 1}
}
func (c coord) west() coord {
	return coord{i: c.i, j: c.j - 1}
}

func parse(input string) (map[coord]byte, coord) {
	m := make(map[coord]byte)
	var start coord
	matrix := strings.Split(input, "\n")
	for i := range matrix {
		for j := range matrix[0] {
			c := coord{i: i, j: j}
			m[c] = matrix[i][j]
			if m[c] == 'S' {
				start = c
			}
		}
	}
	return m, start
}

func getNextPipesToVisitFromS(m map[coord]byte, start coord) []coord {
	nextPipes := make([]coord, 0)
	// up
	north := start.north()
	n, ok := m[north]
	northpipes := map[byte]bool{'|': true, '7': true, 'F': true}
	if ok {
		_, ok := northpipes[n]
		if ok {
			nextPipes = append(nextPipes, north)
		}
	}

	// down
	south := start.south()
	s, ok := m[south]
	southPipes := map[byte]bool{'|': true, 'J': true, 'L': true}
	if ok {
		_, ok := southPipes[s]
		if ok {
			nextPipes = append(nextPipes, south)
		}
	}

	// left
	east := start.east()
	w, ok := m[east]
	eastPipes := map[byte]bool{'-': true, 'J': true, '7': true}
	if ok {
		_, ok := eastPipes[w]
		if ok {
			nextPipes = append(nextPipes, east)
		}
	}

	// right
	west := start.west()
	e, ok := m[west]
	westPipes := map[byte]bool{'-': true, 'L': true, 'F': true}
	if ok {
		_, ok := westPipes[e]
		if ok {
			nextPipes = append(nextPipes, west)
		}
	}
	if len(nextPipes) != 2 {
		panic(fmt.Sprintf("Failed to find next pipes for S at i:%d, j:%d\nNumber of next pipes should be 2 but got %d.", start.i, start.j, len(nextPipes)))
	}
	return nextPipes
}

func getNextPipesToVisit(m map[coord]byte, c coord, visited map[coord]bool) []coord {
	next := []coord{}
	n, s, w, e := c.north(), c.south(), c.west(), c.east()
	nV, nOk := m[n]
	sV, sOk := m[s]
	wV, wOk := m[w]
	eV, eOk := m[e]
	_, nVisited := visited[n]
	_, sVisited := visited[s]
	_, wVisited := visited[w]
	_, eVisited := visited[e]
	nOk = nOk && !nVisited && nV != 'S'
	sOk = sOk && !sVisited && sV != 'S'
	wOk = wOk && !wVisited && wV != 'S'
	eOk = eOk && !eVisited && eV != 'S'

	switch m[c] {
	case '|':
		if nOk {
			next = append(next, n)
		}
		if sOk {
			next = append(next, s)
		}
	case '7':
		if wOk {
			next = append(next, w)
		}
		if sOk {
			next = append(next, s)
		}
	case 'F':
		if eOk {
			next = append(next, e)
		}
		if sOk {
			next = append(next, s)
		}
	case 'J':
		if wOk {
			next = append(next, w)
		}
		if nOk {
			next = append(next, n)
		}
	case 'L':
		if eOk {
			next = append(next, e)
		}
		if nOk {
			next = append(next, n)
		}
	case '-':
		if wOk {
			next = append(next, w)
		}
		if eOk {
			next = append(next, e)
		}
	}
	if len(next) > 2 {
		fmt.Println("next to be visited: ", next)
		panic(fmt.Sprintf("Too many next elements, %d.\n", len(next)))
	}
	return next
}

func enqueue(queue []coord, elements []coord) []coord {
	for _, element := range elements {

		queue = append(queue, element) // Simply append to enqueue.
		// fmt.Println("Enqueued:", element)
	}
	return queue
}

func dequeue(queue []coord) (coord, []coord) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []coord{}
		return element, tmp

	}
	return element, queue[1:] // Slice off the element once it is dequeued.
}

// 7-F7-
// .FJ|7
// SJLL7
// |F--J
// LJ.LJ
func part1(m map[coord]byte, start coord) int {
	toVisit := getNextPipesToVisitFromS(m, start)
	visited := make(map[coord]bool)
	var next coord
	for len(toVisit) > 0 {
		// fmt.Println("toVisit", toVisit)
		// fmt.Println("visited", visited)
		next, toVisit = dequeue(toVisit)
		nextPipes := getNextPipesToVisit(m, next, visited)
		visited[next] = true
		toVisit = enqueue(toVisit, nextPipes)
	}
	// fmt.Println("visited", visited)
	return (len(visited) + 1) / 2
}

func main() {

	input := utils.ReadInput()
	startTime := time.Now()

	m, start := parse(input)
	res1 := part1(m, start)
	endTime := time.Now()
	fmt.Println("Part1:", res1)
	fmt.Printf("Part1 took %v\n", endTime.Sub(startTime))

	// startTime = time.Now()
	// res2 := part2(input)
	// endTime = time.Now()
	// fmt.Println("Part2:", res2)
	// fmt.Printf("Part2 took %v\n", endTime.Sub(startTime))
}
