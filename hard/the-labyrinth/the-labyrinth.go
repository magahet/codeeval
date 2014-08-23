package main

import (
    "io/ioutil"
	"fmt"
	"os"
	"path"
	"strings"
)

type Maze struct {
    M, N int
    Value string
}

func (a *Maze) String() string {
    result := ""
    for i := 0; i < a.M; i++ {
        result += a.Value[i*a.M:i*a.M+a.N] + "\n"
    }
    return result
}

func (a *Maze) Start() Point {
    x := strings.Index(a.Value[:a.N], " ")
    s := Point{x, 0, 0, 0, 0, nil}
    h := getDist(s, a.End())
    s.H = h
    s.F = h
    return s
}

func (a *Maze) End() Point {
    i := strings.Index(a.Value[(a.M-1)*a.N:], " ")
    return Point{a.M-1, i, 0, 0, 0, nil}
}

func getNeighborIndexies (i, m, n int) <-chan int {
    out := make(chan int)
    go func () {
        defer close(out)
        for d := -1; d <= 1; d += 2 {
            for _, j := range []int{i+d, i+d*n} {
                if j >= 0 && j < len(m*n) {
                    out <- j
                }
            }
        }
    }()
    return out
}

func (a *Maze) getAdjecent(p *Point) <- chan Point {
    out := make(chan Point)
    go func() {
        defer close(out)
        i := p.Y * a.N + p.X
        var nP Point
        for j := range getNeighborIndexies(i, a.M, a.N) {
            if a.Value[i+di] == uint8('*') {
                continue
            }
            nP = Point{(j)%a.N, j/a.N, 0, 0, 0, p}
            nP.H = getDist(nP, a.End())
            nP.G = p.G + 1
            nP.F = nP.H + nP.G
            out <- nP
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

func getDist(p1, p2 Point) int {
    return Abs(p1.X - p2.X) + Abs(p1.Y - p2.Y)
}

func NewMaze(file string) Maze {
    rows := strings.Count(file, "\n") + 1
    cols := strings.Index(file, "\n")
    return Maze{rows, cols, strings.Replace(file, "\n", "", -1)}
}

type Point struct {
    X, Y, F, G, H int
    Parent *Point
}

func getMinF(list []Point) Point {
	index := 0
	for i, p := range list {
		if (i > 0) && (p.F <= list[index].F) {
			index = i
		}
	}
	return list[index]
}

func getPath(a Maze) []Point {
    open := []Point{a.Start()}
    var current Point
    for len(open) > 0 {
        current = getMinF(open)
        for _, p := a.getAdjecent(&current) {
            open = append(open, p)
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
	//fmt.Println(maze.Start(), maze.End())
	path := findPath(maze)
	fmt.Println(path)
}
