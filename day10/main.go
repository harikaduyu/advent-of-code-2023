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

func parse(input string) (map[coord]byte, coord, int, int) {
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
	return m, start, len(matrix), len(matrix[0])
}

func getNextPipesToVisitFromS(m map[coord]byte, start coord) []coord {
	nextPipes := make([]coord, 0)
	// up
	north := start.north()
	n, N := m[north]
	N = N && (n == '|' || n == '7' || n == 'F')
	if N {
		nextPipes = append(nextPipes, north)
	}

	// down
	south := start.south()
	s, S := m[south]
	S = S && (s == '|' || s == 'J' || s == 'L')
	if S {
		nextPipes = append(nextPipes, south)
	}

	// left
	east := start.east()
	e, E := m[east]
	E = E && (e == '-' || e == 'J' || e == '7')
	if E {
		nextPipes = append(nextPipes, east)
	}

	// right
	west := start.west()
	w, W := m[west]
	W = W && (w == '-' || w == 'L' || w == 'F')
	if W {
		nextPipes = append(nextPipes, west)
	}

	if len(nextPipes) != 2 {
		panic(fmt.Sprintf("Failed to find next pipes for S at i:%d, j:%d\nNumber of next pipes should be 2 but got %d.", start.i, start.j, len(nextPipes)))
	}

	m[start] = shapeOfS(N, S, E, W)
	return nextPipes
}

func shapeOfS(N, S, E, W bool) byte {
	switch {
	case N && S:
		return '|'
	case N && E:
		return 'L'
	case N && W:
		return 'J'
	case S && E:
		return 'F'
	case S && W:
		return '7'
	case W && E:
		return '-'
	default:
		panic("Can't find shape of S")
	}
}

func getNextPipesToVisit(m map[coord]byte, c coord, visited map[coord]bool) []coord {
	next := []coord{}
	n, s, w, e := c.north(), c.south(), c.west(), c.east()
	_, N := m[n]
	_, S := m[s]
	_, W := m[w]
	_, E := m[e]
	_, nVisited := visited[n]
	_, sVisited := visited[s]
	_, wVisited := visited[w]
	_, eVisited := visited[e]
	N = N && !nVisited
	S = S && !sVisited
	W = W && !wVisited
	E = E && !eVisited

	switch m[c] {
	case '|':
		if N {
			next = append(next, n)
		}
		if S {
			next = append(next, s)
		}
	case '7':
		if W {
			next = append(next, w)
		}
		if S {
			next = append(next, s)
		}
	case 'F':
		if E {
			next = append(next, e)
		}
		if S {
			next = append(next, s)
		}
	case 'J':
		if W {
			next = append(next, w)
		}
		if N {
			next = append(next, n)
		}
	case 'L':
		if E {
			next = append(next, e)
		}
		if N {
			next = append(next, n)
		}
	case '-':
		if W {
			next = append(next, w)
		}
		if E {
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
		queue = append(queue, element)
	}
	return queue
}

func dequeue(queue []coord) (coord, []coord) {
	element := queue[0]
	return element, queue[1:]
}

func part1(m map[coord]byte, start coord) (int, map[coord]byte, map[coord]bool) {
	toVisit := getNextPipesToVisitFromS(m, start)
	visited := map[coord]bool{start: true}
	var next coord
	for len(toVisit) > 0 {
		next, toVisit = dequeue(toVisit)
		nextPipes := getNextPipesToVisit(m, next, visited)
		visited[next] = true
		toVisit = enqueue(toVisit, nextPipes)
	}
	return (len(visited)) / 2, m, visited
}

func part2(m map[coord]byte, visited map[coord]bool, cL, rL int) int {
	// make everything that is not the loop, ground '.'
	for c := range m {
		_, inLoop := visited[c]
		if !inLoop {
			m[c] = '.'
		}
	}
	insideLoop := 0
	for i := 0; i < cL; i++ {
		in := false
		for j := 0; j < rL; j++ {
			loc := coord{i: i, j: j}
			val, exist := m[loc]
			if !exist {
				panic("Something's wrong with the loop in part2")
			}
			if val == '|' || val == 'L' || val == 'J' {
				in = !in
			} else if val == '.' && in {
				insideLoop++
			}
		}
	}
	return insideLoop
}

func main() {

	input := utils.ReadInput()
	startTime := time.Now()

	m, start, c, r := parse(input)
	res1, m, visited := part1(m, start)
	endTime := time.Now()
	fmt.Println("Part1:", res1)
	fmt.Printf("Part1 took %v\n", endTime.Sub(startTime))

	startTime = time.Now()
	res2 := part2(m, visited, c, r)
	endTime = time.Now()
	fmt.Println("Part2:", res2)
	fmt.Printf("Part2 took %v\n", endTime.Sub(startTime))
}
