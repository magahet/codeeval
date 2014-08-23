package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Maze struct {
	M, N  int
	Value string
}

func (a *Maze) String() string {
	result := ""
	for i := 0; i < a.M; i++ {
		result += a.Value[i*a.M:i*a.M+a.N] + "\n"
	}
	return strings.Trim(result, "\n")
}

func (a *Maze) getStart() int {
	return strings.Index(a.Value[:a.N], " ")
}

func (a *Maze) getEnd() int {
	return strings.Index(a.Value[(a.M-1)*a.N:], " ") + ((a.M - 1) * a.N)
}

func (a *Maze) getAdjecent(i int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for d := -1; d <= 1; d += 2 {
			for _, j := range []int{i + d, i + d*a.N} {
				if j >= 0 && j < a.M*a.N && a.Value[j] == uint8(' ') {
					out <- j
				}
			}
		}
	}()
	return out
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (a *Maze) getDist(p1, p2 int) int {
	x1 := p1 % a.N
	y1 := p1 / a.N
	x2 := p2 % a.N
	y2 := p2 / a.N
	return Abs(x1-x2) + Abs(y1-y2)
}

func NewMaze(file string) Maze {
	rows := strings.Count(file, "\n")
	cols := strings.Index(file, "\n")
	return Maze{rows, cols, strings.Replace(file, "\n", "", -1)}
}

func getMinF(open map[int]bool, fScores map[int]int) int {
	minIndex := 0
	minF := 99999
	for i := range open {
		if fScores[i] < minF {
			minF = fScores[i]
			minIndex = i
		}
	}
	return minIndex
}

func (a *Maze) fillPath(cameFrom map[int]int) {
	path := make(map[int]bool)
	current := a.getEnd()
	path[current] = true
	path[a.getStart()] = true
	for current, exists := cameFrom[current]; exists; current, exists = cameFrom[current] {
		path[current] = true
	}
	newValue := ""
	for i, char := range a.Value {
		if _, isInPath := path[i]; isInPath {
			newValue += "+"
		} else {
			newValue += (string)(char)
		}
	}
	a.Value = newValue
}

func (a *Maze) findPath() {
	start := a.getStart()
	goal := a.getEnd()
	closed := make(map[int]bool)
	open := make(map[int]bool)
	open[start] = true
	cameFrom := make(map[int]int)

	gScore := make(map[int]int)
	fScore := make(map[int]int)
	gScore[start] = 0
	fScore[start] = a.getDist(start, goal)

	var current int
	for len(open) > 0 {
		current = getMinF(open, fScore)
		if current == goal {
			a.fillPath(cameFrom)
			return
		}

		delete(open, current)
		closed[current] = true

		for n := range a.getAdjecent(current) {
			if _, exists := closed[n]; exists {
				continue
			}

			if _, exists := open[n]; !exists || gScore[current]+1 < gScore[n] {
				cameFrom[n] = current
				gScore[n] = gScore[current] + 1
				fScore[n] = gScore[n] + a.getDist(n, goal)
				open[n] = true
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", path.Base(os.Args[0]), "file")
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}

	maze := NewMaze(string(file))
	maze.findPath()
	fmt.Println(maze.String())
}
