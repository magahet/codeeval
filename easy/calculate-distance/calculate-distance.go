package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

type Point struct {
	x, y float64
}

func (p *Point) Dist(p2 *Point) int {
	dx := p.x - p2.x
	dy := p.y - p2.y
	return int(math.Sqrt(dx*dx + dy*dy))
}

func processLine(line string) int {
	parts := strings.Split(line, ") (")
	points := make([]*Point, 2)
	for i, pointStr := range parts {
		pointStr = strings.Trim(pointStr, "( )")
		coords := strings.Split(pointStr, ", ")
		x, _ := strconv.ParseFloat(coords[0], 64)
		y, _ := strconv.ParseFloat(coords[1], 64)
		points[i] = &Point{x, y}
	}

	return points[0].Dist(points[1])
}

func readLine(file *os.File) <-chan string {
	out := make(chan string)
	go func() {
		in := bufio.NewReader(file)
		linePartial := ""
		for {
			bytes, isPrefix, err := in.ReadLine()
			if err != nil {
				break
			} else if isPrefix {
				linePartial += string(bytes)
			} else {
				out <- linePartial + string(bytes)
				linePartial = ""
			}
		}
		close(out)
	}()
	return out
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", path.Base(os.Args[0]), "file")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	defer file.Close()

	if err != nil {
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}

	for line := range readLine(file) {
		fmt.Println(processLine(line))
	}
}
